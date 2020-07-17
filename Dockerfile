FROM go:latest

ADD . $GOPATH/src/github.com/polipopoliko/todo/todo2

WORKDIR $GOPATH/src/github.com/polipopoliko/todo/todo2

CMD ["go", "run", "main.go"]
