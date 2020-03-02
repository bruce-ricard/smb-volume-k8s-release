package integration_tests_test

import (
	local_k8s_cluster "code.cloudfoundry.org/smb-volume-k8s-local-cluster"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegrationTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTests Suite")
}

var _ = BeforeSuite(func() {
	SetDefaultEventuallyTimeout(10 * time.Minute)

	local_k8s_cluster.CreateK8sCluster("test", "/tmp/kubeconfig")
	local_k8s_cluster.CreateKpackImageResource()

	By("deploying the smb broker into the k8s cluster", func() {
		local_k8s_cluster.Kubectl("create", "namespace", "cf-workloads")
		local_k8s_cluster.Helm("install", "smb-broker", "../smb-broker/helm", "--set", "smbBrokerUsername=foo", "--set", "smbBrokerPassword=bar", "--set", "targetNamespace=cf-workloads", "--set", "ingress.hosts[0].host=localhost", "--set", "ingress.hosts[0].paths={/v2}", "--set", "ingress.enabled=true", "--set", "image.repository=registry:5000/cfpersi/smb-broker", "--set", "image.tag=local-test")
	})

	By("deploying smb csi driver into the k8s cluster", func() {
		local_k8s_cluster.Kubectl("apply", "--kustomize", "../smb-csi-driver/base")
		Eventually(func()string{
			return local_k8s_cluster.Kubectl("get", "pod", "-l", "app=csi-nodeplugin-smbplugin")
		}, 10 * time.Minute, 1 * time.Second).Should(ContainSubstring("Running"))
	})
})

var _ = AfterSuite(func() {
	local_k8s_cluster.DeleteK8sCluster("test", "/tmp/kubeconfig")
})
