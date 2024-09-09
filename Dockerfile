FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o bin/drone-plugin .
ENTRYPOINT ["/bin/drone-plugin"]
