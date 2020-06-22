FROM golang:1.13-alpine3.10 AS builder
ARG VERSION
RUN apk update && mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o dist/hal9k -a -ldflags "-w -s -X hal9k/pkg/version.Version=${VERSION}" ./cmd/

FROM alpine:3.10

ARG REPO_URL
ARG BRANCH
ARG COMMIT_REF

ENV LUCK_FILE "/etc/hal9k/luck_data.json"
ENV IMAGE_PATH "/data"

LABEL repo-url=${REPO_URL}
LABEL branch=${BRANCH}
LABEL commit-ref=${COMMIT_REF}

RUN apk update && mkdir /app && mkdir -pv /etc/hal9k && mkdir -pv /data

WORKDIR /app
COPY --from=builder /app/dist/hal9k /app/hal9k
COPY --from=builder /app/config/luck_data.json /etc/hal9k/luck_data.json

VOLUME /data
EXPOSE 8443

CMD ["/app/hal9k"]