FROM alpine

WORKDIR '/app'
COPY go-poc .

CMD ./go-poc