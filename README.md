# schedule-docker-tasks
Service to schedule simple docker-based tasks

## Architecture

* REST API service that essentially validates API requests and configures the
crontab
* tasks are simple `(schedule, image, command)` tuples
* They are run as: `docker run --entrypoint command image` on the given
schedule

## Notes

* only simple, short-running docker tasks should be scheduled
* designed to be run in kubernetes with dind. Deploy on k8s like so:

```
kubectl apply -f k8s_template.yml
```

* the cli is deployed in the pod. Access it like so:
```
kubectl exec -it <POD_NAME> -c cli -h
```
It can also be run from outside the pod:
```
docker run -e "$SCHEDULER_HOST_IP" -e "$SCHEDULER_HOST_PORT" pratikmallya/scheduler-cli -h
```
as environment variables pointing to the server. Note that the k8s template
deploys a NodePort service so make sure to use that port instead of `8080`.
