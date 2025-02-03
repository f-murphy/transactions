package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	logger "bank/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}

	logFile, err := logger.InitLogger()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading logrus")
	}
	logrus.Info("logFile initialized successfully")
	defer logFile.Close()

	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		logrus.WithError(err).Fatal("failed to initialize db")
	}
	logrus.Info("Database connected successfully")
	defer conn.Close(context.Background())
	
	repo := repository.NewRepository(conn)
    service := service.NewTransactionService(repo)
    handler := handler.NewTransactionHandler(service)

	r := gin.Default()
	r.POST("/transferMoney", handler.TransferMoney)
	r.POST("/replenishment", handler.Replenishment)
	r.GET("/transactions/:userID", handler.GetLatestTransactions)
	
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); 
		err != nil && err != http.ErrServerClosed {
			logrus.Fatal("error - ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); 
	err != nil {
		logrus.WithError(err).Error("Failed to shut down server")
	}
	logrus.Info("Server shut down")
}
