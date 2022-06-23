package operator

import (
	"fmt"
	"testing"

	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestInstallPlanApprove(t *testing.T) {

	//init fake client
	installPlanList := operatorv1alpha1.InstallPlanList{
		Items: []operatorv1alpha1.InstallPlan{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "installPlan1",
					Namespace: "fakens1",
				},

				Spec: operatorv1alpha1.InstallPlanSpec{
					Approved: false,
				},
			},
		},
	}
	testData := []runtime.Object{&installPlanList}
	fakeOpcapClient, _ := newFakeClient(testData)

	//call with fake client
	err := fakeOpcapClient.InstallPlanApprove("fakens1")
	if err != nil {
		t.Fatal("installPlanApprove failed")
	}
	t.Log("InstallPlanApprove PASS")
	err = nil
}

func TestInstallPlanApproveFail(t *testing.T) {

	//init fake client
	installPlanList := operatorv1alpha1.InstallPlanList{
		Items: []operatorv1alpha1.InstallPlan{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "installPlan1",
					Namespace: "fakens1",
				},

				Spec: operatorv1alpha1.InstallPlanSpec{
					Approved: false,
				},
			},
		},
	}
	testData := []runtime.Object{&installPlanList}
	fakeOpcapClient, _ := newFakeClient(testData)

	//call with fake client
	err := fakeOpcapClient.InstallPlanApprove("fakens2")
	if err != nil {
		t.Log("InstallPlanApprove FAILED due to wrong namespace name")
	}

	err = nil
}

func newFakeClient(testData []runtime.Object) (Opcap, error) {

	scheme := runtime.NewScheme()

	if err := operatorv1alpha1.AddToScheme(scheme); err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := fake.NewClientBuilder().WithRuntimeObjects(testData...).WithScheme(scheme).Build()
	var operatorClient Opcap = &opcapClient{
		Client: client,
	}
	return operatorClient, nil
}
