# AdGuardHome Swarm Dns Updater

## Features
- Alpine 3.12.4 based
- 

## Environment
- DNS_CONTAINER: dns container name (Default: unbound)
- AGH_CONTAINER: adguardhome container name (Default: adguard)
- AGH_SECURE: use https for adguardhome api requests (Default: false)
- AGH_PORT: adguardhome api port (Default: 80)
- AGH_USER: adguardhome username (Default: admin)
- AGH_PASSWORD: adguardhome password (Default: password)
- TIMER_LOOP: seconds beetween ip validation (Default: 30)
- CONTAINER_ONLY: remove all other upstream dns resolver (Default: false)
- VERBOSE: more logmessages (Default: false)
