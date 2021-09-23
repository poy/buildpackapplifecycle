// +build windows,windows2012R2

package main_test

import (
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func buildLauncher() string {
	launcherPath, err := gexec.Build("github.com/poy/buildpackapplifecycle/launcher", "-race", "-tags=windows2012R2")
	Expect(err).NotTo(HaveOccurred())

	return launcherPath
}
