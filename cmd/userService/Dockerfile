FROM golang:latest
ENV PATH /go/bin:$PATH
ADD . /go/src/awesomeProject

WORKDIR /go/src/awesomeProject

# Set GO111MODULE=on variable to activate module support
ENV GO111MODULE on
RUN go build -o /go/bin/userService -i ./cmd/userService
#ENV MONGO_CONNECT_STRING mongodb://localhost/wmsight-devdb-copy
ARG MONGO_CONNECT_STRING
RUN echo $MONGO_CONNECT_STRING
ENV MONGO_CONNECT_STRING $MONGO_CONNECT_STRING

CMD ["userService"]