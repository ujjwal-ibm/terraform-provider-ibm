---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Server Profile"
description: |-
  Manages IBM Cloud Bare Metal Server Profile.
---

# ibm\_is_bare_metal_server_profile

Import the details of an existing IBM Cloud Bare Metal Server profile as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "a7a0626c-f97e-4180-afbe-0331ec62f32a"
  profile = "bc1-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.testacc_subnet.id
  }

  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
}

data "ibm_is_bare_metal_server_profile" "ds_bmsprofile" {
  name        = "${ibm_is_instance.testacc_instance.name}"
  private_key = file("~/.ssh/id_rsa")
  passphrase  = ""
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name for this profile .

## Attribute Reference

The following attributes can be exported:

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