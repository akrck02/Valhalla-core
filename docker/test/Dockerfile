# Download base image ubuntu 22.04
FROM ubuntu:22.04

# LABEL about the custom image
LABEL maintainer="akrck02@gmail.com"
LABEL version="0.1"
LABEL description="This is a custom Docker Image for go test execution"

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Update Ubuntu Software repository
RUN apt update
RUN apt upgrade -y

# Install curl
RUN apt install -y curl
    
# Install go (1.20.3) 
RUN curl -OL https://golang.org/dl/go1.20.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz 

# Add go to the PATH environment variable
ENV PATH=$PATH:/usr/local/go/bin

# Add goPath to the PATH environment variable
RUN mkdir /cache
ENV GOPATH=/cache

# Make compile directiory
RUN mkdir /compile

# Volume configuration
VOLUME ["/compile", "/cache"]

# Copy start.sh script and define default command for the container
COPY start.sh /start.sh

CMD ["sh","/start.sh"]
