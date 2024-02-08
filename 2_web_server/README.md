# Simple Web Server
As the second step for Golang learning, a simple web server is a proper mini-project.

The needed files for this project are a Dockerfile and a simple golang sample.

## How to run
You can build the application on local or either use docker to ease the process. Here, we will only cover the Dockerfile-through process.

Place yourself on the terminal, at the same level as this directory, and run:
```
docker build -t <simple project name> .
docker run -p 8080:8080 <simple project name>
```

## How to test
You can simply run a test on your terminal (PowerShell/CMD/Unix-based/etc) through a simple curl:
```
curl http://localhost:8080
```

HTTP response code should be 200, and the content should be the one we added on main.go (line 12). The following is an example:

```
StatusCode        : 200
StatusDescription : OK
Content           : Your Golang-based Web Server is working properly!
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 49
                    Content-Type: text/plain; charset=utf-8
                    Date: Thu, 08 Feb 2024 23:00:14 GMT

                    Your Golang-based Web Server is working properly!
```

## Reference Index

For this mini-project we need [net/http package](https://pkg.go.dev/net/http) for handling HTTP connections, and [log](https://pkg.go.dev/log) for handling logging. For this reason, we import it, as well as fmt as explained on the previous module. [Ref. 1]

The basic structure for this example is explained on the official documentation for [HTTP Server](https://pkg.go.dev/net/http#hdr-Servers). Nonetheless, we modified the example to make it easier to understand _HandlerFunc_. Below is how it's used on the official Golang Documentation:
```go
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
```

We use a function (simply called 'handler') for simplifying the HandlerFunc method [Ref. 2]:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your Golang-based Web Server is working properly!")
}
```

This way we only use _HandlerFunc_ as follows [Ref. 3]:
```go
http.HandleFunc("/", handler)
```
