//go:build mage
// +build mage

package main

import (
	"os"

	"get.porter.sh/porter/mage/mixins"
	"get.porter.sh/porter/mage/releases"

	// Import common targets that all mixins should expose to the user
	// mage:import
	_ "get.porter.sh/porter/mage"
)

const (
	mixinName    = "ansible"
	mixinPackage = "get.porter.sh/mixin/ansible"
	mixinBin     = "bin/mixins/" + mixinName
)

var magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)

// Build the mixin
func Build() {
	magefile.Build()
}

// Cross-compile the mixin before a release
func XBuildAll() {
	magefile.XBuildAll()
}

// Run unit tests
func TestUnit() {
	magefile.TestUnit()
}

func Test() {
	magefile.Test()
}

// Publish the mixin to github
func Publish() {
	// You can test out publishing locally by overriding PORTER_RELEASE_REPOSITORY and PORTER_PACKAGES_REMOTE
	if _, overridden := os.LookupEnv(releases.ReleaseRepository); !overridden {
		os.Setenv(releases.ReleaseRepository, "github.com/LvffY/ansible")
	}
	magefile.PublishBinaries()

	// TODO: uncomment out the lines below to publish a mixin feed
	// Set PORTER_PACKAGES_REMOTE to a repository that will contain your mixin feed, similar to github.com/getporter/packages
	//if _, overridden := os.LookupEnv(releases.PackagesRemote); !overridden {
	//	os.Setenv("PORTER_PACKAGES_REMOTE", "git@github.com:LvffY/YOUR_PACKAGES_REPOSITORY")
	//}
	//magefile.PublishMixinFeed()
}

// Install the mixin
func Install() {
	magefile.Install()
}

// Remove generated build files
func Clean() {
	magefile.Clean()
}
