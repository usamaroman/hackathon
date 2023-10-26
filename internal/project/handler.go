package project

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/usamaroman/hackathon/internal/task"
	"github.com/usamaroman/hackathon/pkg/jwt"
)

const DDMMYYYY = "02.01.2006"

type handler struct {
	client *pgxpool.Pool
}

func NewHandler(pool *pgxpool.Pool) *handler {
	return &handler{
		client: pool,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.Handle(http.MethodPost, "/projects", jwt.Middleware(h.CreateProject))
	router.Handle(http.MethodGet, "/projects", jwt.Middleware(h.GetAllProjects))
	router.Handle(http.MethodGet, "/projects/:id", jwt.Middleware(h.GetProject))
	router.Handle(http.MethodPost, "/projects/:project_id/tasks/:task_id", jwt.Middleware(h.linkTask))
	router.Handle(http.MethodGet, "/projects/:id/tasks", jwt.Middleware(h.getAllProjectTasks))
}

func (h *handler) CreateProject(ctx *gin.Context) {
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

	var dto CreateProjectDto

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		log.Println("error during binding json", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse(DDMMYYYY, dto.Start)
	if err != nil {
		log.Println("time parsing")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	endDate, err := time.Parse(DDMMYYYY, dto.End)
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
	log.Println(startDate, endDate)

	query := `
		INSERT INTO projects (id, title, description, start, "end")
		VALUES ($1, $2, $3, $4, $5)
`

	exec, err := h.client.Exec(ctx, query, uuid.New().String(), dto.Title, dto.Description, startDate, endDate)
	if err != nil {
		log.Println("error during execution", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(exec.RowsAffected())

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "ok",
	})
}

func (h *handler) GetAllProjects(ctx *gin.Context) {
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

	query := `
		SELECT id, title, description, start, "end" FROM projects
	`

	rows, err := h.client.Query(ctx, query)
	if err != nil {
		return
	}

	var res []*Project

	for rows.Next() {
		var dto Project
		var start time.Time
		var end time.Time

		err := rows.Scan(&dto.Id, &dto.Title, &dto.Description, &start, &end)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		dto.Start = start.Format("02.01.2006")
		dto.End = end.Format("02.01.2006")

		res = append(res, &dto)
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) linkTask(ctx *gin.Context) {
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

	projectID := ctx.Param("project_id")
	taskID := ctx.Param("task_id")
	log.Println(projectID, taskID)

	_, err := h.client.Exec(ctx, "INSERT INTO projects_tasks (project_id, task_id) VALUES($1, $2)", projectID, taskID)
	if err != nil {
		log.Println("Error while adding task to project ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Задача добавлена в проект"})
}

func (h *handler) getAllProjectTasks(ctx *gin.Context) {
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
	var res []*task.Task
	log.Println(id)

	query := `
		SELECT id, title, description, start, "end", difficulty, priority, status
		FROM tasks 
		WHERE id in (
		    SELECT task_id FROM projects_tasks WHERE project_id = $1
		)
	`

	rows, err := h.client.Query(ctx, query, id)
	if err != nil {
		log.Println("Error while adding task to project ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var dto task.Task
		var start time.Time
		var end time.Time

		err := rows.Scan(&dto.Id, &dto.Title, &dto.Description, &start, &end, &dto.Difficulty, &dto.Priority, &dto.Start)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		dto.Start = start.Format(DDMMYYYY)
		dto.End = end.Format(DDMMYYYY)

		res = append(res, &dto)
		log.Println(res)
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) GetProject(ctx *gin.Context) {
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
	var project Project
	var start time.Time
	var end time.Time

	query := `
		SELECT id, title, description, start, "end" FROM projects WHERE id = $1
	`

	err := h.client.QueryRow(ctx, query, id).Scan(&project.Id, &project.Title, &project.Description, &start, &end)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	project.Start = start.Format(DDMMYYYY)
	project.End = end.Format(DDMMYYYY)

	ctx.JSON(http.StatusOK, project)
}
