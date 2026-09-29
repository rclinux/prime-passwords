PREFIX  ?= $(HOME)/.local
BINDIR  := $(PREFIX)/bin
APPDIR  := $(PREFIX)/share/applications
ICONDIR := $(PREFIX)/share/icons/hicolor

.PHONY: build test install uninstall clean

build:
	go build -tags wayland -trimpath -ldflags "-s -w" -o prime-passwords .

test:
	go vet ./...
	go test ./...

install: build
	install -Dm755 prime-passwords $(BINDIR)/prime-passwords
	install -Dm644 assets/icon.png $(ICONDIR)/512x512/apps/prime-passwords.png
	install -Dm644 assets/icon.svg $(ICONDIR)/scalable/apps/prime-passwords.svg
	install -d $(APPDIR)
	sed 's|@BINDIR@|$(BINDIR)|' packaging/prime-passwords.desktop > $(APPDIR)/prime-passwords.desktop
	-update-desktop-database $(APPDIR) 2>/dev/null
	if [ -f $(ICONDIR)/index.theme ]; then gtk-update-icon-cache -q $(ICONDIR); fi

uninstall:
	rm -f $(BINDIR)/prime-passwords $(APPDIR)/prime-passwords.desktop \
	      $(ICONDIR)/512x512/apps/prime-passwords.png \
	      $(ICONDIR)/scalable/apps/prime-passwords.svg

clean:
	rm -f prime-passwords
