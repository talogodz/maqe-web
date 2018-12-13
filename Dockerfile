FROM golang:1.11.1-alpine3.8 as build-env
RUN mkdir /maqe
WORKDIR /maqe
COPY go.mod . 
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/maqe
FROM scratch 
COPY --from=build-env /go/bin/maqe /go/bin/maqe
ENTRYPOINT ["/go/bin/maqe"]