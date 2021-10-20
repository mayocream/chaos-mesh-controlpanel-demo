package controlpanel

import (
	"context"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func applyPodKill(name, namespace string, labels map[string]string) error {
	cli := mgr.GetClient()

	cr := &v1alpha1.PodChaos{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name,
			Namespace:    namespace,
		},
		Spec: v1alpha1.PodChaosSpec{
			Action: v1alpha1.PodKillAction,
			ContainerSelector: v1alpha1.ContainerSelector{
				PodSelector: v1alpha1.PodSelector{
					Mode: v1alpha1.OnePodMode,
					Selector: v1alpha1.PodSelectorSpec{
						Namespaces:     []string{namespace},
						LabelSelectors: labels,
					},
				},
			},
		},
	}

	if err := cli.Create(context.Background(), cr); err != nil {
		return errors.Wrap(err, "create podkill")
	}

	return nil
}
