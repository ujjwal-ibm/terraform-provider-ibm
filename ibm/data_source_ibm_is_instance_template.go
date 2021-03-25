// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMISInstanceTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceTemplateRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isInstanceTemplatesName, "identifier"},
				ValidateFunc: validateISName,
			},
			isInstanceTemplatesName: {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isInstanceTemplatesName, "identifier"},
				Computed:     true,
			},
			isInstanceTemplatesHref: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplatesCrn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateVPC: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateZone: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateProfile: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateKeys: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			isInstanceTemplateVolumeAttachments: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplatesDeleteVol: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						isInstanceTemplatesName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplatesVol: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			isInstanceTemplatePrimaryNetworkInterface: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateNicName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicSecurityGroups: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceTemplateNicSubnet: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isInstanceTemplateNetworkInterfaces: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateNicName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicSecurityGroups: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceTemplateNicSubnet: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			isInstanceTemplateUserData: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateImage: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplatesBootVolumeAttachment: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplatesDeleteVol: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						isInstanceTemplatesName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplatesVol: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateBootSize: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceTemplateBootProfile: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			isInstanceTemplateResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISInstanceTemplateRead(d *schema.ResourceData, meta interface{}) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}
	name := d.Get(isInstanceTemplatesName).(string)
	id := d.Get("identifier").(string)
	listInstanceTemplatesOptions := &vpcv1.ListInstanceTemplatesOptions{}
	availableTemplates, _, err := instanceC.ListInstanceTemplates(listInstanceTemplatesOptions)
	if err != nil {
		return err
	}

	for _, instTempl := range availableTemplates.Templates {
		instance := instTempl.(*vpcv1.InstanceTemplate)
		if *instance.Name == name || *instance.ID == id {
			d.Set(isInstanceTemplatesHref, instance.Href)
			d.Set(isInstanceTemplatesCrn, instance.CRN)
			d.Set(isInstanceTemplateName, instance.Name)
			d.Set(isInstanceTemplateUserData, instance.UserData)

			if instance.Keys != nil {
				keys := []string{}
				for _, intfc := range instance.Keys {
					instanceKeyIntf := intfc.(*vpcv1.KeyIdentity)
					keys = append(keys, *instanceKeyIntf.ID)
				}
				d.Set(isInstanceTemplateKeys, keys)
			}
			if instance.Profile != nil {
				instanceProfileIntf := instance.Profile
				identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
				d.Set(isInstanceTemplateProfile, identity.Name)
			}
			if instance.PrimaryNetworkInterface != nil {
				interfaceList := make([]map[string]interface{}, 0)
				currentPrimNic := map[string]interface{}{}
				currentPrimNic[isInstanceTemplateNicName] = *instance.PrimaryNetworkInterface.Name
				if instance.PrimaryNetworkInterface.PrimaryIpv4Address != nil {
					currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *instance.PrimaryNetworkInterface.PrimaryIpv4Address
				}
				subInf := instance.PrimaryNetworkInterface.Subnet
				subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
				currentPrimNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID

				if len(instance.PrimaryNetworkInterface.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(instance.PrimaryNetworkInterface.SecurityGroups); i++ {
						secGrpInf := instance.PrimaryNetworkInterface.SecurityGroups[i]
						secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
						secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
					}
					currentPrimNic[isInstanceTemplateNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
				}
				interfaceList = append(interfaceList, currentPrimNic)
				d.Set(isInstanceTemplatePrimaryNetworkInterface, interfaceList)
			}

			if instance.NetworkInterfaces != nil {
				interfacesList := make([]map[string]interface{}, 0)
				for _, intfc := range instance.NetworkInterfaces {
					currentNic := map[string]interface{}{}
					currentNic[isInstanceTemplateNicName] = *intfc.Name
					if intfc.PrimaryIpv4Address != nil {
						currentNic[isInstanceTemplateNicPrimaryIpv4Address] = *intfc.PrimaryIpv4Address
					}
					//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
					subInf := intfc.Subnet
					subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
					currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
					if len(intfc.SecurityGroups) != 0 {
						secgrpList := []string{}
						for i := 0; i < len(intfc.SecurityGroups); i++ {
							secGrpInf := intfc.SecurityGroups[i]
							secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
							secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
						}
						currentNic[isInstanceTemplateNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
					}

					interfacesList = append(interfacesList, currentNic)
				}
				d.Set(isInstanceTemplateNetworkInterfaces, interfacesList)
			}

			if instance.Image != nil {
				imageInf := instance.Image
				imageIdentity := imageInf.(*vpcv1.ImageIdentity)
				d.Set(isInstanceTemplateImage, imageIdentity.ID)
			}

			if instance.VPC != nil {
				vpcInf := instance.VPC
				vpcRef := vpcInf.(*vpcv1.VPCIdentity)
				d.Set(isInstanceTemplateVPC, vpcRef.ID)
			}

			if instance.Zone != nil {
				zoneInf := instance.Zone
				zone := zoneInf.(*vpcv1.ZoneIdentity)
				d.Set(isInstanceTemplateZone, zone.Name)
			}

			interfacesList := make([]map[string]interface{}, 0)
			if instance.VolumeAttachments != nil {
				for _, volume := range instance.VolumeAttachments {
					volumeAttach := map[string]interface{}{}
					volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
					volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
					volumeIntf := volume.Volume
					volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentVolumePrototypeInstanceContext)
					volumeAttach[isInstanceTemplateVolAttVolume] = volumeInst.Name
					interfacesList = append(interfacesList, volumeAttach)
				}
				d.Set(isInstanceTemplateVolumeAttachments, interfacesList)
			}

			if instance.BootVolumeAttachment != nil {
				bootVolList := make([]map[string]interface{}, 0)
				bootVol := map[string]interface{}{}

				bootVol[isInstanceTemplatesDeleteVol] = *instance.BootVolumeAttachment.DeleteVolumeOnInstanceDelete
				if instance.BootVolumeAttachment.Volume != nil {
					volumeIntf := instance.BootVolumeAttachment.Volume
					bootVol[isInstanceTemplatesName] = volumeIntf.Name
					bootVol[isInstanceTemplatesVol] = volumeIntf.Name
					bootVol[isInstanceTemplateBootSize] = volumeIntf.Capacity
					if instance.BootVolumeAttachment.Volume.Profile != nil {
						volProfIntf := instance.BootVolumeAttachment.Volume.Profile
						volProfInst := volProfIntf.(*vpcv1.VolumeProfileIdentity)
						bootVol[isInstanceTemplateBootProfile] = volProfInst.Name
					}
				}
				bootVolList = append(bootVolList, bootVol)
				d.Set(isInstanceTemplatesBootVolumeAttachment, bootVolList)
			}

			if instance.ResourceGroup != nil {
				rg := instance.ResourceGroup
				d.Set(isInstanceTemplateResourceGroup, rg.ID)
			}

		}
	}
	return fmt.Errorf("No instance template found with %s %s", name, id)
}
