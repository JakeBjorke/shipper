# consignment-cli

FROM golang:1.9.2 as builder
WORKDIR /go/src/github.com/jakebjorke/shipper/consignment-cli
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo . 

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jakebjorke/shipper/consignment-cli/consignment-cli .
#copy in the json file....probably won't need this for long...
COPY --from=builder /go/src/github.com/jakebjorke/shipper/consignment-cli/consignment.json .

CMD [ "./consignment-cli" ]