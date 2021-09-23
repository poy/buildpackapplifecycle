package profile_test

import (
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var getenv string

func TestProfile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buildpack-Lifecycle-Launcher-Profile Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	getenvPath, err := gexec.Build("github.com/poy/buildpackapplifecycle/getenv")
	Expect(err).NotTo(HaveOccurred())
	return []byte(getenvPath)
}, func(getenvExe []byte) {
	if runtime.GOOS == "windows" {
		getenv = string(getenvExe)
	}
})

var _ = SynchronizedAfterSuite(func() {
	//noop
}, func() {
	gexec.CleanupBuildArtifacts()
})
