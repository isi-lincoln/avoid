FROM golang:bullseye as build

RUN apt update && apt upgrade -qy
RUN apt install -y \
    build-essential \
    golang \
    make \
    ca-certificates \
    protobuf-compiler \
    vim

RUN mkdir -p /go
ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH "$PATH:$GOBIN"
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY . /avoid
WORKDIR /avoid

RUN rm -rf build
RUN make dns

FROM ubuntu:22.04 as final

RUN apt update && apt install -qy \
    iproute2

COPY --from=build /avoid/build/avoid* /usr/bin/

CMD /usr/bin/avoid-dns
