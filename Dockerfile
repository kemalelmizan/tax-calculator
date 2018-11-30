FROM golang:1.10

WORKDIR /go/src/tax-calculator
COPY ./src .

RUN go get -d -v ./...
RUN go build -o build/tax-calculator ./

ENV migrations_dir=file:///go/src/tax-calculator/migrations/
ENV db_username=postgres
ENV db_password=dummypassword
ENV db_name=tax_calculator
ENV db_host=localhost
ENV db_post=5432


CMD ["./build/tax-calculator", "start"]