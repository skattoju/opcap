package operator

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ = Describe("Install Plan Approval", func() {

	var installPlanList operatorv1alpha1.InstallPlanList
	var fakeClient OpcapClient

	Context("When InstallPlan Approval is manual", func() {

		BeforeEach(func() {
			installPlanList = operatorv1alpha1.InstallPlanList{
				Items: []operatorv1alpha1.InstallPlan{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "installPlan",
							Namespace: "fakeNS",
						},

						Spec: operatorv1alpha1.InstallPlanSpec{
							Approval: operatorv1alpha1.ApprovalManual,
						},
					},
				},
			}
			testData := []runtime.Object{&installPlanList}
			fakeClient = NewFakeClient(testData)
		})

		It("Should set Spec.approved", func() {
			err := fakeClient.InstallPlanApprove("fakeNS")
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
