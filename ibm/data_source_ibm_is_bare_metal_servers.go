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
	isBareMetalServers = "servers"
)

func dataSourceIBMISBMSs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBMSsRead,

		Schema: map[string]*schema.Schema{

			isBareMetalServers: {
				Type:        schema.TypeList,
				Description: "List of Bare Metal Servers",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: InvokeDataSourceValidator("ibm_is_snapshot", "identifier"),
						},

						isBMSBandwidth: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total bandwidth (in megabits per second)",
						},

						isBMSKeys: {
							Type:             schema.TypeSet,
							Computed:         true,
							Elem:             &schema.Schema{Type: schema.TypeString},
							Set:              schema.HashString,
							DiffSuppressFunc: applyOnce,
							Description:      "SSH key Ids for the bare metal server",
						},

						isBMSImage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "image name",
						},

						isBMSUserData: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User data given for the bare metal server",
						},

						isBMSPrimaryNetworkInterface: {
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Computed:    true,
							Description: "Primary Network interface info",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									isBMSNicAllowIPSpoofing: {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether IP spoofing is allowed on this interface.",
									},
									isBMSNicName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isBMSNicPortSpeed: {
										Type:       schema.TypeInt,
										Computed:   true,
										Deprecated: "This field is deprected",
									},
									isBMSNicIPs: {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									isBMSNicSecurityGroups: {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									isBMSNicSubnet: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						isBMSNetworkInterfaces: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									isBMSNicAllowIPSpoofing: {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether IP spoofing is allowed on this interface.",
									},
									isBMSNicName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isBMSNicIPs: {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									isBMSNicSecurityGroups: {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									isBMSNicSubnet: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						isBMSZone: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Zone name",
						},

						isBMSVPC: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPC the bare metal server is to be a part of",
						},

						isBMSResourceGroup: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource group name",
						},

						isBMSCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CRN value for the Bare metal server",
						},
						isBMSStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bare metal server status",
						},
						isBMSCPU: {
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "CPU info",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBMSCPUArchitecture: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isBMSCPUCoreCount: {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether IP spoofing is allowed on this interface.",
									},
									isBMSCpuSocketCount: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isBMSCpuThreadPerCore: {
										Type:             schema.TypeInt,
										Computed:         true,
										DiffSuppressFunc: applyOnce,
										Deprecated:       "This field is deprected",
									},
								},
							},
						},

						isBMSStatusReasons: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBMSStatusReasonsCode: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A snake case string succinctly identifying the status reason",
									},

									isBMSStatusReasonsMessage: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the status reason",
									},
								},
							},
						},

						isBMSTags: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: InvokeValidator("ibm_is_bare_metal_server", "tag")},
							Set:         resourceIBMVPCHash,
							Description: "Tags for the Bare metal server",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISBMSsRead(d *schema.ResourceData, meta interface{}) error {

	err := volumeProfilesList(d, meta)
	if err != nil {
		return err
	}
	return nil
}

func bmsList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listBareMetalServersOptions := &vpcv1.ListBareMetalServersOptions{}
	availableServers, response, err := sess.ListBareMetalServers(listBareMetalServersOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching Bare Metal Servers %s\n%s", err, response)
	}
	serversInfo := make([]map[string]interface{}, 0)
	for _, bms := range availableServers.BareMetalServers {

		l := map[string]interface{}{
			isBMSName: *bms.Name,
		}

		serversInfo = append(serversInfo, l)
	}
	d.SetId(dataSourceIBMISBMSsID(d))
	d.Set(isBareMetalServers, serversInfo)
	return nil
}

// dataSourceIBMISBMSsID returns a reasonable ID for a Bare Metal Servers list.
func dataSourceIBMISBMSsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
