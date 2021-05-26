# hasty-challenge-manager

This repository refers to backend hasty test `challenge/Hasty_BE_Challange.pdf`

----

## Technologies

+ [golang](https://golang.org/doc/install/source?download=go1.16.4.src.tar.gz) ⚡
+ [postgres](https://www.postgresql.org/)🐘
+ [make](https://www.gnu.org/software/make/) ❤️
+ [docker](https://www.docker.com/) 🐋

## Architecture

TODO

## Environment

| Env         |  URL                                           |
|-------------|----------------------------------------------- |
| Local       |  http://localhost:9000                         |
| Kubernetes  |  `Run your cluster`                            |

## Install

``` bash
$ make install
```

## DB Dependencies

```sh
$ make docker/up
```

## Run API

This will up an API in port 9000 by default

``` bash
$ make run
```

## Schedule Checker

This will find jobs that has timeout or should be retried or resumed.

### Configure crontab locally

Use `crontab –e` to every five minutes `*/5 * * * *`.

### Run

``` bash
$ make run-schedule
```

## Tests

``` bash
$ make test
```

## Deployment

### Docker

Generate and push a docker image to registry.

``` bash
$ make docker/registry
```

## Kubernetes
The deploy directory contains yaml files to deploy to a kubernetes cluster. These yaml files are validated for continuous integration, but not deployed.

## CI/CD

This project has a simple integration with github actions to run automated tests and validate kubernetes yaml file.
