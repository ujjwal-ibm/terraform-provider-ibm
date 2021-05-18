// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBareMetalServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBMSRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:         schema.TypeString,
				Required:     true,
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
							Type:             schema.TypeInt,
							Computed:         true,
							DiffSuppressFunc: applyOnce,
							Deprecated:       "This field is deprected",
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
				Computed:    true,
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
	}
}

func dataSourceIBMISBMSValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "identifier",
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isBMSName,
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString})

	ibmISBMSDataSourceValidator := ResourceValidator{ResourceName: "ibm_is_snapshot", Schema: validateSchema}
	return &ibmISBMSDataSourceValidator
}

func dataSourceIBMISBMSRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Get("identifier").(string)
	err := bmsGetById(d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func bmsGetById(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetBareMetalServerOptions{
		ID: &id,
	}
	bms, response, err := sess.GetBareMetalServer(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Bare Metal Server (%s): %s\n%s", id, err, response)
	}
	d.SetId(*bms.ID)
	d.Set(isBMSName, *bms.Name)
	d.Set(isBMSBandwidth, *bms.Bandwidth)
	cpuList := make([]map[string]interface{}, 0)
	if bms.Cpu != nil {
		currentCPU := map[string]interface{}{}
		currentCPU[isBMSCPUArchitecture] = *bms.Cpu.Architecture
		cpuList = append(cpuList, currentCPU)
	}
	d.Set(isBMSCPU, cpuList)
	bmsBootTargetIntf := bms.BootTarget.(*vpcv1.BareMetalServerBootTarget)
	bmsBootTarget := bmsBootTargetIntf.ID
	d.Set(isBMSBootTarget, bmsBootTarget)
	d.Set(isBMSCrn, *bms.CRN)
	d.Set(isBMSEnableSecureBoot, *bms.EnableSecureBoot)
	d.Set(isBMSHref, *bms.Href)
	d.Set(isBMSMemory, *bms.Memory)
	d.Set(isBMSProfile, *bms.Profile.Name)
	d.Set(isBMSResourceGroup, *bms.ResourceGroup.ID)
	d.Set(isBMSStatus, *bms.Status)
	// d.Set(isBMSStatusReasons, *bms
	d.Set(isBMSTrustedPlatformModule, *bms.TrustedPlatformModule.Enabled)
	d.Set(isBMSVPC, *bms.VPC.ID)
	d.Set(isBMSZone, *bms.Zone.Name)
	// d.Set(isBMSStatusReasonsCode, *bms.
	// d.Set(isBMSStatusReasonsMessage, *bms.
	// d.Set(isBMSNicSubnet, *bms.
	d.Set(isBMSProfile, *bms.Profile.Name)
	d.Set(isVolumeZone, *bms.Zone.Name)
	// if bms.EncryptionKey != nil {
	// 	d.Set(isVolumeEncryptionKey, bms.EncryptionKey.CRN)
	// }
	// d.Set(isVolumeIops, *bms.Iops)
	// d.Set(isVolumeCapacity, *bms.Capacity)
	//set the status reasons
	if bms.StatusReasons != nil {
		statusReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range bms.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isBMSStatusReasonsCode] = *sr.Code
				currentSR[isBMSStatusReasonsMessage] = *sr.Message
				statusReasonsList = append(statusReasonsList, currentSR)
			}
		}
		d.Set(isBMSStatusReasons, statusReasonsList)
	}
	tags, err := GetTagsUsingCRN(meta, *bms.CRN)
	if err != nil {
		log.Printf(
			"Error on get of resource bare metal server (%s) tags: %s", d.Id(), err)
	}
	d.Set(isBMSTags, tags)
	if bms.ResourceGroup != nil {
		d.Set(isBMSResourceGroup, *bms.ResourceGroup.ID)
	}
	return nil
}
