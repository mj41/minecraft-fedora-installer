# AI Agent Instructions for mc-desktop

This file contains instructions for AI coding agents working on the mc-desktop project.

## Project Overview

mc-desktop is a Go application that installs the Minecraft launcher per-user on Fedora Linux with automatic GPU detection and desktop integration.

## Release Process Checklist

When preparing a new release, follow these steps in order:

### Pre-Release

- [ ] All code changes are committed and tested
- [ ] Run `make build` to verify compilation succeeds
- [ ] Test binary locally with `./mc-installer --version` and `./mc-installer --help`
- [ ] Verify GPU detection works (check code in `detectGPU()` function)

### Version Updates

- [ ] Update version in `CHANGELOG.md`:
  - Move changes from `[Unreleased]` to new version section
  - Set release date (format: `YYYY-MM-DD`)
  - Add version comparison link at bottom
- [ ] Update man page version in `docs/mc-installer.1` (first line)
- [ ] Create release notes in `temp/release-notes/vX.Y.Z.md` (copy template from latest)
- [ ] Update any version-specific references in documentation

### File Checklist

Files that typically need updates:
- `CHANGELOG.md` - Version, date, changes
- `docs/mc-installer.1` - Version number in header
- `temp/release-notes/vX.Y.Z.md` - New file for release

Files to verify:
- `.github/workflows/release.yml` - Should use official GitHub CLI
- `readme.md` - Install URLs should point to correct repository
- `install.sh` - URL in comments should be correct
- `.github/RELEASE_TEMPLATE.md` - Template for future releases

### Repository URLs

Correct repository name: `minecraft-fedora-installer` (not `mc-desktop`)

URLs should use:
- `https://github.com/mj41/minecraft-fedora-installer`
- `https://raw.githubusercontent.com/mj41/minecraft-fedora-installer/refs/heads/main/install.sh`

### Git Operations

```bash
# 1. Commit version updates
git add CHANGELOG.md docs/mc-installer.1 temp/release-notes/vX.Y.Z.md
git commit -m "Prepare vX.Y.Z release"
git push origin main

# 2. Create and push tag
git tag -a vX.Y.Z -m "Release version X.Y.Z"
git push origin vX.Y.Z

# 3. Wait for GitHub Actions to complete
# 4. Create release on GitHub using notes from temp/release-notes/vX.Y.Z.md
```

### If Release Fails

```bash
# Delete failed release on GitHub UI first, then:
git tag -d vX.Y.Z
git push --delete origin vX.Y.Z

# Fix issues, commit, then recreate tag
git tag -a vX.Y.Z -m "Release version X.Y.Z"
git push origin vX.Y.Z
```

## Common Release Issues

### Build Failures

- Check `.github/workflows/release.yml` has `permissions: contents: write`
- Verify workflow uses official GitHub CLI (`gh`) for uploads
- Ensure binaries are built with proper ldflags for version info

### URL Issues

- Repository is `minecraft-fedora-installer` not `mc-desktop`
- Install script URL: `https://raw.githubusercontent.com/mj41/minecraft-fedora-installer/refs/heads/main/install.sh`
- Use `/refs/heads/main/` not just `/main/`

### Version Numbering

Follow semantic versioning:
- MAJOR.MINOR.PATCH (e.g., 0.3.4)
- Prefix with `v` in tags (e.g., `v0.3.4`)
- Current version series: 0.3.x (not stable 1.0 yet)

## Code Guidelines

- Binary must be statically linked (CGO_ENABLED=0)
- GPU detection uses `lspci` and `/sys/module/nvidia`
- Desktop template is embedded via `//go:embed`
- Icon is embedded via `//go:embed`
- All user-facing paths use XDG Base Directory spec

## Testing

```bash
# Build
make build

# Test version
./mc-installer --version

# Test help
./mc-installer --help

# Test installation (will actually install)
./mc-installer

# Force reinstall
./mc-installer --force
```

## Documentation

- Man page: `docs/mc-installer.1`
- Release notes: `temp/release-notes/vX.Y.Z.md`
- Changelog: `CHANGELOG.md`
- Main readme: `readme.md`
