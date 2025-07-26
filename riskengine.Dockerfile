FROM golang:1.24

WORKDIR /app
COPY mock_riskengine.go .
RUN go build -o riskengine mock_riskengine.go

EXPOSE 9090
CMD ["./riskengine"]
