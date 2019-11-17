FROM golang:1.13.4-alpine as builder

WORKDIR /src
COPY . .
RUN go mod download

WORKDIR /src/cmd
RUN CGO_ENABLED=0 go build -o ./server/server ./server/server.go 
RUN CGO_ENABLED=0 go build -o ./healthcheck/healthcheck ./healthcheck/healthcheck.go
RUN chmod +x ./healthcheck

FROM scratch
COPY --from=builder /src/cmd/server/server .
COPY --from=builder /src/cmd/healthcheck/healthcheck .

EXPOSE 8081
HEALTHCHECK --interval=30s CMD ["/healthcheck"]
ENTRYPOINT [ "/server" ]
