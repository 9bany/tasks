# tasks

## Improves
- Log/span: When an API error occurs, this is helpful for tracing bugs.
- Monitor: Add metrics to the get_site API.
- SiteData structure: If the Iframely API interface is updated to a new version, this code will break.
## Run
### 1. Docker
Linux 
```
docker run --add-host host.docker.internal:host-gateway -p 80:8080 ghcr.io/9bany/tasks:latest
```