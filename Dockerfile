FROM arm64v8/alpine
COPY . /app
WORKDIR /app
CMD ["/app/server"]