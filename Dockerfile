## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /todo-api-app

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /todo-api-app /todo-api-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/todo-api-app"]