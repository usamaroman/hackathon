package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/usamaroman/hackathon/pkg/client/postgresql"
	"log"
	"net/http"
	"time"
)

// Priority
type priority string

const (
	LowPriority  = "LowPriority"
	MidPriority  = "MidPriority"
	HighPriority = "LowPriority"
)

// Status
type status string

const (
	notStarted = "notStarted"
	started    = "started"
	completed  = "completed"
	postponed  = "postponed"
)

type task struct {
	Id          int    `json:"task_id"`
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
	ctx.Group("/task")
	{
		ctx.POST("/create", h.createTask)
	}
}

func (h *handler) createTask(ctx *gin.Context) {
	var t task
	err := ctx.ShouldBindJSON(t)
	if err != nil {
		log.Println("Error while serializing json ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Валидация: верю Роберту

	h.storage.Query(ctx, postgresql.FormatQuery(
		`INSERT INTO tasks(task_id, title, description, difficulty, priority, status, start, end)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8`))
	if err != nil {
		log.Println("Error while writing to database ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Задача создана"})
}
