package controlpanel

import (
	"testing"
)

func Test_applyPodKill(t *testing.T) {
	err := applyPodKill("podkill", "msp", map[string]string{
		"app.kubernetes.io/instance": "gw-kic",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("apply podkill")
}
