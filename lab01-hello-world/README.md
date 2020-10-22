# Hello, World

```sh
docker run hello-world
```

```sh
docker container run hello-world
```

```sh
docker image ls
```

```sh
docker container ls -a
```

## Run with -d

```sh
docker container run boyone:hello-world
```

```sh
docker container run -d boyone:hello-world
```

```sh
docker container ls
```

```sh
docker container logs <container id| container name>
```

```sh
docker container logs -f <container id| container name>
```

## Start, Stop, and Remove Containers

```sh
docker container stop <container id| container name>
```

```sh
docker container start <container id| container name>
```

```sh
docker container rm <container id| container name>
```

```sh
docker container prune
```

## Docker Create Your Own Image with Tag

```sh
docker container run -d boyone:hello-world
```

```sh
docker container cp ./hello <container id| container name>:/root/
```

```sh
docker container commit <container id| container name> image_name:tag
```

## Push Your Own Images to Docker Hub

- Sign Up [https://hub.docker.com](https://hub.docker.com)
- Login

```sh
docker login
```

- Push Image to hub.docker.com

```sh
docker image push image_name:tag
```

## Remove All Images

```sh
docker image rm image_name
```

## Run Your Own Images

```sh
docker container run image_name
```

## Artifacts: <name>/image_name:tags

- name/hello-world:0.0.1
- name/hello-world:0.0.2