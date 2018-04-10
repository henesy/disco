GOOS=plan9
GOARCH=$objtype

all: 
	go build

install: all
	cp ./disco $home/bin/$GOARCH/disco

bins: 
	arch=(amd64 386 arm)
	mkdir bin
	for(a in $arch){
		mkdir bin/$a
		GOARCH=$a go build
		cp disco bin/^$a^/
		rm disco
	}
