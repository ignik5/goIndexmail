FROM golang:1.18-alpine AS builder

WORKDIR /build
COPY go.mod  ./
RUN go mod download
COPY ./src ./src
COPY ./text ./text
RUN cd ./src && go build -ldflags "-s -w" -o exec

FROM alpine
WORKDIR  app
COPY --from=builder /build/text ./text

COPY --from=builder /build/src/exec ./exec
CMD ["./exec"]
