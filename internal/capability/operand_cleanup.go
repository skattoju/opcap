package capability

import (
	"context"
	"fmt"

	"github.com/opdev/opcap/internal/logger"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// OperandCleanup removes the operand from the OCP cluster in the ca.namespace
func operandCleanup(ctx context.Context, opts ...auditOption) auditCleanupFn {
	var options auditOptions
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return func(_ context.Context) error {
				return fmt.Errorf("option failed: %v", err)
			}
		}
	}

	return func(ctx context.Context) error {
		logger.Debugw("cleaningUp operand for operator", "package", options.subscription.Package, "channel", options.subscription.Channel, "installmode",
			options.subscription.InstallModeType)

		if len(options.customResources) > 0 {
			for _, cr := range options.customResources {
				obj := &unstructured.Unstructured{Object: cr}

				// extract name from CustomResource object and delete it
				name := obj.Object["metadata"].(map[string]interface{})["name"].(string)

				// check if CR exists, only then cleanup the operand
				err := options.client.GetUnstructured(ctx, options.namespace, name, obj)
				if !apierrors.IsNotFound(err) {
					// Actual error. Return it
					return fmt.Errorf("could not get operaand: %v", err)
				}
				if obj == nil || apierrors.IsNotFound(err) {
					// Did not find it. Somehow already gone.
					// Not an error condition, but no point in
					// continuing.
					return nil
				}

				// delete the resource using the dynamic client
				if err := options.client.DeleteUnstructured(ctx, obj); err != nil {
					logger.Debugf("failed operandCleanUp: package: %s error: %s\n", options.subscription.Package, err.Error())
					return err
				}

				// Forcing cleanup of finalizers
				err = options.client.GetUnstructured(ctx, options.namespace, name, obj)
				if apierrors.IsNotFound(err) {
					return nil
				}

				obj.SetFinalizers([]string{})

				if err := options.client.UpdateUnstructured(ctx, obj); err != nil {
					return err
				}

				if err := options.client.GetUnstructured(ctx, options.namespace, name, obj); err != nil && !apierrors.IsNotFound(err) {
					return fmt.Errorf("error cleaning up operand after deleting finalizer: %v", err)
				}
				return nil
			}
		}

		return nil
	}
}
