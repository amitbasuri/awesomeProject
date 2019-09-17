FROM golang:latest
ENV PATH /go/bin:$PATH
ADD . /go/src/awesomeProject

WORKDIR /go/src/awesomeProject

# Set GO111MODULE=on variable to activate module support
ENV GO111MODULE on
ARG MONGO_CONNECT_STRING
RUN go build -o /go/bin/userService -i ./cmd/userService

CMD ["userService"]