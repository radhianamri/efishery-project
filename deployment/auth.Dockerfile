FROM golang:1.12-alpine AS build-env
ENV GO111MODULE=on

WORKDIR /go/src/auth-go
COPY auth-go/ /go/src/auth-go
RUN apk update \
&& apk add --no-cache git \
&& rm -rf /var/cache/apk/* \
&& rm -rf /tmp/* \ 
&& go get -u github.com/swaggo/swag/cmd/swag \
&& swag init
RUN go build -o main

# final stage
FROM alpine
WORKDIR /app
RUN apk update && apk --no-cache add ca-certificates && update-ca-certificates \
&& rm -rf /var/cache/apk/* \
&& rm -rf /tmp/*
RUN mkdir config
RUN mkdir docs
COPY --from=build-env /go/src/auth-go/main /app
COPY --from=build-env /go/src/auth-go/config/config.toml /app/config
COPY --from=build-env /go/src/auth-go/docs/swagger.json /app/docs/
EXPOSE 7000
ENTRYPOINT deployment_type=PRODUCTION ./main
