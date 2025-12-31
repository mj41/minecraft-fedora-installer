# Release Process

This document describes how to create a new release of mc-desktop.

## Prerequisites

- Push access to the repository
- All changes committed and pushed to main branch
- Clean working directory
- CHANGELOG.md updated with release information

## Release Steps

### 1. Update CHANGELOG

Edit `CHANGELOG.md`:

```bash
# Update [Unreleased] section with the new version
# Change date from YYYY-MM-DD to actual date
# Move unreleased changes to the version section
```

Example:
```markdown
## [0.3.0] - 2025-12-31

### Added
- Feature 1
- Feature 2
```

Commit the changes:
```bash
git add CHANGELOG.md
git commit -m "Update CHANGELOG for v0.3.0"
git push origin main
```

### 2. Create and Push Tag

```bash
# Ensure you're on main branch and up to date
git checkout main
git pull origin main

# Create a new tag (use semantic versioning: v0.3.0, v0.4.0, v1.0.0, etc.)
git tag -a v0.3.0 -m "Release version 0.3.0"

# Push the tag to GitHub
git push origin v0.3.0
```

### 3. GitHub Actions Workflow

Once you push the tag, GitHub Actions will automatically:

1. Trigger the release workflow (`.github/workflows/release.yml`)
2. Build binaries on Fedora 43 for both architectures:
   - `mc-installer-amd64` (x86_64)
   - `mc-installer-arm64` (ARM64)
3. Embed version information via ldflags:
   - Version: from the git tag (e.g., `v1.0.0`)
   - Git commit: short commit hash
   - Build date: UTC timestamp
4. Upload binaries as release assets

### 4. Create GitHub Release

After the workflow completes:

1. Go to https://github.com/mj41/mc-desktop/releases
2. Click "Draft a new release"
3. Select the tag you just created (e.g., `v0.3.0`)
4. The release notes will be pre-filled from `.github/RELEASE_TEMPLATE.md`
5. Update the template with actual changes:
   - Fill in new features, improvements, bug fixes
   - Update version-specific information
6. The binaries will already be attached automatically
7. Click "Publish release"

Note: GitHub automatically uses `.github/RELEASE_TEMPLATE.md` when creating releases.

## Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version (v2.0.0): Incompatible changes
- **MINOR** version (v1.1.0): New features, backwards compatible
- **PATCH** version (v1.0.1): Bug fixes, backwards compatible

## Testing a Release Locally

Before creating an official release, test the build process:

```bash
# Build with version info
make build VERSION=v0.3.0

# Check version
./mc-installer --version

# Test installation
./mc-installer
```

## Release Checklist

- [ ] CHANGELOG.md updated with release notes
- [ ] All tests pass
- [ ] Documentation is up to date (README.md, man page)
- [ ] Desktop file is valid (can test with `desktop-file-validate`)
- [ ] Version tag follows semantic versioning
- [ ] Tag is annotated with meaningful message
- [ ] Workflow completes successfully
- [ ] Both binaries are attached to the release
- [ ] Release notes updated from template with actual changes
- [ ] Binaries are tested on target systems

## Rollback

If a release has issues:

1. Delete the GitHub release
2. Delete the tag from GitHub: `git push --delete origin v0.3.0`
3. Delete local tag: `git tag -d v0.3.0`
4. Fix issues and create a new release

## Example Release Notes Template
Release Template

The release notes template is stored in `.github/RELEASE_TEMPLATE.md` and is automatically used when creating new releases on GitHub. Update this template if you want to change the default release notes format.