all: build

build :
	go get -v
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" .

test : build
	go test -v
	golint > lint.txt
	go tool vet -v . > vet.txt
	gocov test | gocov-xml > coverage.xml
	go test -bench=. -test.benchmem -v | gobench2plot > benchmarks.xml

prepare : build
	cp storemeta Docker/storemeta

docker-devel : prepare
	-@sudo docker rmi -f bayugyug/storemeta 2>/dev/null || true
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/storemeta .

docker-wheezy: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/storemeta -f  wheezy/Dockerfile .

docker-scratch: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/storemeta:scratch -f  scratch/Dockerfile .

docker-alpine: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/storemeta:alpine  -f  alpine/Dockerfile .

clean:
	rm -f storemeta Docker/storemeta
	rm -f benchmarks.xml coverage.xml vet.txt lint.txt

re: clean all

