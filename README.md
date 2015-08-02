Go URL Tracker
==============

This is a test for [Go](https://golang.org/) in a web environment.

This web app will redirect people to the provided link in `url` param logging
the request into `access.log`.

## Build

```
go build tracker.go
```

## Run

```
./tracker
```

Access http://localhost:8080/r?url=http%3A%2F%2Fwww.google.com%2F
