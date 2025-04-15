// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCloud OrchestAwstemplateAutoscaleExistingVpc

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpc() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcCreate,
		Read:   resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcRead,
		Update: resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcUpdate,
		Delete: resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcDelete,

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
			"private_subnet_route_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_subnet1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_subnet2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_subnet1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_subnet2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vpc_endpoint_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestAwstemplateAutoscaleExistingVpc(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCloudOrchestAwstemplateAutoscaleExistingVpc(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcRead(d, m)
}

func resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestAwstemplateAutoscaleExistingVpc(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCloudOrchestAwstemplateAutoscaleExistingVpc(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcRead(d, m)
}

func resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCloudOrchestAwstemplateAutoscaleExistingVpc(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCloudOrchestAwstemplateAutoscaleExistingVpcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCloudOrchestAwstemplateAutoscaleExistingVpc(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCloudOrchestAwstemplateAutoscaleExistingVpc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwstemplateAutoscaleExistingVpc resource from API: %v", err)
	}
	return nil
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetContainer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazAutoscaleAdminUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazCustomPrivateIpaddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazInstanceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazIntegrationOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtInstanceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtPskSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgCoolDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityPayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgHealthCheckGracePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizeByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizePayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizeByol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizePayg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleInThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleOutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFosVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcGetLicenseGracePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatDelayAllowance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatLossCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancerDnsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancingOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalTargetGroupHealthCheckPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcKeyPairName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLifecycleHookTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingHealthCheckThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcNotificationEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrimaryElectionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnetRouteTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcResourceTagPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3BucketName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3KeyPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcSyncRecoveryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcTerminateUnhealthyVm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcUseCustomAssetLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcEndpointId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCloudOrchestAwstemplateAutoscaleExistingVpc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("custom_asset_container", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetContainer(o["custom-asset-container"], d, "custom_asset_container")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-asset-container"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-CustomAssetContainer"); ok {
			if err = d.Set("custom_asset_container", vv); err != nil {
				return fmt.Errorf("Error reading custom_asset_container: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_asset_container: %v", err)
		}
	}

	if err = d.Set("custom_asset_directory", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetDirectory(o["custom-asset-directory"], d, "custom_asset_directory")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-asset-directory"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-CustomAssetDirectory"); ok {
			if err = d.Set("custom_asset_directory", vv); err != nil {
				return fmt.Errorf("Error reading custom_asset_directory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_asset_directory: %v", err)
		}
	}

	if err = d.Set("custom_identifier", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomIdentifier(o["custom-identifier"], d, "custom_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-identifier"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-CustomIdentifier"); ok {
			if err = d.Set("custom_identifier", vv); err != nil {
				return fmt.Errorf("Error reading custom_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_identifier: %v", err)
		}
	}

	if err = d.Set("faz_autoscale_admin_username", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazAutoscaleAdminUsername(o["faz-autoscale-admin-username"], d, "faz_autoscale_admin_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-autoscale-admin-username"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FazAutoscaleAdminUsername"); ok {
			if err = d.Set("faz_autoscale_admin_username", vv); err != nil {
				return fmt.Errorf("Error reading faz_autoscale_admin_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_autoscale_admin_username: %v", err)
		}
	}

	if err = d.Set("faz_custom_private_ipaddress", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazCustomPrivateIpaddress(o["faz-custom-private-ipaddress"], d, "faz_custom_private_ipaddress")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-custom-private-ipaddress"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FazCustomPrivateIpaddress"); ok {
			if err = d.Set("faz_custom_private_ipaddress", vv); err != nil {
				return fmt.Errorf("Error reading faz_custom_private_ipaddress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_custom_private_ipaddress: %v", err)
		}
	}

	if err = d.Set("faz_instance_type", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazInstanceType(o["faz-instance-type"], d, "faz_instance_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-instance-type"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FazInstanceType"); ok {
			if err = d.Set("faz_instance_type", vv); err != nil {
				return fmt.Errorf("Error reading faz_instance_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_instance_type: %v", err)
		}
	}

	if err = d.Set("faz_integration_options", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazIntegrationOptions(o["faz-integration-options"], d, "faz_integration_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-integration-options"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FazIntegrationOptions"); ok {
			if err = d.Set("faz_integration_options", vv); err != nil {
				return fmt.Errorf("Error reading faz_integration_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_integration_options: %v", err)
		}
	}

	if err = d.Set("faz_version", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazVersion(o["faz-version"], d, "faz_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-version"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FazVersion"); ok {
			if err = d.Set("faz_version", vv); err != nil {
				return fmt.Errorf("Error reading faz_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_version: %v", err)
		}
	}

	if err = d.Set("fgt_admin_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminCidr(o["fgt-admin-cidr"], d, "fgt_admin_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-admin-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtAdminCidr"); ok {
			if err = d.Set("fgt_admin_cidr", vv); err != nil {
				return fmt.Errorf("Error reading fgt_admin_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_admin_cidr: %v", err)
		}
	}

	if err = d.Set("fgt_admin_port", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminPort(o["fgt-admin-port"], d, "fgt_admin_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-admin-port"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtAdminPort"); ok {
			if err = d.Set("fgt_admin_port", vv); err != nil {
				return fmt.Errorf("Error reading fgt_admin_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_admin_port: %v", err)
		}
	}

	if err = d.Set("fgt_instance_type", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtInstanceType(o["fgt-instance-type"], d, "fgt_instance_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-instance-type"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtInstanceType"); ok {
			if err = d.Set("fgt_instance_type", vv); err != nil {
				return fmt.Errorf("Error reading fgt_instance_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_instance_type: %v", err)
		}
	}

	if err = d.Set("fgt_psk_secret", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtPskSecret(o["fgt-psk-secret"], d, "fgt_psk_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-psk-secret"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtPskSecret"); ok {
			if err = d.Set("fgt_psk_secret", vv); err != nil {
				return fmt.Errorf("Error reading fgt_psk_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_psk_secret: %v", err)
		}
	}

	if err = d.Set("fgtasg_cool_down", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgCoolDown(o["fgtasg-cool-down"], d, "fgtasg_cool_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-cool-down"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgCoolDown"); ok {
			if err = d.Set("fgtasg_cool_down", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_cool_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_cool_down: %v", err)
		}
	}

	if err = d.Set("fgtasg_desired_capacity_byol", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityByol(o["fgtasg-desired-capacity-byol"], d, "fgtasg_desired_capacity_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-desired-capacity-byol"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgDesiredCapacityByol"); ok {
			if err = d.Set("fgtasg_desired_capacity_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_desired_capacity_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_desired_capacity_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_desired_capacity_payg", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityPayg(o["fgtasg-desired-capacity-payg"], d, "fgtasg_desired_capacity_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-desired-capacity-payg"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgDesiredCapacityPayg"); ok {
			if err = d.Set("fgtasg_desired_capacity_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_desired_capacity_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_desired_capacity_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_health_check_grace_period", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgHealthCheckGracePeriod(o["fgtasg-health-check-grace-period"], d, "fgtasg_health_check_grace_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-health-check-grace-period"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgHealthCheckGracePeriod"); ok {
			if err = d.Set("fgtasg_health_check_grace_period", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_health_check_grace_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_health_check_grace_period: %v", err)
		}
	}

	if err = d.Set("fgtasg_max_size_byol", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizeByol(o["fgtasg-max-size-byol"], d, "fgtasg_max_size_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-max-size-byol"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgMaxSizeByol"); ok {
			if err = d.Set("fgtasg_max_size_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_max_size_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_max_size_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_max_size_payg", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizePayg(o["fgtasg-max-size-payg"], d, "fgtasg_max_size_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-max-size-payg"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgMaxSizePayg"); ok {
			if err = d.Set("fgtasg_max_size_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_max_size_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_max_size_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_min_size_byol", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizeByol(o["fgtasg-min-size-byol"], d, "fgtasg_min_size_byol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-min-size-byol"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgMinSizeByol"); ok {
			if err = d.Set("fgtasg_min_size_byol", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_min_size_byol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_min_size_byol: %v", err)
		}
	}

	if err = d.Set("fgtasg_min_size_payg", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizePayg(o["fgtasg-min-size-payg"], d, "fgtasg_min_size_payg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-min-size-payg"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgMinSizePayg"); ok {
			if err = d.Set("fgtasg_min_size_payg", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_min_size_payg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_min_size_payg: %v", err)
		}
	}

	if err = d.Set("fgtasg_scale_in_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleInThreshold(o["fgtasg-scale-in-threshold"], d, "fgtasg_scale_in_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-scale-in-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgScaleInThreshold"); ok {
			if err = d.Set("fgtasg_scale_in_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_scale_in_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_scale_in_threshold: %v", err)
		}
	}

	if err = d.Set("fgtasg_scale_out_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleOutThreshold(o["fgtasg-scale-out-threshold"], d, "fgtasg_scale_out_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgtasg-scale-out-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FgtasgScaleOutThreshold"); ok {
			if err = d.Set("fgtasg_scale_out_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fgtasg_scale_out_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgtasg_scale_out_threshold: %v", err)
		}
	}

	if err = d.Set("fos_version", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFosVersion(o["fos-version"], d, "fos_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["fos-version"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-FosVersion"); ok {
			if err = d.Set("fos_version", vv); err != nil {
				return fmt.Errorf("Error reading fos_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fos_version: %v", err)
		}
	}

	if err = d.Set("get_license_grace_period", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcGetLicenseGracePeriod(o["get-license-grace-period"], d, "get_license_grace_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["get-license-grace-period"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-GetLicenseGracePeriod"); ok {
			if err = d.Set("get_license_grace_period", vv); err != nil {
				return fmt.Errorf("Error reading get_license_grace_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading get_license_grace_period: %v", err)
		}
	}

	if err = d.Set("heartbeat_delay_allowance", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatDelayAllowance(o["heartbeat-delay-allowance"], d, "heartbeat_delay_allowance")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-delay-allowance"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-HeartbeatDelayAllowance"); ok {
			if err = d.Set("heartbeat_delay_allowance", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_delay_allowance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_delay_allowance: %v", err)
		}
	}

	if err = d.Set("heartbeat_interval", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatInterval(o["heartbeat-interval"], d, "heartbeat_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-interval"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-HeartbeatInterval"); ok {
			if err = d.Set("heartbeat_interval", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("heartbeat_loss_count", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatLossCount(o["heartbeat-loss-count"], d, "heartbeat_loss_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["heartbeat-loss-count"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-HeartbeatLossCount"); ok {
			if err = d.Set("heartbeat_loss_count", vv); err != nil {
				return fmt.Errorf("Error reading heartbeat_loss_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heartbeat_loss_count: %v", err)
		}
	}

	if err = d.Set("internal_balancer_dns_name", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancerDnsName(o["internal-balancer-dns-name"], d, "internal_balancer_dns_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-balancer-dns-name"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-InternalBalancerDnsName"); ok {
			if err = d.Set("internal_balancer_dns_name", vv); err != nil {
				return fmt.Errorf("Error reading internal_balancer_dns_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_balancer_dns_name: %v", err)
		}
	}

	if err = d.Set("internal_balancing_options", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancingOptions(o["internal-balancing-options"], d, "internal_balancing_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-balancing-options"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-InternalBalancingOptions"); ok {
			if err = d.Set("internal_balancing_options", vv); err != nil {
				return fmt.Errorf("Error reading internal_balancing_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_balancing_options: %v", err)
		}
	}

	if err = d.Set("internal_target_group_health_check_path", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalTargetGroupHealthCheckPath(o["internal-target-group-health-check-path"], d, "internal_target_group_health_check_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-target-group-health-check-path"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-InternalTargetGroupHealthCheckPath"); ok {
			if err = d.Set("internal_target_group_health_check_path", vv); err != nil {
				return fmt.Errorf("Error reading internal_target_group_health_check_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_target_group_health_check_path: %v", err)
		}
	}

	if err = d.Set("key_pair_name", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcKeyPairName(o["key-pair-name"], d, "key_pair_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-pair-name"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-KeyPairName"); ok {
			if err = d.Set("key_pair_name", vv); err != nil {
				return fmt.Errorf("Error reading key_pair_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_pair_name: %v", err)
		}
	}

	if err = d.Set("lifecycle_hook_timeout", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLifecycleHookTimeout(o["lifecycle-hook-timeout"], d, "lifecycle_hook_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["lifecycle-hook-timeout"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-LifecycleHookTimeout"); ok {
			if err = d.Set("lifecycle_hook_timeout", vv); err != nil {
				return fmt.Errorf("Error reading lifecycle_hook_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lifecycle_hook_timeout: %v", err)
		}
	}

	if err = d.Set("loadbalancing_health_check_threshold", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingHealthCheckThreshold(o["loadbalancing-health-check-threshold"], d, "loadbalancing_health_check_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["loadbalancing-health-check-threshold"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-LoadbalancingHealthCheckThreshold"); ok {
			if err = d.Set("loadbalancing_health_check_threshold", vv); err != nil {
				return fmt.Errorf("Error reading loadbalancing_health_check_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loadbalancing_health_check_threshold: %v", err)
		}
	}

	if err = d.Set("loadbalancing_traffic_port", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficPort(o["loadbalancing-traffic-port"], d, "loadbalancing_traffic_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["loadbalancing-traffic-port"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-LoadbalancingTrafficPort"); ok {
			if err = d.Set("loadbalancing_traffic_port", vv); err != nil {
				return fmt.Errorf("Error reading loadbalancing_traffic_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loadbalancing_traffic_port: %v", err)
		}
	}

	if err = d.Set("loadbalancing_traffic_protocol", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficProtocol(o["loadbalancing-traffic-protocol"], d, "loadbalancing_traffic_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["loadbalancing-traffic-protocol"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-LoadbalancingTrafficProtocol"); ok {
			if err = d.Set("loadbalancing_traffic_protocol", vv); err != nil {
				return fmt.Errorf("Error reading loadbalancing_traffic_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loadbalancing_traffic_protocol: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("notification_email", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcNotificationEmail(o["notification-email"], d, "notification_email")); err != nil {
		if vv, ok := fortiAPIPatch(o["notification-email"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-NotificationEmail"); ok {
			if err = d.Set("notification_email", vv); err != nil {
				return fmt.Errorf("Error reading notification_email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notification_email: %v", err)
		}
	}

	if err = d.Set("primary_election_timeout", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrimaryElectionTimeout(o["primary-election-timeout"], d, "primary_election_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-election-timeout"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-PrimaryElectionTimeout"); ok {
			if err = d.Set("primary_election_timeout", vv); err != nil {
				return fmt.Errorf("Error reading primary_election_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_election_timeout: %v", err)
		}
	}

	if err = d.Set("private_subnet_route_table", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnetRouteTable(o["private-subnet-route-table"], d, "private_subnet_route_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-subnet-route-table"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-PrivateSubnetRouteTable"); ok {
			if err = d.Set("private_subnet_route_table", vv); err != nil {
				return fmt.Errorf("Error reading private_subnet_route_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_subnet_route_table: %v", err)
		}
	}

	if err = d.Set("private_subnet1", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet1(o["private-subnet1"], d, "private_subnet1")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-subnet1"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-PrivateSubnet1"); ok {
			if err = d.Set("private_subnet1", vv); err != nil {
				return fmt.Errorf("Error reading private_subnet1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_subnet1: %v", err)
		}
	}

	if err = d.Set("private_subnet2", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet2(o["private-subnet2"], d, "private_subnet2")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-subnet2"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-PrivateSubnet2"); ok {
			if err = d.Set("private_subnet2", vv); err != nil {
				return fmt.Errorf("Error reading private_subnet2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_subnet2: %v", err)
		}
	}

	if err = d.Set("public_subnet1", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet1(o["public-subnet1"], d, "public_subnet1")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-subnet1"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-PublicSubnet1"); ok {
			if err = d.Set("public_subnet1", vv); err != nil {
				return fmt.Errorf("Error reading public_subnet1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_subnet1: %v", err)
		}
	}

	if err = d.Set("public_subnet2", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet2(o["public-subnet2"], d, "public_subnet2")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-subnet2"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-PublicSubnet2"); ok {
			if err = d.Set("public_subnet2", vv); err != nil {
				return fmt.Errorf("Error reading public_subnet2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_subnet2: %v", err)
		}
	}

	if err = d.Set("resource_tag_prefix", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcResourceTagPrefix(o["resource-tag-prefix"], d, "resource_tag_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-tag-prefix"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-ResourceTagPrefix"); ok {
			if err = d.Set("resource_tag_prefix", vv); err != nil {
				return fmt.Errorf("Error reading resource_tag_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_tag_prefix: %v", err)
		}
	}

	if err = d.Set("s3_bucket_name", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3BucketName(o["s3-bucket-name"], d, "s3_bucket_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["s3-bucket-name"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-S3BucketName"); ok {
			if err = d.Set("s3_bucket_name", vv); err != nil {
				return fmt.Errorf("Error reading s3_bucket_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading s3_bucket_name: %v", err)
		}
	}

	if err = d.Set("s3_key_prefix", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3KeyPrefix(o["s3-key-prefix"], d, "s3_key_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["s3-key-prefix"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-S3KeyPrefix"); ok {
			if err = d.Set("s3_key_prefix", vv); err != nil {
				return fmt.Errorf("Error reading s3_key_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading s3_key_prefix: %v", err)
		}
	}

	if err = d.Set("sync_recovery_count", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcSyncRecoveryCount(o["sync-recovery-count"], d, "sync_recovery_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-recovery-count"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-SyncRecoveryCount"); ok {
			if err = d.Set("sync_recovery_count", vv); err != nil {
				return fmt.Errorf("Error reading sync_recovery_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_recovery_count: %v", err)
		}
	}

	if err = d.Set("terminate_unhealthy_vm", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcTerminateUnhealthyVm(o["terminate-unhealthy-vm"], d, "terminate_unhealthy_vm")); err != nil {
		if vv, ok := fortiAPIPatch(o["terminate-unhealthy-vm"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-TerminateUnhealthyVm"); ok {
			if err = d.Set("terminate_unhealthy_vm", vv); err != nil {
				return fmt.Errorf("Error reading terminate_unhealthy_vm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading terminate_unhealthy_vm: %v", err)
		}
	}

	if err = d.Set("use_custom_asset_location", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcUseCustomAssetLocation(o["use-custom-asset-location"], d, "use_custom_asset_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-custom-asset-location"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-UseCustomAssetLocation"); ok {
			if err = d.Set("use_custom_asset_location", vv); err != nil {
				return fmt.Errorf("Error reading use_custom_asset_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_custom_asset_location: %v", err)
		}
	}

	if err = d.Set("vpc_cidr", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcCidr(o["vpc-cidr"], d, "vpc_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-cidr"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-VpcCidr"); ok {
			if err = d.Set("vpc_cidr", vv); err != nil {
				return fmt.Errorf("Error reading vpc_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_cidr: %v", err)
		}
	}

	if err = d.Set("vpc_endpoint_id", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcEndpointId(o["vpc-endpoint-id"], d, "vpc_endpoint_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-endpoint-id"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-VpcEndpointId"); ok {
			if err = d.Set("vpc_endpoint_id", vv); err != nil {
				return fmt.Errorf("Error reading vpc_endpoint_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_endpoint_id: %v", err)
		}
	}

	if err = d.Set("vpc_id", flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcId(o["vpc-id"], d, "vpc_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-id"], "ObjectCloudOrchestAwstemplateAutoscaleExistingVpc-VpcId"); ok {
			if err = d.Set("vpc_id", vv); err != nil {
				return fmt.Errorf("Error reading vpc_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_id: %v", err)
		}
	}

	return nil
}

func flattenObjectCloudOrchestAwstemplateAutoscaleExistingVpcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetContainer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazAutoscaleAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazAutoscaleAdminUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazCustomPrivateIpaddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazInstanceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazIntegrationOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtInstanceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtPskSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgCoolDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityPayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgHealthCheckGracePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizeByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizePayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizeByol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizePayg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleInThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleOutThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFosVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcGetLicenseGracePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatDelayAllowance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatLossCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancerDnsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancingOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalTargetGroupHealthCheckPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcKeyPairName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLifecycleHookTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingHealthCheckThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcNotificationEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrimaryElectionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnetRouteTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcResourceTagPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3BucketName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3KeyPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcSyncRecoveryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcTerminateUnhealthyVm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcUseCustomAssetLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcEndpointId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCloudOrchestAwstemplateAutoscaleExistingVpc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom_asset_container"); ok || d.HasChange("custom_asset_container") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetContainer(d, v, "custom_asset_container")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-asset-container"] = t
		}
	}

	if v, ok := d.GetOk("custom_asset_directory"); ok || d.HasChange("custom_asset_directory") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomAssetDirectory(d, v, "custom_asset_directory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-asset-directory"] = t
		}
	}

	if v, ok := d.GetOk("custom_identifier"); ok || d.HasChange("custom_identifier") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcCustomIdentifier(d, v, "custom_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-identifier"] = t
		}
	}

	if v, ok := d.GetOk("faz_autoscale_admin_password"); ok || d.HasChange("faz_autoscale_admin_password") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazAutoscaleAdminPassword(d, v, "faz_autoscale_admin_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-autoscale-admin-password"] = t
		}
	}

	if v, ok := d.GetOk("faz_autoscale_admin_username"); ok || d.HasChange("faz_autoscale_admin_username") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazAutoscaleAdminUsername(d, v, "faz_autoscale_admin_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-autoscale-admin-username"] = t
		}
	}

	if v, ok := d.GetOk("faz_custom_private_ipaddress"); ok || d.HasChange("faz_custom_private_ipaddress") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazCustomPrivateIpaddress(d, v, "faz_custom_private_ipaddress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-custom-private-ipaddress"] = t
		}
	}

	if v, ok := d.GetOk("faz_instance_type"); ok || d.HasChange("faz_instance_type") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazInstanceType(d, v, "faz_instance_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-instance-type"] = t
		}
	}

	if v, ok := d.GetOk("faz_integration_options"); ok || d.HasChange("faz_integration_options") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazIntegrationOptions(d, v, "faz_integration_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-integration-options"] = t
		}
	}

	if v, ok := d.GetOk("faz_version"); ok || d.HasChange("faz_version") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFazVersion(d, v, "faz_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-version"] = t
		}
	}

	if v, ok := d.GetOk("fgt_admin_cidr"); ok || d.HasChange("fgt_admin_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminCidr(d, v, "fgt_admin_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-admin-cidr"] = t
		}
	}

	if v, ok := d.GetOk("fgt_admin_port"); ok || d.HasChange("fgt_admin_port") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtAdminPort(d, v, "fgt_admin_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-admin-port"] = t
		}
	}

	if v, ok := d.GetOk("fgt_instance_type"); ok || d.HasChange("fgt_instance_type") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtInstanceType(d, v, "fgt_instance_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-instance-type"] = t
		}
	}

	if v, ok := d.GetOk("fgt_psk_secret"); ok || d.HasChange("fgt_psk_secret") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtPskSecret(d, v, "fgt_psk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-psk-secret"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_cool_down"); ok || d.HasChange("fgtasg_cool_down") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgCoolDown(d, v, "fgtasg_cool_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-cool-down"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_desired_capacity_byol"); ok || d.HasChange("fgtasg_desired_capacity_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityByol(d, v, "fgtasg_desired_capacity_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-desired-capacity-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_desired_capacity_payg"); ok || d.HasChange("fgtasg_desired_capacity_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgDesiredCapacityPayg(d, v, "fgtasg_desired_capacity_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-desired-capacity-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_health_check_grace_period"); ok || d.HasChange("fgtasg_health_check_grace_period") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgHealthCheckGracePeriod(d, v, "fgtasg_health_check_grace_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-health-check-grace-period"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_max_size_byol"); ok || d.HasChange("fgtasg_max_size_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizeByol(d, v, "fgtasg_max_size_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-max-size-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_max_size_payg"); ok || d.HasChange("fgtasg_max_size_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMaxSizePayg(d, v, "fgtasg_max_size_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-max-size-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_min_size_byol"); ok || d.HasChange("fgtasg_min_size_byol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizeByol(d, v, "fgtasg_min_size_byol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-min-size-byol"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_min_size_payg"); ok || d.HasChange("fgtasg_min_size_payg") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgMinSizePayg(d, v, "fgtasg_min_size_payg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-min-size-payg"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_scale_in_threshold"); ok || d.HasChange("fgtasg_scale_in_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleInThreshold(d, v, "fgtasg_scale_in_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-scale-in-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fgtasg_scale_out_threshold"); ok || d.HasChange("fgtasg_scale_out_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFgtasgScaleOutThreshold(d, v, "fgtasg_scale_out_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgtasg-scale-out-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fos_version"); ok || d.HasChange("fos_version") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcFosVersion(d, v, "fos_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fos-version"] = t
		}
	}

	if v, ok := d.GetOk("get_license_grace_period"); ok || d.HasChange("get_license_grace_period") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcGetLicenseGracePeriod(d, v, "get_license_grace_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["get-license-grace-period"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_delay_allowance"); ok || d.HasChange("heartbeat_delay_allowance") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatDelayAllowance(d, v, "heartbeat_delay_allowance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-delay-allowance"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_interval"); ok || d.HasChange("heartbeat_interval") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatInterval(d, v, "heartbeat_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("heartbeat_loss_count"); ok || d.HasChange("heartbeat_loss_count") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcHeartbeatLossCount(d, v, "heartbeat_loss_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heartbeat-loss-count"] = t
		}
	}

	if v, ok := d.GetOk("internal_balancer_dns_name"); ok || d.HasChange("internal_balancer_dns_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancerDnsName(d, v, "internal_balancer_dns_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-balancer-dns-name"] = t
		}
	}

	if v, ok := d.GetOk("internal_balancing_options"); ok || d.HasChange("internal_balancing_options") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalBalancingOptions(d, v, "internal_balancing_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-balancing-options"] = t
		}
	}

	if v, ok := d.GetOk("internal_target_group_health_check_path"); ok || d.HasChange("internal_target_group_health_check_path") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcInternalTargetGroupHealthCheckPath(d, v, "internal_target_group_health_check_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-target-group-health-check-path"] = t
		}
	}

	if v, ok := d.GetOk("key_pair_name"); ok || d.HasChange("key_pair_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcKeyPairName(d, v, "key_pair_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-pair-name"] = t
		}
	}

	if v, ok := d.GetOk("lifecycle_hook_timeout"); ok || d.HasChange("lifecycle_hook_timeout") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLifecycleHookTimeout(d, v, "lifecycle_hook_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lifecycle-hook-timeout"] = t
		}
	}

	if v, ok := d.GetOk("loadbalancing_health_check_threshold"); ok || d.HasChange("loadbalancing_health_check_threshold") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingHealthCheckThreshold(d, v, "loadbalancing_health_check_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loadbalancing-health-check-threshold"] = t
		}
	}

	if v, ok := d.GetOk("loadbalancing_traffic_port"); ok || d.HasChange("loadbalancing_traffic_port") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficPort(d, v, "loadbalancing_traffic_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loadbalancing-traffic-port"] = t
		}
	}

	if v, ok := d.GetOk("loadbalancing_traffic_protocol"); ok || d.HasChange("loadbalancing_traffic_protocol") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcLoadbalancingTrafficProtocol(d, v, "loadbalancing_traffic_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loadbalancing-traffic-protocol"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notification_email"); ok || d.HasChange("notification_email") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcNotificationEmail(d, v, "notification_email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification-email"] = t
		}
	}

	if v, ok := d.GetOk("primary_election_timeout"); ok || d.HasChange("primary_election_timeout") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrimaryElectionTimeout(d, v, "primary_election_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-election-timeout"] = t
		}
	}

	if v, ok := d.GetOk("private_subnet_route_table"); ok || d.HasChange("private_subnet_route_table") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnetRouteTable(d, v, "private_subnet_route_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-subnet-route-table"] = t
		}
	}

	if v, ok := d.GetOk("private_subnet1"); ok || d.HasChange("private_subnet1") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet1(d, v, "private_subnet1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-subnet1"] = t
		}
	}

	if v, ok := d.GetOk("private_subnet2"); ok || d.HasChange("private_subnet2") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPrivateSubnet2(d, v, "private_subnet2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-subnet2"] = t
		}
	}

	if v, ok := d.GetOk("public_subnet1"); ok || d.HasChange("public_subnet1") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet1(d, v, "public_subnet1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-subnet1"] = t
		}
	}

	if v, ok := d.GetOk("public_subnet2"); ok || d.HasChange("public_subnet2") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcPublicSubnet2(d, v, "public_subnet2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-subnet2"] = t
		}
	}

	if v, ok := d.GetOk("resource_tag_prefix"); ok || d.HasChange("resource_tag_prefix") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcResourceTagPrefix(d, v, "resource_tag_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-tag-prefix"] = t
		}
	}

	if v, ok := d.GetOk("s3_bucket_name"); ok || d.HasChange("s3_bucket_name") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3BucketName(d, v, "s3_bucket_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["s3-bucket-name"] = t
		}
	}

	if v, ok := d.GetOk("s3_key_prefix"); ok || d.HasChange("s3_key_prefix") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcS3KeyPrefix(d, v, "s3_key_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["s3-key-prefix"] = t
		}
	}

	if v, ok := d.GetOk("sync_recovery_count"); ok || d.HasChange("sync_recovery_count") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcSyncRecoveryCount(d, v, "sync_recovery_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-recovery-count"] = t
		}
	}

	if v, ok := d.GetOk("terminate_unhealthy_vm"); ok || d.HasChange("terminate_unhealthy_vm") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcTerminateUnhealthyVm(d, v, "terminate_unhealthy_vm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["terminate-unhealthy-vm"] = t
		}
	}

	if v, ok := d.GetOk("use_custom_asset_location"); ok || d.HasChange("use_custom_asset_location") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcUseCustomAssetLocation(d, v, "use_custom_asset_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-custom-asset-location"] = t
		}
	}

	if v, ok := d.GetOk("vpc_cidr"); ok || d.HasChange("vpc_cidr") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcCidr(d, v, "vpc_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-cidr"] = t
		}
	}

	if v, ok := d.GetOk("vpc_endpoint_id"); ok || d.HasChange("vpc_endpoint_id") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcEndpointId(d, v, "vpc_endpoint_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-endpoint-id"] = t
		}
	}

	if v, ok := d.GetOk("vpc_id"); ok || d.HasChange("vpc_id") {
		t, err := expandObjectCloudOrchestAwstemplateAutoscaleExistingVpcVpcId(d, v, "vpc_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-id"] = t
		}
	}

	return &obj, nil
}
