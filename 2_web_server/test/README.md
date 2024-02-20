# HTTP Testing in Go

## How to run
You can build the application on local or either use docker to ease the process. Here, we will only cover the Dockerfile-through process.

Once running the main web server, place yourself on the terminal, at the same level as this directory, and run:
```
docker build -f Dockerfile.test -t test .
docker run --rm --network="host" test
```
