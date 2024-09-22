FROM ubuntu:latest
LABEL authors="evan"

ENTRYPOINT ["top", "-b"]