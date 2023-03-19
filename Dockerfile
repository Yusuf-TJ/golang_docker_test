FROM golang:1.20-alpine3.17
#  AS build

WORKDIR /app
COPY . .
RUN go build -o app

# FROM alpine:latest
# WORKDIR /app
# COPY --from=build /app/app .
EXPOSE 8000
CMD ["./app"]
