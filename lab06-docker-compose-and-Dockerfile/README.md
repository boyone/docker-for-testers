# docker-compose and Dockerfile

## Build -> Start Steps (Up to technology stacks and programming language)

### Steps

- Compile
- Test
- Package
- Start

### Golang

```sh
# Download dependencies
go mod download

# Package
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start app
./app
```

### Node.js

```sh
# Download dependencies
npm install

# Start app
npm start
```

## Docker Compose

```sh
docker-compose up
```

```sh
docker-compose up -d
```

```sh
docker-compose up --build --force-recreate
```

## Call API

```sh
curl http://localhost:8080/hello
```

## Check hello-world logs in Database

```sh
# Login to container
docker container exec -it hello-world-db sh

# Login to db
mysql -uroot -p

# Enter password: `password`
Enter password: password

# Tell db to use helloworld database
use helloworld;

# Query data
select * from access_log;

# Exit db
quit;

# Exit container
exit
```
