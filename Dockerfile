FROM golang:alpine as builder

WORKDIR /app

COPY . .

ARG svcVersion
ARG svcName

RUN cd cmd/rest && CGO_ENABLED=0 GOOS=linux go build -a -o /app/bin/depto \
    --ldflags "-extldflags \"static\" -X main.svcName=${svcName} -X main.svcVersion=${svcVersion}" \
    -tags netgo -installsuffix netgo

RUN cd cmd/migrations && CGO_ENABLED=0 GOOS=linux go build -a -o /app/bin/goose

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /usr/bin

COPY bin/* /usr/bin/

COPY --from=builder /app/bin/depto .
COPY --from=builder /app/bin/goose .

EXPOSE 8000

CMD ["/bin/sh", "-l", "-c", "wait-pg-db && goose -dsn=${DATABASE_URL} up && depto"]
