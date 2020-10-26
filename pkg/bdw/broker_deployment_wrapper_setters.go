package bdw
/* This file contains configuration setter methods for BrokerDeploymentWrapper
 */

import (
	brokerclientset "github.com/artemiscloud/activemq-artemis-operator/pkg/client/clientset/versioned"
	"github.com/rh-messaging/shipshape/pkg/framework"
)

// WithWait sets if shipshape would wait for completion
func (bdw *BrokerDeploymentWrapper) WithWait(wait bool) *BrokerDeploymentWrapper {
	bdw.wait = wait
	return bdw
}

func (bdw *BrokerDeploymentWrapper) WithName(name string) *BrokerDeploymentWrapper {
	bdw.name = name
	return bdw
}

// WithBrokerClient sets broker kubernetes client to use
func (bdw *BrokerDeploymentWrapper) WithBrokerClient(brokerClient brokerclientset.Interface) *BrokerDeploymentWrapper {
	bdw.brokerClient = brokerClient
	return bdw
}

// WithContext sets shipshape context
func (bdw *BrokerDeploymentWrapper) WithContext(ctx1 *framework.ContextData) *BrokerDeploymentWrapper {
	bdw.ctx1 = ctx1
	return bdw
}

// WithCustomImage wets Broker Image to be used
func (bdw *BrokerDeploymentWrapper) WithCustomImage(image string) *BrokerDeploymentWrapper {
	bdw.customImage = image
	return bdw
}

// WithMigration Sets Migration parameter (controls message migration availability)
func (bdw *BrokerDeploymentWrapper) WithMigration(migration bool) *BrokerDeploymentWrapper {
	bdw.migration = migration
	return bdw
}

// WithPersistence Sets Persistence parameter (controls creationf of PVCs)
func (bdw *BrokerDeploymentWrapper) WithPersistence(persistence bool) *BrokerDeploymentWrapper {
	bdw.persistence = persistence
	return bdw
}

func (bdw *BrokerDeploymentWrapper) WithSsl(ssl bool) *BrokerDeploymentWrapper {
	bdw.sslEnabled = ssl
	return bdw
}

func (bdw *BrokerDeploymentWrapper) WithConsoleExposure(expose bool) *BrokerDeploymentWrapper {
	bdw.exposeConsole = expose
	return bdw
}

func (bdw *BrokerDeploymentWrapper) WithLts(lts bool) *BrokerDeploymentWrapper {
	bdw.isLtsDeployment = lts
	return bdw
}

