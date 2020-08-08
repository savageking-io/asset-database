CC=go
VERSION=$(shell cat VERSION)
OS=$(shell uname -s)
ARCH=$(shell uname -m)
BINARY=assets-database
include config.make

linux: db-linux 
darwin: db-darwin 
windows: db-windows 

db-linux: bin/linux/$(BINARY)
db-windows: bin/windows/$(BINARY).exe
db-darwin: bin/darwin/$(BINARY)

all: linux windows darwin test

SOURCES=main.go \
		config.go

test:
	$(CC) test . -v

bin/linux/$(BINARY): $(SOURCES)
	GOOS=linux $(CC) build -ldflags="-w -s -X main.AppVersion=$(APP_VERSION)" -o $@ -v $^

bin/windows/$(BINARY).exe: $(SOURCES)
	GOOS=windows $(CC) build -ldflags="-w -s -X main.AppVersion=$(APP_VERSION)" -o $@ -v $^
	
bin/darwin/$(BINARY): $(SOURCES)
	GOOS=darwin $(CC) build -ldflags="-w -s -X main.AppVersion=$(APP_VERSION)" -o $@ -v $^


clean:
	-rm -f bin/linux/$(BINARY)
	-rm -f bin/windows/$(BINARY).exe
	-rm -f bin/darwin/$(BINARY)

mrproper:
	-rm -rf bin
	-rm -f config.make
