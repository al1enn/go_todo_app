FROM golang:1.24.3

RUN go version
ENV GOPATH=/


COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o go_todo_app ./cmd/main.go

CMD ["./go_todo_app"]