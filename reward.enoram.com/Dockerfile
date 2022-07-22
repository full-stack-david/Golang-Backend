# step 1: build
FROM golang:1.12-alpine3.10 as build-step

# for go mod download
RUN apk add --update --no-cache ca-certificates git
RUN mkdir /reward.enoram.com

WORKDIR /reward.enoram.com
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

#ENTRYPOINT go run main.go
#EXPOSE 8080

RUN CGO_ENABLED=0 go build -o /deploy/reward.enoram.com
# -----------------------------------------------------------------------------
# step 2: exec
FROM scratch

COPY --from=build-step /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-step /deploy/reward.enoram.com /deploy/reward.enoram.com

COPY --from=build-step /reward.enoram.com/config/*.json /deploy
COPY --from=build-step /reward.enoram.com/*.json /deploy/


ENTRYPOINT ["/deploy/reward.enoram.com", "/deploy/"]
EXPOSE 8080