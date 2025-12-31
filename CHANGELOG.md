# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.1] - 2025-12-31

### Added
- Quick install script for one-line installation via curl
- Support for passing command-line arguments through install script

### Changed
- Updated documentation with quick install instructions

## [0.3.0] - 2025-12-31

### Added
- Per-user installation to standard XDG directories
- Automatic download of Minecraft launcher from Mojang
- Desktop integration with application menu and icon
- Automatic GPU detection (AMD/Intel/NVIDIA) with appropriate configuration
- Idempotent installation (skips already completed steps)
- Self-contained binary with embedded icon and desktop template
- Multi-architecture support (AMD64 and ARM64)
- Command-line options: `--help`, `--version`, `--force`
- Automatic cleanup of temporary download files
- GitHub Actions workflow for automated builds on Fedora 43

### Technical Details
- Uses `lspci` to detect GPU type
- AMD/Intel/Nouveau GPUs: configured with `__GLX_VENDOR_LIBRARY_NAME=mesa`
- NVIDIA proprietary drivers: configured without Mesa override

### Installation Locations
- Application: `~/opt/minecraft-launcher/`
- Desktop file: `~/.local/share/applications/minecraft.desktop`
- Icon: `~/.local/share/icons/minecraft.png`

[Unreleased]: https://github.com/mj41/mc-desktop/compare/v0.3.1...HEAD
[0.3.1]: https://github.com/mj41/mc-desktop/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/mj41/mc-desktop/releases/tag/v0.3.0
