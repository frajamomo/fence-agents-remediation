package dast

import (
	"fmt"

	g "github.com/onsi/ginkgo/v2"
	// exutil "github.com/openshift/openshift-tests-private/test/extended/util"
)

var _ = g.Describe("FAR DAST", func() {
	defer g.GinkgoRecover()

	var (
	//		buildPruningBaseDir string
	)

	//	g.BeforeEach(func() {
	//		buildPruningBaseDir = exutil.FixturePath("testdata", "far-operator")
	//	})

	g.It("Author:frajamomo-High-XXXXX-FAR operator should pass DAST test", func() {
		//		configFile := filepath.Join(buildPruningBaseDir, "rapidast/data_rapidastconfig_compliance_v1alpha1.yaml")
		//		policyFile := filepath.Join(buildPruningBaseDir, "rapidast/customscan.policy")

		fmt.Printf("JAVIER TESTCASE\n")

		g.By("JAVIER")
		//		_, err := rapidastScan(oc, oc.Namespace(), configFile, policyFile, "compliance.openshift.io_v1alpha1")
		//		o.Expect(err).NotTo(o.HaveOccurred())
	})

})
