// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCloud OrchestAwstemplateAutoscaleTgwNewVpc

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCreate,
		Read:   resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcRead,
		Update: resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcUpdate,
		Delete: resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"availability_zones": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bgp_asn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_asset_container": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_asset_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz_autoscale_admin_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"faz_autoscale_admin_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"faz_custom_private_ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"faz_instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz_integration_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgt_admin_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgt_admin_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgt_instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgt_psk_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgtasg_cool_down": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_desired_capacity_byol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_desired_capacity_payg": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fgtasg_health_check_grace_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_max_size_byol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_max_size_payg": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_min_size_byol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_min_size_payg": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fgtasg_scale_in_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgtasg_scale_out_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fos_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"get_license_grace_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heartbeat_delay_allowance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heartbeat_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heartbeat_loss_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"key_pair_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_hook_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"notification_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_election_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"public_subnet1_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_subnet2_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_tag_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"s3_bucket_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"s3_key_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_recovery_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"terminate_unhealthy_vm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transit_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transit_gateway_support_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_custom_asset_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcRead(d, m)
}

func resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcRead(d, m)
}

func resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc resource from API: %v", err)
	}
	return nil
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcAvailabilityZones(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcBgpAsn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetContainer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazAutoscaleAdminUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazCustomPrivateIpaddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazInstanceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazIntegrationOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtInstanceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtPskSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgCoolDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityPayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgHealthCheckGracePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizeByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizePayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizeByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizePayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleInThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleOutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFosVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcGetLicenseGracePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatDelayAllowance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatLossCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcKeyPairName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcLifecycleHookTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcNotificationEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPrimaryElectionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet1Cidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet2Cidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcResourceTagPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3BucketName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3KeyPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcSyncRecoveryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTerminateUnhealthyVm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewayId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewaySupportOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcUseCustomAssetLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcVpcCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("availability_zones", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcAvailabilityZones(o["availability-zones"], d, "availability_zones")); err != nil {
		if vv, ok := fortiAPIPatch(o["availability-zones"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-AvailabilityZones"); ok {
			if err = d.Set("availability_zones", vv); err != nil {
				return fmt.Errorf("Error reading availability_zones: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading availability_zones: %v", err)
		}
	}

	if err = d.Set("bgp_asn", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcBgpAsn(o["bgp-asn"], d, "bgp_asn")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-asn"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-BgpAsn"); ok {
			if err = d.Set("bgp_asn", vv); err != nil {
				return fmt.Errorf("Error reading bgp_asn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_asn: %v", err)
		}
	}

	if err = d.Set("custom_asset_container", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetContainer(o["custom-asset-container"], d, "custom_asset_container")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-asset-container"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-CustomAssetContainer"); ok {
			if err = d.Set("custom_asset_container", vv); err != nil {
				return fmt.Errorf("Error reading custom_asset_container: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_asset_container: %v", err)
		}
	}

	if err = d.Set("custom_asset_directory", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetDirectory(o["custom-asset-directory"], d, "custom_asset_directory")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-asset-directory"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-CustomAssetDirectory"); ok {
			if err = d.Set("custom_asset_directory", vv); err != nil {
				return fmt.Errorf("Error reading custom_asset_directory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_asset_directory: %v", err)
		}
	}

	if err = d.Set("custom_identifier", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomIdentifier(o["custom-identifier"], d, "custom_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-identifier"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-CustomIdentifier"); ok {
			if err = d.Set("custom_identifier", vv); err != nil {
				return fmt.Errorf("Error reading custom_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_identifier: %v", err)
		}
	}

	if err = d.Set("faz_autoscale_admin_username", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazAutoscaleAdminUsername(o["faz-autoscale-admin-username"], d, "faz_autoscale_admin_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-autoscale-admin-username"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FazAutoscaleAdminUsername"); ok {
			if err = d.Set("faz_autoscale_admin_username", vv); err != nil {
				return fmt.Errorf("Error reading faz_autoscale_admin_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_autoscale_admin_username: %v", err)
		}
	}

	if err = d.Set("faz_custom_private_ipaddress", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazCustomPrivateIpaddress(o["faz-custom-private-ipaddress"], d, "faz_custom_private_ipaddress")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-custom-private-ipaddress"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FazCustomPrivateIpaddress"); ok {
			if err = d.Set("faz_custom_private_ipaddress", vv); err != nil {
				return fmt.Errorf("Error reading faz_custom_private_ipaddress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_custom_private_ipaddress: %v", err)
		}
	}

	if err = d.Set("faz_instance_type", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazInstanceType(o["faz-instance-type"], d, "faz_instance_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-instance-type"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FazInstanceType"); ok {
			if err = d.Set("faz_instance_type", vv); err != nil {
				return fmt.Errorf("Error reading faz_instance_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_instance_type: %v", err)
		}
	}

	if err = d.Set("faz_integration_options", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazIntegrationOptions(o["faz-integration-options"], d, "faz_integration_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-integration-options"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FazIntegrationOptions"); ok {
			if err = d.Set("faz_integration_options", vv); err != nil {
				return fmt.Errorf("Error reading faz_integration_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_integration_options: %v", err)
		}
	}

	if err = d.Set("faz_version", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazVersion(o["faz-version"], d, "faz_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-version"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FazVersion"); ok {
			if err = d.Set("faz_version", vv); err != nil {
				return fmt.Errorf("Error reading faz_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_version: %v", err)
		}
	}

	if err = d.Set("fgt_admin_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminCidr(o["fgt-admin-cidr"], d, "fgt_admin_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-admin-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtAdminCidr"); ok {
			if err = d.Set("fgt_admin_cidr", vv); err != nil {
				return fmt.Errorf("Error reading fgt_admin_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_admin_cidr: %v", err)
		}
	}

	if err = d.Set("fgt_admin_port", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminPort(o["fgt-admin-port"], d, "fgt_admin_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-admin-port"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtAdminPort"); ok {
			if err = d.Set("fgt_admin_port", vv); err != nil {
				return fmt.Errorf("Error reading fgt_admin_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_admin_port: %v", err)
		}
	}

	if err = d.Set("fgt_instance_type", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtInstanceType(o["fgt-instance-type"], d, "fgt_instance_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-instance-type"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtInstanceType"); ok {
			if err = d.Set("fgt_instance_type", vv); err != nil {
				return fmt.Errorf("Error reading fgt_instance_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_instance_type: %v", err)
		}
	}

	if err = d.Set("fgt_psk_secret", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtPskSecret(o["fgt-psk-secret"], d, "fgt_psk_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-psk-secret"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtPskSecret"); ok {
			if err = d.Set("fgt_psk_secret", vv); err != nil {
				return fmt.Errorf("Error reading fgt_psk_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_psk_secret: %v", err)
		}
	}

	if err = d.Set("fgtasg_cool_down", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgCoolDown(o["fgtasg-cool-down"], d, "fgtasg_cool_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-cool-down"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgCoolDown"); ok {
			if err = d.Set("fgtasg_cool_down", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_cool_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_cool_down: %v", err)
		}
	}

	if err = d.Set("fgtasg_desired_capacity_byol", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityByol(o["fgtasg-desired-capacity-byol"], d, "fgtasg_desired_capacity_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-desired-capacity-byol"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgDesiredCapacityByol"); ok {
			if err = d.Set("fgtasg_desired_capacity_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_desired_capacity_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_desired_capacity_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_desired_capacity_payg", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityPayg(o["fgtasg-desired-capacity-payg"], d, "fgtasg_desired_capacity_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-desired-capacity-payg"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgDesiredCapacityPayg"); ok {
			if err = d.Set("fgtasg_desired_capacity_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_desired_capacity_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_desired_capacity_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_health_check_grace_period", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgHealthCheckGracePeriod(o["fgtasg-health-check-grace-period"], d, "fgtasg_health_check_grace_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-health-check-grace-period"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgHealthCheckGracePeriod"); ok {
			if err = d.Set("fgtasg_health_check_grace_period", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_health_check_grace_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_health_check_grace_period: %v", err)
		}
	}

	if err = d.Set("fgtasg_max_size_byol", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizeByol(o["fgtasg-max-size-byol"], d, "fgtasg_max_size_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-max-size-byol"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgMaxSizeByol"); ok {
			if err = d.Set("fgtasg_max_size_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_max_size_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_max_size_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_max_size_payg", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizePayg(o["fgtasg-max-size-payg"], d, "fgtasg_max_size_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-max-size-payg"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgMaxSizePayg"); ok {
			if err = d.Set("fgtasg_max_size_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_max_size_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_max_size_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_min_size_byol", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizeByol(o["fgtasg-min-size-byol"], d, "fgtasg_min_size_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-min-size-byol"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgMinSizeByol"); ok {
			if err = d.Set("fgtasg_min_size_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_min_size_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_min_size_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_min_size_payg", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizePayg(o["fgtasg-min-size-payg"], d, "fgtasg_min_size_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-min-size-payg"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgMinSizePayg"); ok {
			if err = d.Set("fgtasg_min_size_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_min_size_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_min_size_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_scale_in_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleInThreshold(o["fgtasg-scale-in-threshold"], d, "fgtasg_scale_in_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-scale-in-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgScaleInThreshold"); ok {
			if err = d.Set("fgtasg_scale_in_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_scale_in_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_scale_in_threshold: %v", err)
		}
	}

	if err = d.Set("fgtasg_scale_out_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleOutThreshold(o["fgtasg-scale-out-threshold"], d, "fgtasg_scale_out_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-scale-out-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FgtasgScaleOutThreshold"); ok {
			if err = d.Set("fgtasg_scale_out_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_scale_out_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_scale_out_threshold: %v", err)
		}
	}

	if err = d.Set("fos_version", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFosVersion(o["fos-version"], d, "fos_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["fos-version"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-FosVersion"); ok {
			if err = d.Set("fos_version", vv); err != nil {
				return fmt.Errorf("Error reading fos_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fos_version: %v", err)
		}
	}

	if err = d.Set("get_license_grace_period", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcGetLicenseGracePeriod(o["get-license-grace-period"], d, "get_license_grace_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["get-license-grace-period"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-GetLicenseGracePeriod"); ok {
			if err = d.Set("get_license_grace_period", vv); err != nil {
				return fmt.Errorf("Error reading get_license_grace_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading get_license_grace_period: %v", err)
		}
	}

	if err = d.Set("heartbeat_delay_allowance", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatDelayAllowance(o["heartbeat-delay-allowance"], d, "heartbeat_delay_allowance")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-delay-allowance"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-HeartbeatDelayAllowance"); ok {
			if err = d.Set("heartbeat_delay_allowance", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_delay_allowance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_delay_allowance: %v", err)
		}
	}

	if err = d.Set("heartbeat_interval", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatInterval(o["heartbeat-interval"], d, "heartbeat_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-interval"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-HeartbeatInterval"); ok {
			if err = d.Set("heartbeat_interval", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("heartbeat_loss_count", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatLossCount(o["heartbeat-loss-count"], d, "heartbeat_loss_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-loss-count"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-HeartbeatLossCount"); ok {
			if err = d.Set("heartbeat_loss_count", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_loss_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_loss_count: %v", err)
		}
	}

	if err = d.Set("key_pair_name", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcKeyPairName(o["key-pair-name"], d, "key_pair_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-pair-name"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-KeyPairName"); ok {
			if err = d.Set("key_pair_name", vv); err != nil {
				return fmt.Errorf("Error reading key_pair_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_pair_name: %v", err)
		}
	}

	if err = d.Set("lifecycle_hook_timeout", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcLifecycleHookTimeout(o["lifecycle-hook-timeout"], d, "lifecycle_hook_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["lifecycle-hook-timeout"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-LifecycleHookTimeout"); ok {
			if err = d.Set("lifecycle_hook_timeout", vv); err != nil {
				return fmt.Errorf("Error reading lifecycle_hook_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lifecycle_hook_timeout: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("notification_email", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcNotificationEmail(o["notification-email"], d, "notification_email")); err != nil {
		if vv, ok := fortiAPIPatch(o["notification-email"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-NotificationEmail"); ok {
			if err = d.Set("notification_email", vv); err != nil {
				return fmt.Errorf("Error reading notification_email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notification_email: %v", err)
		}
	}

	if err = d.Set("primary_election_timeout", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPrimaryElectionTimeout(o["primary-election-timeout"], d, "primary_election_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-election-timeout"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-PrimaryElectionTimeout"); ok {
			if err = d.Set("primary_election_timeout", vv); err != nil {
				return fmt.Errorf("Error reading primary_election_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_election_timeout: %v", err)
		}
	}

	if err = d.Set("public_subnet1_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet1Cidr(o["public-subnet1-cidr"], d, "public_subnet1_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-subnet1-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-PublicSubnet1Cidr"); ok {
			if err = d.Set("public_subnet1_cidr", vv); err != nil {
				return fmt.Errorf("Error reading public_subnet1_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_subnet1_cidr: %v", err)
		}
	}

	if err = d.Set("public_subnet2_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet2Cidr(o["public-subnet2-cidr"], d, "public_subnet2_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-subnet2-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-PublicSubnet2Cidr"); ok {
			if err = d.Set("public_subnet2_cidr", vv); err != nil {
				return fmt.Errorf("Error reading public_subnet2_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_subnet2_cidr: %v", err)
		}
	}

	if err = d.Set("resource_tag_prefix", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcResourceTagPrefix(o["resource-tag-prefix"], d, "resource_tag_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-tag-prefix"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-ResourceTagPrefix"); ok {
			if err = d.Set("resource_tag_prefix", vv); err != nil {
				return fmt.Errorf("Error reading resource_tag_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_tag_prefix: %v", err)
		}
	}

	if err = d.Set("s3_bucket_name", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3BucketName(o["s3-bucket-name"], d, "s3_bucket_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["s3-bucket-name"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-S3BucketName"); ok {
			if err = d.Set("s3_bucket_name", vv); err != nil {
				return fmt.Errorf("Error reading s3_bucket_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading s3_bucket_name: %v", err)
		}
	}

	if err = d.Set("s3_key_prefix", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3KeyPrefix(o["s3-key-prefix"], d, "s3_key_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["s3-key-prefix"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-S3KeyPrefix"); ok {
			if err = d.Set("s3_key_prefix", vv); err != nil {
				return fmt.Errorf("Error reading s3_key_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading s3_key_prefix: %v", err)
		}
	}

	if err = d.Set("sync_recovery_count", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcSyncRecoveryCount(o["sync-recovery-count"], d, "sync_recovery_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-recovery-count"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-SyncRecoveryCount"); ok {
			if err = d.Set("sync_recovery_count", vv); err != nil {
				return fmt.Errorf("Error reading sync_recovery_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_recovery_count: %v", err)
		}
	}

	if err = d.Set("terminate_unhealthy_vm", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTerminateUnhealthyVm(o["terminate-unhealthy-vm"], d, "terminate_unhealthy_vm")); err != nil {
		if vv, ok := fortiAPIPatch(o["terminate-unhealthy-vm"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-TerminateUnhealthyVm"); ok {
			if err = d.Set("terminate_unhealthy_vm", vv); err != nil {
				return fmt.Errorf("Error reading terminate_unhealthy_vm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading terminate_unhealthy_vm: %v", err)
		}
	}

	if err = d.Set("transit_gateway_id", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewayId(o["transit-gateway-id"], d, "transit_gateway_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["transit-gateway-id"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-TransitGatewayId"); ok {
			if err = d.Set("transit_gateway_id", vv); err != nil {
				return fmt.Errorf("Error reading transit_gateway_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transit_gateway_id: %v", err)
		}
	}

	if err = d.Set("transit_gateway_support_options", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewaySupportOptions(o["transit-gateway-support-options"], d, "transit_gateway_support_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["transit-gateway-support-options"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-TransitGatewaySupportOptions"); ok {
			if err = d.Set("transit_gateway_support_options", vv); err != nil {
				return fmt.Errorf("Error reading transit_gateway_support_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transit_gateway_support_options: %v", err)
		}
	}

	if err = d.Set("use_custom_asset_location", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcUseCustomAssetLocation(o["use-custom-asset-location"], d, "use_custom_asset_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-custom-asset-location"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-UseCustomAssetLocation"); ok {
			if err = d.Set("use_custom_asset_location", vv); err != nil {
				return fmt.Errorf("Error reading use_custom_asset_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_custom_asset_location: %v", err)
		}
	}

	if err = d.Set("vpc_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcVpcCidr(o["vpc-cidr"], d, "vpc_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc-VpcCidr"); ok {
			if err = d.Set("vpc_cidr", vv); err != nil {
				return fmt.Errorf("Error reading vpc_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_cidr: %v", err)
		}
	}

	return nil
}

func flattenObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcAvailabilityZones(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcBgpAsn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetContainer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazAutoscaleAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazAutoscaleAdminUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazCustomPrivateIpaddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazInstanceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazIntegrationOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtInstanceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtPskSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgCoolDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityPayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgHealthCheckGracePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizeByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizePayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizeByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizePayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleInThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleOutThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFosVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcGetLicenseGracePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatDelayAllowance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatLossCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcKeyPairName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcLifecycleHookTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcNotificationEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPrimaryElectionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet1Cidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet2Cidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcResourceTagPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3BucketName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3KeyPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcSyncRecoveryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTerminateUnhealthyVm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewayId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewaySupportOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcUseCustomAssetLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcVpcCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCloudOrchestAwstemplateAutoscaleTgwNewVpc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("availability_zones"); ok || d.HasChange("availability_zones") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcAvailabilityZones(d, v, "availability_zones")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["availability-zones"] = t
		}
	}

	if v, ok := d.GetOk("bgp_asn"); ok || d.HasChange("bgp_asn") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcBgpAsn(d, v, "bgp_asn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-asn"] = t
		}
	}

	if v, ok := d.GetOk("custom_asset_container"); ok || d.HasChange("custom_asset_container") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetContainer(d, v, "custom_asset_container")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-asset-container"] = t
		}
	}

	if v, ok := d.GetOk("custom_asset_directory"); ok || d.HasChange("custom_asset_directory") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomAssetDirectory(d, v, "custom_asset_directory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-asset-directory"] = t
		}
	}

	if v, ok := d.GetOk("custom_identifier"); ok || d.HasChange("custom_identifier") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcCustomIdentifier(d, v, "custom_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-identifier"] = t
		}
	}

	if v, ok := d.GetOk("faz_autoscale_admin_password"); ok || d.HasChange("faz_autoscale_admin_password") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazAutoscaleAdminPassword(d, v, "faz_autoscale_admin_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-autoscale-admin-password"] = t
		}
	}

	if v, ok := d.GetOk("faz_autoscale_admin_username"); ok || d.HasChange("faz_autoscale_admin_username") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazAutoscaleAdminUsername(d, v, "faz_autoscale_admin_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-autoscale-admin-username"] = t
		}
	}

	if v, ok := d.GetOk("faz_custom_private_ipaddress"); ok || d.HasChange("faz_custom_private_ipaddress") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazCustomPrivateIpaddress(d, v, "faz_custom_private_ipaddress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-custom-private-ipaddress"] = t
		}
	}

	if v, ok := d.GetOk("faz_instance_type"); ok || d.HasChange("faz_instance_type") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazInstanceType(d, v, "faz_instance_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-instance-type"] = t
		}
	}

	if v, ok := d.GetOk("faz_integration_options"); ok || d.HasChange("faz_integration_options") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazIntegrationOptions(d, v, "faz_integration_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-integration-options"] = t
		}
	}

	if v, ok := d.GetOk("faz_version"); ok || d.HasChange("faz_version") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFazVersion(d, v, "faz_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-version"] = t
		}
	}

	if v, ok := d.GetOk("fgt_admin_cidr"); ok || d.HasChange("fgt_admin_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminCidr(d, v, "fgt_admin_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-admin-cidr"] = t
		}
	}

	if v, ok := d.GetOk("fgt_admin_port"); ok || d.HasChange("fgt_admin_port") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtAdminPort(d, v, "fgt_admin_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-admin-port"] = t
		}
	}

	if v, ok := d.GetOk("fgt_instance_type"); ok || d.HasChange("fgt_instance_type") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtInstanceType(d, v, "fgt_instance_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-instance-type"] = t
		}
	}

	if v, ok := d.GetOk("fgt_psk_secret"); ok || d.HasChange("fgt_psk_secret") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtPskSecret(d, v, "fgt_psk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-psk-secret"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_cool_down"); ok || d.HasChange("fgtasg_cool_down") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgCoolDown(d, v, "fgtasg_cool_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-cool-down"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_desired_capacity_byol"); ok || d.HasChange("fgtasg_desired_capacity_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityByol(d, v, "fgtasg_desired_capacity_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-desired-capacity-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_desired_capacity_payg"); ok || d.HasChange("fgtasg_desired_capacity_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgDesiredCapacityPayg(d, v, "fgtasg_desired_capacity_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-desired-capacity-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_health_check_grace_period"); ok || d.HasChange("fgtasg_health_check_grace_period") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgHealthCheckGracePeriod(d, v, "fgtasg_health_check_grace_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-health-check-grace-period"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_max_size_byol"); ok || d.HasChange("fgtasg_max_size_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizeByol(d, v, "fgtasg_max_size_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-max-size-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_max_size_payg"); ok || d.HasChange("fgtasg_max_size_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMaxSizePayg(d, v, "fgtasg_max_size_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-max-size-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_min_size_byol"); ok || d.HasChange("fgtasg_min_size_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizeByol(d, v, "fgtasg_min_size_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-min-size-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_min_size_payg"); ok || d.HasChange("fgtasg_min_size_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgMinSizePayg(d, v, "fgtasg_min_size_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-min-size-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_scale_in_threshold"); ok || d.HasChange("fgtasg_scale_in_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleInThreshold(d, v, "fgtasg_scale_in_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-scale-in-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_scale_out_threshold"); ok || d.HasChange("fgtasg_scale_out_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFgtasgScaleOutThreshold(d, v, "fgtasg_scale_out_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-scale-out-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fos_version"); ok || d.HasChange("fos_version") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcFosVersion(d, v, "fos_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fos-version"] = t
		}
	}

	if v, ok := d.GetOk("get_license_grace_period"); ok || d.HasChange("get_license_grace_period") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcGetLicenseGracePeriod(d, v, "get_license_grace_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["get-license-grace-period"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_delay_allowance"); ok || d.HasChange("heartbeat_delay_allowance") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatDelayAllowance(d, v, "heartbeat_delay_allowance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-delay-allowance"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_interval"); ok || d.HasChange("heartbeat_interval") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatInterval(d, v, "heartbeat_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_loss_count"); ok || d.HasChange("heartbeat_loss_count") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcHeartbeatLossCount(d, v, "heartbeat_loss_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-loss-count"] = t
		}
	}

	if v, ok := d.GetOk("key_pair_name"); ok || d.HasChange("key_pair_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcKeyPairName(d, v, "key_pair_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-pair-name"] = t
		}
	}

	if v, ok := d.GetOk("lifecycle_hook_timeout"); ok || d.HasChange("lifecycle_hook_timeout") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcLifecycleHookTimeout(d, v, "lifecycle_hook_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lifecycle-hook-timeout"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notification_email"); ok || d.HasChange("notification_email") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcNotificationEmail(d, v, "notification_email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification-email"] = t
		}
	}

	if v, ok := d.GetOk("primary_election_timeout"); ok || d.HasChange("primary_election_timeout") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPrimaryElectionTimeout(d, v, "primary_election_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-election-timeout"] = t
		}
	}

	if v, ok := d.GetOk("public_subnet1_cidr"); ok || d.HasChange("public_subnet1_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet1Cidr(d, v, "public_subnet1_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-subnet1-cidr"] = t
		}
	}

	if v, ok := d.GetOk("public_subnet2_cidr"); ok || d.HasChange("public_subnet2_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcPublicSubnet2Cidr(d, v, "public_subnet2_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-subnet2-cidr"] = t
		}
	}

	if v, ok := d.GetOk("resource_tag_prefix"); ok || d.HasChange("resource_tag_prefix") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcResourceTagPrefix(d, v, "resource_tag_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-tag-prefix"] = t
		}
	}

	if v, ok := d.GetOk("s3_bucket_name"); ok || d.HasChange("s3_bucket_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3BucketName(d, v, "s3_bucket_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["s3-bucket-name"] = t
		}
	}

	if v, ok := d.GetOk("s3_key_prefix"); ok || d.HasChange("s3_key_prefix") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcS3KeyPrefix(d, v, "s3_key_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["s3-key-prefix"] = t
		}
	}

	if v, ok := d.GetOk("sync_recovery_count"); ok || d.HasChange("sync_recovery_count") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcSyncRecoveryCount(d, v, "sync_recovery_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-recovery-count"] = t
		}
	}

	if v, ok := d.GetOk("terminate_unhealthy_vm"); ok || d.HasChange("terminate_unhealthy_vm") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTerminateUnhealthyVm(d, v, "terminate_unhealthy_vm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["terminate-unhealthy-vm"] = t
		}
	}

	if v, ok := d.GetOk("transit_gateway_id"); ok || d.HasChange("transit_gateway_id") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewayId(d, v, "transit_gateway_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transit-gateway-id"] = t
		}
	}

	if v, ok := d.GetOk("transit_gateway_support_options"); ok || d.HasChange("transit_gateway_support_options") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcTransitGatewaySupportOptions(d, v, "transit_gateway_support_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transit-gateway-support-options"] = t
		}
	}

	if v, ok := d.GetOk("use_custom_asset_location"); ok || d.HasChange("use_custom_asset_location") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcUseCustomAssetLocation(d, v, "use_custom_asset_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-custom-asset-location"] = t
		}
	}

	if v, ok := d.GetOk("vpc_cidr"); ok || d.HasChange("vpc_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleTgwNewVpcVpcCidr(d, v, "vpc_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-cidr"] = t
		}
	}

	return &obj, nil
}
