FROM golang:1.23-alpine3.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /ims ./cmd/app/
RUN set -a && . /app/.env && set +a
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
EXPOSE 8080
CMD ["/ims"]
