GOCMD=go
GOBUILD=$(GOCMD) build

install:
	$(GOBUILD) -o build/labqueue.exe cmd/labqueue/main.go

clean:
	rm -rf build/*.exe