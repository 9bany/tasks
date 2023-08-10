# tasks

## Improves
- [migration] move migration tool in code
- [consider] performance: random api key query  
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

## Api Keys setup

1. Exec into container
```
docker exec -it <container_app_name> /bin/sh
```
2. Insert your key

```
./task keys create --key=<your_key>
```
3. Log key info

```
./task keys get --key=<your_key>
```
