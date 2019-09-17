FROM golang:latest
ENV PATH /go/bin:$PATH
ADD . /go/src/awesomeProject

WORKDIR /go/src/awesomeProject

# Set GO111MODULE=on variable to activate module support
ENV GO111MODULE on

RUN go build -o /go/bin/gqlServer -i ./cmd/gqlServer

CMD ["gqlServer"]