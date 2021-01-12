FROM golang:1.15-alpine AS build

RUN apk add --update \
  git \
  build-base \
  ca-certificates \
  bash

WORKDIR /var/app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/server main.go

# production image
FROM alpine:3.11

RUN apk add --no-cache ca-certificates

WORKDIR /var/app

COPY --from=build /bin/* ./

EXPOSE 8080
# COPY entrypoint.sh vault.sh ./

# ENTRYPOINT ["/var/app/entrypoint.sh"]
# CMD ["/var/app/server"]