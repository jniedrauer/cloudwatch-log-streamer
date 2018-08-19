# Executables
GO = ~/go/bin/go1.11beta3
GOLINT = golint
XUNIT = go2xunit

# Build config
BUILDFLAGS := -v -race -ldflags "-extldflags '-static'"
GOOS = linux
CGO_ENABLED = 0
GOARCH = x86_64
TESTFLAGS = -v -race
BUILDDIR = build
REPORTDIR = $(BUILDDIR)/test-reports
CMDDIR = cmd
PKGS = $(shell $(GO) list ./... | grep -v /vendor/)

all: clean lint test build

build: $(BUILDDIR)/log-streamer

$(BUILDDIR)/%: $(CMDDIR)/%
	$(GO) build $(BUILDFLAGS) -o $@ ./$<

lint:
	@# TODO: Fix this once golint supports modules
	@# https://github.com/golang/lint/issues/409
	$(GOLINT) -set_exit_status ./...

test:
	@mkdir -p $(REPORTDIR)
	$(GO) test $(TESTFLAGS) $(PKGS) \
		| tee -i /dev/stderr \
		| $(XUNIT) -output $(REPORTDIR)/unit.xml
clean:
	$(GO) clean -cache -testcache
	rm -f $(BUILDDIR)/log-streamer
	rm -f $(BUILDDIR)/*.xml

.PHONY: all build lint test
