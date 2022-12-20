FROM golang:1.19-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /app ./...


FROM alpine:latest

WORKDIR /
COPY --from=build /app/app /app
USER nobody:nobody

CMD ["/app"]
