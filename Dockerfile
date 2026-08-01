FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /out/fcpd ./cmd/fcpd

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/fcpd /usr/local/bin/fcpd
COPY examples/basic-provider/context.json /data/context.json
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/fcpd"]
CMD ["-listen", ":8080", "-catalog", "/data/context.json"]
