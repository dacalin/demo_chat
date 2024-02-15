############################
# STEP 1 build executable binary
############################
FROM golang:1.21 AS builder

RUN apt-get update


WORKDIR /src
COPY ./src /src
# Fetch dependencies.
# Using go get.
RUN go clean -modcache
RUN go mod tidy
RUN go mod download
# Build the binary.
RUN env GOOS=linux GOARCH=arm64 go build -o /go/bin/app ./main.go


############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/app /go/bin/app
# Run the binary.
CMD ["/go/bin/app"]


