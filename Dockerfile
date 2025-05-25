FROM golang:1.24.3

RUN go version
ENV GOPATH=/


COPY ./ ./

# build go app
RUN go mod download
RUN go build -o go_todo_app ./cmd/main.go

CMD ["./go_todo_app"]