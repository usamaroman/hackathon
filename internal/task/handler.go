package task

import (
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/usamaroman/hackathon/pkg/jwt"
	"log"
	"net/http"
	"time"
)

// Priority
type priority string

const DDMMYYYY = "02.01.2006"

const (
	PriorityLow  = "низкий"
	PriorityMid  = "средний"
	PriorityHigh = "высокий"
)

// Status
type status string

const (
	StatusNotStarted = "надо сделать"
	StatusInProcess  = "в процессе"
	StatusCompleted  = "выполнено"
)

type Task struct {
	Id          string   `json:"id"`
	Title       string   `json:"title" required:"true"`
	Description string   `json:"description"`
	Difficulty  int      `json:"difficulty"`
	Priority    priority `json:"priority"`
	Status      status   `json:"status"`
	Start       string   `json:"start" `
	End         string   `json:"end"`
}

type handler struct {
	storage *pgxpool.Pool
}

func New(storage *pgxpool.Pool) *handler {
	return &handler{storage}
}

func (h *handler) Register(ctx *gin.Engine) {
	task := ctx.Group("/tasks")
	{
		task.GET("/", jwt.Middleware(h.getAllTasks))
		task.POST("/", jwt.Middleware(h.createTask))
		task.PATCH("/done/:id", h.taskDone)
		task.DELETE("/:id", h.deleteTask)
		task.POST("/taskToProj", h.taskToProject)
	}
}

func (h *handler) getAllTasks(ctx *gin.Context) {
	value, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, "не авторизован")
		return
	}
	token, ok := value.(jwt2.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, "не авторизован")
		return
	}
	role := token["role"]
	log.Println(role)

	rows, err := h.storage.Query(ctx, `SELECT id, title, description, start, "end", difficulty, priority, status FROM tasks`)
	if err != nil {
		log.Println("Error while querying tasks", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res []*Task

	for rows.Next() {
		var t Task
		var start time.Time
		var end time.Time

		err := rows.Scan(&t.Id, &t.Title, &t.Description, &start, &end, &t.Difficulty, &t.Priority, &t.Status)
		if err != nil {
			log.Println("Error while scanning task row", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		t.Start = start.Format(DDMMYYYY)
		t.End = end.Format(DDMMYYYY)

		res = append(res, &t)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error while iterating over task rows", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) createTask(ctx *gin.Context) {
	value, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, "не авторизован")
		return
	}
	token, ok := value.(jwt2.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, "не авторизован")
		return
	}
	role := token["role"]
	log.Println(role)

	var t Task
	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		log.Println("Error while serializing JSON ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse(DDMMYYYY, t.Start)
	if err != nil {
		log.Println("time parsing")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	endDate, err := time.Parse(DDMMYYYY, t.End)
	if err != nil {
		log.Println("time parsing")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	now := time.Now()
	if startDate.Before(now) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "u wanna to rent car in the past",
		})
		return
	}

	_, err = h.storage.Exec(ctx,
		`INSERT INTO tasks (id, title, description, start, "end", difficulty, priority, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		uuid.New().String(), t.Title, t.Description, startDate, endDate, t.Difficulty, t.Priority, StatusNotStarted)
	if err != nil {
		log.Println("Error while writing to the database ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Задача создана"})
}

func (h *handler) deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	// Выполнение SQL-запроса
	_, err := h.storage.Exec(ctx, `DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		log.Println("Error while deleting from the database ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}

func (h *handler) taskDone(ctx *gin.Context) {
	id := ctx.Param("id")

	// Выполнение SQL-запроса
	_, err := h.storage.Exec(ctx,
		`UPDATE tasks
			SET status = $1
			WHERE id = $2`,
		StatusCompleted, id)
	if err != nil {
		log.Println("Error while writing to the database ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Задача выполнена"})
}

func (h *handler) taskToProject(ctx *gin.Context) {
	taskID := ctx.Query("taskID")
	projID := ctx.Query("projID")

	_, err := h.storage.Exec(ctx, "INSERT INTO project_task (project_id, task_id) VALUES($1, $2)", projID, taskID)
	if err != nil {
		log.Println("Error while adding task to project ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Задача добавлена в проект"})
}
