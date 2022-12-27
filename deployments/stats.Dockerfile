FROM golang:1.19-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o ./shorturl ./main.go


FROM alpine:latest

WORKDIR /
COPY --from=build /build/shorturl /shorturl
USER nobody:nobody

CMD ["/shorturl", "consume-stats"]
