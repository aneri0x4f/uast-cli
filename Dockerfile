# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git make
WORKDIR /usr/src
COPY . .
RUN make

# final stage
FROM gcr.io/distroless/base:latest
COPY --from=builder /usr/src/bin/uast /
CMD ["/uast"]
