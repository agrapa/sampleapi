FROM golang:1.13.6 AS builder
MAINTAINER arie.grapa@machinify.com
COPY ./ /sampleapi
RUN go get -u golang.org/x/lint/golint
RUN go get -u github.com/mgechev/revive
RUN cd /sampleapi && CGO_ENABLED=0 GOOS=linux make build

FROM scratch
COPY --from=builder /sampleapi/dist/sampleapi /
ENTRYPOINT ["/sampleapi"]

