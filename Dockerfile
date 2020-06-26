FROM golang:1.13.1 AS builder
LABEL maintainer="ericchou19831101@msn.com"

ARG version="local"
ARG application="mattermost"

ENV GOOS=linux \
    GO111MODULE="on" \
    CGO_ENABLED=0

WORKDIR /src
COPY . ./

RUN go build -ldflags "-w -s -X main.version=${version} -X main.author=WenZhou" -o ${application}

RUN curl -H "X-JFrog-Art-Api:mySecretToken" \ 
    --progress-bar --upload-file ${application} \ 
    "https://artifactory.mycompany.com/artifactory/repo/${application}/${version}/${application}"




