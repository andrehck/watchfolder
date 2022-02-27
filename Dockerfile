FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN apk add build-base
RUN go mod download

COPY . ./

RUN go build -o /spybond

CMD [ "/spybond" ]
