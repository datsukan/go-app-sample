package infrastructure

import (
	"github.com/gin-gonic/gin"

	"portfolio/app/interfaces/controllers"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users", func(c *gin.Context) { usersController.Index(c) })
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Show(c) })
	r.Gin.POST("/users", func(c *gin.Context) { usersController.Store(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
