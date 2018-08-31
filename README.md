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

Use the MESSAGE environment variable to add a customizable string to each response, useful when you deploy multiple instances and want to be able to differentiate them:

```
docker run -p 8080:8080 -d -e MESSAGE="service 1" golang-http-server 
docker run -p 8081:8080 -d -e MESSAGE="service 2" golang-http-server
```