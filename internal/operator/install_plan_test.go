package operator_test

import (
	"context"
	"fmt"
	"opcap/internal/operator"
	"testing"

	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestInstallPlanApprove(t *testing.T) {

	//init fake client
	installPlan := operatorv1alpha1.InstallPlan{

		ObjectMeta: metav1.ObjectMeta{
			Name:      "installPlan1",
			Namespace: "fakens1",
		},
		Spec: operatorv1alpha1.InstallPlanSpec{
			Approved: false,
			Approval: operatorv1alpha1.ApprovalManual,
		},
	}

	installPlanList := operatorv1alpha1.InstallPlanList{
		Items: []operatorv1alpha1.InstallPlan{
			installPlan,
		},
	}

	testData := []runtime.Object{&installPlanList}
	fakeOpcapClient, _ := newFakeClient(testData)

	//call with fake client
	err := fakeOpcapClient.InstallPlanApprove("fakens1")
	if err != nil {
		t.Log("InstallPlanApprove should not error when install plan is found")
		t.Log(err)
		t.Fail()
	}
	var updatedIp operatorv1alpha1.InstallPlan
	fakeOpcapClient.Client.Get(context.TODO(), types.NamespacedName{
		Namespace: "fakens1",
		Name:      "installPlan1",
	}, &updatedIp)

	fmt.Println("Did it work?!", updatedIp.Spec.Approved)
}

func TestInstallPlanApproveEmpty(t *testing.T) {

	//init fake client
	installPlanList := operatorv1alpha1.InstallPlanList{
		Items: []operatorv1alpha1.InstallPlan{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "installPlan1",
					Namespace: "fakens1",
				},

				Spec: operatorv1alpha1.InstallPlanSpec{
					Approval: operatorv1alpha1.ApprovalManual,
					Approved: false,
				},
			},
		},
	}
	testData := []runtime.Object{&installPlanList}
	fakeOpcapClient, _ := newFakeClient(testData)

	//call with fake client
	err := fakeOpcapClient.InstallPlanApprove("fakens2")
	if err == nil {
		t.Log("InstallPlanApprove should error when install plan is not found")
		t.Fail()
	}
}

func newFakeClient(testData []runtime.Object) (*operator.OpcapClient, error) {

	scheme := runtime.NewScheme()

	if err := operatorv1alpha1.AddToScheme(scheme); err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := fake.NewClientBuilder().WithRuntimeObjects(testData...).WithScheme(scheme).Build()
	// here

	return &operator.OpcapClient{
		Client: client,
	}, nil
}
