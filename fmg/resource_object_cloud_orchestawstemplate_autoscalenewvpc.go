// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCloud OrchestAwstemplateAutoscaleNewVpc

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCloudOrchestAwstemplateAutoscaleNewVpc() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcCreate,
		Read:   resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcRead,
		Update: resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcUpdate,
		Delete: resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcDelete,

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
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"internal_balancer_dns_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_balancing_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal_target_group_health_check_path": &schema.Schema{
				Type:     schema.TypeString,
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
			"loadbalancing_health_check_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"loadbalancing_traffic_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"loadbalancing_traffic_protocol": &schema.Schema{
				Type:     schema.TypeString,
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
			"private_subnet1_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_subnet2_cidr": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCloudOrchestAwstemplateAutoscaleNewVpc(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCloudOrchestAwstemplateAutoscaleNewVpc(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcRead(d, m)
}

func resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestAwstemplateAutoscaleNewVpc(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCloudOrchestAwstemplateAutoscaleNewVpc(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcRead(d, m)
}

func resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCloudOrchestAwstemplateAutoscaleNewVpc(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCloudOrchestAwstemplateAutoscaleNewVpcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCloudOrchestAwstemplateAutoscaleNewVpc(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCloudOrchestAwstemplateAutoscaleNewVpc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwstemplateAutoscaleNewVpc resource from API: %v", err)
	}
	return nil
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcAvailabilityZones(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetContainer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazCustomPrivateIpaddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazInstanceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazIntegrationOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtInstanceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtPskSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgCoolDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityPayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgHealthCheckGracePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizeByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizePayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizeByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizePayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleInThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleOutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFosVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcGetLicenseGracePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatDelayAllowance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatLossCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancerDnsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancingOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalTargetGroupHealthCheckPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcKeyPairName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLifecycleHookTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingHealthCheckThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcNotificationEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPrimaryElectionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet1Cidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet2Cidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet1Cidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet2Cidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcResourceTagPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcS3BucketName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcS3KeyPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcSyncRecoveryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcTerminateUnhealthyVm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcUseCustomAssetLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcVpcCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCloudOrchestAwstemplateAutoscaleNewVpc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("availability_zones", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcAvailabilityZones(o["availability-zones"], d, "availability_zones")); err != nil {
		if vv, ok := fortiAPIPatch(o["availability-zones"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-AvailabilityZones"); ok {
			if err = d.Set("availability_zones", vv); err != nil {
				return fmt.Errorf("Error reading availability_zones: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading availability_zones: %v", err)
		}
	}

	if err = d.Set("custom_asset_container", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetContainer(o["custom-asset-container"], d, "custom_asset_container")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-asset-container"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-CustomAssetContainer"); ok {
			if err = d.Set("custom_asset_container", vv); err != nil {
				return fmt.Errorf("Error reading custom_asset_container: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_asset_container: %v", err)
		}
	}

	if err = d.Set("custom_asset_directory", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetDirectory(o["custom-asset-directory"], d, "custom_asset_directory")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-asset-directory"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-CustomAssetDirectory"); ok {
			if err = d.Set("custom_asset_directory", vv); err != nil {
				return fmt.Errorf("Error reading custom_asset_directory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_asset_directory: %v", err)
		}
	}

	if err = d.Set("custom_identifier", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomIdentifier(o["custom-identifier"], d, "custom_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-identifier"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-CustomIdentifier"); ok {
			if err = d.Set("custom_identifier", vv); err != nil {
				return fmt.Errorf("Error reading custom_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_identifier: %v", err)
		}
	}

	if err = d.Set("faz_autoscale_admin_password", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminPassword(o["faz-autoscale-admin-password"], d, "faz_autoscale_admin_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-autoscale-admin-password"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FazAutoscaleAdminPassword"); ok {
			if err = d.Set("faz_autoscale_admin_password", vv); err != nil {
				return fmt.Errorf("Error reading faz_autoscale_admin_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_autoscale_admin_password: %v", err)
		}
	}

	if err = d.Set("faz_autoscale_admin_username", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminUsername(o["faz-autoscale-admin-username"], d, "faz_autoscale_admin_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-autoscale-admin-username"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FazAutoscaleAdminUsername"); ok {
			if err = d.Set("faz_autoscale_admin_username", vv); err != nil {
				return fmt.Errorf("Error reading faz_autoscale_admin_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_autoscale_admin_username: %v", err)
		}
	}

	if err = d.Set("faz_custom_private_ipaddress", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazCustomPrivateIpaddress(o["faz-custom-private-ipaddress"], d, "faz_custom_private_ipaddress")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-custom-private-ipaddress"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FazCustomPrivateIpaddress"); ok {
			if err = d.Set("faz_custom_private_ipaddress", vv); err != nil {
				return fmt.Errorf("Error reading faz_custom_private_ipaddress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_custom_private_ipaddress: %v", err)
		}
	}

	if err = d.Set("faz_instance_type", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazInstanceType(o["faz-instance-type"], d, "faz_instance_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-instance-type"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FazInstanceType"); ok {
			if err = d.Set("faz_instance_type", vv); err != nil {
				return fmt.Errorf("Error reading faz_instance_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_instance_type: %v", err)
		}
	}

	if err = d.Set("faz_integration_options", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazIntegrationOptions(o["faz-integration-options"], d, "faz_integration_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-integration-options"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FazIntegrationOptions"); ok {
			if err = d.Set("faz_integration_options", vv); err != nil {
				return fmt.Errorf("Error reading faz_integration_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_integration_options: %v", err)
		}
	}

	if err = d.Set("faz_version", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFazVersion(o["faz-version"], d, "faz_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-version"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FazVersion"); ok {
			if err = d.Set("faz_version", vv); err != nil {
				return fmt.Errorf("Error reading faz_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_version: %v", err)
		}
	}

	if err = d.Set("fgt_admin_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminCidr(o["fgt-admin-cidr"], d, "fgt_admin_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-admin-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtAdminCidr"); ok {
			if err = d.Set("fgt_admin_cidr", vv); err != nil {
				return fmt.Errorf("Error reading fgt_admin_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_admin_cidr: %v", err)
		}
	}

	if err = d.Set("fgt_admin_port", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminPort(o["fgt-admin-port"], d, "fgt_admin_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-admin-port"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtAdminPort"); ok {
			if err = d.Set("fgt_admin_port", vv); err != nil {
				return fmt.Errorf("Error reading fgt_admin_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_admin_port: %v", err)
		}
	}

	if err = d.Set("fgt_instance_type", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtInstanceType(o["fgt-instance-type"], d, "fgt_instance_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-instance-type"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtInstanceType"); ok {
			if err = d.Set("fgt_instance_type", vv); err != nil {
				return fmt.Errorf("Error reading fgt_instance_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_instance_type: %v", err)
		}
	}

	if err = d.Set("fgt_psk_secret", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtPskSecret(o["fgt-psk-secret"], d, "fgt_psk_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-psk-secret"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtPskSecret"); ok {
			if err = d.Set("fgt_psk_secret", vv); err != nil {
				return fmt.Errorf("Error reading fgt_psk_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_psk_secret: %v", err)
		}
	}

	if err = d.Set("fgtasg_cool_down", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgCoolDown(o["fgtasg-cool-down"], d, "fgtasg_cool_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-cool-down"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgCoolDown"); ok {
			if err = d.Set("fgtasg_cool_down", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_cool_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_cool_down: %v", err)
		}
	}

	if err = d.Set("fgtasg_desired_capacity_byol", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityByol(o["fgtasg-desired-capacity-byol"], d, "fgtasg_desired_capacity_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-desired-capacity-byol"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgDesiredCapacityByol"); ok {
			if err = d.Set("fgtasg_desired_capacity_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_desired_capacity_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_desired_capacity_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_desired_capacity_payg", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityPayg(o["fgtasg-desired-capacity-payg"], d, "fgtasg_desired_capacity_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-desired-capacity-payg"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgDesiredCapacityPayg"); ok {
			if err = d.Set("fgtasg_desired_capacity_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_desired_capacity_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_desired_capacity_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_health_check_grace_period", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgHealthCheckGracePeriod(o["fgtasg-health-check-grace-period"], d, "fgtasg_health_check_grace_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-health-check-grace-period"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgHealthCheckGracePeriod"); ok {
			if err = d.Set("fgtasg_health_check_grace_period", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_health_check_grace_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_health_check_grace_period: %v", err)
		}
	}

	if err = d.Set("fgtasg_max_size_byol", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizeByol(o["fgtasg-max-size-byol"], d, "fgtasg_max_size_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-max-size-byol"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgMaxSizeByol"); ok {
			if err = d.Set("fgtasg_max_size_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_max_size_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_max_size_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_max_size_payg", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizePayg(o["fgtasg-max-size-payg"], d, "fgtasg_max_size_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-max-size-payg"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgMaxSizePayg"); ok {
			if err = d.Set("fgtasg_max_size_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_max_size_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_max_size_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_min_size_byol", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizeByol(o["fgtasg-min-size-byol"], d, "fgtasg_min_size_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-min-size-byol"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgMinSizeByol"); ok {
			if err = d.Set("fgtasg_min_size_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_min_size_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_min_size_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_min_size_payg", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizePayg(o["fgtasg-min-size-payg"], d, "fgtasg_min_size_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-min-size-payg"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgMinSizePayg"); ok {
			if err = d.Set("fgtasg_min_size_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_min_size_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_min_size_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_scale_in_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleInThreshold(o["fgtasg-scale-in-threshold"], d, "fgtasg_scale_in_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-scale-in-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgScaleInThreshold"); ok {
			if err = d.Set("fgtasg_scale_in_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_scale_in_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_scale_in_threshold: %v", err)
		}
	}

	if err = d.Set("fgtasg_scale_out_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleOutThreshold(o["fgtasg-scale-out-threshold"], d, "fgtasg_scale_out_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-scale-out-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FgtasgScaleOutThreshold"); ok {
			if err = d.Set("fgtasg_scale_out_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_scale_out_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_scale_out_threshold: %v", err)
		}
	}

	if err = d.Set("fos_version", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFosVersion(o["fos-version"], d, "fos_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["fos-version"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-FosVersion"); ok {
			if err = d.Set("fos_version", vv); err != nil {
				return fmt.Errorf("Error reading fos_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fos_version: %v", err)
		}
	}

	if err = d.Set("get_license_grace_period", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcGetLicenseGracePeriod(o["get-license-grace-period"], d, "get_license_grace_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["get-license-grace-period"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-GetLicenseGracePeriod"); ok {
			if err = d.Set("get_license_grace_period", vv); err != nil {
				return fmt.Errorf("Error reading get_license_grace_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading get_license_grace_period: %v", err)
		}
	}

	if err = d.Set("heartbeat_delay_allowance", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatDelayAllowance(o["heartbeat-delay-allowance"], d, "heartbeat_delay_allowance")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-delay-allowance"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-HeartbeatDelayAllowance"); ok {
			if err = d.Set("heartbeat_delay_allowance", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_delay_allowance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_delay_allowance: %v", err)
		}
	}

	if err = d.Set("heartbeat_interval", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatInterval(o["heartbeat-interval"], d, "heartbeat_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-interval"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-HeartbeatInterval"); ok {
			if err = d.Set("heartbeat_interval", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("heartbeat_loss_count", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatLossCount(o["heartbeat-loss-count"], d, "heartbeat_loss_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-loss-count"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-HeartbeatLossCount"); ok {
			if err = d.Set("heartbeat_loss_count", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_loss_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_loss_count: %v", err)
		}
	}

	if err = d.Set("internal_balancer_dns_name", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancerDnsName(o["internal-balancer-dns-name"], d, "internal_balancer_dns_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-balancer-dns-name"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-InternalBalancerDnsName"); ok {
			if err = d.Set("internal_balancer_dns_name", vv); err != nil {
				return fmt.Errorf("Error reading internal_balancer_dns_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_balancer_dns_name: %v", err)
		}
	}

	if err = d.Set("internal_balancing_options", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancingOptions(o["internal-balancing-options"], d, "internal_balancing_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-balancing-options"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-InternalBalancingOptions"); ok {
			if err = d.Set("internal_balancing_options", vv); err != nil {
				return fmt.Errorf("Error reading internal_balancing_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_balancing_options: %v", err)
		}
	}

	if err = d.Set("internal_target_group_health_check_path", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalTargetGroupHealthCheckPath(o["internal-target-group-health-check-path"], d, "internal_target_group_health_check_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-target-group-health-check-path"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-InternalTargetGroupHealthCheckPath"); ok {
			if err = d.Set("internal_target_group_health_check_path", vv); err != nil {
				return fmt.Errorf("Error reading internal_target_group_health_check_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_target_group_health_check_path: %v", err)
		}
	}

	if err = d.Set("key_pair_name", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcKeyPairName(o["key-pair-name"], d, "key_pair_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-pair-name"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-KeyPairName"); ok {
			if err = d.Set("key_pair_name", vv); err != nil {
				return fmt.Errorf("Error reading key_pair_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_pair_name: %v", err)
		}
	}

	if err = d.Set("lifecycle_hook_timeout", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLifecycleHookTimeout(o["lifecycle-hook-timeout"], d, "lifecycle_hook_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["lifecycle-hook-timeout"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-LifecycleHookTimeout"); ok {
			if err = d.Set("lifecycle_hook_timeout", vv); err != nil {
				return fmt.Errorf("Error reading lifecycle_hook_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lifecycle_hook_timeout: %v", err)
		}
	}

	if err = d.Set("loadbalancing_health_check_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingHealthCheckThreshold(o["loadbalancing-health-check-threshold"], d, "loadbalancing_health_check_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["loadbalancing-health-check-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-LoadbalancingHealthCheckThreshold"); ok {
			if err = d.Set("loadbalancing_health_check_threshold", vv); err != nil {
				return fmt.Errorf("Error reading loadbalancing_health_check_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loadbalancing_health_check_threshold: %v", err)
		}
	}

	if err = d.Set("loadbalancing_traffic_port", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficPort(o["loadbalancing-traffic-port"], d, "loadbalancing_traffic_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["loadbalancing-traffic-port"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-LoadbalancingTrafficPort"); ok {
			if err = d.Set("loadbalancing_traffic_port", vv); err != nil {
				return fmt.Errorf("Error reading loadbalancing_traffic_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loadbalancing_traffic_port: %v", err)
		}
	}

	if err = d.Set("loadbalancing_traffic_protocol", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficProtocol(o["loadbalancing-traffic-protocol"], d, "loadbalancing_traffic_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["loadbalancing-traffic-protocol"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-LoadbalancingTrafficProtocol"); ok {
			if err = d.Set("loadbalancing_traffic_protocol", vv); err != nil {
				return fmt.Errorf("Error reading loadbalancing_traffic_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loadbalancing_traffic_protocol: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("notification_email", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcNotificationEmail(o["notification-email"], d, "notification_email")); err != nil {
		if vv, ok := fortiAPIPatch(o["notification-email"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-NotificationEmail"); ok {
			if err = d.Set("notification_email", vv); err != nil {
				return fmt.Errorf("Error reading notification_email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notification_email: %v", err)
		}
	}

	if err = d.Set("primary_election_timeout", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPrimaryElectionTimeout(o["primary-election-timeout"], d, "primary_election_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-election-timeout"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-PrimaryElectionTimeout"); ok {
			if err = d.Set("primary_election_timeout", vv); err != nil {
				return fmt.Errorf("Error reading primary_election_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_election_timeout: %v", err)
		}
	}

	if err = d.Set("private_subnet1_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet1Cidr(o["private-subnet1-cidr"], d, "private_subnet1_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-subnet1-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-PrivateSubnet1Cidr"); ok {
			if err = d.Set("private_subnet1_cidr", vv); err != nil {
				return fmt.Errorf("Error reading private_subnet1_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_subnet1_cidr: %v", err)
		}
	}

	if err = d.Set("private_subnet2_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet2Cidr(o["private-subnet2-cidr"], d, "private_subnet2_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-subnet2-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-PrivateSubnet2Cidr"); ok {
			if err = d.Set("private_subnet2_cidr", vv); err != nil {
				return fmt.Errorf("Error reading private_subnet2_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_subnet2_cidr: %v", err)
		}
	}

	if err = d.Set("public_subnet1_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet1Cidr(o["public-subnet1-cidr"], d, "public_subnet1_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-subnet1-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-PublicSubnet1Cidr"); ok {
			if err = d.Set("public_subnet1_cidr", vv); err != nil {
				return fmt.Errorf("Error reading public_subnet1_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_subnet1_cidr: %v", err)
		}
	}

	if err = d.Set("public_subnet2_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet2Cidr(o["public-subnet2-cidr"], d, "public_subnet2_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-subnet2-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-PublicSubnet2Cidr"); ok {
			if err = d.Set("public_subnet2_cidr", vv); err != nil {
				return fmt.Errorf("Error reading public_subnet2_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_subnet2_cidr: %v", err)
		}
	}

	if err = d.Set("resource_tag_prefix", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcResourceTagPrefix(o["resource-tag-prefix"], d, "resource_tag_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-tag-prefix"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-ResourceTagPrefix"); ok {
			if err = d.Set("resource_tag_prefix", vv); err != nil {
				return fmt.Errorf("Error reading resource_tag_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_tag_prefix: %v", err)
		}
	}

	if err = d.Set("s3_bucket_name", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcS3BucketName(o["s3-bucket-name"], d, "s3_bucket_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["s3-bucket-name"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-S3BucketName"); ok {
			if err = d.Set("s3_bucket_name", vv); err != nil {
				return fmt.Errorf("Error reading s3_bucket_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading s3_bucket_name: %v", err)
		}
	}

	if err = d.Set("s3_key_prefix", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcS3KeyPrefix(o["s3-key-prefix"], d, "s3_key_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["s3-key-prefix"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-S3KeyPrefix"); ok {
			if err = d.Set("s3_key_prefix", vv); err != nil {
				return fmt.Errorf("Error reading s3_key_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading s3_key_prefix: %v", err)
		}
	}

	if err = d.Set("sync_recovery_count", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcSyncRecoveryCount(o["sync-recovery-count"], d, "sync_recovery_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-recovery-count"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-SyncRecoveryCount"); ok {
			if err = d.Set("sync_recovery_count", vv); err != nil {
				return fmt.Errorf("Error reading sync_recovery_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_recovery_count: %v", err)
		}
	}

	if err = d.Set("terminate_unhealthy_vm", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcTerminateUnhealthyVm(o["terminate-unhealthy-vm"], d, "terminate_unhealthy_vm")); err != nil {
		if vv, ok := fortiAPIPatch(o["terminate-unhealthy-vm"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-TerminateUnhealthyVm"); ok {
			if err = d.Set("terminate_unhealthy_vm", vv); err != nil {
				return fmt.Errorf("Error reading terminate_unhealthy_vm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading terminate_unhealthy_vm: %v", err)
		}
	}

	if err = d.Set("use_custom_asset_location", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcUseCustomAssetLocation(o["use-custom-asset-location"], d, "use_custom_asset_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-custom-asset-location"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-UseCustomAssetLocation"); ok {
			if err = d.Set("use_custom_asset_location", vv); err != nil {
				return fmt.Errorf("Error reading use_custom_asset_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_custom_asset_location: %v", err)
		}
	}

	if err = d.Set("vpc_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcVpcCidr(o["vpc-cidr"], d, "vpc_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleNewVpc-VpcCidr"); ok {
			if err = d.Set("vpc_cidr", vv); err != nil {
				return fmt.Errorf("Error reading vpc_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_cidr: %v", err)
		}
	}

	return nil
}

func flattenObjectCloudOrchestAwstemplateAutoscaleNewVpcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcAvailabilityZones(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetContainer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazCustomPrivateIpaddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazInstanceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazIntegrationOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtInstanceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtPskSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgCoolDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityPayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgHealthCheckGracePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizeByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizePayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizeByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizePayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleInThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleOutThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFosVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcGetLicenseGracePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatDelayAllowance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatLossCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancerDnsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancingOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalTargetGroupHealthCheckPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcKeyPairName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLifecycleHookTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingHealthCheckThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcNotificationEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPrimaryElectionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet1Cidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet2Cidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet1Cidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet2Cidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcResourceTagPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcS3BucketName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcS3KeyPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcSyncRecoveryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcTerminateUnhealthyVm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcUseCustomAssetLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleNewVpcVpcCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCloudOrchestAwstemplateAutoscaleNewVpc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("availability_zones"); ok || d.HasChange("availability_zones") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcAvailabilityZones(d, v, "availability_zones")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["availability-zones"] = t
		}
	}

	if v, ok := d.GetOk("custom_asset_container"); ok || d.HasChange("custom_asset_container") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetContainer(d, v, "custom_asset_container")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-asset-container"] = t
		}
	}

	if v, ok := d.GetOk("custom_asset_directory"); ok || d.HasChange("custom_asset_directory") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomAssetDirectory(d, v, "custom_asset_directory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-asset-directory"] = t
		}
	}

	if v, ok := d.GetOk("custom_identifier"); ok || d.HasChange("custom_identifier") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcCustomIdentifier(d, v, "custom_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-identifier"] = t
		}
	}

	if v, ok := d.GetOk("faz_autoscale_admin_password"); ok || d.HasChange("faz_autoscale_admin_password") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminPassword(d, v, "faz_autoscale_admin_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-autoscale-admin-password"] = t
		}
	}

	if v, ok := d.GetOk("faz_autoscale_admin_username"); ok || d.HasChange("faz_autoscale_admin_username") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazAutoscaleAdminUsername(d, v, "faz_autoscale_admin_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-autoscale-admin-username"] = t
		}
	}

	if v, ok := d.GetOk("faz_custom_private_ipaddress"); ok || d.HasChange("faz_custom_private_ipaddress") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazCustomPrivateIpaddress(d, v, "faz_custom_private_ipaddress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-custom-private-ipaddress"] = t
		}
	}

	if v, ok := d.GetOk("faz_instance_type"); ok || d.HasChange("faz_instance_type") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazInstanceType(d, v, "faz_instance_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-instance-type"] = t
		}
	}

	if v, ok := d.GetOk("faz_integration_options"); ok || d.HasChange("faz_integration_options") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazIntegrationOptions(d, v, "faz_integration_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-integration-options"] = t
		}
	}

	if v, ok := d.GetOk("faz_version"); ok || d.HasChange("faz_version") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFazVersion(d, v, "faz_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-version"] = t
		}
	}

	if v, ok := d.GetOk("fgt_admin_cidr"); ok || d.HasChange("fgt_admin_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminCidr(d, v, "fgt_admin_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-admin-cidr"] = t
		}
	}

	if v, ok := d.GetOk("fgt_admin_port"); ok || d.HasChange("fgt_admin_port") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtAdminPort(d, v, "fgt_admin_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-admin-port"] = t
		}
	}

	if v, ok := d.GetOk("fgt_instance_type"); ok || d.HasChange("fgt_instance_type") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtInstanceType(d, v, "fgt_instance_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-instance-type"] = t
		}
	}

	if v, ok := d.GetOk("fgt_psk_secret"); ok || d.HasChange("fgt_psk_secret") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtPskSecret(d, v, "fgt_psk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-psk-secret"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_cool_down"); ok || d.HasChange("fgtasg_cool_down") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgCoolDown(d, v, "fgtasg_cool_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-cool-down"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_desired_capacity_byol"); ok || d.HasChange("fgtasg_desired_capacity_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityByol(d, v, "fgtasg_desired_capacity_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-desired-capacity-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_desired_capacity_payg"); ok || d.HasChange("fgtasg_desired_capacity_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgDesiredCapacityPayg(d, v, "fgtasg_desired_capacity_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-desired-capacity-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_health_check_grace_period"); ok || d.HasChange("fgtasg_health_check_grace_period") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgHealthCheckGracePeriod(d, v, "fgtasg_health_check_grace_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-health-check-grace-period"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_max_size_byol"); ok || d.HasChange("fgtasg_max_size_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizeByol(d, v, "fgtasg_max_size_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-max-size-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_max_size_payg"); ok || d.HasChange("fgtasg_max_size_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMaxSizePayg(d, v, "fgtasg_max_size_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-max-size-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_min_size_byol"); ok || d.HasChange("fgtasg_min_size_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizeByol(d, v, "fgtasg_min_size_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-min-size-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_min_size_payg"); ok || d.HasChange("fgtasg_min_size_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgMinSizePayg(d, v, "fgtasg_min_size_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-min-size-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_scale_in_threshold"); ok || d.HasChange("fgtasg_scale_in_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleInThreshold(d, v, "fgtasg_scale_in_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-scale-in-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_scale_out_threshold"); ok || d.HasChange("fgtasg_scale_out_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFgtasgScaleOutThreshold(d, v, "fgtasg_scale_out_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-scale-out-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fos_version"); ok || d.HasChange("fos_version") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcFosVersion(d, v, "fos_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fos-version"] = t
		}
	}

	if v, ok := d.GetOk("get_license_grace_period"); ok || d.HasChange("get_license_grace_period") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcGetLicenseGracePeriod(d, v, "get_license_grace_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["get-license-grace-period"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_delay_allowance"); ok || d.HasChange("heartbeat_delay_allowance") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatDelayAllowance(d, v, "heartbeat_delay_allowance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-delay-allowance"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_interval"); ok || d.HasChange("heartbeat_interval") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatInterval(d, v, "heartbeat_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_loss_count"); ok || d.HasChange("heartbeat_loss_count") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcHeartbeatLossCount(d, v, "heartbeat_loss_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-loss-count"] = t
		}
	}

	if v, ok := d.GetOk("internal_balancer_dns_name"); ok || d.HasChange("internal_balancer_dns_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancerDnsName(d, v, "internal_balancer_dns_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-balancer-dns-name"] = t
		}
	}

	if v, ok := d.GetOk("internal_balancing_options"); ok || d.HasChange("internal_balancing_options") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalBalancingOptions(d, v, "internal_balancing_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-balancing-options"] = t
		}
	}

	if v, ok := d.GetOk("internal_target_group_health_check_path"); ok || d.HasChange("internal_target_group_health_check_path") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcInternalTargetGroupHealthCheckPath(d, v, "internal_target_group_health_check_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-target-group-health-check-path"] = t
		}
	}

	if v, ok := d.GetOk("key_pair_name"); ok || d.HasChange("key_pair_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcKeyPairName(d, v, "key_pair_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-pair-name"] = t
		}
	}

	if v, ok := d.GetOk("lifecycle_hook_timeout"); ok || d.HasChange("lifecycle_hook_timeout") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLifecycleHookTimeout(d, v, "lifecycle_hook_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lifecycle-hook-timeout"] = t
		}
	}

	if v, ok := d.GetOk("loadbalancing_health_check_threshold"); ok || d.HasChange("loadbalancing_health_check_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingHealthCheckThreshold(d, v, "loadbalancing_health_check_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loadbalancing-health-check-threshold"] = t
		}
	}

	if v, ok := d.GetOk("loadbalancing_traffic_port"); ok || d.HasChange("loadbalancing_traffic_port") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficPort(d, v, "loadbalancing_traffic_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loadbalancing-traffic-port"] = t
		}
	}

	if v, ok := d.GetOk("loadbalancing_traffic_protocol"); ok || d.HasChange("loadbalancing_traffic_protocol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcLoadbalancingTrafficProtocol(d, v, "loadbalancing_traffic_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loadbalancing-traffic-protocol"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notification_email"); ok || d.HasChange("notification_email") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcNotificationEmail(d, v, "notification_email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification-email"] = t
		}
	}

	if v, ok := d.GetOk("primary_election_timeout"); ok || d.HasChange("primary_election_timeout") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPrimaryElectionTimeout(d, v, "primary_election_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-election-timeout"] = t
		}
	}

	if v, ok := d.GetOk("private_subnet1_cidr"); ok || d.HasChange("private_subnet1_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet1Cidr(d, v, "private_subnet1_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-subnet1-cidr"] = t
		}
	}

	if v, ok := d.GetOk("private_subnet2_cidr"); ok || d.HasChange("private_subnet2_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPrivateSubnet2Cidr(d, v, "private_subnet2_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-subnet2-cidr"] = t
		}
	}

	if v, ok := d.GetOk("public_subnet1_cidr"); ok || d.HasChange("public_subnet1_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet1Cidr(d, v, "public_subnet1_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-subnet1-cidr"] = t
		}
	}

	if v, ok := d.GetOk("public_subnet2_cidr"); ok || d.HasChange("public_subnet2_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcPublicSubnet2Cidr(d, v, "public_subnet2_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-subnet2-cidr"] = t
		}
	}

	if v, ok := d.GetOk("resource_tag_prefix"); ok || d.HasChange("resource_tag_prefix") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcResourceTagPrefix(d, v, "resource_tag_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-tag-prefix"] = t
		}
	}

	if v, ok := d.GetOk("s3_bucket_name"); ok || d.HasChange("s3_bucket_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcS3BucketName(d, v, "s3_bucket_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["s3-bucket-name"] = t
		}
	}

	if v, ok := d.GetOk("s3_key_prefix"); ok || d.HasChange("s3_key_prefix") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcS3KeyPrefix(d, v, "s3_key_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["s3-key-prefix"] = t
		}
	}

	if v, ok := d.GetOk("sync_recovery_count"); ok || d.HasChange("sync_recovery_count") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcSyncRecoveryCount(d, v, "sync_recovery_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-recovery-count"] = t
		}
	}

	if v, ok := d.GetOk("terminate_unhealthy_vm"); ok || d.HasChange("terminate_unhealthy_vm") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcTerminateUnhealthyVm(d, v, "terminate_unhealthy_vm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["terminate-unhealthy-vm"] = t
		}
	}

	if v, ok := d.GetOk("use_custom_asset_location"); ok || d.HasChange("use_custom_asset_location") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcUseCustomAssetLocation(d, v, "use_custom_asset_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-custom-asset-location"] = t
		}
	}

	if v, ok := d.GetOk("vpc_cidr"); ok || d.HasChange("vpc_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleNewVpcVpcCidr(d, v, "vpc_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-cidr"] = t
		}
	}

	return &obj, nil
}
