FROM golang:alpine AS builder
RUN apk add --no-cache git make
WORKDIR /usr/src
COPY . .
RUN make

FROM scratch
COPY --from=builder /usr/src/bin/uast /
CMD ["/uast"]
