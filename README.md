# schedule-docker-tasks
Service to schedule simple docker-based tasks

## Architecture

* REST API service that essentially validates API requests and schedules task
* tasks are simple `(schedule, image, command)` tuples
* They are run as: `docker run --rm --entrypoint command image` on the given
schedule
* only simple, short-running docker tasks should be scheduled


## Running

### On Kubernetes
This is the recommended way. Deploy on k8s like so:

```
kubectl apply -f k8s_template.yml

```

### Locally
* The container essentially needs to talk to a docker daemon, so you can
run it locally as well. You do need to bind mount the socket file used to talk
the the daemon. Warning: this might lead to flood of docker containers on your
machine if many tasks are being scheduled

* Run it like so:

```
docker run --rm -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock pratikmallya/scheduler
```

## Talking to the Server
* Currently runs with no auth
* the cli can be run like so:

```
$ docker run pratikmallya/scheduler-cli -h
Use this cli to talk to the scheduler

Usage:
  task [command]

Available Commands:
  create      create a new task
  delete      Delete task with id
  help        Help about any command
  list        List all available tasks

Flags:
  -h, --help          help for task
  -i, --ip string     IP of scheduler (default "0.0.0.0")
  -p, --port string   Port of scheduler (default "8080")

Use "task [command] --help" for more information about a command.
```

Note that the k8s template deploys a NodePort service so make sure to use that
port instead of `8080`.
