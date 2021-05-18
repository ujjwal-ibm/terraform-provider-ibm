// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBMSProfileName            = "name"
	isBMSProfileBandwidth       = "bandwidth"
	isBMSProfileType            = "type"
	isBMSProfileValue           = "value"
	isBMSProfileCPUArchitecture = "cpu_architecture"
	isBMSProfileCPUCoreCount    = "cpu_core_count"
	isBMSProfileCPUSocketCount  = "cpu_socket_count"
	isBMSProfileDisks           = "disks"
	isBMSProfileDiskQuantity    = "quantity"
	isBMSProfileDiskSize        = "size"
	isBMSProfileDiskSITs        = "supported_interface_types"
	isBMSProfileFamily          = "family"
	isBMSProfileHref            = "href"
	isBMSProfileMemory          = "memory"
	isBMSProfileOS              = "os_architecture"
	isBMSProfileValues          = "values"
	isBMSProfileRT              = "resource_type"
	isBMSProfileSIFs            = "supported_image_flags"
	isBMSProfileSTPMMs          = "supported_trusted_platform_module_modes"
)

func dataSourceBareMetalServerProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBMSProfileRead,

		Schema: map[string]*schema.Schema{
			isBMSProfileName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for this bare metal server profile",
			},

			isBMSProfileFamily: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The product family this bare metal server profile belongs to",
			},
			isBMSProfileHref: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The URL for this bare metal server profile",
			},

			isBMSProfileBandwidth: {
				Type:        schema.TypeList,
				Computed:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "The total bandwidth (in megabits per second)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBMSProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBMSProfileValue: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},

			isBMSProfileCPUArchitecture: {
				Type:     schema.TypeList,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBMSProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBMSProfileValue: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},

			isBMSProfileCPUSocketCount: {
				Type:     schema.TypeList,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBMSProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBMSProfileValue: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},

			isBMSProfileDisks: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBMSProfileDiskQuantity: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk quantity",
						},

						isBMSProfileDiskSize: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk size",
						},
						isBMSProfileDiskSITs: {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Disk supported interfaces",
						},
					},
				},
			},

			isBMSProfileOS: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The supported OS architecture(s) for a bare metal server with this profile",
			},

			isBMSProfileSTPMMs: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "An array of supported trusted platform module (TPM) modes for this bare metal server profile",
			},
			isBMSProfileSIFs: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "An array of flags supported by this bare metal server profile",
			},
		},
	}
}

func dataSourceIBMISBMSProfileRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	err := bmsProfileGetByName(d, meta, name)
	if err != nil {
		return err
	}
	return nil
}

func bmsProfileGetByName(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetBareMetalServerProfileOptions{
		Name: &name,
	}
	bmsProfile, response, err := sess.GetBareMetalServerProfile(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Bare Metal Server Profile (%s): %s\n%s", name, err, response)
	}
	// d.SetId(*bmsProfile.ID)
	d.Set(isBMSProfileName, *bmsProfile.Name)
	d.Set(isBMSProfileHref, *bmsProfile.Href)
	d.Set(isBMSProfileRT, *bmsProfile.ResourceType)
	var siflist []string
	for _, item := range bmsProfile.SupportedImageFlags {
		siflist = append(siflist, item)
	}
	d.Set(isBMSProfileSIFs, siflist)
	var stpmmlist []string
	for _, item := range bmsProfile.SupportedTrustedPlatformModuleModes.Values {
		stpmmlist = append(stpmmlist, item)
	}
	d.Set(isBMSProfileSTPMMs, stpmmlist)

	return nil
}
