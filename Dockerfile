FROM alpine:latest
WORKDIR /app
COPY ./bin/linux/server ./server
EXPOSE 8080
CMD ["./server"]
