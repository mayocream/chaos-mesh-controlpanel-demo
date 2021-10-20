package controlpanel

import (
	"testing"
)

func TestGetPodChaos(t *testing.T) {
	_, err := GetPodChaos("podkillvjn77", "msp")
	if err != nil {
		t.Fatal(err)
	}
}

func TestListPodChaos(t *testing.T) {
	_, err := ListPodChaos("msp", nil)
	if err != nil {
		t.Fatal(err)
	}
}
