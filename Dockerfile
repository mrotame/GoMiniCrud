FROM golang:1.20.6-bullseye
WORKDIR /app
COPY . .
RUN go mod download
RUN go build
EXPOSE 8001
ENTRYPOINT ["./GoCrudApi"]
# ENTRYPOINT ["tail", "-f", "/dev/null"]
