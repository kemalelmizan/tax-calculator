FROM golang:1.10

WORKDIR /go/src/tax-calculator
COPY . .

RUN go get -d -v ./...
RUN go build -o build/tax-calculator ./src
RUN ./build/tax-calculator migrate

CMD ["./build/tax-calculator", "start"]