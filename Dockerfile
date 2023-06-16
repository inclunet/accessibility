FROM ubuntu:latest
LABEL maintainer="inclunet"
WORKDIR /accessbot
COPY ./reports/templates/*.* /accessbot/reports/templates/*.*
COPY ./cmd/accessbot/accessbot /usr/bin/accessbot
# CMD accessbot
