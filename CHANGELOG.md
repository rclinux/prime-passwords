# Changelog

All notable changes are listed here. Versions follow [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added
- Desktop window (Fyne) showing three passwords: 64 hex, 63 printable ASCII,
  63 letters and digits, each with its strength in bits.
- Refresh button, F5 and Ctrl+R for new passwords.
- Copy buttons; clipboard is cleared after 30 seconds unless something else
  was copied since.
- About page explaining how the passwords are made and which to use.
- Terminal mode: `--cli`, `--hex`, `--ascii`, `--alnum`, `--version`.
- Padlock app icon, desktop launcher entry, `make install` / `make uninstall`.
- Tests for lengths, character sets, even character odds (chi-square) and
  the window's buttons.
