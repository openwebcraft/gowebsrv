# gowebsrv

Single purpose [Go][golang.org] web server, serving a user interface.

By design, they have no dependencies outside the standard library and require no additional data files at run time; for a web server, it's about as simple as it gets. 

Use `go get` to fetch and install e.g. the `hostname` in your Go workspace:

    $ go get github.com/openwebcraft/gowebsrv/hostname

Then compile it with the `go` tool: 

    $ go install github.com/openwebcraft/gowebsrv/hostname

Execute the command to start the server:

    $GOPATH/bin/hostname
    
Finally, point your favourite browser to <http://localhost:8080>...

## Build and run Go server with Docker

Invoke [Docker][docker.com] from the `hostname` package directory to build an image using the included `Dockerfile`:

    $ docker build -t hostname .

To run a container from the resulting image:

    $ docker run --publish 80:8080 --name myhostname --rm hostname

And point your favourite browser to <http://localdocker/> - or to which ever IP or DNS your Docker env resolves...

## Resources

- The [outyet](http://godoc.org/github.com/golang/example/outyet) program from the [Go examples repository](https://github.com/golang/example).
- The Go Blog: [Deploying Go servers with Docker](https://blog.golang.org/docker)


[golang.org]: https://golang.org/ "Go"
[docker.com]: https://www.docker.com/ "Docker"