package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()

  // create task
  router.POST("/tasks", func(c *gin.Context) {
  })

  // list tasks
  router.GET("/tasks", func(c *gin.Context) {
  })

  // delete specific task
  router.DELETE("/tasks/:task", func(c *gin.Context) {
  })

  // delete all tasks
  router.DELETE("/tasks", func(c *gin.Context) {
  })

	r.Run() // listen and serve on 0.0.0.0:8080
}
