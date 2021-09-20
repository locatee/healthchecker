# Healthchecker

This is a very small standalone binary which allows to fetch HEALTHCHECK targets over HTTP/HTTPS and exits with 0 or 1 depending on the HTTP status of the target.

This allows to replace wget/curl in container environments where the image does not provide either, e.g. distroless images. How to use it in a Dockerfile:

```
FROM docker.io/library/golang:latest as healthcheck
RUN go install -tags netgo github.com/locatee/healthchecker@latest


FROM xyz
COPY --from=healthcheck /go/bin/healthchecker /usr/bin/healthchecker
```


## Usage

```
/usr/bin/healtchecker http://localhost:9090/health
``` 

will call http://localhost:9090/health and exit with 1 if the GET request failed. See [https://pkg.go.dev/net/http#Client.Get](https://pkg.go.dev/net/http#Client.Get) for further details about error behavior.
