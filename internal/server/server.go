package server

import (
	// import third-party packages
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/butterfly/internal/db"
	"github.com/impossible98/butterfly/internal/server/api"
)

func InitServer() {
	db.InitDb()
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/version", api.Version)
	router.Run()
}
