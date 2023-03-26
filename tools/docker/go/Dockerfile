FROM golang:1.20.2-alpine3.16

RUN apk add --no-cache gcc musl-dev

ENV CGO_ENABLED=1
ENV CC=gcc

ENV WEAVER_VERSION=v0.1.5
ENV WEAVER_GKE_VERSION=v0.1.6
ENV WEAVER_GKE_LOCAL_VERSION="${WEAVER_GKE_VERSION}"

RUN go install github.com/ServiceWeaver/weaver/cmd/weaver@"${WEAVER_VERSION}"
RUN go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke@"${WEAVER_GKE_VERSION}"
RUN go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke-local@"${WEAVER_GKE_LOCAL_VERSION}"
