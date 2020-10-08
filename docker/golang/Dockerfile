# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.14

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
# RUN apt-get update && apt-get  upgrade && \
#   apt-get add --virtual build-dependencies build-base gcc wget && \
#   apt-get add --no-cache bash git openssh


RUN apt-get update && apt-get install -y \
  build-essential \
  gcc \
  wget \
  zip \
  unzip \
  git \
  curl \
  vim  

# Add Maintainer Info
LABEL maintainer="j3ns789@gmail.com"

WORKDIR /app

ENV GOPATH /app/vendor
ENV GOBIN ${GOPATH}/bin
ENV PATH ${GOBIN}:${PATH}

ENV TZ=Asia/Ho_Chi_Minh

RUN go get -u github.com/cosmtrek/air