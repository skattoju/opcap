package operator_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"opcap/internal/operator"

	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
)

var _ = Describe("Install Plan Approval", func() {

	var installPlanList operatorv1alpha1.InstallPlanList
	var namespacedName types.NamespacedName
	var client operator.OpcapClient

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
							Approved: false,
						},
					},
				},
			}

			namespacedName = types.NamespacedName{
				Name:      "installPlan",
				Namespace: "fakeNS",
			}

			client = newFakeClient([]runtime.Object{&installPlanList})

		})

		It("Should set Spec.approved to true", func() {
			client.InstallPlanApprove("fakeNS")
			updatedInstallPlanList := operatorv1alpha1.InstallPlanList{
				Items: []operatorv1alpha1.InstallPlan{
					{},
				},
			}
			client.Client.Get(context.Background(), namespacedName, &runtime.Object{&installPlanList})
			Expect(updatedInstallPlanList.Items[0].Spec.Approved).To(BeTrue())
		})
	})

})
