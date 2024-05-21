FROM golang:1.22-alpine3.19 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /apriori

FROM alpine:3.19 AS build-release-stage

WORKDIR /

COPY --from=build-stage /apriori /apriori

EXPOSE 3007

ENTRYPOINT ["./apriori"]
