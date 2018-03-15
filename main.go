package main

import (
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type crontab struct {
	Schedule string `json:"schedule" binding:"required"`
	Command  string `json:"command" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

func main() {
	router := gin.Default()

	// create task
	router.POST("/v1/tasks", func(c *gin.Context) {
		var crontabIN crontab
		err := c.BindJSON(&crontabIN)
		if err != nil {
			panic(err)
		}
		crontabf := base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s.%s.%s", crontabIN.Schedule, crontabIN.Image, crontabIN.Command)))
		cronEntry := fmt.Sprintf("%s root docker run --entrypoint %s %s >>/opt/crond/logs/%s.log 2>&1", crontabIN.Schedule, crontabIN.Command, crontabIN.Image, crontabf)
		_, err = os.Stat(crontabf)
		if err == nil {
			c.JSON(200, gin.H{})
			return
		}
		err = ioutil.WriteFile(fmt.Sprintf("/opt/crond/crontabs/%s", crontabf), []byte(cronEntry), 0644)
		if err != nil {
			panic(err)
		}
		c.JSON(201, gin.H{"id": crontabf})
	})

	// list tasks
	router.GET("/v1/tasks", func(c *gin.Context) {
		taskFiles, err := ioutil.ReadDir("/opt/crond/crontabs/")
		if err != nil {
			panic(err)
		}
		var tasks []string
		for i := range taskFiles {
			tasks = append(tasks, taskFiles[i].Name())
		}
		c.JSON(201, gin.H{"tasks": tasks})
	})

	// delete specific task
	router.DELETE("/v1/tasks/:taskid", func(c *gin.Context) {
		taskID := c.Param("taskid")
		err := os.Remove(fmt.Sprintf("/opt/crond/crontabs/%s", taskID))
		if err != nil {
			panic(err)
		}
		c.JSON(204, gin.H{})
	})

	// delete all tasks
	router.DELETE("/v1/tasks", func(c *gin.Context) {
		err := os.RemoveAll("/opt/crond/crontabs/")
		if err != nil {
			panic(err)
		}
		c.JSON(204, gin.H{})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
