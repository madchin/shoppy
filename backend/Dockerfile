FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY .. --exclude=docker* Docker* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-bin

CMD ["/api-bin"]

FROM gcr.io/distroless/base-debian12 as build-release-stage

WORKDIR /

COPY --from=build-stage /api-bin /api-bin

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT [ "/api-bin" ]

