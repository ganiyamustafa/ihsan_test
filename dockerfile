FROM golang:1.19.2-alpine3.15 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY . ./

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o main .

## Distribution Development
FROM alpine:latest as dev-env

RUN apk add --no-cache --upgrade ca-certificates tzdata bash 

WORKDIR /app

COPY --from=builder /app/main ./

CMD [ "/app/main" ]

## Distribution Production
# FROM alpine:latest as prod-env

# RUN apk add --no-cache ca-certificates tzdata

# WORKDIR /app

# COPY --from=builder /app/main ./
# COPY --from=builder /app/config/app/production.json ./config/app/
# COPY --from=builder /app/rds-combined-ca-bundle.pem ./

# ENV APP_ENV=production

# CMD [ "/app/main" ]