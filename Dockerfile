FROM golang:1.15.6
WORKDIR /api
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
EXPOSE 5000
COPY main.go .
COPY go.mod .
COPY ./handlers ./handlers
COPY ./helpers ./helpers
RUN go build main.go
CMD ["./main"]
