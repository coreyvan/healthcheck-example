FROM golang:1.13.4-alpine as builder

WORKDIR /src
COPY . .
RUN go mod download

WORKDIR /src/cmd
RUN CGO_ENABLED=0 go build -o server server.go 
RUN CGO_ENABLED=0 go build -o healthcheck healthcheck.go
RUN chmod +x ./healthcheck
RUN ls

# FROM scratch
# COPY --from=builder /src/cmd/server .
# COPY --from=builder /src/cmd/healthcheck .

EXPOSE 8081
HEALTHCHECK --interval=5s CMD ["/src/cmd/healthcheck"]
ENTRYPOINT [ "/src/cmd/server" ]
