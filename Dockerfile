FROM golang

WORKDIR /go/src/app
COPY . .

ENV TZ America/Recife

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]