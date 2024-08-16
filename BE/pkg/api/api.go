package api

import (
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/api/healthCheck"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(r *gin.Engine, db *mongo.Database, client *mongo.Client) {
	g := r.Group("/api")
	{
		healthCheck.APIHealthCheckServices(g)
	}
}
