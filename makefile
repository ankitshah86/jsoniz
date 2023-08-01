# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
BINARY_NAME := jsoniz

# Target to build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go

# Target to run tests
test:
	$(GOTEST) -v ./...

# Target to clean up built files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Target to fetch dependencies
deps:
	$(GOGET) ./...

# Target to build the binary and run tests
all: deps test build

# Target to install the binary to $GOPATH/bin
install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) cmd/main.go

# Target to uninstall the binary from $GOPATH/bin
uninstall:
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

# Target to run the binary
run:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go
	./$(BINARY_NAME)

# Set the default target to 'all'
default: all
