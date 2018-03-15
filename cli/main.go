package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "bytes"

  "github.com/spf13/cobra"
  "gopkg.in/robfig/cron.v2"

  "github.com/pratikmallya/schedule-docker-tasks/lib"
)

var hostIP string
var hostPort string

var cmd = &cobra.Command{
	Use:   "task",
	Short: "Task scheduler CLI",
	Long: "Use this cli to talk to the scheduler",
}

var create = &cobra.Command{
  Use:   "create [schedule] [command] [image]",
  Short: "create a new task",
  Args: cobra.ExactArgs(3),
  Run: func(cmd *cobra.Command, args []string) {
    _, err := cron.Parse(args[0])
    if err != nil {
      panic(err)
    }
    b, err := json.Marshal(lib.Task{
      Schedule: args[0],
      Command: args[1],
      Image: args[2],
    })
    if err != nil {
      panic(err)
    }
    req, err := http.NewRequest("POST", fmt.Sprintf("http://{hostIP}:{hostPort}/v1/tasks", hostIP, hostPort), bytes.NewBuffer(b))
    if err != nil {
      panic(err)
    }
    req.Header.Add("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(body)
  },
}

var list = &cobra.Command{
  Use:   "list",
  Short: "List all available tasks",
  Args: cobra.NoArgs,
  Run: func(cmd *cobra.Command, args []string) {
    resp, err := http.Get(fmt.Sprintf("http://{hostIP}:{hostPort}/v1/tasks", hostIP, hostPort))
    if err != nil {
      panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      panic(err)
    }
    fmt.Println(body)
  },
}

var delete = &cobra.Command{
  Use:   "delete [id]",
  Short: "Delete task with id",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Deleting task with id: %s", args[0])
    req, err := http.NewRequest("DELETE", fmt.Sprintf("http://{hostIP}:{hostPort}/v1/tasks/{id}",hostIP, hostPort, args[0]), nil)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      panic(err)
    }
    if resp.StatusCode != http.StatusOK {
      panic(fmt.Sprintf("Got %d from server", resp.StatusCode))
    }
    fmt.Println("OK")
  },
}

func main() {
  cmd.PersistentFlags().StringVarP(&hostIP, "ip", "i", "0.0.0.0", "IP of scheduler")
  cmd.PersistentFlags().StringVarP(&hostPort, "port", "p", "8080", "Port of scheduler")
  cmd.AddCommand(create, list, delete)
  cmd.Execute()
}
