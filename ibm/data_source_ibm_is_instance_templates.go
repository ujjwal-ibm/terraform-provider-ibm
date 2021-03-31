// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceTemplates                     = "templates"
	isInstanceTemplatesFirst                = "first"
	isInstanceTemplatesHref                 = "href"
	isInstanceTemplatesCrn                  = "crn"
	isInstanceTemplatesLimit                = "limit"
	isInstanceTemplatesNext                 = "next"
	isInstanceTemplatesTotalCount           = "total_count"
	isInstanceTemplatesName                 = "name"
	isInstanceTemplatesPortSpeed            = "port_speed"
	isInstanceTemplatesPortType             = "type"
	isInstanceTemplatesPortValue            = "value"
	isInstanceTemplatesDeleteVol            = "delete_volume_on_instance_delete"
	isInstanceTemplatesVol                  = "volume"
	isInstanceTemplatesMemory               = "memory"
	isInstanceTemplatesMemoryValue          = "value"
	isInstanceTemplatesMemoryType           = "type"
	isInstanceTemplatesMemoryValues         = "values"
	isInstanceTemplatesMemoryDefault        = "default"
	isInstanceTemplatesMemoryMin            = "min"
	isInstanceTemplatesMemoryMax            = "max"
	isInstanceTemplatesMemoryStep           = "step"
	isInstanceTemplatesSocketCount          = "socket_count"
	isInstanceTemplatesSocketValue          = "value"
	isInstanceTemplatesSocketType           = "type"
	isInstanceTemplatesSocketValues         = "values"
	isInstanceTemplatesSocketDefault        = "default"
	isInstanceTemplatesSocketMin            = "min"
	isInstanceTemplatesSocketMax            = "max"
	isInstanceTemplatesSocketStep           = "step"
	isInstanceTemplatesVcpuArch             = "vcpu_architecture"
	isInstanceTemplatesVcpuArchType         = "type"
	isInstanceTemplatesVcpuArchValue        = "value"
	isInstanceTemplatesVcpuCount            = "vcpu_count"
	isInstanceTemplatesVcpuCountValue       = "value"
	isInstanceTemplatesVcpuCountType        = "type"
	isInstanceTemplatesVcpuCountValues      = "values"
	isInstanceTemplatesVcpuCountDefault     = "default"
	isInstanceTemplatesVcpuCountMin         = "min"
	isInstanceTemplatesVcpuCountMax         = "max"
	isInstanceTemplatesVcpuCountStep        = "step"
	isInstanceTemplatesStart                = "start"
	isInstanceTemplatesVersion              = "version"
	isInstanceTemplatesGeneration           = "generation"
	isInstanceTemplatesBootVolumeAttachment = "boot_volume_attachment"

	isInstanceTemplateVPC                     = "vpc"
	isInstanceTemplateZone                    = "zone"
	isInstanceTemplateProfile                 = "profile"
	isInstanceTemplateKeys                    = "keys"
	isInstanceTemplateVolumeAttachments       = "volume_attachments"
	isInstanceTemplateNetworkInterfaces       = "network_interfaces"
	isInstanceTemplatePrimaryNetworkInterface = "primary_network_interface"
	isInstanceTemplateNicName                 = "name"
	isInstanceTemplateNicPortSpeed            = "port_speed"
	isInstanceTemplateNicAllowIPSpoofing      = "allow_ip_spoofing"
	isInstanceTemplateNicPrimaryIpv4Address   = "primary_ipv4_address"
	isInstanceTemplateNicPrimaryIpv6Address   = "primary_ipv6_address"
	isInstanceTemplateNicSecondaryAddress     = "secondary_addresses"
	isInstanceTemplateNicSecurityGroups       = "security_groups"
	isInstanceTemplateNicSubnet               = "subnet"
	isInstanceTemplateNicFloatingIPs          = "floating_ips"
	isInstanceTemplateUserData                = "user_data"
	isInstanceTemplateGeneration              = "generation"
	isInstanceTemplateImage                   = "image"
	isInstanceTemplateResourceGroup           = "resource_group"
	isInstanceTemplateName                    = "name"
	isInstanceTemplateDeleteVolume            = "delete_volume_on_instance_delete"
	isInstanceTemplateVolAttName              = "name"
	isInstanceTemplateVolAttVolume            = "volume"
)

func dataSourceIBMISInstanceTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceTemplatesRead,
		Schema: map[string]*schema.Schema{
			isInstanceTemplates: {
				Type:        schema.TypeList,
				Description: "Collection of instance templates",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplatesName: {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceIBMISInstanceTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listInstanceTemplatesOptions := &vpcv1.ListInstanceTemplatesOptions{}
	availableTemplates, _, err := instanceC.ListInstanceTemplates(listInstanceTemplatesOptions)
	if err != nil {
		return err
	}
	templates := make([]map[string]interface{}, 0)
	for _, instTempl := range availableTemplates.Templates {
		template := map[string]interface{}{}
		instance := instTempl.(*vpcv1.InstanceTemplate)
		template["id"] = instance.ID
		template[isInstanceTemplatesHref] = instance.Href
		template[isInstanceTemplatesCrn] = instance.CRN
		template[isInstanceTemplateName] = instance.Name
		template[isInstanceTemplateUserData] = instance.UserData

		if instance.Keys != nil {
			keys := []string{}
			for _, intfc := range instance.Keys {
				instanceKeyIntf := intfc.(*vpcv1.KeyIdentity)
				keys = append(keys, *instanceKeyIntf.ID)
			}
			template[isInstanceTemplateKeys] = keys
		}
		if instance.Profile != nil {
			instanceProfileIntf := instance.Profile
			identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
			template[isInstanceTemplateProfile] = identity.Name
		}
		if instance.PrimaryNetworkInterface != nil {
			interfaceList := make([]map[string]interface{}, 0)
			currentPrimNic := map[string]interface{}{}
			currentPrimNic[isInstanceTemplateNicName] = *instance.PrimaryNetworkInterface.Name
			if instance.PrimaryNetworkInterface.PrimaryIP != nil {
				ipIntf := instance.PrimaryNetworkInterface.PrimaryIP
				ipAdd := ipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
				currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *ipAdd.Address
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
			template[isInstanceTemplatePrimaryNetworkInterface] = interfaceList
		}

		if instance.NetworkInterfaces != nil {
			interfacesList := make([]map[string]interface{}, 0)
			for _, intfc := range instance.NetworkInterfaces {
				currentNic := map[string]interface{}{}
				currentNic[isInstanceTemplateNicName] = *intfc.Name
				if intfc.PrimaryIP != nil {
					ipIntf := intfc.PrimaryIP
					ipAdd := ipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
					currentNic[isInstanceTemplateNicPrimaryIpv4Address] = *ipAdd.Address
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
			template[isInstanceTemplateNetworkInterfaces] = interfacesList
		}

		if instance.Image != nil {
			imageInf := instance.Image
			imageIdentity := imageInf.(*vpcv1.ImageIdentity)
			template[isInstanceTemplateImage] = imageIdentity.ID
		}

		if instance.VPC != nil {
			vpcInf := instance.VPC
			vpcRef := vpcInf.(*vpcv1.VPCIdentity)
			template[isInstanceTemplateVPC] = vpcRef.ID
		}

		if instance.Zone != nil {
			zoneInf := instance.Zone
			zone := zoneInf.(*vpcv1.ZoneIdentity)
			template[isInstanceTemplateZone] = zone.Name
		}

		interfacesList := make([]map[string]interface{}, 0)
		if instance.VolumeAttachments != nil {
			for _, volume := range instance.VolumeAttachments {
				volumeAttach := map[string]interface{}{}
				volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
				volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
				volumeIntf := volume.Volume
				volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeVolume)
				volumeAttach[isInstanceTemplateVolAttVolume] = volumeInst.Name
				interfacesList = append(interfacesList, volumeAttach)
			}
			template[isInstanceTemplateVolumeAttachments] = interfacesList
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
			template[isInstanceTemplatesBootVolumeAttachment] = bootVolList
		}

		if instance.ResourceGroup != nil {
			rg := instance.ResourceGroup
			template[isInstanceTemplateResourceGroup] = rg.ID
		}

		templates = append(templates, template)
	}
	d.SetId(dataSourceIBMISInstanceTemplatesID(d))
	d.Set(isInstanceTemplates, templates)
	return nil
}

// dataSourceIBMISInstanceTemplatesID returns a reasonable ID for a instance templates list.
func dataSourceIBMISInstanceTemplatesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
