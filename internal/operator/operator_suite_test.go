package operator_test

import (
	"opcap/internal/operator"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestOperator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operator Suite")
}

func newFakeClient(testData []runtime.Object) operator.Opcap {
	
	scheme := runtime.NewScheme()
	operatorv1alpha1.AddToScheme(scheme)
	client := fake.NewClientBuilder().WithRuntimeObjects(testData...).WithScheme(scheme).Build()
	var operatorClient operator.Opcap = &operator.OpcapClient{
		Client: client,
	}
	return operatorClient
}
