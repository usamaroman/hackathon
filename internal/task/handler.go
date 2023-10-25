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
	priority_low  = "low"
	priority_mid  = "mid"
	priority_high = "high"
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
	task := ctx.Group("/task")
	{
		task.POST("/create", h.createTask)
		task.POST("/done/:id", h.taskDone)
		task.POST("/delete/:id", h.deleteTask)
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
		t.Title, t.Description, t.Difficulty, t.Priority, status_notStarted, t.Start, t.End).Scan(&insertedID)
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
		status_completed, id)
	if err != nil {
		log.Println("Error while writing to the database ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Задача выполнена"})
}
