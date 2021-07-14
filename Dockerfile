FROM golang:alpine3.13 AS build

WORKDIR /go/src/app
COPY . .

RUN go build -v ./...
RUN go build -o bin/sql-isready cmd/sql-isready/main.go

FROM alpine:latest AS final
WORKDIR /app
COPY --from=build /go/src/app/bin/sql-isready .
CMD "/app/sql-isready"