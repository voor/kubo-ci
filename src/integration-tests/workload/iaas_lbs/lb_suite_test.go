package iaas_lbs_test

import (
	"integration-tests/test_helpers"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIAASLb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IAASLb Suite")
}

var (
	runner        *test_helpers.KubectlRunner
	nginxSpec     = test_helpers.PathFromRoot("specs/nginx.yml")
	workerAddress string
)

var _ = BeforeSuite(func() {
	workerAddress = os.Getenv("WORKLOAD_ADDRESS")
	if workerAddress == "" {
		Fail("WORKLOAD_ADDRESS is not set")
	}

	runner = test_helpers.NewKubectlRunner()
	runner.RunKubectlCommand("create", "namespace", runner.Namespace()).Wait("60s")
})

var _ = AfterSuite(func() {
	if runner != nil {
		runner.RunKubectlCommand("delete", "namespace", runner.Namespace()).Wait("60s")
	}
})