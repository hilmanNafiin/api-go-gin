FROM golang:1.22.3
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bin
EXPOSE 8080
ENTRYPOINT ["./bin"]


# docker build . -t hilmannafiin009/api-go:latest -f Dockerfile --platform=linux/amd64