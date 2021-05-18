// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISBMSProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISBMSProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", ""),
				),
			},
		},
	})
}

func testAccCheckIBMISBMSProfileDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_bare_metal_server_profile" "test1" {
		  name=""
      }`)
}
