package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"online-learning-platform/internal/config"
	"online-learning-platform/internal/db"
	"online-learning-platform/internal/repository"
	"online-learning-platform/internal/rest/handlers"
	"online-learning-platform/internal/rest/routers"
	"online-learning-platform/pkg/compose"
	"online-learning-platform/pkg/logger"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func initializeDB() config.Database {
	dbConfig := config.Database{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
		Name:     os.Getenv("POSTGRES_NAME"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	return dbConfig
}

func initializeSession() config.Session {
	sessionConfig := config.Session{
		Secret: os.Getenv("SESSION_SECRET_KEY"),
		Name:   os.Getenv("SESSION_NAME"),
		Key:    os.Getenv("SESSION_KEY"),
	}

	return sessionConfig
}

func initializeRedis() config.RedisSessionConfig {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error converting REDIS_DB to int: %s", err)
	}

	sessionConfig := initializeSession()

	redisConfig := config.RedisSessionConfig{
		ConnectionTimeoutSeconds: os.Getenv("REDIS_CONNECTION_TIMEOUT_SECONDS"),
		NetworkType:              os.Getenv("REDIS_NETWORK_TYPE"),
		Host:                     os.Getenv("REDIS_HOST"),
		Port:                     os.Getenv("REDIS_PORT"),
		Password:                 os.Getenv("REDIS_PASSWORD"),
		DB:                       redisDB,
		Session:                  sessionConfig,
	}

	return redisConfig
}

var appConfig config.App

func main() {
	logger.InitLogger()
	err := compose.StartDockerComposeService()
	if err != nil {
		logger.GetLogger().Fatal("Error starting Docker Compose service:", err)
	}
	logger.GetLogger().Info("Docker Compose service started successfully!")

	appConfig = config.App{
		PORT:  os.Getenv("APP_PORT"),
		DB:    initializeDB(),
		Redis: initializeRedis(),
	}

	dbInstance, err := db.GetDBInstance(appConfig.DB)
	if err != nil {
		logger.GetLogger().Fatal("Error initializing DB:", err)
	}

	userRepo := repository.NewUserRepository(dbInstance)
	//courseRepo := repository.NewCourseRepository(dbInstance)
	//lessonRepo := repository.NewLessonRepository(dbInstance)

	authHandlers := handlers.NewAuthHandlers(userRepo)
	//newsHandlers := handlers.NewNewsHandlers(newsRepo, tagRepo)
	//tagHandlers := handlers.NewTagHandlers(tagRepo, newsRepo)

	r := gin.Default()

	router := routers.NewRouters(*authHandlers)
	router.SetupRoutes(r)
	r.Use(rateLimitMiddleware())

	server := &http.Server{
		Addr:    ":" + appConfig.PORT,
		Handler: r,
	}

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		logger.GetLogger().Info("Server is shutting down...")

		// TODO: Uncomment these lines when the project finished
		//if err := compose.StopDockerComposeService(); err != nil {
		//	logger.GetLogger().Error("Error stopping Docker Compose service:", err)
		//}
		//logger.GetLogger().Info("Docker Compose service stopped successfully!")

		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.GetLogger().Fatal("Server shutdown error:", err)
		}

		logger.GetLogger().Info("Server has gracefully stopped")
		os.Exit(0)
	}()

	logger.GetLogger().Info("Server is running on :" + appConfig.PORT)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.GetLogger().Fatal("Error starting server:", err)
	}
}

func rateLimitMiddleware() gin.HandlerFunc {
	limiter := time.Tick(time.Second)

	return func(c *gin.Context) {
		select {
		case <-limiter:
			c.Next()
		default:
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
		}
	}
}
