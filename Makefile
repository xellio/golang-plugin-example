TARGET = modules

GO      	= go
GOLINT  	= $(GOPATH)/bin/golint
GO_SUBPKGS 	= $(shell $(GO) list ./... | grep -v /vendor/)
MAKE 		= make

BINDIR 		= ./bin/
MODULESDIR 	= ./modules/

all: $(TARGET)

$(TARGET): build

build: clean $(BINDIR) build_plugins
	$(GO) build -ldflags="-s -w" -o $(BINDIR)$(TARGET) ./cmd/cli/*.go

build_plugins: 
	$(MAKE) -C $(MODULESDIR)

clean:
	rm -f $(BINDIR)*

$(BINDIR):
	mkdir -p $(BINDIR)

$(GOLINT):
	$(GO) get -u golang.org/x/lint/golint

test:
	$(GO) test -race $$($(GO) list ./...)

lint: $(GOLINT)
	@for f in $(GO_SUBPKGS) ; do $(GOLINT) $$f ; done

.PHONY:test lint clean build_plugins