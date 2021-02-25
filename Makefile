.POSIX:
.SUFFIXES:

GO = go
RM = rm
GOFLAGS =
PREFIX = /usr/local
BINDIR = $(PREFIX)/bin

goflags = $(GOFLAGS)

all: tent

tent:
	$(GO) build $(goflags)

clean:
	$(RM) -f tent

install: all
	mkdir -p $(DESTDIR)$(BINDIR)
	cp -f tent $(DESTDIR)$(BINDIR)

