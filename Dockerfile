FROM golang:latest as build

WORKDIR /go/src/go.lafronz.com/fileServer

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o runtime /go/src/go.lafronz.com/fileServer

FROM scratch
COPY --from=build /go/src/go.lafronz.com/fileServer/runtime /app
COPY --from=build /go/src/go.lafronz.com/fileServer/static /static
ENTRYPOINT ["/app"]