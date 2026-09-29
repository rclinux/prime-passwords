# Changelog

All notable changes are listed here. Versions follow [Semantic Versioning](https://semver.org/).

## [Unreleased]

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
- Padlock app icon, desktop launcher entry, `make install` / `make uninstall`.
- Linux hardening: process is marked non-dumpable, so passwords never land
  in a crash dump and other same-user programs cannot read its memory.
- `make audit`: module checksums, vet, staticcheck, gosec, govulncheck and
  race-detector tests. SECURITY.md describes what is and is not protected.
- Tests for lengths, character sets, even character odds (chi-square) and
  the window's buttons.
