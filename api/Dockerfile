FROM golang:latest

ENV APP_HOME /src/
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 3333
CMD go build;GIN_MODE=release ./valhalla-core