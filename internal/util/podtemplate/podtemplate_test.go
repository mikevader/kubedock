package podtemplate

import (
	"testing"
)

func TestPodFromFile(t *testing.T) {
	pod, err := PodFromFile("test/test_pod.yaml")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if pod == nil {
		t.Error("error unmarshalling pod")
	}
	if pod != nil && pod.Spec.ServiceAccountName != "kubedock" {
		t.Error("invalid serviceAccountName")
	}

	pod, err = PodFromFile("test/notfound.yaml")
	if pod != nil {
		t.Error("unexpected pod object")
	}
	if err == nil {
		t.Error("expected an error when file is not available")
	}

	pod, err = PodFromFile("test/test_invalid_kind.yaml")
	if pod != nil {
		t.Error("unexpected pod object")
	}
	if err == nil {
		t.Error("expected an error when kind is not a pod")
	}

	pod, err = PodFromFile("test/test_invalid.yaml")
	if pod != nil {
		t.Error("unexpected pod object")
	}
	if err == nil {
		t.Error("expected an error when file is invalid yaml")
	}
}
