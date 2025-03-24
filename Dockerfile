#docker build -t go-oauth .
#docker run -dit --name go-oauth -p 5100:5100 go-oauth sleep infinity


FROM golang:1.23.3 As builder

RUN apt-get update && apt-get install bash && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /app
COPY . .
RUN go mod tidy

WORKDIR /app/cmd
RUN go build -o go-oauth -ldflags '-linkmode external -w -extldflags "-static"'

FROM alpine

WORKDIR /app
COPY --from=builder /app/cmd/go-oauth .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/assets/certs/server-private.key ../assets/certs/
COPY --from=builder /app/assets/certs/server-public.key ../assets/certs/
COPY --from=builder /app/assets/certs/crl-ca.crl ../assets/certs/

CMD ["/app/go-oauth"]