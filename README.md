# AdGuardHome Swarm Dns Updater

## ***Discontinued!***
An alternative solution is to insert '127.0.0.11' as bootstrap dns server and 'tcp://*container name*' as upstream dns server.

---

## Description
Service for updating upstream dns resolver in AdGuardHome container.<br />
Container addresses are obtained through the Docker DNS resolver.<br />
Upstream DNS resolver are only updated if a address change is detected.<br />
The DNS resolver container can purely run inside an overlay network without exposed ports.

## Features
- Golang static build
- small size through scratch image with only necessarry executable
- all configurations done throug environment variables or docker secrets(ENV_*secret name*)

## Environment
- DNS_CONTAINER: dns container name (Default: unbound)
- AGH_CONTAINER: adguardhome container name (Default: adguard)
- AGH_SECURE: use https for adguardhome api requests (Default: false)
- AGH_PORT: adguardhome api port (Default: 80)
- AGH_USER: adguardhome username (Default: admin)
- AGH_PASSWORD: adguardhome password (Default: password)
- TIMER_LOOP: seconds beetween ip lookups (Default: 30)
- CONTAINER_ONLY: remove all other upstream dns resolver without a zone restriction (Default: false)
- VERBOSE: more logmessages (Default: false)

## Example-Stack
```YAML
version: "3.8"

services:
  unbound:
    image: kwitsch/unbound:v1.2.0
  aghsdu:
    image: kwitsch/aghsdu:v1.0.0
    environment:
      - AGH_PASSWORD=your_password
      - CONTAINER_ONLY=true
  adguard:
    image: adguard/adguardhome
    ports:
      - 53:53/tcp
      - 53:53/udp
      - 80:80/tcp
```
