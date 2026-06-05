FROM golang:1.25.3-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o app-bin

FROM alpine
WORKDIR /app
COPY --from=build /app/app-bin .
EXPOSE 8080
CMD ["./app-bin"]

