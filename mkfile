GOOS=plan9
GOARCH=$objtype

all: 
	go build

install: all
	cp ./disco $home/bin/$GOARCH/disco

deps: 
	mkdir -p $GOPATH/src/bitbucket.org/henesy/disco
	dircp ./ $GOPATH/src/bitbucket.org/henesy/disco

crypto:
	orig=`{pwd}
	mkdir -p $GOPATH/src/golang.org/x/crypto
	cd $GOPATH/src/golang.org/x
	hget https://github.com/golang/crypto/archive/master.zip > crypto.zip; unzip -f crypto.zip
	dircp crypto-master crypto; rm -r crypto-master
	cd $orig

rebuild: 
	go install bitbucket.org/henesy/disco
	mk install

bins: 
	arch=(amd64 386 arm)
	for(a in $arch){
		GOARCH=$a go build
		cp disco bin/^$a^/
	}
