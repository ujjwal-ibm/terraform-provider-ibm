// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBMSBandwidth               = "bandwidth"
	isBMSCPU                     = "cpu"
	isBMSBootTarget              = "boot_target"
	isBMSCPUArchitecture         = "architecture"
	isBMSCPUCoreCount            = "core_count"
	isBMSCpuSocketCount          = "socket_count"
	isBMSCpuThreadPerCore        = "threads_per_core"
	isBMSCrn                     = "crn"
	isBMSEnableSecureBoot        = "enable_secure_boot"
	isBMSHref                    = "href"
	isBMSTags                    = "tags"
	isBMSMemory                  = "memory"
	isBMSName                    = "name"
	isBMSNetworkInterfaces       = "network_interfaces"
	isBMSPrimaryNetworkInterface = "primary_network_interface"
	isBMSProfile                 = "profile"
	isBMSResourceGroup           = "resource_group"
	isBMSStatus                  = "status"
	isBMSStatusReasons           = "status_reasons"
	isBMSTrustedPlatformModule   = "trusted_platform_module"
	isBMSVPC                     = "vpc"
	isBMSZone                    = "zone"
	isBMSStatusReasonsCode       = "code"
	isBMSStatusReasonsMessage    = "message"
	isBMSImage                   = "image"
	isBMSKeys                    = "keys"
	isBMSUserData                = "user_data"
	isBMSNicName                 = "name"
	isBMSNicPortSpeed            = "port_speed"
	isBMSNicAllowIPSpoofing      = "allow_ip_spoofing"
	isBMSNicIPs                  = "ips"
	isBMSNicSecurityGroups       = "security_groups"
	isBMSNicSubnet               = "subnet"
	isBMSActionDeleting          = "deleting"
	isBMSActionDeleted           = "deleted"
	isBMSActionStatusStopping    = "stopping"

	isBMSActionStatusStopped = "stopped"
	isBMSStatusPending       = "pending"
	isBMSStatusRunning       = "running"
	isBMSStatusFailed        = "failed"
	// isBMSResourceGroup        = "trusted_platform_module"
	// isBMSResourceGroup        = "trusted_platform_module"
)

func resourceIBMisBMS() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISBMSCreate,
		Read:     resourceIBMISBMSRead,
		Update:   resourceIBMISBMSUpdate,
		Delete:   resourceIBMISBMSDelete,
		Exists:   resourceIBMISBMSExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{

			isBMSName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_bare_metal_server", isBMSName),
				Description:  "Bare metal server name",
			},

			isBMSBandwidth: {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The total bandwidth (in megabits per second)",
			},

			isBMSKeys: {
				Type:             schema.TypeSet,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
				Description:      "SSH key Ids for the bare metal server",
			},

			isBMSImage: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "image name",
			},

			isBMSUserData: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "User data given for the bare metal server",
			},

			isBMSPrimaryNetworkInterface: {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Primary Network interface info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBMSNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBMSNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isBMSNicPortSpeed: {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
							Deprecated:       "This field is deprected",
						},
						isBMSNicIPs: {
							Type:     schema.TypeSet,
							ForceNew: true,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBMSNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBMSNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			isBMSNetworkInterfaces: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBMSNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBMSNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isBMSNicIPs: {
							Type:     schema.TypeSet,
							ForceNew: true,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBMSNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBMSNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			isBMSZone: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone name",
			},

			isBMSVPC: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The VPC the bare metal server is to be a part of",
			},

			isBMSResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							Optional:    true,
							Default:     false,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBMSCpuSocketCount: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isBMSCpuThreadPerCore: {
							Type:             schema.TypeInt,
							Optional:         true,
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
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: InvokeValidator("ibm_is_bare_metal_server", "tag")},
				Set:         resourceIBMVPCHash,
				Description: "Tags for the Bare metal server",
			},
		},
	}
}

func resourceIBMISBMSValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isBMSName,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "tag",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISBMSResourceValidator := ResourceValidator{ResourceName: "ibm_is_bare_metal_server", Schema: validateSchema}
	return &ibmISBMSResourceValidator
}

func resourceIBMISBMSCreate(d *schema.ResourceData, meta interface{}) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateBareMetalServerOptions{}
	var imageStr string
	if image, ok := d.GetOk(isBMSImage); ok {
		imageStr = image.(string)
	}
	keySet := d.Get(isInstanceKeys).(*schema.Set)
	if keySet.Len() != 0 {
		keyobjs := make([]vpcv1.KeyIdentityIntf, keySet.Len())
		for i, key := range keySet.List() {
			keystr := key.(string)
			keyobjs[i] = &vpcv1.KeyIdentity{
				ID: &keystr,
			}
		}
		options.Initialization = &vpcv1.BareMetalServerInitializationPrototype{
			Image: &vpcv1.ImageIdentity{
				ID: &imageStr,
			},
			Keys: keyobjs,
		}
		if userdata, ok := d.GetOk(isBMSUserData); ok {
			userdatastr := userdata.(string)
			options.Initialization.UserData = &userdatastr
		}
	}

	if name, ok := d.GetOk(isBMSName); ok {
		nameStr := name.(string)
		options.Name = &nameStr
	}

	if rgrp, ok := d.GetOk(isBMSResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	if p, ok := d.GetOk(isBMSProfile); ok {
		profile := p.(string)
		options.Profile = &vpcv1.BareMetalServerProfileIdentity{
			Name: &profile,
		}
	}

	if z, ok := d.GetOk(isBMSZone); ok {
		zone := z.(string)
		options.Zone = &vpcv1.ZoneIdentity{
			Name: &zone,
		}
	}

	if v, ok := d.GetOk(isBMSVPC); ok {
		vpc := v.(string)
		options.VPC = &vpcv1.VPCIdentity{
			ID: &vpc,
		}
	}

	if primnicintf, ok := d.GetOk(isBMSPrimaryNetworkInterface); ok {
		primnic := primnicintf.([]interface{})[0].(map[string]interface{})
		subnetintf, _ := primnic[isBMSNicSubnet]
		subnetintfstr := subnetintf.(string)
		var primnicobj = &vpcv1.BareMetalServerPrimaryNetworkInterfacePrototype{}
		primnicobj.Subnet = &vpcv1.SubnetIdentity{
			ID: &subnetintfstr,
		}
		name, _ := primnic[isBMSNicName]
		namestr := name.(string)
		if namestr != "" {
			primnicobj.Name = &namestr
		}
		if ips, ok := primnic[isBMSNicIPs]; ok {
			ipSet := ips.(*schema.Set)
			if ipSet.Len() != 0 {
				var ipobjs = make([]vpcv1.NetworkInterfaceIPPrototypeIntf, ipSet.Len())
				for i, ipIntf := range ipSet.List() {
					ipIntfstr := ipIntf.(string)
					ipobjs[i] = &vpcv1.NetworkInterfaceIPPrototype{
						Address: &ipIntfstr,
					}
				}
				primnicobj.Ips = ipobjs
			}
		}

		allowIPSpoofing, ok := primnic[isBMSNicAllowIPSpoofing]
		allowIPSpoofingbool := allowIPSpoofing.(bool)
		if ok {
			primnicobj.AllowIPSpoofing = &allowIPSpoofingbool
		}
		secgrpintf, ok := primnic[isBMSNicSecurityGroups]
		if ok {
			secgrpSet := secgrpintf.(*schema.Set)
			if secgrpSet.Len() != 0 {
				var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
				for i, secgrpIntf := range secgrpSet.List() {
					secgrpIntfstr := secgrpIntf.(string)
					secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
						ID: &secgrpIntfstr,
					}
				}
				primnicobj.SecurityGroups = secgrpobjs
			}
		}
		options.PrimaryNetworkInterface = primnicobj
	}

	bms, response, err := sess.CreateBareMetalServer(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create bare metal server err %s\n%s", err, response)
	}
	d.SetId(*bms.ID)
	log.Printf("[INFO] Bare Metal Server : %s", *bms.ID)
	_, err = isWaitForBMSAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVolumeTags); ok || v != "" {
		oldList, newList := d.GetChange(isBMSTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *bms.CRN)
		if err != nil {
			log.Printf(
				"Error on create of resource bare metal server (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISBMSRead(d, meta)
}

func resourceIBMISBMSRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	err := bmsGet(d, meta, id)
	if err != nil {
		return err
	}

	return nil
}

func bmsGet(d *schema.ResourceData, meta interface{}, id string) error {
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

func resourceIBMISBMSUpdate(d *schema.ResourceData, meta interface{}) error {

	id := d.Id()

	err := bmsUpdate(d, meta, id)
	if err != nil {
		return err
	}

	return resourceIBMISVolumeRead(d, meta)
}

func bmsUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	if d.HasChange(isBMSName) {
		nameStr := ""
		if name, ok := d.GetOk(isBMSName); ok {
			nameStr = name.(string)
		}
		options := &vpcv1.UpdateBareMetalServerOptions{
			ID: &id,
		}
		bmsPatchModel := &vpcv1.BareMetalServerPatch{
			Name: &nameStr,
		}
		bmsPatch, err := bmsPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("Error calling asPatch for BareMetalServerPatch: %s", err)
		}
		options.BareMetalServerPatch = bmsPatch
		_, response, err := sess.UpdateBareMetalServer(options)
		if err != nil {
			return fmt.Errorf("Error updating Bare Metal Server: %s\n%s", err, response)
		}
	}

	return nil
}

func resourceIBMISBMSDelete(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()

	err := bmsDelete(d, meta, id)
	if err != nil {
		return err
	}

	return nil
}

func bmsDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getBmsOptions := &vpcv1.GetBareMetalServerOptions{
		ID: &id,
	}
	bms, response, err := sess.GetBareMetalServer(getBmsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error Getting Bare Metal Server (%s): %s\n%s", id, err, response)
	}
	if *bms.Status == "running" {
		options := &vpcv1.CreateBareMetalServerStopOptions{
			ID: bms.ID,
		}
		response, err = sess.CreateBareMetalServerStop(options)
	}

	options := &vpcv1.DeleteBareMetalServerOptions{
		ID: &id,
	}
	response, err = sess.DeleteBareMetalServer(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Bare Metal Server : %s\n%s", err, response)
	}
	_, err = isWaitForBMSDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForBMSDeleted(bmsC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isBMSActionDeleting},
		Target:     []string{"done", ""},
		Refresh:    isBMSDeleteRefreshFunc(bmsC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isBMSDeleteRefreshFunc(bmsC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := bmsC.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return bms, isBMSActionDeleted, nil
			}
			return bms, "", fmt.Errorf("Error Getting Bare Metal Server: %s\n%s", err, response)
		}
		return bms, isBMSActionDeleting, err
	}
}

func resourceIBMISBMSExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	id := d.Id()

	exists, err := bmsExists(d, meta, id)
	return exists, err

}

func bmsExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	options := &vpcv1.GetBareMetalServerOptions{
		ID: &id,
	}
	_, response, err := sess.GetBareMetalServer(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Bare Metal Server: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForBMSAvailable(client *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"restarting", isBMSStatusPending},
		Target:     []string{isBMSStatusRunning, ""},
		Refresh:    isBMSRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isBMSRefreshFunc(client *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := client.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error Getting Bare Metal Server: %s\n%s", err, response)
		}

		if *bms.Status == "running" {
			return bms, "running", nil
		}

		return bms, "pending", nil
	}
}

func isWaitForBMSActionStop(bmsC *vpcv1.VpcV1, timeout time.Duration, id string, d *schema.ResourceData) (interface{}, error) {
	communicator := make(chan interface{})
	stateConf := &resource.StateChangeConf{
		Pending: []string{isBMSStatusRunning, isBMSStatusPending, isBMSActionStatusStopping},
		Target:  []string{isBMSActionStatusStopped, isBMSStatusFailed, ""},
		Refresh: func() (interface{}, string, error) {
			getbmsoptions := &vpcv1.GetBareMetalServerOptions{
				ID: &id,
			}
			bms, response, err := bmsC.GetBareMetalServer(getbmsoptions)
			if err != nil {
				return nil, "", fmt.Errorf("Error Getting Bare Metal Server: %s\n%s", err, response)
			}
			select {
			case data := <-communicator:
				return nil, "", data.(error)
			default:
				fmt.Println("no message sent")
			}
			if *bms.Status == isBMSStatusFailed {
				// let know the isRestartStopAction() to stop
				close(communicator)
				return bms, *bms.Status, fmt.Errorf("The  Bare Metal Server %s failed to stop: %v", id, err)
			}
			return bms, *bms.Status, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	if v, ok := d.GetOk("force_recovery_time"); ok {
		forceTimeout := v.(int)
		go isRestartStopAction(bmsC, id, d, forceTimeout, communicator)
	}

	return stateConf.WaitForState()
}

func isBMSRestartStopAction(bmsC *vpcv1.VpcV1, id string, d *schema.ResourceData, forceTimeout int, communicator chan interface{}) {
	subticker := time.NewTicker(time.Duration(forceTimeout) * time.Minute)
	//subticker := time.NewTicker(time.Duration(forceTimeout) * time.Second)
	for {
		select {

		case <-subticker.C:
			log.Println("Bare Metal Server is still in stopping state, retrying to stop with -force")
			actiontype := "hard"
			createbmssactoptions := &vpcv1.CreateBareMetalServerStopOptions{
				ID:   &id,
				Type: &actiontype,
			}
			response, err := bmsC.CreateBareMetalServerStop(createbmssactoptions)
			if err != nil {
				communicator <- fmt.Errorf("Error retrying Bare Metal Server action stop: %s\n%s", err, response)
				return
			}
		case <-communicator:
			// indicates refresh func is reached target and not proceed with the thread)
			subticker.Stop()
			return

		}
	}
}
