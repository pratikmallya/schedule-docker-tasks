package main

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gopkg.in/robfig/cron.v2"
)

type dockerCrontab struct {
	Schedule string `json:"schedule" binding:"required"`
	Command  string `json:"command" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

func (d dockerCrontab) parse() (cron.Schedule, error) {
	return cron.Parse(d.Schedule)
}

func (d dockerCrontab) Run() {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "/bin/docker", "run", "--rm", "--entrypoint", fmt.Sprintf("%s", d.Command), fmt.Sprintf("%s", d.Image))
	cmd.Output()
}

func main() {
	router := gin.Default()
	crond := cron.New()
	// create task
	router.POST("/v1/tasks", func(c *gin.Context) {
		var crontabIN dockerCrontab

		err := c.BindJSON(&crontabIN)
		if err != nil {
			c.Error(err)
      c.JSON(400, gin.H{"error": err.Error()})
      return
		}
		sched, err := crontabIN.parse()
		if err != nil {
			c.Error(err)
      c.JSON(400, gin.H{"error": err.Error()})
      return
		}
		id := crond.Schedule(sched, crontabIN)
		c.JSON(201, gin.H{"id": id})
	})

	// list tasks
	router.GET("/v1/tasks", func(c *gin.Context) {
		var tasks []string
		entries := crond.Entries()
		for i := range entries {
			crontabIN := entries[i].Job.(dockerCrontab)
			tasks = append(tasks, fmt.Sprintf("%s %s %s", crontabIN.Schedule, crontabIN.Image, crontabIN.Command))
		}
		c.JSON(200, gin.H{"tasks": tasks})
	})

	// delete specific task
	router.DELETE("/v1/tasks/:taskid", func(c *gin.Context) {
		taskIDs := c.Param("taskid")
		taskIDi, err := strconv.Atoi(taskIDs)
		if err != nil {
			c.Error(err)
      c.JSON(400, gin.H{"error": err.Error()})
      return
		}
		crond.Remove(cron.EntryID(taskIDi))
		c.JSON(200, gin.H{})
	})

	crond.Start()
	router.Run() // listen and serve on 0.0.0.0:8080
}
