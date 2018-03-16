[![Build Status](https://travis-ci.org/pratikmallya/schedule-docker-tasks.svg?branch=master)](https://travis-ci.org/pratikmallya/schedule-docker-tasks)

# schedule-docker-tasks
Service to schedule simple docker-based tasks

## Architecture

* REST API service that essentially validates API requests and schedules task
* tasks are simple `(schedule, image, command)` tuples
* They are run as: `docker run --rm --entrypoint command image` on the given
schedule
* only simple, short-running docker tasks should be scheduled


## Running

### Kubernetes
This is the recommended way. Deploy on k8s like so:

```
$ kubectl apply -f k8s_template.yml
```

### Locally
* The container essentially needs to talk to a docker daemon, so you can
run it locally as well. You do need to bind mount the socket file used to talk
the the daemon. Warning: this might lead to flood of docker containers on your
machine if many tasks are being scheduled

* Run it like so:

```
$ docker run --rm -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock pratikmallya/scheduler
```

## CLI

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

### Kubernetes
You will need to specify the public IP of the Service, as well as the nodeport:
```
$ docker run pratikmallya/scheduler-cli -i 169.60.205.31 -p 30004 list
```

### Locally
If running locally, specify `networking` as host:
```
$ docker run --network host pratikmallya/scheduler-cli list
{"tasks":null}
```
