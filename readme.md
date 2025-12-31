# Minecraft Launcher Installer for Fedora Linux

A simple Go application to install the Minecraft launcher per-user on Fedora Linux. Downloads, extracts, and configures the official Minecraft launcher with desktop integration.

## Features

- **No root required**: Per-user installation to standard XDG directories
- **Automatic download**: Fetches the latest launcher from Mojang
- **Desktop integration**: Adds launcher to application menu with icon
- **Auto GPU detection**: Automatically configures for AMD/Intel/NVIDIA GPUs
- **Idempotent**: Skips already completed steps unless `--force` is used
- **Self-contained binary**: Icon and desktop template embedded at build time
- **Multi-architecture**: Pre-built binaries for AMD64 and ARM64

## GPU Compatibility

The installer **automatically detects your GPU** and configures the desktop file accordingly:

- ✅ **AMD GPUs** (Radeon) - Configured with Mesa
- ✅ **Intel GPUs** - Configured with Mesa
- ✅ **NVIDIA proprietary drivers** - Configured without Mesa override
- ✅ **NVIDIA Nouveau** (open source) - Configured with Mesa

No manual configuration needed! The installer uses `lspci` to detect your GPU type.

## Per-User Installation Locations

| Item | Path |
|------|------|
| Application | `~/opt/minecraft-launcher/` |
| Desktop file | `~/.local/share/applications/minecraft.desktop` |
| Icon | `~/.local/share/icons/minecraft.png` |

These are the standard XDG Base Directory locations for per-user installations.

## Download Pre-built Binary

### Quick Install (Recommended)

One-line install - downloads and runs the installer automatically:

```bash
curl -fsSL https://raw.githubusercontent.com/mj41/mc-desktop/main/install.sh | bash
```

For force reinstall:
```bash
curl -fsSL https://raw.githubusercontent.com/mj41/mc-desktop/main/install.sh | bash -s -- --force
```

### Manual Download

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

Binaries are built with Go 1.23 as static executables - compatible with all Linux distributions including Fedora 43+.

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
make run
make run-force
```

The compiled binary is fully self-contained with the icon and desktop template embedded. You can distribute just the `mc-installer` executable.

## Command-Line Options

- `--help` - Show help message and installation paths
- `--version` - Show version information
- `--force` - Force reinstallation even if already installed

## What it does

1. Detects your GPU type (AMD/Intel/NVIDIA)
2. Creates the necessary directories (`~/opt/minecraft-launcher/`, `~/.local/share/applications/`, `~/.local/share/icons/`)
3. Downloads from https://launcher.mojang.com/download/Minecraft.tar.gz to `/tmp/` (skipped if already present)
4. Extracts the launcher to `~/opt/minecraft-launcher/` (strips top-level directory from archive)
5. Installs the embedded icon to `~/.local/share/icons/minecraft.png`
6. Creates a `.desktop` file with GPU-specific configuration
7. Cleans up temporary download file
8. After installation, Minecraft will appear in your application menu

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

## Documentation

- [Release Process](docs/dev/release.md) - How to create new releases
- [Man Page](docs/mc-installer.1) - Manual page (can be viewed with `man docs/mc-installer.1`)
- [Changelog](CHANGELOG.md) - Version history and changes

## Contributing

Contributions are welcome! Please ensure:
- Desktop file passes validation: `desktop-file-validate minecraft.desktop.tmpl`
- CHANGELOG.md is updated for notable changes
- Follow existing code style
