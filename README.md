# Golang example REST application

## REST, RESTful, WTH!!

REST is an acronym for Representational State Transfer. It is a web standards architecture and HTTP Protocol. The REST protocol, decribes six (6) constraints:

* Uniform Interface
* Cacheable
* Client-Server
* Stateless
* Code on Demand
* Layered System

REST is composed of methods such as a base URL, media types, etc. RESTful applicaitons uses HTTP requests to perform the CRUD operations.

## Router

We'll need to use a mux to route requests, so we need a Go package for that (mux stands for HTTP request multiplexer which matches an incoming request to against a list of routes (registered)). In the rest-api directory, let's require the dependency (package rather).

```bash
go get github.com/gorilla/mux
```

## References

* https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
* https://github.com/wilkinsona/spring-boot-guided-tour-goto-amst