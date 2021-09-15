# ToDoX - stay productive

## Powered by GoLang and VueJs

&nbsp;

# Development

During development activities it is possible to run backend and frontend part of application separately.

&nbsp;

## Backend

Application will run on **http://localhost:8080/**

```sh
cd todox-backend/ && \
 go run server.go
```

&nbsp;

## Frontend

Application will run on **http://localhost:8081/**

```sh
cd todox-frontend/ && \
 npm install && \
 npm run serve
```

When application is running on dev-server all requests prefixed with **/api** are redirected to the default backend which is set to **http://localhost:8080/**

&nbsp;

# Docker

## Build image

In order to build image let's execute following command:

```sh
{gitroot}/build-targets/buildImage.sh
```

It will build image with complete application which is runnable via:

```sh
docker run -d -p 8080:8080 todox:1.0.0
```

&nbsp;

# Kubernetes (Helm)

Feture will be available soon...
