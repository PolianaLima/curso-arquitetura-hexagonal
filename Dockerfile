FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"


RUN go get -u github.com/spf13/cobra@latest
RUN go install github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update
RUN apt-get install sqlite3 -y

CMD [ "tail", "-f", "/dev/null" ]
