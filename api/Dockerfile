FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

COPY . ./
RUN go build -o /valhalla-core

EXPOSE 3333

CMD [ "/valhalla-core" ]