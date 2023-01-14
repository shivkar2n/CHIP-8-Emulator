# Dockerfile for setting up emulator

# Get ubuntu image for docker hub
FROM ubuntu:22.04

# Move project directory to app directory
WORKDIR /app
COPY . .

# Initialize dpkg with timezone
RUN DEBIAN_FRONTEND="noninteractive" TZ="Etc/UTC" apt-get update && apt-get install -y tzdata 

# Update certificate for gopkg
RUN apt install -y ca-certificates && update-ca-certificates

# Install For go compiler and package manger
RUN apt install -y golang-go git

# Install SDL2 dependencies
RUN apt install -y libsdl2-dev libsdl2-image-dev libsdl2-mixer-dev libsdl2-ttf-dev libsdl2-gfx-dev

# Build go binary
RUN cd src && go mod download && go build -o ../build
