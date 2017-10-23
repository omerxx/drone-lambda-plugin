FROM alpine

RUN apk update && \
    apk add \
    ca-certificates && \
    rm -rf /var/cache/apk/*

ENV AWS_SDK_LOAD_CONFIG=true

ADD main /bin/

ENTRYPOINT ["/bin/main"]

