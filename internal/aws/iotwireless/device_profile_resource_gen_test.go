// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTWirelessDeviceProfile_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTWireless::DeviceProfile", "awscc_iotwireless_device_profile", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSIoTWirelessDeviceProfile_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTWireless::DeviceProfile", "awscc_iotwireless_device_profile", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
