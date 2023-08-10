# tasks

## Improves
- UnitTest: api request timeout, data stite error format, mete  data site ops error 
- Log/span: When an API error occurs, this is helpful for tracing bugs.
- Monitor: Add metrics to the get_site API.
- SiteData structure: If the Iframely API interface is updated to a new version, this code will break.
## Run
### 1. Docker
Linux 
```
docker run --add-host host.docker.internal:host-gateway -p 8080:8080 -e DB_SOURCE='postgresql://<user>:<pass>@host.docker.internal:5432/task?sslmode=disable' task:latest
```