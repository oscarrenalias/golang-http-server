# What

This is a teeny-tiny Golang-based web server that responds to all requests with the name of the host that served the request, as well as a list of the headers and request fields. It's built as a multi-staged build and packaged as a tiny container (6Mb or so)

# Build

The provided Dockerfile is a multi-stage build:

```
docker build -t golang-http-server .
```

# Running

```
docker run -p 8080 -d golang-http-server
```