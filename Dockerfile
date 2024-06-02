FROM golang:latest

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air", "-c", ".air.toml"]