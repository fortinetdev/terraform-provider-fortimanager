// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure dm.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemDm() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDmUpdate,
		Read:   resourceSystemDmRead,
		Update: resourceSystemDmUpdate,
		Delete: resourceSystemDmDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"concurrent_install_image_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"concurrent_install_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"concurrent_install_script_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"conf_merge_after_script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discover_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dpm_logsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgfm_install_refresh_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgfm_sock_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fgfm_keepalive_itvl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"force_remote_diff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiap_refresh_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortiap_refresh_itvl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortiext_refresh_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"install_image_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"install_tunnel_retry_itvl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_revs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nr_retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retry_intvl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rollback_allow_reboot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script_logsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"skip_scep_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_tunnel_fcp_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_install": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemDm(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDm resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDm(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemDm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemDm")

	return resourceSystemDmRead(d, m)
}

func resourceSystemDmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemDm(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemDm(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemDm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDm resource from API: %v", err)
	}
	return nil
}

func flattenSystemDmConcurrentInstallImageLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmConcurrentInstallLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmConcurrentInstallScriptLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmConfMergeAfterScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemDmDiscoverTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmDpmLogsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmFgfmInstallRefreshCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmFgfmSockTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmFgfmKeepaliveItvl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmForceRemoteDiff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemDmFortiapRefreshCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmFortiapRefreshItvl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmFortiextRefreshCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmInstallImageTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmInstallTunnelRetryItvl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmMaxRevs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmNrRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemDmRetryIntvl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmRollbackAllowReboot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemDmScriptLogsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDmSkipScepCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemDmSkipTunnelFcpReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemDmVerifyInstall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "optimal",
			2: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemDm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("concurrent_install_image_limit", flattenSystemDmConcurrentInstallImageLimit(o["concurrent-install-image-limit"], d, "concurrent_install_image_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-install-image-limit"], "SystemDm-ConcurrentInstallImageLimit"); ok {
			if err = d.Set("concurrent_install_image_limit", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_install_image_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_install_image_limit: %v", err)
		}
	}

	if err = d.Set("concurrent_install_limit", flattenSystemDmConcurrentInstallLimit(o["concurrent-install-limit"], d, "concurrent_install_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-install-limit"], "SystemDm-ConcurrentInstallLimit"); ok {
			if err = d.Set("concurrent_install_limit", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_install_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_install_limit: %v", err)
		}
	}

	if err = d.Set("concurrent_install_script_limit", flattenSystemDmConcurrentInstallScriptLimit(o["concurrent-install-script-limit"], d, "concurrent_install_script_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-install-script-limit"], "SystemDm-ConcurrentInstallScriptLimit"); ok {
			if err = d.Set("concurrent_install_script_limit", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_install_script_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_install_script_limit: %v", err)
		}
	}

	if err = d.Set("conf_merge_after_script", flattenSystemDmConfMergeAfterScript(o["conf-merge-after-script"], d, "conf_merge_after_script")); err != nil {
		if vv, ok := fortiAPIPatch(o["conf-merge-after-script"], "SystemDm-ConfMergeAfterScript"); ok {
			if err = d.Set("conf_merge_after_script", vv); err != nil {
				return fmt.Errorf("Error reading conf_merge_after_script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conf_merge_after_script: %v", err)
		}
	}

	if err = d.Set("discover_timeout", flattenSystemDmDiscoverTimeout(o["discover-timeout"], d, "discover_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["discover-timeout"], "SystemDm-DiscoverTimeout"); ok {
			if err = d.Set("discover_timeout", vv); err != nil {
				return fmt.Errorf("Error reading discover_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discover_timeout: %v", err)
		}
	}

	if err = d.Set("dpm_logsize", flattenSystemDmDpmLogsize(o["dpm-logsize"], d, "dpm_logsize")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpm-logsize"], "SystemDm-DpmLogsize"); ok {
			if err = d.Set("dpm_logsize", vv); err != nil {
				return fmt.Errorf("Error reading dpm_logsize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpm_logsize: %v", err)
		}
	}

	if err = d.Set("fgfm_install_refresh_count", flattenSystemDmFgfmInstallRefreshCount(o["fgfm-install-refresh-count"], d, "fgfm_install_refresh_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-install-refresh-count"], "SystemDm-FgfmInstallRefreshCount"); ok {
			if err = d.Set("fgfm_install_refresh_count", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_install_refresh_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_install_refresh_count: %v", err)
		}
	}

	if err = d.Set("fgfm_sock_timeout", flattenSystemDmFgfmSockTimeout(o["fgfm-sock-timeout"], d, "fgfm_sock_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-sock-timeout"], "SystemDm-FgfmSockTimeout"); ok {
			if err = d.Set("fgfm_sock_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_sock_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_sock_timeout: %v", err)
		}
	}

	if err = d.Set("fgfm_keepalive_itvl", flattenSystemDmFgfmKeepaliveItvl(o["fgfm_keepalive_itvl"], d, "fgfm_keepalive_itvl")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm_keepalive_itvl"], "SystemDm-FgfmKeepaliveItvl"); ok {
			if err = d.Set("fgfm_keepalive_itvl", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_keepalive_itvl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_keepalive_itvl: %v", err)
		}
	}

	if err = d.Set("force_remote_diff", flattenSystemDmForceRemoteDiff(o["force-remote-diff"], d, "force_remote_diff")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-remote-diff"], "SystemDm-ForceRemoteDiff"); ok {
			if err = d.Set("force_remote_diff", vv); err != nil {
				return fmt.Errorf("Error reading force_remote_diff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_remote_diff: %v", err)
		}
	}

	if err = d.Set("fortiap_refresh_cnt", flattenSystemDmFortiapRefreshCnt(o["fortiap-refresh-cnt"], d, "fortiap_refresh_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiap-refresh-cnt"], "SystemDm-FortiapRefreshCnt"); ok {
			if err = d.Set("fortiap_refresh_cnt", vv); err != nil {
				return fmt.Errorf("Error reading fortiap_refresh_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiap_refresh_cnt: %v", err)
		}
	}

	if err = d.Set("fortiap_refresh_itvl", flattenSystemDmFortiapRefreshItvl(o["fortiap-refresh-itvl"], d, "fortiap_refresh_itvl")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiap-refresh-itvl"], "SystemDm-FortiapRefreshItvl"); ok {
			if err = d.Set("fortiap_refresh_itvl", vv); err != nil {
				return fmt.Errorf("Error reading fortiap_refresh_itvl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiap_refresh_itvl: %v", err)
		}
	}

	if err = d.Set("fortiext_refresh_cnt", flattenSystemDmFortiextRefreshCnt(o["fortiext-refresh-cnt"], d, "fortiext_refresh_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiext-refresh-cnt"], "SystemDm-FortiextRefreshCnt"); ok {
			if err = d.Set("fortiext_refresh_cnt", vv); err != nil {
				return fmt.Errorf("Error reading fortiext_refresh_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiext_refresh_cnt: %v", err)
		}
	}

	if err = d.Set("install_image_timeout", flattenSystemDmInstallImageTimeout(o["install-image-timeout"], d, "install_image_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["install-image-timeout"], "SystemDm-InstallImageTimeout"); ok {
			if err = d.Set("install_image_timeout", vv); err != nil {
				return fmt.Errorf("Error reading install_image_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading install_image_timeout: %v", err)
		}
	}

	if err = d.Set("install_tunnel_retry_itvl", flattenSystemDmInstallTunnelRetryItvl(o["install-tunnel-retry-itvl"], d, "install_tunnel_retry_itvl")); err != nil {
		if vv, ok := fortiAPIPatch(o["install-tunnel-retry-itvl"], "SystemDm-InstallTunnelRetryItvl"); ok {
			if err = d.Set("install_tunnel_retry_itvl", vv); err != nil {
				return fmt.Errorf("Error reading install_tunnel_retry_itvl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading install_tunnel_retry_itvl: %v", err)
		}
	}

	if err = d.Set("max_revs", flattenSystemDmMaxRevs(o["max-revs"], d, "max_revs")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-revs"], "SystemDm-MaxRevs"); ok {
			if err = d.Set("max_revs", vv); err != nil {
				return fmt.Errorf("Error reading max_revs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_revs: %v", err)
		}
	}

	if err = d.Set("nr_retry", flattenSystemDmNrRetry(o["nr-retry"], d, "nr_retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["nr-retry"], "SystemDm-NrRetry"); ok {
			if err = d.Set("nr_retry", vv); err != nil {
				return fmt.Errorf("Error reading nr_retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nr_retry: %v", err)
		}
	}

	if err = d.Set("retry", flattenSystemDmRetry(o["retry"], d, "retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry"], "SystemDm-Retry"); ok {
			if err = d.Set("retry", vv); err != nil {
				return fmt.Errorf("Error reading retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("retry_intvl", flattenSystemDmRetryIntvl(o["retry-intvl"], d, "retry_intvl")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry-intvl"], "SystemDm-RetryIntvl"); ok {
			if err = d.Set("retry_intvl", vv); err != nil {
				return fmt.Errorf("Error reading retry_intvl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry_intvl: %v", err)
		}
	}

	if err = d.Set("rollback_allow_reboot", flattenSystemDmRollbackAllowReboot(o["rollback-allow-reboot"], d, "rollback_allow_reboot")); err != nil {
		if vv, ok := fortiAPIPatch(o["rollback-allow-reboot"], "SystemDm-RollbackAllowReboot"); ok {
			if err = d.Set("rollback_allow_reboot", vv); err != nil {
				return fmt.Errorf("Error reading rollback_allow_reboot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rollback_allow_reboot: %v", err)
		}
	}

	if err = d.Set("script_logsize", flattenSystemDmScriptLogsize(o["script-logsize"], d, "script_logsize")); err != nil {
		if vv, ok := fortiAPIPatch(o["script-logsize"], "SystemDm-ScriptLogsize"); ok {
			if err = d.Set("script_logsize", vv); err != nil {
				return fmt.Errorf("Error reading script_logsize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script_logsize: %v", err)
		}
	}

	if err = d.Set("skip_scep_check", flattenSystemDmSkipScepCheck(o["skip-scep-check"], d, "skip_scep_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-scep-check"], "SystemDm-SkipScepCheck"); ok {
			if err = d.Set("skip_scep_check", vv); err != nil {
				return fmt.Errorf("Error reading skip_scep_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_scep_check: %v", err)
		}
	}

	if err = d.Set("skip_tunnel_fcp_req", flattenSystemDmSkipTunnelFcpReq(o["skip-tunnel-fcp-req"], d, "skip_tunnel_fcp_req")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-tunnel-fcp-req"], "SystemDm-SkipTunnelFcpReq"); ok {
			if err = d.Set("skip_tunnel_fcp_req", vv); err != nil {
				return fmt.Errorf("Error reading skip_tunnel_fcp_req: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_tunnel_fcp_req: %v", err)
		}
	}

	if err = d.Set("verify_install", flattenSystemDmVerifyInstall(o["verify-install"], d, "verify_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-install"], "SystemDm-VerifyInstall"); ok {
			if err = d.Set("verify_install", vv); err != nil {
				return fmt.Errorf("Error reading verify_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_install: %v", err)
		}
	}

	return nil
}

func flattenSystemDmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDmConcurrentInstallImageLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmConcurrentInstallLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmConcurrentInstallScriptLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmConfMergeAfterScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmDiscoverTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmDpmLogsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmFgfmInstallRefreshCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmFgfmSockTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmFgfmKeepaliveItvl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmForceRemoteDiff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmFortiapRefreshCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmFortiapRefreshItvl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmFortiextRefreshCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmInstallImageTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmInstallTunnelRetryItvl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmMaxRevs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmNrRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmRetryIntvl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmRollbackAllowReboot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmScriptLogsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmSkipScepCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmSkipTunnelFcpReq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDmVerifyInstall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("concurrent_install_image_limit"); ok {
		t, err := expandSystemDmConcurrentInstallImageLimit(d, v, "concurrent_install_image_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-install-image-limit"] = t
		}
	}

	if v, ok := d.GetOk("concurrent_install_limit"); ok {
		t, err := expandSystemDmConcurrentInstallLimit(d, v, "concurrent_install_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-install-limit"] = t
		}
	}

	if v, ok := d.GetOk("concurrent_install_script_limit"); ok {
		t, err := expandSystemDmConcurrentInstallScriptLimit(d, v, "concurrent_install_script_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-install-script-limit"] = t
		}
	}

	if v, ok := d.GetOk("conf_merge_after_script"); ok {
		t, err := expandSystemDmConfMergeAfterScript(d, v, "conf_merge_after_script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conf-merge-after-script"] = t
		}
	}

	if v, ok := d.GetOk("discover_timeout"); ok {
		t, err := expandSystemDmDiscoverTimeout(d, v, "discover_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discover-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dpm_logsize"); ok {
		t, err := expandSystemDmDpmLogsize(d, v, "dpm_logsize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpm-logsize"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_install_refresh_count"); ok {
		t, err := expandSystemDmFgfmInstallRefreshCount(d, v, "fgfm_install_refresh_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-install-refresh-count"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_sock_timeout"); ok {
		t, err := expandSystemDmFgfmSockTimeout(d, v, "fgfm_sock_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-sock-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_keepalive_itvl"); ok {
		t, err := expandSystemDmFgfmKeepaliveItvl(d, v, "fgfm_keepalive_itvl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm_keepalive_itvl"] = t
		}
	}

	if v, ok := d.GetOk("force_remote_diff"); ok {
		t, err := expandSystemDmForceRemoteDiff(d, v, "force_remote_diff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-remote-diff"] = t
		}
	}

	if v, ok := d.GetOk("fortiap_refresh_cnt"); ok {
		t, err := expandSystemDmFortiapRefreshCnt(d, v, "fortiap_refresh_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiap-refresh-cnt"] = t
		}
	}

	if v, ok := d.GetOk("fortiap_refresh_itvl"); ok {
		t, err := expandSystemDmFortiapRefreshItvl(d, v, "fortiap_refresh_itvl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiap-refresh-itvl"] = t
		}
	}

	if v, ok := d.GetOk("fortiext_refresh_cnt"); ok {
		t, err := expandSystemDmFortiextRefreshCnt(d, v, "fortiext_refresh_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiext-refresh-cnt"] = t
		}
	}

	if v, ok := d.GetOk("install_image_timeout"); ok {
		t, err := expandSystemDmInstallImageTimeout(d, v, "install_image_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["install-image-timeout"] = t
		}
	}

	if v, ok := d.GetOk("install_tunnel_retry_itvl"); ok {
		t, err := expandSystemDmInstallTunnelRetryItvl(d, v, "install_tunnel_retry_itvl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["install-tunnel-retry-itvl"] = t
		}
	}

	if v, ok := d.GetOk("max_revs"); ok {
		t, err := expandSystemDmMaxRevs(d, v, "max_revs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-revs"] = t
		}
	}

	if v, ok := d.GetOk("nr_retry"); ok {
		t, err := expandSystemDmNrRetry(d, v, "nr_retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nr-retry"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok {
		t, err := expandSystemDmRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOk("retry_intvl"); ok {
		t, err := expandSystemDmRetryIntvl(d, v, "retry_intvl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry-intvl"] = t
		}
	}

	if v, ok := d.GetOk("rollback_allow_reboot"); ok {
		t, err := expandSystemDmRollbackAllowReboot(d, v, "rollback_allow_reboot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rollback-allow-reboot"] = t
		}
	}

	if v, ok := d.GetOk("script_logsize"); ok {
		t, err := expandSystemDmScriptLogsize(d, v, "script_logsize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script-logsize"] = t
		}
	}

	if v, ok := d.GetOk("skip_scep_check"); ok {
		t, err := expandSystemDmSkipScepCheck(d, v, "skip_scep_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-scep-check"] = t
		}
	}

	if v, ok := d.GetOk("skip_tunnel_fcp_req"); ok {
		t, err := expandSystemDmSkipTunnelFcpReq(d, v, "skip_tunnel_fcp_req")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-tunnel-fcp-req"] = t
		}
	}

	if v, ok := d.GetOk("verify_install"); ok {
		t, err := expandSystemDmVerifyInstall(d, v, "verify_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-install"] = t
		}
	}

	return &obj, nil
}
