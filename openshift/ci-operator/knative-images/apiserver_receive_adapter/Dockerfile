# DO NOT EDIT! Generated Dockerfile for cmd/apiserver_receive_adapter.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG UBI8_MINIMAL=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

COPY . .

RUN mkdir -p /var/run/ko && \
    mkdir -p cmd/apiserver_receive_adapter/kodata && \
    go build -o /usr/bin/main ./cmd/apiserver_receive_adapter && \
    cp -r cmd/apiserver_receive_adapter/kodata /var/run/ko

FROM $UBI8_MINIMAL

# install the missing zoneinfo to ubi-minimal
RUN microdnf install tzdata

USER 65532

COPY --from=builder /usr/bin/main /usr/bin/main
COPY --from=builder /var/run/ko /var/run/ko
ENTRYPOINT ["/usr/bin/main"]
