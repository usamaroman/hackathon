package task

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/usamaroman/hackathon/pkg/jwt"
)

type priority string

const DDMMYYYY = "02.01.2006"

const (
	PriorityLow  = "низкий"
	PriorityMid  = "средний"
	PriorityHigh = "высокий"
)

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

type Comment struct {
	Text string `json:"text"`
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
		task.PATCH("/done/:id", jwt.Middleware(h.taskDone))
		task.DELETE("/:id", jwt.Middleware(h.deleteTask))
		task.POST("/taskToProj", jwt.Middleware(h.taskToProject))
		task.POST("/:id/comments", jwt.Middleware(h.commentTask))
		task.GET("/:id/comments", jwt.Middleware(h.getAllComments))
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

// todo
func (h *handler) taskToProject(ctx *gin.Context) {
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

func (h *handler) commentTask(ctx *gin.Context) {
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

	taskID := ctx.Param("id")
	log.Println(taskID)

	var dto Comment
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		log.Println("Error while serializing JSON ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentID := uuid.New().String()

	query := `
		INSERT INTO comments (id, text, user_id) VALUES ($1, $2, $3)
	`

	exec, err := h.storage.Exec(ctx, query, commentID, dto.Text, token["id"])
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(exec.RowsAffected())

	query = `
		INSERT INTO tasks_comments (task_id, comment_id) VALUES ($1, $2)
	`

	exec, err = h.storage.Exec(ctx, query, taskID, commentID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(exec.RowsAffected())

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "комментарий создан",
	})
}

func (h *handler) getAllComments(ctx *gin.Context) {
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

	taskID := ctx.Param("id")
	log.Println(taskID)

	query := `
		SELECT c.id, c.text, c.created_at, u.email 
		FROM comments as c
		INNER JOIN public.tasks_comments tc on c.id = tc.comment_id
		INNER JOIN public.users as u on u.id = c.user_id
		WHERE tc.task_id = $1
	`

	var res []*GetComment

	rows, err := h.storage.Query(ctx, query, taskID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for rows.Next() {
		var c GetComment
		var createdAt time.Time

		err = rows.Scan(&c.ID, &c.Text, &createdAt, &c.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.CreatedAt = createdAt.Format(DDMMYYYY)

		res = append(res, &c)
	}

	ctx.JSON(http.StatusCreated, res)
}

type GetComment struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
