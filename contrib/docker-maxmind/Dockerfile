FROM --platform=${BUILDPLATFORM} golang:1.21-alpine3.19 AS builder

ENV GEOIP_LOOKUP_SERVICE_URL https://github.com/superstes/geoip-lookup-service/archive/refs/tags/1.0.tar.gz

ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR $GOPATH/src/github.com/superstes/geoip-lookup-service

RUN set -eux; \
        wget -O geoip-lookup-service.tar.gz "$GEOIP_LOOKUP_SERVICE_URL"; \
        tar -xzf geoip-lookup-service.tar.gz -C $GOPATH/src/github.com/superstes/geoip-lookup-service --strip-components=1; \
        rm geoip-lookup-service.tar.gz; \
        case ${TARGETPLATFORM} in \
            "linux/amd64")  GOARCH=amd64  ;; \
            # arm64 and arm64v8 are equivalent in go and do not require a goarm
            # https://github.com/golang/go/wiki/GoArm
            "linux/arm64" | "linux/arm/v8")  GOARCH=arm64  ;; \
            "linux/ppc64le")  GOARCH=ppc64le  ;; \
            "linux/arm/v6") GOARCH=arm GOARM=6  ;; \
            "linux/arm/v7") GOARCH=arm GOARM=7 ;; \
        esac; \
        printf "Building geoip-lookup-service for arch ${GOARCH}\n"; \
        GOARCH=${GOARCH} go build -a -o /geoip-lookup-service github.com/superstes/geoip-lookup-service/main

FROM alpine:3.19

RUN set -eux; \
        addgroup --gid 99 --system geoip-lookup; \
        adduser \
                --disabled-password \
                --home /var/lib/geoip-lookup \
                --ingroup geoip-lookup \
                --no-create-home \
                --system \
                --uid 99 \
                geoip-lookup \
        ; \
        mkdir /var/lib/geoip-lookup; \
        chown geoip-lookup:geoip-lookup /var/lib/geoip-lookup

COPY --from=builder /geoip-lookup-service /usr/local/bin

USER geoip-lookup

WORKDIR /var/lib/geoip-lookup
