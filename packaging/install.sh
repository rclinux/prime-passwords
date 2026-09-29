#!/bin/sh
# Prime Passwords installer for Linux.
# Installs for the current user only (no root needed):
#   ~/.local/bin/prime-passwords
#   ~/.local/share/applications/prime-passwords.desktop  (app launcher entry)
#   ~/.local/share/icons/hicolor/.../prime-passwords.png (icon)
#
#   ./install.sh              install or update
#   ./install.sh --uninstall  remove everything this script installed
set -eu

here=$(cd "$(dirname "$0")" && pwd)
prefix="${PREFIX:-$HOME/.local}"
bindir="$prefix/bin"
appdir="$prefix/share/applications"
icondir="$prefix/share/icons/hicolor"

if [ "${1:-}" = "--uninstall" ]; then
	rm -f "$bindir/prime-passwords" "$appdir/prime-passwords.desktop" \
		"$icondir/512x512/apps/prime-passwords.png" \
		"$icondir/scalable/apps/prime-passwords.svg"
	echo "Prime Passwords removed."
	exit 0
fi

install -Dm755 "$here/prime-passwords" "$bindir/prime-passwords"
install -Dm644 "$here/prime-passwords.png" "$icondir/512x512/apps/prime-passwords.png"
install -Dm644 "$here/prime-passwords.svg" "$icondir/scalable/apps/prime-passwords.svg"
mkdir -p "$appdir"
sed "s|@BINDIR@|$bindir|" "$here/prime-passwords.desktop" > "$appdir/prime-passwords.desktop"
command -v update-desktop-database >/dev/null && update-desktop-database "$appdir" 2>/dev/null || true

echo "Prime Passwords installed."
echo "Open it from your app menu, or run: $bindir/prime-passwords"
case ":$PATH:" in
	*":$bindir:"*) ;;
	*) echo "Note: $bindir is not on your PATH, so type the full path above to use it from a terminal." ;;
esac
