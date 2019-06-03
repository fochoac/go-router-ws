FROM golang
WORKDIR /go/src/github.com/fochoac/go-router-ws
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .
EXPOSE 8888
ENTRYPOINT ./go-route-ws