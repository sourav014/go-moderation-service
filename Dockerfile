FROM golang:1.23.0
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o go-moderation-service .
EXPOSE 8080
CMD [ "./go-moderation-service" ]