package controlpanel

import (
	"context"
	"log"
	"time"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var schema = runtime.NewScheme()
var mgr manager.Manager

func init() {
	v1alpha1.AddToScheme(schema)
	NewManager()
}

func NewManager() {
	var err error
	mgr, err = ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: schema,
	})
	if err != nil {
		log.Fatal(err)
	}
	go mgr.Start(context.Background())
	time.Sleep(2 * time.Second)
}
