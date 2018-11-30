# build on mac -> linux
build : main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pingdog main.go

# clean : 
# 	rm pingdog