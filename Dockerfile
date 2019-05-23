FROM golang:1.12.5-alpine AS build
ADD . $GOPATH/src/github.com/machadolhes/tanker-go
WORKDIR $GOPATH/src/github.com/machadolhes/tanker-go
RUN go install

FROM alpine:3.9
RUN mkdir -p /opt/tanker
WORKDIR /opt/tanker
COPY --from=build /go/bin/tanker-go /opt/tanker/tanker-go
CMD ["/opt/tanker/tanker-go"]