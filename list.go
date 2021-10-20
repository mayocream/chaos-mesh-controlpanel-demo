package controlpanel

import (
	"context"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetPodChaos(name, namespace string) (*v1alpha1.PodChaos, error) {
	cli := mgr.GetClient()

	item := new(v1alpha1.PodChaos)
	if err := cli.Get(context.Background(), client.ObjectKey{Name: name, Namespace: namespace}, item); err != nil {
		return nil, errors.Wrap(err, "get cr")
	}

	return item, nil
}

func ListPodChaos(namespace string, labels map[string]string) ([]v1alpha1.PodChaos, error) {
	cli := mgr.GetClient()

	list := new(v1alpha1.PodChaosList)
	if err := cli.List(context.Background(), list, client.InNamespace(namespace), client.MatchingLabels(labels)); err != nil {
		return nil, err
	}

	return list.Items, nil
}
