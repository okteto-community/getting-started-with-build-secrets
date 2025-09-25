FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY main.go .

# my_name is the secret defined in the okteto.yaml. Okteto will replace the ENV VAR with the corresponding value on container image build time
RUN --mount=type=secret,id=my_name,target=/run/secrets/name.txt cp /run/secrets/name.txt /app/name.txt && \
    go build -o main main.go && \
    # we delete it so it won't be persisted on the rest of the layers
    rm /app/name.txt
    

    FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]