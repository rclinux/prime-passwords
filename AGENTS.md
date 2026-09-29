# Notes for AI coding agents

- Versioning is SemVer. Bump `version` in `main.go` and `Version` in
  `FyneApp.toml` together, and add a CHANGELOG entry.
  - PATCH: bug fixes, wording, look-and-feel tweaks.
  - MINOR: new features that keep existing behaviour and flags.
  - MAJOR: removing or changing existing flags or outputs.
- `internal/gen` must stay standard library only (`crypto/rand`). Never use
  `math/rand` or `byte % n` for password characters.
- The program must never write generated values to disk, logs, or network.
- Run `make test` before any commit.
- Commits use the GitHub noreply email set in this repo's git config.
