FROM golang:latest

RUN mkdir /temp

COPY /cmd /go/src/cmd
COPY /internal /go/src/internal
COPY /ui /go/src/ui
COPY /go.mod /go/src
COPY /go.sum /go/src

RUN go build -C /go/src/cmd/web -o /temp/frank

EXPOSE 8080
EXPOSE 3306

ENTRYPOINT [ "/temp/frank" ]
