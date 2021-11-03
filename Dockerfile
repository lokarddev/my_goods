FROM golang
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
WORKDIR /app/cmd
RUN go build -o goods
EXPOSE 8000

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait


CMD ["/app/cmd/goods"]
