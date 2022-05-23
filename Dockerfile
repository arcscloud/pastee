FROM alpine
COPY . /app
WORKDIR /app
CMD ["/app/server"]
