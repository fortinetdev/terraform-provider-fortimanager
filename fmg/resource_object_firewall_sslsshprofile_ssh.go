// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSH options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallSslSshProfileSsh() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallSslSshProfileSshUpdate,
		Read:   resourceObjectFirewallSslSshProfileSshRead,
		Update: resourceObjectFirewallSslSshProfileSshUpdate,
		Delete: resourceObjectFirewallSslSshProfileSshDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inspect_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"proxy_after_tcp_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_policy_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_tun_policy_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallSslSshProfileSshUpdate(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectObjectFirewallSslSshProfileSsh(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileSsh resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallSslSshProfileSsh(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileSsh resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallSslSshProfileSsh")

	return resourceObjectFirewallSslSshProfileSshRead(d, m)
}

func resourceObjectFirewallSslSshProfileSshDelete(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	err = c.DeleteObjectFirewallSslSshProfileSsh(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallSslSshProfileSsh resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallSslSshProfileSshRead(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	if ssl_ssh_profile == "" {
		ssl_ssh_profile = importOptionChecking(m.(*FortiClient).Cfg, "ssl_ssh_profile")
		if ssl_ssh_profile == "" {
			return fmt.Errorf("Parameter ssl_ssh_profile is missing")
		}
		if err = d.Set("ssl_ssh_profile", ssl_ssh_profile); err != nil {
			return fmt.Errorf("Error set params ssl_ssh_profile: %v", err)
		}
	}
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	o, err := c.ReadObjectFirewallSslSshProfileSsh(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileSsh resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallSslSshProfileSsh(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileSsh resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallSslSshProfileSshInspectAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshSshAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshSshPolicyCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshSshTunPolicyCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshUnsupportedVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallSslSshProfileSsh(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("inspect_all", flattenObjectFirewallSslSshProfileSshInspectAll2edl(o["inspect-all"], d, "inspect_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspect-all"], "ObjectFirewallSslSshProfileSsh-InspectAll"); ok {
			if err = d.Set("inspect_all", vv); err != nil {
				return fmt.Errorf("Error reading inspect_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspect_all: %v", err)
		}
	}

	if err = d.Set("ports", flattenObjectFirewallSslSshProfileSshPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "ObjectFirewallSslSshProfileSsh-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("proxy_after_tcp_handshake", flattenObjectFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(o["proxy-after-tcp-handshake"], d, "proxy_after_tcp_handshake")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-after-tcp-handshake"], "ObjectFirewallSslSshProfileSsh-ProxyAfterTcpHandshake"); ok {
			if err = d.Set("proxy_after_tcp_handshake", vv); err != nil {
				return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
		}
	}

	if err = d.Set("ssh_algorithm", flattenObjectFirewallSslSshProfileSshSshAlgorithm2edl(o["ssh-algorithm"], d, "ssh_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-algorithm"], "ObjectFirewallSslSshProfileSsh-SshAlgorithm"); ok {
			if err = d.Set("ssh_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssh_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_algorithm: %v", err)
		}
	}

	if err = d.Set("ssh_policy_check", flattenObjectFirewallSslSshProfileSshSshPolicyCheck2edl(o["ssh-policy-check"], d, "ssh_policy_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-policy-check"], "ObjectFirewallSslSshProfileSsh-SshPolicyCheck"); ok {
			if err = d.Set("ssh_policy_check", vv); err != nil {
				return fmt.Errorf("Error reading ssh_policy_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_policy_check: %v", err)
		}
	}

	if err = d.Set("ssh_tun_policy_check", flattenObjectFirewallSslSshProfileSshSshTunPolicyCheck2edl(o["ssh-tun-policy-check"], d, "ssh_tun_policy_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-tun-policy-check"], "ObjectFirewallSslSshProfileSsh-SshTunPolicyCheck"); ok {
			if err = d.Set("ssh_tun_policy_check", vv); err != nil {
				return fmt.Errorf("Error reading ssh_tun_policy_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_tun_policy_check: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallSslSshProfileSshStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallSslSshProfileSsh-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("unsupported_version", flattenObjectFirewallSslSshProfileSshUnsupportedVersion2edl(o["unsupported-version"], d, "unsupported_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-version"], "ObjectFirewallSslSshProfileSsh-UnsupportedVersion"); ok {
			if err = d.Set("unsupported_version", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_version: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallSslSshProfileSshFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallSslSshProfileSshInspectAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshSshAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshSshPolicyCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshSshTunPolicyCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshUnsupportedVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallSslSshProfileSsh(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("inspect_all"); ok || d.HasChange("inspect_all") {
		t, err := expandObjectFirewallSslSshProfileSshInspectAll2edl(d, v, "inspect_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspect-all"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandObjectFirewallSslSshProfileSshPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("proxy_after_tcp_handshake"); ok || d.HasChange("proxy_after_tcp_handshake") {
		t, err := expandObjectFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(d, v, "proxy_after_tcp_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-after-tcp-handshake"] = t
		}
	}

	if v, ok := d.GetOk("ssh_algorithm"); ok || d.HasChange("ssh_algorithm") {
		t, err := expandObjectFirewallSslSshProfileSshSshAlgorithm2edl(d, v, "ssh_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_check"); ok || d.HasChange("ssh_policy_check") {
		t, err := expandObjectFirewallSslSshProfileSshSshPolicyCheck2edl(d, v, "ssh_policy_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-check"] = t
		}
	}

	if v, ok := d.GetOk("ssh_tun_policy_check"); ok || d.HasChange("ssh_tun_policy_check") {
		t, err := expandObjectFirewallSslSshProfileSshSshTunPolicyCheck2edl(d, v, "ssh_tun_policy_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-tun-policy-check"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallSslSshProfileSshStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_version"); ok || d.HasChange("unsupported_version") {
		t, err := expandObjectFirewallSslSshProfileSshUnsupportedVersion2edl(d, v, "unsupported_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-version"] = t
		}
	}

	return &obj, nil
}
