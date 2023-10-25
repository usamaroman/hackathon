package project

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
	router.Handle(http.MethodGet, "/projects", h.GetAllProjects)
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
	//h.client.Exec(ctx)
}

//
//func (h *handler) Registration(ctx *gin.Context) {
//	var body CreateProjectDto
//	err := ctx.ShouldBindJSON(&body)
//	if err != nil {
//		log.Println("error during binding json", err)
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = ValidateForEmptyPassword(body.Password)
//	if err != nil {
//		log.Println("error during validating password and full name", err)
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = h.service.Registration(ctx, body)
//	if err != nil {
//		log.Println("error during registration process", err)
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusCreated, gin.H{"message": "успешная регистрация"})
//}
//
//func (h *handler) Login(ctx *gin.Context) {
//	var body LoginDto
//	err := ctx.ShouldBindJSON(&body)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("wrong entered data").Error()})
//		return
//	}
//
//	user, token, err := h.service.Login(ctx, body)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	var res LoginResDto
//	res.User = user
//	res.AccessToken = token
//
//	refreshToken, err := jwt.GenerateRefreshToken(user.Id)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
//		return
//	}
//
//	res.RefreshToken = refreshToken
//
//	ctx.JSON(http.StatusOK, res)
//}
//
//func (h *handler) RefreshToken(ctx *gin.Context) {
//	var reqDto struct {
//		RefreshToken string `json:"refresh_token"`
//	}
//
//	err := ctx.ShouldBindJSON(&reqDto)
//	if err != nil {
//		log.Println(err)
//		ctx.JSON(http.StatusUnauthorized, gin.H{
//			"error": "не авторизован",
//		})
//		return
//	}
//	log.Println(reqDto)
//
//	if reqDto.RefreshToken == "" {
//		fmt.Println("empty")
//		ctx.JSON(http.StatusUnauthorized, gin.H{
//			"error": "не авторизован",
//		})
//		return
//	}
//
//	token, err := jwt.ParseRefreshTokenToken(reqDto.RefreshToken)
//	if err != nil {
//		fmt.Println("hz")
//		ctx.JSON(http.StatusUnauthorized, gin.H{
//			"error": "не авторизован",
//		})
//		return
//	}
//
//	id := token["id"]
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "server error",
//		})
//		return
//	}
//
//	u, err := h.service.repository.GetOneUserById(ctx, fmt.Sprintf("%s", id))
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "server error",
//		})
//		return
//	}
//
//	role, err := h.service.repository.GetRole(ctx, u.Id)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "server error",
//		})
//		return
//	}
//
//	generateAccessToken, err := jwt.GenerateAccessToken(u.Id, u.Email, role)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "server error",
//		})
//		return
//	}
//
//	generateRefreshToken, err := jwt.GenerateRefreshToken(u.Id)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "server error",
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"access_token":  generateAccessToken,
//		"refresh_token": generateRefreshToken,
//	})
//}
//
//func (h *handler) Logout(ctx *gin.Context) {
//	ctx.SetCookie("access_token", "", -1, "/", "127.0.01", false, false)
//	ctx.JSON(http.StatusOK, gin.H{
//		"message": "log out",
//	})
//}
