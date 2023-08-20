# Container Setup

## start the local stack

```bash
podman-compose up -d
```

## connect to adminer

use `127.0.0.1:8080` or `podman inspect <adminer-container-id> |grep IP` and use that IP in the adminer UI
