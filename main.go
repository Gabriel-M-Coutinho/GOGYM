package main

import (
	"log"

	userservice "gym/internal/application/user"
	"gym/internal/infrastructure/database"
	userrepo "gym/internal/infrastructure/persistence/user"
	userhandler "gym/internal/interfaces/http/user"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.NewDB("test.db")
	sqlDB, _ := database.GetSQLDB(db)

	if err != nil {
		log.Fatal("Falha ao conectar:", err)
	}

	// Extrair conexão SQL do GORM

	if err != nil {
		log.Fatal("Falha ao obter conexão SQL:", err)
	}

	// Injeção de dependências
	repo := userrepo.NewRepository(sqlDB)
	service := userservice.NewService(repo)
	handler := userhandler.NewHandler(service)

	r := gin.Default()

	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	r.Run(":8080")
}
