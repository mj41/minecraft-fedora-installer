# Minecraft Launcher Installer for Fedora Linux

A simple Go application to install the Minecraft launcher per-user on Fedora Linux. Downloads, extracts, and configures the official Minecraft launcher with desktop integration.

## Features

- **No root required**: Per-user installation to standard XDG directories
- **Automatic download**: Fetches the latest launcher from Mojang
- **Desktop integration**: Adds launcher to application menu with icon
- **Mesa workaround**: Includes `__GLX_VENDOR_LIBRARY_NAME=mesa` for Fedora compatibility
- **Idempotent**: Skips already completed steps unless `--force` is used
- **Self-contained binary**: Icon and desktop template embedded at build time
- **Multi-architecture**: Pre-built binaries for AMD64 and ARM64

## Per-User Installation Locations

| Item | Path |
|------|------|
| Application | `~/opt/minecraft-launcher/` |
| Desktop file | `~/.local/share/applications/minecraft.desktop` |
| Icon | `~/.local/share/icons/minecraft.png` |

These are the standard XDG Base Directory locations for per-user installations.

## Download Pre-built Binary

Download the latest release from GitHub:

```bash
# For AMD64/x86_64 systems (most common)
wget https://github.com/mj41/mc-desktop/releases/latest/download/mc-installer-amd64
chmod +x mc-installer-amd64
./mc-installer-amd64

# For ARM64 systems
wget https://github.com/mj41/mc-desktop/releases/latest/download/mc-installer-arm64
chmod +x mc-installer-arm64
./mc-installer-arm64
```

Binaries are built on Fedora 43 for maximum compatibility.

## Build and Run

First, save the Minecraft icon as `minecraft-icon.png` in this directory. The icon and desktop template will be embedded directly into the binary.

```bash
# Build (icon and template are embedded into the binary)
go build -o mc-installer main.go
# Or use make
make build

# Show help
./mc-installer --help

# Run installer
./mc-installer

# Force reinstall (even if already installed)
./mc-installer --force

# Using make
make runto `/tmp/` (skipped if already present)
3. Extracts the launcher to `~/opt/minecraft-launcher/` (strips top-level directory from archive)
4. Installs the embedded icon to `~/.local/share/icons/minecraft.png`
5. Creates a `.desktop` file in `~/.local/share/applications/minecraft.desktop` from embedded template
6. Cleans up temporary download file
7he compiled binary is fully self-contained with the icon and desktop template embedded. You can distribute just the `mc-installer` executable.

## Command-Line Options

- `--help` - Show help message and installation paths
- `--force` - Force reinstallation even if already installed

## What it does

1. Creates the necessary directories (`~/opt/minecraft-launcher/`, `~/.local/share/applications/`, `~/.local/share/icons/`)
2. Downloads from https://launcher.mojang.com/download/Minecraft.tar.gz (skipped if already present)
3. Extracts the launcher to `~/opt/minecraft-launcher/`
4. Installs the embedded icon to `~/.local/share/icons/minecraft.png`
5. Creates a `.desktop` file in `~/.local/share/applications/minecraft.desktop` with embedded template
6. After installation, Minecraft will appear in your application menu

The installer will skip steps if already completed, unless `--force` is specified.

## Running Minecraft

After installation:
- Find "Minecraft" in your application menu
- Or run directly: `~/opt/minecraft-launcher/minecraft-launcher`

## Uninstallation

```bash
rm -rf ~/opt/minecraft-launcher
rm ~/.local/share/applications/minecraft.desktop
rm ~/.local/share/icons/minecraft.png
```
