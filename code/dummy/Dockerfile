FROM golang

RUN mkdir /dummy
ADD dummy-google.go /app/
WORKDIR /app
RUN go build -o dummy .

EXPOSE 9001 9002 9003

CMD ["/app/dummy", "9001"]