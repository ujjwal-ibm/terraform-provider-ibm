---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare metal server"
description: |-
  Manages IBM bare metal sever.
---

# ibm\_is_bare_metal_server

Provides a Bare Metal Server resource. This allows Bare Metal Server to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a Bare Metal Server:

```hcl
resource "ibm_is_bare_metal_server" "testacc_bms" {
  name = "test"
}

```

## Timeouts

ibm_is_instance provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 30 minutes) Used for creating Instance.
* `update` - (Default 30 minutes) Used for updating Instance or while attaching it with volume attachments or interfaces.
* `delete` - (Default 30 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The instance name.
* `vpc` - (Required, Forces new resource, string) The vpc id. 
* `zone` - (Required, Forces new resource, string) Name of the zone. 
* `profile` - (Required, Forces new resource, string) The profile name. 
* `image` - (Required, string) ID of the image.
* `boot_volume` - (Optional, list) A block describing the boot volume of this instance.  
`boot_volume` block have the following structure:
  * `name` - (Optional, string) The name of the boot volume.
  * `encryption` -(Optional, string) The encryption of the boot volume.
* `keys` - (Required, list) Comma separated IDs of ssh keys.  
* `primary_network_interface` - (Required, list) A nested block describing the primary network interface of this instance. We can have only one primary network interface.
Nested `primary_network_interface` block have the following structure:
  * `name` - (Optional, string) The name of the network interface.
  * `port_speed` - (Deprecated, int) Speed of the network interface.
  * `primary_ipv4_address` - (Optional, Forces new resource, string) The IPV4 address of the interface
  * `subnet` -  (Required, string) ID of the subnet.
  * `security_groups` - (Optional, list) Comma separated IDs of security groups.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `network_interfaces` - (Optional, Forces new resource, list) A nested block describing the additional network interface of this instance.
Nested `network_interfaces` block have the following structure:
  * `name` - (Optional, string) The name of the network interface.
  * `primary_ipv4_address` - (Optional, Forces new resource, string) The IPV4 address of the interface
  * `subnet` -  (Required, string) ID of the subnet.
  * `security_groups` - (Optional, list) Comma separated IDs of security groups.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `volumes` - (Optional, list) Comma separated IDs of volumes. 
* `auto_delete_volume` - (Optional, bool) If set to true, automatically deletes volumes attached to the instance.  
**Note** Setting this argument may bring some inconsistency in volume resources since the volumes will be destroyed along with instances.
* `user_data` - (Optional, string) User data to transfer to the server instance.
* `resource_group` - (Optional, Forces new resource, string) The resource group ID for this instance.
* `tags` - (Optional, array of strings) Tags associated with the instance.
* `force_recovery_time` - (Optional, int) Define timeout (in minutes), to force the is_instance to recover from a perpetual "starting" state, during provisioning; similarly, to force the is_instance to recover from a perpetual "stopping" state, during deprovisioning.  **Note**: the force_recovery_time is used to retry multiple times until timeout.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the instance.
* `memory` - Memory of the instance.
* `status` - Status of the instance.
* `vcpu` - A nested block describing the VCPU configuration of this instance.
Nested `vcpu` blocks have the following structure:
  * `architecture` - The architecture of the instance.
  * `count` - The number of VCPUs assigned to the instance.
* `gpu` - A nested block describing the gpu of this instance.
Nested `gpu` blocks have the following structure:
  * `cores` - The cores of the gpu.
  * `count` - Count of the gpu.
  * `manufacture` - Manufacture of the gpu.
  * `memory` - Memory of the gpu.
  * `model` - Model of the gpu.
* `primary_network_interface` - A nested block describing the primary network interface of this instance.
Nested `primary_network_interface` blocks have the following structure:
  * `id` - The id of the network interface.
  * `name` - The name of the network interface.
  * `subnet` -  ID of the subnet.
  * `security_groups` -  List of security groups.
  * `primary_ipv4_address` - The primary IPv4 address.
  * `allow_ip_spoofing` - Indicates whether IP spoofing is allowed on this interface.
* `network_interfaces` - A nested block describing the additional network interface of this instance.
Nested `network_interfaces` blocks have the following structure:
  * `id` - The id of the network interface.
  * `name` - The name of the network interface.
  * `subnet` -  ID of the subnet.
  * `security_groups` -  List of security groups.
  * `primary_ipv4_address` - The primary IPv4 address.
  * `allow_ip_spoofing` - Indicates whether IP spoofing is allowed on this interface.
* `boot_volume` - A nested block describing the boot volume.
Nested `boot_volume` blocks have the following structure:
  * `name` - The name of the boot volume.
  * `size` -  Capacity of the volume in GB.
  * `iops` -  Input/Output Operations Per Second for the volume.
  * `profile` - The profile of the volume.
  * `encryption` - The encryption of the boot volume.
* `volume_attachments` - A nested block describing the volume attachments.  
Nested `volume_attachments` block have the following structure:
  * `id` - The id of the volume attachment
  * `name` -  The name of the volume attachment
  * `volume_id` - The id of the volume attachment's volume
  * `volume_name` -  The name of the volume attachment's volume
  * `volume_crn` -  The CRN of the volume attachment's volume

## Import

ibm_is_instance can be imported using instanceID, eg

```
$ terraform import ibm_is_instance.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
