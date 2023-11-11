FROM golang:alpine AS builder
RUN apk --no-cache add tzdata
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o app ./cmd/routines-app/main.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/cmd/routines-app/.env /app/app ./
CMD [ "./app" ]