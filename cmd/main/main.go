package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/usamaroman/hackathon/internal/auth"
	"github.com/usamaroman/hackathon/internal/user"
	"github.com/usamaroman/hackathon/pkg/client/postgresql"
)

func main() {
	ctx := context.Background()

	log.Println("gin init")
	router := gin.Default()
	router.Use(CORSMiddleware())

	//log.Println("minio init")
	//client := minio.New(cfg.Minio.Host, cfg.Minio.Port)
	//log.Println(client)

	log.Println("postgresql config init")
	//pgConfig := postgresql.NewPgConfig(
	//	cfg.PostgresStorage.Username,
	//	cfg.PostgresStorage.Password,
	//	cfg.PostgresStorage.Host,
	//	cfg.PostgresStorage.Port,
	//	cfg.PostgresStorage.Database,
	//)

	pgConfig := postgresql.NewPgConfig(
		"chechyotka",
		"5432",
		"localhost",
		"5432",
		"hackathon",
	)
	pgClient := postgresql.NewClient(ctx, pgConfig)
	_ = pgClient

	userRepository := user.NewRepository(pgClient)

	authService := auth.NewService(userRepository)
	authH := auth.NewHandler(authService)
	authH.Register(router)

	//handler := user2.NewHandler(repository, client)
	//handler.Register(router)
	//
	//grpcClient := grpc.NewCarsManagementClient(cfg.ElasticSearchMicroservice.Host, cfg.ElasticSearchMicroservice.Port)
	//carRepository := car2.NewRepository(pgClient)
	//imgRep := images_storage.NewRepository(pgClient)
	//reservationRep := reservation.NewRepository(pgClient)
	//carHandler := car.NewHandler(carRepository, imgRep, reservationRep, repository, grpcClient, client)
	//carHandler.Register(router)

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "health")
	})

	log.Println("http server init")
	port := fmt.Sprintf(":%d", 8000)

	server := http.Server{
		Handler:      router,
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("docs http://localhost:8000/health")
	log.Fatal(server.ListenAndServe())

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}