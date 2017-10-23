FROM alpine

RUN apk update && \
    apk add \
    ca-certificates && \
    rm -rf /var/cache/apk/*

ENV AWS_SDK_LOAD_CONFIG=true

RUN pwd

ADD /go/src/github.com/omerxx/drone-lambda-plugin/main /bin/

ENTRYPOINT ["/bin/main"]

