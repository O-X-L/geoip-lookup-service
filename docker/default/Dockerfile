FROM golang:alpine AS builder-be

# BUILD: docker build -f Dockerfile -t geoip-lookup --network=host --no-cache ../..
# RUN: docker run -it --rm --name geoip-lookup -p 10000:10000 --volume /var/local/geoip:/var/local/geoip geoip-lookup

COPY . /tmp/repo

RUN cd /tmp/repo/src && \
    go mod tidy && \
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /usr/local/bin/geoip-lookup ./cmd/main.go

FROM alpine:latest

ARG UID=9800
ARG USER=geoip
ENV MODE_DEV=0 \
    MODE_DEBUG=0 \
    MODE_TEST=0

RUN adduser --uid ${UID} --shell /usr/sbin/nologin --disabled-password ${USER} ${USER} && \
    mkdir -p /var/local/geoip && \
    chown -R ${USER}:${USER} /var/local/geoip

COPY --from=builder-be /usr/local/bin/geoip-lookup /usr/local/bin/geoip-lookup

RUN chmod +x /usr/local/bin/geoip-lookup

USER "$USER"

CMD ["/usr/local/bin/geoip-lookup", "-l=0.0.0.0", "-p=10000", "-t=ipinfo", "-lite=/var/local/geoip/ipinfo_lite.mmdb"]
# CMD ["/usr/local/bin/geoip-lookup", "-l=0.0.0.0", "-p=10000", "-t=ipinfo", "-country=/var/local/geoip/ipinfo_country.mmdb", "-asn=/var/local/geoip/ipinfo_asn.mmdb"]
# CMD ["/usr/local/bin/geoip-lookup", "-l=0.0.0.0", "-p=10000", "-t=maxmind", "-country=/var/local/geoip/maxmind_country.mmdb", "-asn=/var/local/geoip/maxmind_asn.mmdb"]
