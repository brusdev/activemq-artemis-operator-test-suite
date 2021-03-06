package addresssettings

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/rh-messaging/shipshape/pkg/framework"
	"github.com/artemiscloud/activemq-artemis-operator-test-suite/pkg/bdw"
	"github.com/artemiscloud/activemq-artemis-operator-test-suite/pkg/test_helpers"
	"github.com/artemiscloud/activemq-artemis-operator-test-suite/test"
)

var _ = ginkgo.Describe("AddressSettingMiscTest", func() {

	var (
		ctx1 *framework.ContextData
		//brokerClient brokerclientset.Interface
		brokerDeployer *bdw.BrokerDeploymentWrapper
		//url      string
	)

	var (
		AddressBit  = "someQueue"
		ExpectedURL = "wconsj"
		hw          = test_helpers.NewWrapper()
	)

	ginkgo.BeforeEach(func() {
		if brokerDeployer != nil {
			brokerDeployer.PurgeAddressSettings()
		}
	})

	// PrepareNamespace after framework has been created
	ginkgo.JustBeforeEach(func() {
		ctx1 = sw.Framework.GetFirstContext()
		brokerDeployer = &bdw.BrokerDeploymentWrapper{}
		brokerDeployer.WithWait(true).
			WithBrokerClient(sw.BrokerClient).
			WithContext(ctx1).
			WithCustomImage(test.Config.BrokerImageName).
			WithName(DeployName).
			WithLts(!test.Config.NeedsLatestCR).
			WithConsoleExposure(true)
		brokerDeployer.SetUpDefaultAddressSettings(AddressBit)
	})

	ginkgo.It("DLQPrefix check", func() {
		err := brokerDeployer.WithDlqPrefix(AddressBit, "prefix").DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsString(address, AddressBit, "deadLetterQueuePrefix", "prefix", hw)
	})

	ginkgo.It("DLQSuffix check", func() {
		err := brokerDeployer.WithDlqSuffix(AddressBit, "suffix").DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsString(address, AddressBit, "deadLetterQueueSuffix", "suffix", hw)
	})

	ginkgo.It("DLQAddress check", func() {
		err := brokerDeployer.WithDeadLetterAddress(AddressBit, "DLqQ").DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsString(address, AddressBit, "deadLetterAddress", "DLqQ", hw)
	})

	ginkgo.It("AddressFullPolicy check", func() {
		err := brokerDeployer.WithAddressFullPolicy(AddressBit, bdw.DropPolicy).DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsString(address, AddressBit, "addressFullPolicy", "DROP", hw)
	})

	ginkgo.It("MetricsCheck check", func() {
		err := brokerDeployer.WithEnableMetrics(AddressBit, true).DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsBool(address, AddressBit, "enableMetrics", true, hw)
	})

	/*
		ginkgo.It("MetricsCheck check", func() {
			err := brokerDeployer.WithManagementBrowsePageSize(AddressBit, 101).DeployBrokers(1)
			gomega.Expect(err).To(gomega.BeNil())

			urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
			address := urls[0]
			verifyAddressSettingsBool(address, AddressBit, "enableMetrics",true, hw)
		})
	*/

	ginkgo.It("SlowConsumerCheck check", func() {
		err := brokerDeployer.WithSlowConsumerCheckPeriod(AddressBit, 10).DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsInt(address, AddressBit, "slowConsumerCheckPeriod", 10, hw)
	})

	ginkgo.It("SlowConsumerPolicy check", func() {
		err := brokerDeployer.WithSlowConsumerPolicy(AddressBit, bdw.Notify).DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsString(address, AddressBit, "slowConsumerPolicy", bdw.NOTIFY, hw)
	})

	ginkgo.It("SlowConsumerThreshold check", func() {
		err := brokerDeployer.WithSlowConsumerThreshold(AddressBit, 320).DeployBrokers(1)
		gomega.Expect(err).To(gomega.BeNil())

		urls, err := brokerDeployer.GetExternalUrls(ExpectedURL, 0)
		address := urls[0]
		verifyAddressSettingsInt(address, AddressBit, "slowConsumerThreshold", 320, hw)
	})
})
