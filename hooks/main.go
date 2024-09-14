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

type ProfileError struct {
	Profile Profile
	Error   error
}

func (pe ProfileError) String() string {
	return fmt.Sprintf(
		"{\n\tProfile: %q,\n\tSource: %q,\n\tTarget: %q,\n\tError: %q\n}",
		pe.Profile.Label, pe.Profile.Source, pe.Profile.Target, pe.Error,
	)
}

const sharedHooksSourceDir = "scripts/shared"

func apiProfile() Profile {
	return Profile{
		Label:  "api submodule",
		Source: "scripts/api",
		Target: "../.git/modules/api/hooks",
	}
}

func webProfile() Profile {
	return Profile{
		Label:  "web submodule",
		Source: "scripts/web",
		Target: "../web/.husky/_",
	}
}

func monoProfile() Profile {
	return Profile{
		Label:  "mono",
		Source: "scripts/mono",
		Target: "../.git/hooks",
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		CopyAllHooks()
		return
	}

	switch args[0] {
	case "api":
		if err := CopyHooks(apiProfile()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "web":
		if err := CopyHooks(webProfile()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "mono":
		if err := CopyHooks(monoProfile()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid argument:", args[0])
		os.Exit(1)
	}
}

func CopyAllHooks() {
	var waitGroup sync.WaitGroup

	profiles := []Profile{
		apiProfile(),
		webProfile(),
		monoProfile(),
	}

	errors := make(chan ProfileError, len(profiles))

	for _, profile := range profiles {
		waitGroup.Add(1)
		go func(p Profile) {
			defer waitGroup.Done()
			if err := CopyHooks(p); err != nil {
				errors <- ProfileError{
					Profile: p,
					Error:   err,
				}
			}
		}(profile)
	}

	waitGroup.Wait()
	close(errors)

	if len(errors) > 0 {
		fmt.Println("❌Errors occurred while installing hooks:")

		for err := range errors {
			fmt.Println(err)
		}

		os.Exit(1)
	}
}

func CopyHooks(profile Profile) error {
	// Check if the shared directory exists
	if _, err := os.Stat(sharedHooksSourceDir); os.IsNotExist(err) {
		return fmt.Errorf("Shared directory not found: %v", err)
	}

	// Check if the source directory exists
	if _, err := os.Stat(profile.Source); os.IsNotExist(err) {
		return fmt.Errorf("Source directory not found: %v", err)
	}

	// Create the target directory if it doesn't exist
	if _, err := os.Stat(profile.Target); os.IsNotExist(err) {
		if err := os.MkdirAll(profile.Target, 0755); err != nil {
			return fmt.Errorf("Error creating target directory: %v", err)
		}
	}

	// Remove existing hooks in the target directory if any
	if err := clearDir(profile.Target); err != nil {
		return fmt.Errorf("Error removing existing hooks: %v", err)
	}

	// Copy shared hooks to the target directory
	if err := copyAll(sharedHooksSourceDir, profile.Target); err != nil {
		return fmt.Errorf("Error installing shared hooks: %v", err)
	}

	// Copy hooks to the target directory
	if err := copyAll(profile.Source, profile.Target); err != nil {
		return fmt.Errorf("Error installing profile hooks: %v", err)
	}

	fmt.Println("✅Hooks installed successfully for", profile.Label)
	return nil
}

func clearDir(dir string) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if err = os.Remove(path); err != nil {
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
			if err = os.WriteFile(destPath, input, 0755); err != nil {
				return err
			}
		}
		return nil
	})
}
