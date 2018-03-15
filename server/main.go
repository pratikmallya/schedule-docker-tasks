package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/robfig/cron.v2"

  "github.com/pratikmallya/schedule-docker-tasks/lib"
)



func main() {
	router := gin.Default()
	crond := cron.New()

  // create task
	router.POST("/v1/tasks", func(c *gin.Context) {

    var task lib.Task
		err := c.BindJSON(&task)
		if err != nil {
			c.Error(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
		}
		sched, err := task.Parse()
		if err != nil {
			c.Error(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
		}
		id := crond.Schedule(sched, task)
		c.JSON(201, gin.H{"id": id})

  })


  // list tasks
	router.GET("/v1/tasks", func(c *gin.Context) {

  	var tasks []string
		entries := crond.Entries()
		for i := range entries {
			task := entries[i].Job.(lib.Task)
			tasks = append(tasks, fmt.Sprintf("%s %s %s", task.Schedule, task.Image, task.Command))
		}
		c.JSON(http.StatusOK, gin.H{"tasks": tasks})

  })

	// delete specific task
	router.DELETE("/v1/tasks/:taskid", func(c *gin.Context) {
		taskIDs := c.Param("taskid")
		taskIDi, err := strconv.Atoi(taskIDs)
		if err != nil {
			c.Error(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
		}
		crond.Remove(cron.EntryID(taskIDi))
		c.JSON(http.StatusOK, gin.H{})
	})


	crond.Start()
	router.Run() // listen and serve on 0.0.0.0:8080
}
