#checkov:skip=CKV_DOCKER_2:We don't actually use this user.
#checkov:skip=CKV_DOCKER_3:We don't use _any_ user other than 0 because we don't use libc.
#checkov:skip=CKV_DOCKER_7:The alpine image is just to fetch the certificates.
# hadolint global ignore=DL3007,DL3018

FROM alpine:latest AS certificates
RUN apk add --no-cache ca-certificates

FROM scratch AS release
ENV COLUMNS 80
EXPOSE 80
COPY --from=certificates /etc/ssl/ /etc/ssl/
COPY go-importd /go-importd
ENTRYPOINT ["/go-importd"]
