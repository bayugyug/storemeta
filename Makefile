all: build

build :
	go get -v
	go build -v -ldflags "-X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`"

test : build
	go test -v
	golint > lint.txt
	go tool vet -v . > vet.txt
	gocov test | gocov-xml > coverage.xml
	go test -bench=. -test.benchmem -v | gobench2plot > benchmarks.xml

clean:
	rm -f storemeta
	rm -f benchmarks.xml coverage.xml vet.txt lint.txt

re: clean all

