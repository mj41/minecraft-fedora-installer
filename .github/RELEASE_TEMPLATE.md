## What's New

<!-- Describe new features, improvements, and bug fixes -->

### Added
- Feature:

### Fixed
- Fix:

### Changed
- Improvement:

## Installation

### Quick Install (Recommended)

One-line install:
```bash
curl -fsSL https://raw.githubusercontent.com/mj41/minecraft-fedora-installer/refs/heads/main/install.sh | bash
```

For force reinstall:
```bash
curl -fsSL https://raw.githubusercontent.com/mj41/minecraft-fedora-installer/refs/heads/main/install.sh | bash -s -- --force
```

### Manual Download

Download the appropriate binary for your architecture:

- **AMD64/x86_64** (most common): `mc-installer-amd64`
- **ARM64**: `mc-installer-arm64`

Make it executable and run:
```bash
chmod +x mc-installer-amd64
./mc-installer-amd64
```

## Installation Locations

- Application: `~/opt/minecraft-launcher/`
- Desktop file: `~/.local/share/applications/minecraft.desktop`
- Icon: `~/.local/share/icons/minecraft.png`

## Requirements

- Compatible with all Linux distributions (static binary)
- Tested on Fedora 43
- Internet connection for downloading Minecraft launcher
- `lspci` command (from pciutils package) for GPU detection

## What it does

1. Detects your GPU type automatically
2. Creates necessary directories
3. Downloads Minecraft launcher from Mojang to `/tmp/`
4. Extracts to `~/opt/minecraft-launcher/`
5. Installs embedded icon
6. Creates desktop file with correct GPU configuration
7. Cleans up temporary files

After installation, find "Minecraft" in your application menu!

## Uninstallation

```bash
rm -rf ~/opt/minecraft-launcher
rm ~/.local/share/applications/minecraft.desktop
rm ~/.local/share/icons/minecraft.png
```
