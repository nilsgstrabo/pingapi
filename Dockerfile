FROM golang:1.19-alpine as builder

RUN apk update && \
    apk add bash jq alpine-sdk sed gawk git ca-certificates curl mc && \
    apk add --no-cache gcc musl-dev 

RUN go install honnef.co/go/tools/cmd/staticcheck@v0.3.3 

WORKDIR /go/src/github.com/nilsgstrabo/pingapi/

# get dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy api code
COPY . .

# lint and unit tests
RUN staticcheck ./... && \
    go vet `go list ./...` && \
    CGO_ENABLED=0 GOOS=linux go test `go list ./...`

# Build radix vulnerability scanner API go project
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/pingpong

RUN addgroup -S -g 1000 radix-vuln-scanner
RUN adduser -S -u 1000 -G radix-vuln-scanner radix-vuln-scanner

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/pingpong /usr/local/bin/pingpong
COPY --from=builder /etc/passwd /etc/passwd
USER 1000
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/pingpong"]