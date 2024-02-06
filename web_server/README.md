# Simple Web Server
As the second step for Golang learning, a simple web server is a proper mini-project.

The needed files for this project are a Dockerfile and a simple golang sample.

# How to run
You can build the application on local or either use docker to ease the process. Here, we will only cover the Dockerfile-through process.

Place yourself on the terminal, at the same level as this directory, and run:
```
docker build -t <simple project name> .
docker run -p 8080:8080 <simple project name>
```
