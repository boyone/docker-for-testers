# Working with API

## Run Store Service

```sh
docker container run boyone/store-service-stub
```

## Store Service Dockerfile

```dockerfile
FROM andyrbell/mountebank:2.3.2
WORKDIR /imposters
COPY ./imposters .
EXPOSE 2525
EXPOSE 8000
CMD [ "mb", "start", "--configfile", "/imposters/store-service.json", "--allowInjection" ]
```

## Run Store Service with -p

```sh
docker container run -p 2525:2525 -p 8000:8000 boyone/store-service-stub
```

## When Static files has been modified

## Create New Images

## Run Plain Container with -v

```sh
docker container run -p 2525:2525 andyrbell/mountebank:2.3.2
```

```sh
docker container run -p 2525:2525 -d --name mb andyrbell/mountebank:2.3.2
```

```sh
docker container run -v `pwd`/imposters:/imposters -p 2525:2525 -p 8000:8000 -d --name mb andyrbell/mountebank:2.3.2 mb start --configfile /imposters/store-service.json --allowInjection
```

```sh
docker container exec -it mb sh
```

## Docker compose

```sh
docker-compose up
```

```sh
docker-compose -f docker-compose-build.yml up
```
