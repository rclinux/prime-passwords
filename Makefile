PREFIX  ?= $(HOME)/.local
BINDIR  := $(PREFIX)/bin
APPDIR  := $(PREFIX)/share/applications
ICONDIR := $(PREFIX)/share/icons/hicolor
GOBIN   ?= $(shell go env GOPATH)/bin
VERSION := $(shell sed -n 's/^const version = "\(.*\)"/\1/p' main.go)

.PHONY: build linux-release windows test audit install uninstall clean

# Default build runs on X11 and Wayland (via XWayland).
# For a native Wayland window: make install TAGS=wayland
TAGS ?=

build:
	go build -tags "$(TAGS)" -trimpath -ldflags "-s -w" -o prime-passwords .

# Linux download: binary + icon + launcher entry + install.sh, as .tar.gz.
# Built without -tags wayland so it runs on both X11 and Wayland desktops.
LINUXPKG := prime-passwords-$(VERSION)-linux-amd64
linux-release:
	rm -rf dist/$(LINUXPKG)
	mkdir -p dist/$(LINUXPKG)
	go build -trimpath -ldflags "-s -w" -o dist/$(LINUXPKG)/prime-passwords .
	cp assets/icon.png dist/$(LINUXPKG)/prime-passwords.png
	cp assets/icon.svg dist/$(LINUXPKG)/prime-passwords.svg
	cp packaging/prime-passwords.desktop packaging/install.sh README.md LICENSE SECURITY.md dist/$(LINUXPKG)/
	tar -C dist -czf dist/$(LINUXPKG).tar.gz $(LINUXPKG)
	rm -rf dist/$(LINUXPKG)

# Windows .exe with embedded icon, cross-compiled from Linux.
# Needs mingw-w64-gcc and: go install fyne.io/tools/cmd/fyne@latest
windows:
	cp FyneApp.toml FyneApp.toml.bak   # fyne package bumps Build; keep tree clean
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ \
		$(GOBIN)/fyne package --os windows --release; \
		status=$$?; mv FyneApp.toml.bak FyneApp.toml; exit $$status
	mkdir -p dist
	mv "Prime Passwords.exe" dist/prime-passwords-$(VERSION)-windows-amd64.exe

test:
	go vet ./...
	go test ./...

# Needs: go install golang.org/x/vuln/cmd/govulncheck@latest
#        go install github.com/securego/gosec/v2/cmd/gosec@latest
#        go install honnef.co/go/tools/cmd/staticcheck@latest
audit:
	go mod verify
	go vet ./...
	$(GOBIN)/staticcheck ./...
	$(GOBIN)/gosec -quiet ./...
	$(GOBIN)/govulncheck ./...
	go test -race -count=1 ./...

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
	rm -rf prime-passwords dist
