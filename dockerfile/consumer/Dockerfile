# === Builder ===

FROM golang:1.12.7 AS builder

# Install required packages
RUN apt-get update && apt-get install -y apt-utils gettext-base

# Reconstruct source tree inside docker
WORKDIR /code
ADD . .

# ./vendor is excluded in .dockerignore, reconstruct it with 'mod' tool
RUN GO111MODULE=on go mod tidy
RUN GO111MODULE=on go mod vendor

# Before start build process, we need to ensure go test is OK
RUN ./dev/run_test.sh

# Build operator binary with explicitly specified output
RUN OUTPUT_BIN=/tmp/consumer ./dev/go_build_consumer.sh

# === Image ===

FROM alpine:3.10 AS image

RUN apk add --no-cache ca-certificates

WORKDIR /

# Add config files from local source dir into image
ADD config/consumer.yaml   /etc/atlas/

# Copy binary into result image from builder
COPY --from=builder /tmp/consumer .

# Run binary
# We can specify additional options, such as:
#   --config=/path/to/config
#   --kube-config=/path/to/kubeconf
ENTRYPOINT ["/consumer"]
CMD ["--config=/etc/atlas/consumer.yaml"]
#CMD ["-logtostderr=true", "-v=1"]
