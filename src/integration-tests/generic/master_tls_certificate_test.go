package generic_test

import (
	"integration-tests/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MasterTlsCertificate", func() {

	var (
		runner *test_helpers.KubectlRunner
	)

	BeforeEach(func() {
		runner = test_helpers.NewKubectlRunner()
	})

	It("should be valid for kube-dns names for master", func() {
		session := runner.RunKubectlCommand("run", "test-master-cert-via-curl", "--image=governmentpaas/curl-ssl", "--restart=Never", "-ti", "--rm", "--", "curl", "https://kubernetes.default.svc.cluster.local", "--cacert", "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt")
		<-session.Exited
	  stdo := string(session.Out.Contents())
		Expect(stdo).To(ContainSubstring("User \"system:anonymous\" cannot get path \"/\".: \"No policy matched.\""))
	})
})
