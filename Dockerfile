
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY model/*.go ./model/
COPY controller/*.go ./controller/
COPY view/* ./view/

#RUN go env -w GO111MODULE=off

RUN go build -o /example-api

EXPOSE 8080

CMD [ "/example-api" ]