package lib

import (
  "os/exec"
  "fmt"

  "gopkg.in/robfig/cron.v2"
  "golang.org/x/net/context"
)
// Task is a struct for holding task create data
type Task struct {
	Schedule string `json:"schedule" binding:"required"`
	Command  string `json:"command" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

// Parse parses crontab schedule
func (d Task) Parse() (cron.Schedule, error) {
	return cron.Parse(d.Schedule)
}

// Run executes the Task
func (d Task) Run() {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "/bin/docker", "run", "--rm", "--entrypoint", fmt.Sprintf("%s", d.Command), fmt.Sprintf("%s", d.Image))
	cmd.Output()
}
