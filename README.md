# Prime Passwords

Cryptographically secure random passwords, made on your own computer.

Prime Passwords is a small desktop app inspired by Steve Gibson's
[Perfect Passwords](https://www.grc.com/passwords.htm) page at GRC. Open it
and it shows three brand-new passwords. Press **Refresh** for more.

| Password | Strength | Good for |
|---|---|---|
| 64 hexadecimal characters (`0-9 A-F`) | 256 bits | Raw WPA/WPA2 Wi-Fi keys |
| 63 printable ASCII characters | ~413 bits | Wi-Fi passphrases, shared secrets |
| 63 letters and digits (`a-z A-Z 0-9`) | ~375 bits | Devices or sites that reject symbols |

## Why run it locally?

- **Nothing leaves your computer.** No network code at all.
- **Nothing is saved.** No history, no logs.
- **Real randomness.** Every character comes from the operating system's
  secure random generator (Go's `crypto/rand`), and every character in a set
  is exactly as likely as any other.
- **Clipboard is cleaned up.** Copied passwords are cleared after 30 seconds.

Any part of a password is just as random as the whole, so if something only
accepts 20 characters, take any 20 in a row.

## Install (Linux)

Needs Go 1.26 or newer, a C compiler, and the usual OpenGL/X11/Wayland
development libraries (on Arch: `go gcc`; on Debian/Ubuntu:
`golang gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev`).

```sh
git clone https://github.com/rclinux/prime-passwords
cd prime-passwords
make install      # installs to ~/.local, adds a launcher icon
```

Then open **Prime Passwords** from your app launcher. `make uninstall`
removes it.

## Terminal use

```sh
prime-passwords --cli     # all three
prime-passwords --hex     # just the 64 hex characters
prime-passwords --ascii   # just the 63 printable ASCII characters
prime-passwords --alnum   # just the 63 letters and digits
```

## Keyboard

**F5** or **Ctrl+R** makes new passwords.

## Credits

Idea from GRC's Perfect Passwords page by Steve Gibson. Prime Passwords is an
independent project and is not affiliated with GRC. Built with
[Fyne](https://fyne.io).

## License

MIT — see [LICENSE](LICENSE).
