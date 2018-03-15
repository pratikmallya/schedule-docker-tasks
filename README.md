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
run it locally as well. Warning: this might lead to flood of docker containers
if many tasks are being scheduled

* Run it like so:

```
docker run --rm pratikmallya/scheduler -d
```

## Talking to the Server
* Currently runs with no auth
* the cli can be run like so:

```
docker run pratikmallya/scheduler-cli -h
```
as environment variables pointing to the server. Note that the k8s template
deploys a NodePort service so make sure to use that port instead of `8080`.
