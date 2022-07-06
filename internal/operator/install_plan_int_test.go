package operator

import (
	"testing"

	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ip = operatorv1alpha1.InstallPlan{

	ObjectMeta: metav1.ObjectMeta{
		Name:      "installPlan1",
		Namespace: "fakens1",
	},
	Spec: operatorv1alpha1.InstallPlanSpec{
		Approved: false,
		Approval: operatorv1alpha1.ApprovalManual,
	},
}

func TestFooBar(t *testing.T) {
	itsapproved := approvedInstallPlan(ip)
	if itsapproved.Spec.Approved != true {
		t.Log("Failed")
		t.Fail()
	}
}
