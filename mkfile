GOOS=plan9
GOARCH=$objtype

all: 
	go build

install: all
	cp ./disco $home/bin/$GOARCH/disco

bins: 
	plat=(plan9 linux windows)
	arch=(amd64 386 arm)
	for(p in $plat){
		for(a in $arch){
			mkdir -p bin/$p/$a
			GOOS=$p GOARCH=$a go build
			cp disco bin/$p/$a/
			rm disco
		}
	}
