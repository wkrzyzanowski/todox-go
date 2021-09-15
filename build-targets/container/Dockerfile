FROM node:12.16-alpine3.9 as todox-frontend

LABEL stage=todox-frontend-build

WORKDIR /tmp/todox-frontend

COPY ./todox-frontend /tmp/todox-frontend/

RUN ls -lai

RUN npm install

RUN npm rebuild node-sass

RUN npm run lint

RUN npm run build

FROM golang:1.17-alpine as todox-backend

LABEL stage=todox-backend-build

RUN apk add --no-cache git

WORKDIR /tmp/todox-go

COPY ./todox-backend/go.mod .

COPY ./todox-backend/go.sum .

RUN go mod download

COPY ./todox-backend .

RUN go build -o ./out/todox-go .

FROM alpine:3.9

COPY --from=todox-frontend /tmp/webapp /app/webapp

COPY --from=todox-backend /tmp/todox-go/out/todox-go /app/api/todox-go

EXPOSE 8080

WORKDIR /app/api

CMD ["./todox-go"]