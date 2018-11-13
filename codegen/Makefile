all: sunrpc.go v40.go v41.go 

sunrpc.go: sunrpc.y
	goyacc -o sunrpc.go -p "yy" -v /dev/null sunrpc.y

v40.go: sunrpc.go v40.x
	go run main/main.go v40.x v40.go

v41.go: sunrpc.go v41.x
	go run main/main.go v41.x v41.go

v42.go: sunrpc.go v42.x
	go run main/main.go v42.x v42.go

cleanall: clean
	rm -f sunrpc.go

clean:
	rm -f v42.go v41.go v40.go
