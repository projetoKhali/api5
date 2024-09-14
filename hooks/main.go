package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Profile struct {
	Label  string
	Source string
	Target string
}

const sharedHooks = "scripts/shared"

func api() Profile {
	return Profile{
		Label:  "api",
		Source: "scripts/api",
		Target: "../.git/modules/api/hooks",
	}
}

func web() Profile {
	return Profile{
		Label:  "web",
		Source: "scripts/web",
		Target: "../.git/modules/web/hooks",
	}
}

func mono() Profile {
	return Profile{
		Label:  "mono",
		Source: "scripts/mono",
		Target: "../.git/hooks",
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		var waitGroup sync.WaitGroup

		profiles := []Profile{
			api(),
			web(),
			mono(),
		}

		for _, profile := range profiles {
			waitGroup.Add(1)
			go func(p Profile) {
				defer waitGroup.Done()
				CopyHooks(p)
			}(profile)
		}

		waitGroup.Wait()

		return
	}

	switch args[0] {
	case "api":
		CopyHooks(api())
	case "web":
		CopyHooks(web())
	case "mono":
		CopyHooks(mono())
	default:
		fmt.Println("Invalid argument:", args[0])
		os.Exit(1)
	}
}

func CopyHooks(
	profile Profile,
) {
	// Check if the shared directory exists
	if _, err := os.Stat(sharedHooks); os.IsNotExist(err) {
		fmt.Println("Shared directory not found")
		os.Exit(1)
	}

	// Check if the source directory exists
	if _, err := os.Stat(profile.Source); os.IsNotExist(err) {
		fmt.Println("Source directory not found for profile", profile.Label)
		os.Exit(1)
	}

	// Create the target directory if it doesn't exist
	if _, err := os.Stat(profile.Target); os.IsNotExist(err) {
		err := os.MkdirAll(profile.Target, 0755)
		if err != nil {
			fmt.Println("Error creating target directory for profile", profile.Label, ":", err)
			os.Exit(1)
		}
	}

	// Remove existing hooks in the target directory if any
	err := clearDir(profile.Target)
	if err != nil {
		fmt.Println("Error removing existing hooks for profile", profile.Label, ":", err)
		os.Exit(1)
	}

	fmt.Println("Installing hooks for", profile.Label, "...")

	// Copy shared hooks to the target directory
	err = copyAll(sharedHooks, profile.Target)
	if err != nil {
		fmt.Println("Error installing hooks for profile", profile.Label, ":", err)
		os.Exit(1)
	}

	// Copy hooks to the target directory
	err = copyAll(profile.Source, profile.Target)
	if err != nil {
		fmt.Println("Error installing hooks for profile", profile.Label, ":", err)
		os.Exit(1)
	}

	fmt.Println("Hooks installed successfully for", profile.Label)
}

func clearDir(dir string) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func copyAll(sourceDir string, targetDir string) error {
	return filepath.WalkDir(sourceDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			relativePath, _ := filepath.Rel(sourceDir, path)
			destPath := filepath.Join(targetDir, relativePath)
			input, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			err = os.WriteFile(destPath, input, 0755)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
