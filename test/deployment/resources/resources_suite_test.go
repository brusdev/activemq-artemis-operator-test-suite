package resources

import (
	"testing"

	"gitlab.cee.redhat.com/msgqe/openshift-broker-suite-golang/test"
)

func TestMessaging(t *testing.T) {
	test.PrepareNamespace(t, "resources", "Resources Limitation Suite")
}

func TestMain(m *testing.M) {
	test.Initialize(m)
}