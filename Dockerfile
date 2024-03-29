FROM golang AS build-env

RUN apt update
RUN apt install gcc
ADD src /src
RUN cd /src && go build -ldflags "-linkmode external -extldflags -static" -o aghsdu


FROM scratch
COPY --from=build-env /src/aghsdu /aghsdu

ENV DNS_CONTAINER=unbound \
    AGH_CONTAINER=adguard \
    AGH_SECURE=false \
    AGH_PORT=80 \
    AGH_USER=admin \
    AGH_PASSWORD=password \
    TIMER_LOOP=30 \
    CONTAINER_ONLY=false \
    VERBOSE=false

ENTRYPOINT ["/aghsdu"]