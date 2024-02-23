# HTTP Testing in Go

## How to run
You can build and run the test on local or either using docker to ease the process. Here, we will only cover the Dockerfile-through process.

Once running the main web server, place yourself on the terminal, at the same level as this directory, and run:
```
docker build -f Dockerfile.test -t test .
docker run --rm --network="host" test
```

This will execute a simple HTTP request to `http://localhost:8080` (as defined on the web server). If either connection or response reading fails, error pops up. Nonetheless, if connection is correct, what will pop up is the body of the request ().