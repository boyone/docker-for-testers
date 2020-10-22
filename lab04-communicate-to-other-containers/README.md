# Docker Network 101

## Build Hello-World-API and Time-API Images

```sh
# time-api
cd time
docker build -t time-api:0.0.1 .
```

```sh
# hello-world-api
cd helloworld
docker build -t hello-world-api:0.0.1 .
```

## Vai Host (Without Docker Compose)

```sh
docker container run --rm -p 8000:8000 -e PORT=8000 --name time time-api:0.0.1
```

```sh
# Start time-api
docker container run -d -p 8000:8000 -e PORT=8000 --name time time-api:0.0.1
```

```sh
# Start hello-world-api
docker container run -d -p 8080:8080 --name helloworld hello-world-api:0.0.1
```

```sh
# List running containers
docker container ls
```

```sh
ifconfig
```

```sh
# Start hello-world-api with -e TIME_SERVER
docker container run -d -p 8080:8080 -e TIME_SERVER=172.20.10.3 --name helloworld hello-world-api:0.0.1
```

```sh
curl http://localhost:8080/hello
```

```sh
docker container rm -f  `docker container ls -aq`
docker container rm -f  $(docker container ls -aq)
```
