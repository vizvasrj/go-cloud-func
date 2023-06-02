FROM golang:1.19

WORKDIR /app
COPY . .


CMD [ "go", "run", "main.go" ]