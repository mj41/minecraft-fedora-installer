// Copyright (c) 2025. All rights reserved.
// Minecraft Desktop Installer for Fedora Linux
// Installs per-user: desktop file, icon, and downloads/extracts the launcher.

package main

import (
	"archive/tar"
	"compress/gzip"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed minecraft-icon.png
var embeddedIcon []byte

//go:embed minecraft.desktop.tmpl
var desktopTemplate string

// Version information (set via ldflags during build)
var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

const (
	// Download URL for Minecraft launcher
	minecraftURL = "https://launcher.mojang.com/download/Minecraft.tar.gz"
)

// Per-user installation paths
type InstallPaths struct {
	Home        string // User home directory
	OptDir      string // ~/opt/minecraft-launcher - where the app is installed
	DesktopDir  string // ~/.local/share/applications - desktop files
	IconDir     string // ~/.local/share/icons - icons
	DesktopFile string // Full path to .desktop file
	IconFile    string // Full path to icon file
}

func getInstallPaths() (*InstallPaths, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	return &InstallPaths{
		Home:        home,
		OptDir:      filepath.Join(home, "opt", "minecraft-launcher"),
		DesktopDir:  filepath.Join(home, ".local", "share", "applications"),
		IconDir:     filepath.Join(home, ".local", "share", "icons"),
		DesktopFile: filepath.Join(home, ".local", "share", "applications", "minecraft.desktop"),
		IconFile:    filepath.Join(home, ".local", "share", "icons", "minecraft.png"),
	}, nil
}

func main() {
	// Parse command-line flags
	force := flag.Bool("force", false, "Force reinstallation even if already installed")
	help := flag.Bool("help", false, "Show help message")
	version := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if *version {
		fmt.Printf("mc-installer version %s\n", Version)
		fmt.Printf("Git commit: %s\n", GitCommit)
		fmt.Printf("Build date: %s\n", BuildDate)
		return
	}

	if *help {
		fmt.Println("Minecraft Launcher Installer for Fedora Linux")
		fmt.Printf("Version: %s\n", Version)
		fmt.Println()
		fmt.Println("Usage: mc-installer [options]")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  --force      Force reinstallation even if already installed")
		fmt.Println("  --help       Show this help message")
		fmt.Println("  --version    Show version information")
		fmt.Println()
		fmt.Println("Installation locations (per-user):")
		fmt.Println("  Application:   ~/opt/minecraft-launcher/")
		fmt.Println("  Desktop file:  ~/.local/share/applications/minecraft.desktop")
		fmt.Println("  Icon:          ~/.local/share/icons/minecraft.png")
		return
	}

	fmt.Println("=== Minecraft Launcher Installer for Fedora Linux ===")
	fmt.Println()

	paths, err := getInstallPaths()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Installation paths (per-user):")
	fmt.Printf("  Application: %s\n", paths.OptDir)
	fmt.Printf("  Desktop file: %s\n", paths.DesktopFile)
	fmt.Printf("  Icon: %s\n", paths.IconFile)
	fmt.Println()

	// Check if already installed
	if !*force && isAlreadyInstalled(paths) {
		fmt.Println("Minecraft launcher is already installed!")
		fmt.Println("Use --force to reinstall")
		return
	}

	// Create directories
	if err := createDirectories(paths); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directories: %v\n", err)
		os.Exit(1)
	}

	// Download and extract
	if err := downloadAndExtract(minecraftURL, paths.OptDir, *force); err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading/extracting: %v\n", err)
		os.Exit(1)
	}

	// Install icon (bundled or from extracted files)
	if err := installIcon(paths, *force); err != nil {
		fmt.Printf("Warning: Could not install icon: %v\n", err)
	}

	// Create desktop file
	if err := createDesktopFile(paths, *force); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating desktop file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("=== Installation complete! ===")
	fmt.Println("You can now find Minecraft in your application menu.")
	fmt.Println("Or run directly:", filepath.Join(paths.OptDir, "minecraft-launcher"))
}

func isAlreadyInstalled(paths *InstallPaths) bool {
	// Check if launcher executable exists
	launcherPath := filepath.Join(paths.OptDir, "minecraft-launcher")
	if _, err := os.Stat(launcherPath); err != nil {
		return false
	}

	// Check if desktop file exists
	if _, err := os.Stat(paths.DesktopFile); err != nil {
		return false
	}

	return true
}

func createDirectories(paths *InstallPaths) error {
	dirs := []string{
		paths.OptDir,
		paths.DesktopDir,
		paths.IconDir,
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Creating directory: %s\n", dir)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create %s: %w", dir, err)
			}
		}
	}
	return nil
}

func downloadAndExtract(url, destDir string, force bool) error {
	// Check if already extracted
	launcherPath := filepath.Join(destDir, "minecraft-launcher")
	if !force {
		if _, err := os.Stat(launcherPath); err == nil {
			fmt.Println("Launcher already downloaded and extracted (skipping)")
			return nil
		}
	}

	// Download to /tmp/
	tmpFile := filepath.Join("/tmp", "minecraft-launcher.tar.gz")
	fmt.Printf("Downloading to: %s\n", tmpFile)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %s", resp.Status)
	}

	// Write to temporary file
	out, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile) // Clean up

	if _, err := io.Copy(out, resp.Body); err != nil {
		out.Close()
		return fmt.Errorf("failed to save download: %w", err)
	}
	out.Close()

	// Extract from temp file
	fmt.Println("Extracting archive...")
	tmpFileReader, err := os.Open(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to open temp file: %w", err)
	}
	defer tmpFileReader.Close()

	return extractTarGz(tmpFileReader, destDir)
}

func extractTarGz(r io.Reader, destDir string) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar read error: %w", err)
		}

		// Remove the top-level directory from the path (e.g., "minecraft-launcher/")
		name := header.Name
		parts := strings.SplitN(name, "/", 2)
		if len(parts) < 2 || parts[1] == "" {
			continue // Skip top-level directory itself
		}
		name = parts[1]

		target := filepath.Join(destDir, name)

		// Prevent path traversal attacks
		if !strings.HasPrefix(filepath.Clean(target), filepath.Clean(destDir)) {
			return fmt.Errorf("invalid file path: %s", name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", target, err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", target, err)
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return fmt.Errorf("failed to write file %s: %w", target, err)
			}
			f.Close()
			fmt.Printf("  Extracted: %s\n", name)
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				// Ignore symlink errors, might already exist
				fmt.Printf("  Symlink skipped: %s\n", name)
			}
		}
	}

	return nil
}

func installIcon(paths *InstallPaths, force bool) error {
	// Check if icon already exists
	if !force {
		if _, err := os.Stat(paths.IconFile); err == nil {
			fmt.Println("Icon already installed (skipping)")
			return nil
		}
	}

	// Use embedded icon (compiled into binary)
	if len(embeddedIcon) > 0 {
		fmt.Println("Installing embedded icon...")
		if err := os.WriteFile(paths.IconFile, embeddedIcon, 0644); err != nil {
			return fmt.Errorf("failed to write embedded icon: %w", err)
		}
		return nil
	}

	// Fallback: look for icon in extracted files
	possibleIcons := []string{
		filepath.Join(paths.OptDir, "minecraft-launcher.png"),
		filepath.Join(paths.OptDir, "icon.png"),
		filepath.Join(paths.OptDir, "minecraft.png"),
	}

	var srcIcon string
	for _, icon := range possibleIcons {
		if _, err := os.Stat(icon); err == nil {
			srcIcon = icon
			break
		}
	}

	if srcIcon == "" {
		return fmt.Errorf("no icon found (neither embedded nor in extracted files)")
	}

	fmt.Printf("Copying icon from: %s\n", srcIcon)

	src, err := os.Open(srcIcon)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(paths.IconFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func createDesktopFile(paths *InstallPaths, force bool) error {
	// Check if desktop file already exists
	if !force {
		if _, err := os.Stat(paths.DesktopFile); err == nil {
			fmt.Println("Desktop file already exists (skipping)")
			return nil
		}
	}

	fmt.Printf("Creating desktop file: %s\n", paths.DesktopFile)

	content := fmt.Sprintf(desktopTemplate, paths.OptDir, paths.IconFile)

	if err := os.WriteFile(paths.DesktopFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write desktop file: %w", err)
	}

	return nil
}
