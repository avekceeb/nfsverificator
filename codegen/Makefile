all: sunrpc.go out.go

sunrpc.go: sunrpc.y
	goyacc -o sunrpc.go -p "yy" -v /dev/null sunrpc.y

out.go: sunrpc.go in.x
	go run main/main.go

cleanall: clean
	rm -f sunrpc.go

clean:
	rm -f out.go
