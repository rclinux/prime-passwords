# Security

## What Prime Passwords protects

- **Randomness.** Every character comes from Go's `crypto/rand`, which reads
  the operating system's secure random generator. Characters are chosen with
  `rand.Int`, so every character in a set has exactly equal odds. If the
  system generator ever fails, the program stops rather than showing a weak
  password.
- **No network.** Prime Passwords' own code opens no network connections.
  (The Fyne toolkit links in HTTP code for loading remote images; it is never
  used here. A system-call trace of the running app shows only local Unix
  sockets for the display and desktop bus.)
- **No storage.** Generated passwords are never written to disk, logs or
  settings. Fyne creates an empty `preferences.json` under
  `~/.config/fyne/`; it contains no passwords.
- **Clipboard.** Copied passwords are cleared after 30 seconds unless
  something else was copied since.
- **Memory (Linux).** The process marks itself non-dumpable at start-up, so a
  crash never writes passwords into a core dump and other programs running as
  the same user cannot read its memory.

## What it does not protect against

- **Clipboard history tools.** If a clipboard manager (cliphist, CopyQ,
  Klipper and so on) is running, it may save copied passwords in its history.
  Type or select the password instead, or clear that history.
- **A compromised computer.** Malware running as root, a keylogger or a
  screen recorder can see anything on screen.
- **Memory after use.** Go cannot reliably wipe strings, so old passwords
  may stay in RAM (or swap) until the memory is reused.
- **Screen visibility.** Passwords are shown in full on screen.

## Checking it yourself

```sh
make test     # unit and window tests
make audit    # checksums, vet, staticcheck, gosec, govulncheck, race tests
```

## Reporting a problem

Please open a GitHub issue. For anything sensitive, use GitHub's private
vulnerability reporting ("Report a vulnerability" on the Security tab).
