FROM alpine:3.2

RUN apk update && \
    apk add \
    ca-certificates && \
    rm -rf /var/cache/apk/*

ADD drone-ecs /bin/
ENTRYPOINT ["/bin/drone-ecs"]
