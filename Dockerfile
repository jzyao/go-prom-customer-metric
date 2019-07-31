FROM golang:1.7.3
WORKDIR /go/src/github.com/yilei/go-prom-customer-metrics/
RUN go get -d -v github.com/prometheus/client_golang/prometheus
RUN go get -d -v github.com/prometheus/client_golang/prometheus/promauto
RUN go get -d -v github.com/prometheus/client_golang/prometheus/promhttp

COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/yilei/go-prom-customer-metrics/app .
CMD ["./app"] 
