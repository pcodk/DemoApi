FROM golang:1.11-alpine as builder
RUN apk add gcc git
COPY ./ /workdir
WORKDIR /workdir

ARG GIT_REVISION
RUN CGO_ENABLED=0 GOOS=linux go install -tags netgo ./...


FROM alpine
RUN apk update && apk add ca-certificates
LABEL vendor="Peytz & Co A/S"
LABEL author="Jan-Erik Revsbech"
COPY --from=builder /go/bin/apiserver .
COPY ./config ./config/
CMD ["./apiserver"]
