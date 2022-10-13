// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package greengrassv2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGreengrassV2DeploymentDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GreengrassV2::Deployment", "awscc_greengrassv2_deployment", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSGreengrassV2DeploymentDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GreengrassV2::Deployment", "awscc_greengrassv2_deployment", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
