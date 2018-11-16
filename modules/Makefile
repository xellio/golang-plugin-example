PLUGIN_FIRST 	= first_plugin
PLUGIN_SECOND 	= second_plugin
PLUGIN_THIRD 	= third_plugin

go			= go
GO_SUBPKGS 	= $(shell $(GO) list ./... | grep -v /vendor/)

PLUGINDIR 	= ./../plugins/

all: build

build: clean $(PLUGINDIR) plugins

plugins: $(PLUGIN_FIRST) $(PLUGIN_SECOND) $(PLUGIN_THIRD)

clean:
	rm -f $(PLUGINDIR)*

$(PLUGINDIR):
	mkdir -p $(PLUGINDIR)

$(PLUGIN_FIRST): $(PLUGINDIR)
	$(go) build -buildmode=plugin -o $(PLUGINDIR)$@.so ./first/*.go

$(PLUGIN_SECOND): $(PLUGINDIR)
	$(go) build -buildmode=plugin -o $(PLUGINDIR)$@.so ./second/*.go

$(PLUGIN_THIRD): $(PLUGINDIR)
	$(go) build -buildmode=plugin -o $(PLUGINDIR)$@.so ./third/*.go

.PHONY:clean plugins $(PLUGIN_FIRST) $(PLUGIN_SECOND) $(PLUGIN_THIRD)