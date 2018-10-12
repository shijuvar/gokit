// +build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	binaryOutput = "./cmd/appd"
	packageName  = "github.com/shijuvar/gokit/examples/http-app/cmd/appd"
	goexe        = "go"
)

func getDep() error {
	return sh.Run(goexe, "get", "-u", "github.com/golang/dep/cmd/dep")
}

// Install Go Dep and sync vendored dependencies
func Vendor() error {
	mg.Deps(getDep)
	return sh.Run("dep", "ensure")
}

// Build the app
func Build() error {
	mg.Deps(Vendor)
	return sh.Run(goexe, "build", "-o", binaryOutput, packageName)
}

// Build binary with race detector enabled
func BuildRace() error {
	mg.Deps(Vendor)
	return sh.Run(goexe, "build", "-race", "-o", binaryOutput, packageName)
}

// Install the binary
func Install() error {
	mg.Deps(Vendor)
	return sh.Run(goexe, "install", packageName)
}

var docker = sh.RunCmd("docker")

// Build Docker container
func Docker() error {
	if err := docker("build", "-t", "http-app", "."); err != nil {
		return err
	}
	return nil
}

// Run tests and linters
func Check() {
	if strings.Contains(runtime.Version(), "1.8") {
		// Go 1.8 doesn't play along with go test ./... and /vendor.
		fmt.Printf("Skip Check on %s\n", runtime.Version())
		return
	}
	mg.Deps(TestRace)
}

// Run tests
func Test() error {
	mg.Deps(getDep)
	return sh.Run(goexe, "test", "./...")
}

// Run tests with race detector
func TestRace() error {
	mg.Deps(getDep)
	return sh.Run(goexe, "test", "-race", "./...")
}

// Run gofmt linter
func Fmt() error {
	pkgs, err := httpAppPackages()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// gofmt doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("gofmt", "-l", f)
			if err != nil {
				fmt.Printf("ERROR: running gofmt on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not gofmt'ed:")
					first = false
				}
				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

var pkgPrefixLen = len("github.com/shijuvar/gokit/examples/http-app")

func httpAppPackages() ([]string, error) {
	mg.Deps(getDep)
	s, err := sh.Output(goexe, "list", "./...")
	if err != nil {
		return nil, err
	}
	pkgs := strings.Split(s, "\n")
	for i := range pkgs {
		pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
	}
	return pkgs, nil
}

// Run golint linter
func Lint() error {
	pkgs, err := httpAppPackages()
	if err != nil {
		return err
	}
	failed := false
	for _, pkg := range pkgs {
		// We don't actually want to fail this target if we find golint errors,
		// so we don't pass -set_exit_status, but we still print out any failures.
		if _, err := sh.Exec(nil, os.Stderr, nil, "golint", pkg); err != nil {
			fmt.Printf("ERROR: running go lint on %q: %v\n", pkg, err)
			failed = true
		}
	}
	if failed {
		return errors.New("errors running golint")
	}
	return nil
}

//  Run go vet linter
func Vet() error {
	mg.Deps(getDep)
	if err := sh.Run(goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("error running govendor: %v", err)
	}
	return nil
}
