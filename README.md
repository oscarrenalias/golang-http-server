# Build

The provided Dockerfile is a multi-stage build:

```
docker build -t golang-http-server .
```

# Running

```
docker run -p 8080 -d golang-http-server
```