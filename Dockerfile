FROM golang
MAINTAINER "Daniel Simone Beira"
ENV PORT=8887
COPY . /opt
WORKDIR /opt
RUN go build server/main.go
ENTRYPOINT ./main
EXPOSE $PORT
