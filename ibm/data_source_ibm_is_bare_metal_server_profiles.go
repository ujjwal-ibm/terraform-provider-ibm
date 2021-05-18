// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBMSProfiles = "profiles"
)

func dataSourceIBMISBMSProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBMSProfilesRead,

		Schema: map[string]*schema.Schema{

			isVolumeProfiles: {
				Type:        schema.TypeList,
				Description: "List of Volume profile maps",
				Computed:    true,
				Elem: &schema.Resource{
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
				},
			},
		},
	}
}

func dataSourceIBMISBMSProfilesRead(d *schema.ResourceData, meta interface{}) error {

	err := bmsProfilesList(d, meta)
	if err != nil {
		return err
	}
	return nil
}

func bmsProfilesList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listBMSProfilesOptions := &vpcv1.ListBareMetalServerProfilesOptions{}
	availableProfiles, response, err := sess.ListBareMetalServerProfiles(listBMSProfilesOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching Bare Metal Server Profiles %s\n%s", err, response)
	}
	profilesInfo := make([]map[string]interface{}, 0)
	for _, profile := range availableProfiles.Profiles {

		l := map[string]interface{}{
			isBMSProfileName:   *profile.Name,
			isBMSProfileFamily: *profile.Family,
		}
		l[isBMSProfileHref] = *profile.Href
		l[isBMSProfileRT] = *profile.ResourceType
		var siflist []string
		for _, item := range profile.SupportedImageFlags {
			siflist = append(siflist, item)
		}
		l[isBMSProfileSIFs] = siflist
		var stpmmlist []string
		for _, item := range profile.SupportedTrustedPlatformModuleModes.Values {
			stpmmlist = append(stpmmlist, item)
		}
		l[isBMSProfileSTPMMs] = stpmmlist
		profilesInfo = append(profilesInfo, l)
	}
	d.SetId(dataSourceIBMISBMSProfilesID(d))
	d.Set(isBMSProfiles, profilesInfo)
	return nil
}

// dataSourceIBMISBMSProfilesID returns a reasonable ID for a BMS Profile list.
func dataSourceIBMISBMSProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
