#Build stage
FROM golang:1.20.2-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
       
#Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
RUN chmod +x start.sh
RUN chmod +x wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]

    

