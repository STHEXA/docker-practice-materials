FROM ubuntu:latest

RUN apt-get update && \
    apt-get install cowsay -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/usr/games/cowsay", "Hello, Docker!"]
