# Hello World
As the first step for Golang learning, a hello world application is expected.

The needed files for this project are a Dockerfile and a simple golang sample.

## How to run

You can build the application on local or either use docker to ease the process. Here, we will only cover the Dockerfile-through process.

Place yourself on the terminal, at the same level as this directory, and run:
```
docker build -t <simple project name> .
docker run <simple project name>
```

## Reference Index

For this mini-project we need [fmt package](https://pkg.go.dev/fmt), which implements formatted I/O functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler. This way we will handle output for our _Hello World_.

To import this package we simply do the following:
```
import "fmt"
```

We also use our _main_ function for printing our text through the [_fmt.Println_ function](https://pkg.go.dev/fmt#Println).
