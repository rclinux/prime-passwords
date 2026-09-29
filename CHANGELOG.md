# Changelog

All notable changes are listed here. Versions follow [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added
- Linux download package (`make linux-release`) with `install.sh`
  (per-user install, `--uninstall`), no root or Go needed.
- GitHub Actions: CI runs `make audit`; version tags build the Linux and
  Windows downloads plus SHA256SUMS into a draft release.
- Full README: screenshots, Windows and Linux guides, usage, FAQ,
  troubleshooting, download verification.
- `SCREENSHOTS=1 go test -run TestScreenshots .` renders README images.

### Changed
- Window subtitle shortened to "Cryptographically secure passwords".
- `make build`/`make install` default to a build that runs on both X11 and
  Wayland; `TAGS=wayland` gives a native Wayland window.

## [0.1.0] - 2026-09-29

### Security
- Updated golang.org/x/image to v0.46.0 and golang.org/x/net to v0.59.0 to
  remove known vulnerabilities in Fyne's dependencies (minimum Go is now 1.26).

### Added
- Desktop window (Fyne) showing three passwords: 64 hex, 63 printable ASCII,
  63 letters and digits, each with its strength in bits.
- Refresh button, F5 and Ctrl+R for new passwords.
- Copy buttons; clipboard is cleared after 30 seconds unless something else
  was copied since.
- About page explaining how the passwords are made and which to use.
- Terminal mode: `--cli`, `--hex`, `--ascii`, `--alnum`, `--version`.
- Windows build (`make windows`): single .exe with embedded icon; terminal
  flags print to the console they were started from.
- Padlock app icon, desktop launcher entry, `make install` / `make uninstall`.
- Linux hardening: process is marked non-dumpable, so passwords never land
  in a crash dump and other same-user programs cannot read its memory.
- `make audit`: module checksums, vet, staticcheck, gosec, govulncheck and
  race-detector tests. SECURITY.md describes what is and is not protected.
- Tests for lengths, character sets, even character odds (chi-square) and
  the window's buttons.
