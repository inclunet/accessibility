FROM ubuntu:latest
LABEL maintainer="inclunet"
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /accessbot
COPY ./lang/* /accessbot/lang/*
COPY ./cmd/accessbot/accessbot /usr/bin/accessbot
RUN mkdir reports
CMD accessbot
