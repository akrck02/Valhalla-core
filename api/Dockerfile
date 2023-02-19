FROM golang:latest

ENV PORT=$PORT

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

COPY . ./
RUN go build -o /valhalla-core

EXPOSE $PORT

CMD [ "/valhalla-core" ]