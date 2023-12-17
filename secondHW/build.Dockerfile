FROM golang:1.21.1

RUN apt-get update -y && apt-get upgrade -y \
	&& apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/
