FROM alpine:3.14
COPY . /app
CMD ["/app/server"]
