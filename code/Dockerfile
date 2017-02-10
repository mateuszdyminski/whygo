FROM golang

RUN mkdir /fanout
ADD fan-out-docker.go /app/
WORKDIR /app
RUN go build -o fanout .

EXPOSE 9000

CMD ["/app/fanout"]