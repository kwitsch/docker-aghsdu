FROM golang:alpine3.12 AS build-env
ADD src /src
RUN cd /src && go build -o aghsdu

FROM alpine:3.12
WORKDIR /app
COPY --from=build-env /src/aghsdu /app/

LABEL org.label-schema.name="aghsdu"
LABEL org.label-schema.vendor="kwitsch"
LABEL org.label-schema.vcs-url="https://github.com/kwitsch/docker-aghsdu"

ENV DNS_CONTAINER=unbound
ENV AGH_CONTAINER=adguard
ENV AGH_SECURE=false
ENV AGH_PORT=80
ENV AGH_USER=admin
ENV AGH_PASSWORD=password
ENV TIMER_LOOP=30
ENV CONTAINER_ONLY=false
ENV VERBOSE=false

ENTRYPOINT ["./aghsdu"]