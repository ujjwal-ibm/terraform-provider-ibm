---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Server Profiles"
description: |-
  Manages IBM Cloud Bare Metal Server Profiles.
---

# ibm\_is_bare_metal_server_profiles

Import the details of an existing IBM Cloud virtual server instances as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_bare_metal_server_profiles" "ds_bmsprofiles" {
}

```

## Attribute Reference

The following attributes are exported:

* `profiles` - List of all bare metal server profiles in the IBM Cloud Infrastructure.
  * `name` - The name of the profile.
  * `bandwidth` - Bandwidth of the profile.
  * `cpu_architecture` - CPU Architecture of the profile.
  * `cpu_core_count` - CPU core count of the profile.
  * `cpu_socket_count` - CPU socket count of the profile.
  * `family` - The product family this bare metal server profile belongs to.
  * `href` - The URL for this bare metal server profile.
  * `memory` - The memory (in gibibytes) for a bare metal server with this profile.
  * `os_architecture` - The supported OS architecture(s) for a bare metal server with this profile.
  * `resource_type` - The resource type.
  * `supported_image_flags` - The image flags supported by this profile.
  * `supported_trusted_platform_module_modes` - An array of supported trusted platform module (TPM) modes for this bare metal server profile.
  * `disks` - A nested block describing the collection of the bare metal server profile's disks.
  Nested `disk` blocks have the following profile:
    * `quantity` - The number of disks of this configuration for a bare metal server with this profile.
    * `size` - The size of the disk in GB (gigabytes).
    * `supported_interface_types` - The disk interface used for attaching the disk.