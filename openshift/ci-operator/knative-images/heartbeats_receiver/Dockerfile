# DO NOT EDIT! Generated Dockerfile for cmd/heartbeats_receiver.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG UBI8_MINIMAL=registry.access.redhat.com/ubi8/ubi-minimal
ARG VERSION=knative-v1.15

FROM $GO_BUILDER as builder

COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/heartbeats_receiver

FROM $UBI8_MINIMAL

# Install the missing zoneinfo to ubi-minimal
RUN microdnf install tzdata

USER 65532

COPY --from=builder /usr/bin/main /usr/bin/main

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-heartbeats-receiver-rhel8-container" \
      name="openshift-serverless-1/eventing-heartbeats-receiver-rhel8" \
      version="$VERSION" \
      summary="Red Hat OpenShift Serverless 1 Eventing Heartbeats Receiver" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing Heartbeats Receiver" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing Heartbeats Receiver"

ENTRYPOINT ["/usr/bin/main"]
