package messaging

import (
	"github.com/onsi/gomega"
	"github.com/rh-messaging/shipshape/pkg/framework/ginkgowrapper"
	"gitlab.cee.redhat.com/msgqe/openshift-broker-suite-golang/test"
	"testing"
)

func TestMessaging(t *testing.T) {

	gomega.RegisterFailHandler(ginkgowrapper.Fail)
	test.PrepareNamespace(t, "messaging", "Messaging Suite")
}
