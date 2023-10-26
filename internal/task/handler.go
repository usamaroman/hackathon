package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"time"
)

// Priority
type priority string

const (
	PriorityLow  = "low"
	PriorityMid  = "mid"
	PriorityHigh = "high"
)

// Status
type status string

const (
	StatusNotStarted = "notStarted"
	StatusInProcess  = "inProcess"
	StatusCompleted  = "completed"
	StatusPostponed  = "postponed"
)

type task struct {
	Id          int    `json:"id"`
	Title       string `json:"title" required:"true"`
	Description string `json:"description"`
	//Comment     string    `json:"comment"`
	Difficulty int       `json:"difficulty"`
	Priority   priority  `json:"priority"`
	Status     status    `json:"status"`
	Start      time.Time `json:"start" `
	End        time.Time `json:"end"`
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
		task.POST("/create", h.createTask)
		task.PATCH("/done/:id", h.taskDone)
		task.DELETE("/:id", h.deleteTask)
		task.POST("/taskToProj", h.taskToProject)
	}
}

func (h *handler) createTask(ctx *gin.Context) {
	var t task
	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		log.Println("Error while serializing JSON ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Валидация: верю Роберту

	// Выполнение SQL-запроса
	var insertedID int
	err = h.storage.QueryRow(ctx,
		`INSERT INTO tasks(title, description, difficulty, priority, status, start, "end")
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		t.Title, t.Description, t.Difficulty, t.Priority, StatusNotStarted, t.Start, t.End).Scan(&insertedID)
	if err != nil {
		log.Println("Error while writing to the database ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": insertedID, "message": "Задача создана"})
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
