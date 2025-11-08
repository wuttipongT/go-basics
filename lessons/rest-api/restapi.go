package restapi

import (
	"example.com/jeff/go-basics/lessons/rest-api/db"
	"example.com/jeff/go-basics/lessons/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func New() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":3000") // localhost:8080
}

func main() {

}
