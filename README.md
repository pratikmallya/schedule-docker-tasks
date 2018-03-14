# schedule-docker-tasks
Service to schedule simple docker-based tasks

## Architecture

* REST API service that essentially validates API requests and configures the crontab
* `crond` daemon to read the crontab and schedule tasks
* tasks are simple `(image, command)` tuples
* They are run as: `docker run --entrypoint command image` by `crond`

## Notes

* only simple, short-running docker tasks should be scheduled
* designed to be run in kubernetes with dind. See the deployment file for details
