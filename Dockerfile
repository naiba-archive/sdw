FROM golang:alpine AS binarybuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git linux-pam-dev
WORKDIR /sdw/
COPY . .
RUN CGO_ENABLED=true go build -o sdw -ldflags="-s -w -X github.com/naiba/sdw.BuildVersion=`git rev-parse HEAD`" app/web/main.go

FROM alpine:latest
RUN echo http://dl-2.alpinelinux.org/alpine/edge/community/ >>/etc/apk/repositories && apk --no-cache --no-progress add \
  tzdata \
  libstdc++ \
  ca-certificates
# Copy binary to container
WORKDIR /sdw
COPY resource ./resource
COPY --from=binarybuilder /sdw/sdw .
# Configure Docker Container
VOLUME ["/sdw/data"]
EXPOSE 8080
CMD ["/sdw/sdw"]
