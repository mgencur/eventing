# DO NOT EDIT! Generated Dockerfile for cmd/apiserver_receive_adapter.

FROM brew.registry.redhat.io/rh-osbs/openshift-golang-builder:rhel_8_golang_1.22@sha256:7b60d4fc25b0f183adf917bf4091cd4bf192f42ebebd096ca4648c0c48605c8f as golang_builder
#FROM brew.registry.redhat.io/rh-osbs/openshift-golang-builder:rhel_8_golang_1.22 as builder

COPY . .

RUN mkdir -p /var/run/ko && \
    mkdir -p cmd/apiserver_receive_adapter/kodata && \
    go build -o /usr/bin/main ./cmd/apiserver_receive_adapter && \
    cp -r cmd/apiserver_receive_adapter/kodata /var/run/ko

#FROM registry.access.redhat.com/ubi8/ubi-minimal:latest@sha256:a47c89f02b39a98290f88204ed3d162845db0a0c464b319c2596cfd1e94b444e
FROM registry.access.redhat.com/ubi8/ubi-minimal

# install the missing zoneinfo to ubi-minimal
RUN microdnf install tzdata

USER 65532

COPY --from=golang_builder /usr/bin/main /usr/bin/main
COPY --from=golang_builder /var/run/ko /var/run/ko
ENTRYPOINT ["/usr/bin/main"]
