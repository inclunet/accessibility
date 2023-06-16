FROM ubuntu:latest
LABEL maintainer="inclunet"
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /accessbot
COPY ./reports/templates/*.* /accessbot/reports/templates/*.*
COPY ./cmd/accessbot/accessbot /usr/bin/accessbot
# CMD accessbot
