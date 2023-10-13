// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ACL classifiers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerAclIngressClassifier() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerAclIngressClassifierUpdate,
		Read:   resourceObjectSwitchControllerAclIngressClassifierRead,
		Update: resourceObjectSwitchControllerAclIngressClassifierUpdate,
		Delete: resourceObjectSwitchControllerAclIngressClassifierDelete,

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
			"ingress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dst_ip_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_ip_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerAclIngressClassifierUpdate(d *schema.ResourceData, m interface{}) error {
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

	ingress := d.Get("ingress").(string)
	paradict["ingress"] = ingress

	obj, err := getObjectObjectSwitchControllerAclIngressClassifier(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclIngressClassifier resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerAclIngressClassifier(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclIngressClassifier resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSwitchControllerAclIngressClassifier")

	return resourceObjectSwitchControllerAclIngressClassifierRead(d, m)
}

func resourceObjectSwitchControllerAclIngressClassifierDelete(d *schema.ResourceData, m interface{}) error {
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

	ingress := d.Get("ingress").(string)
	paradict["ingress"] = ingress

	err = c.DeleteObjectSwitchControllerAclIngressClassifier(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerAclIngressClassifier resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerAclIngressClassifierRead(d *schema.ResourceData, m interface{}) error {
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

	ingress := d.Get("ingress").(string)
	if ingress == "" {
		ingress = importOptionChecking(m.(*FortiClient).Cfg, "ingress")
		if ingress == "" {
			return fmt.Errorf("Parameter ingress is missing")
		}
		if err = d.Set("ingress", ingress); err != nil {
			return fmt.Errorf("Error set params ingress: %v", err)
		}
	}
	paradict["ingress"] = ingress

	o, err := c.ReadObjectSwitchControllerAclIngressClassifier(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclIngressClassifier resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerAclIngressClassifier(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclIngressClassifier resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerAclIngressClassifierDstIpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectSwitchControllerAclIngressClassifierDstMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressClassifierSrcIpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectSwitchControllerAclIngressClassifierSrcMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressClassifierVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerAclIngressClassifier(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dst_ip_prefix", flattenObjectSwitchControllerAclIngressClassifierDstIpPrefix2edl(o["dst-ip-prefix"], d, "dst_ip_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-ip-prefix"], "ObjectSwitchControllerAclIngressClassifier-DstIpPrefix"); ok {
			if err = d.Set("dst_ip_prefix", vv); err != nil {
				return fmt.Errorf("Error reading dst_ip_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_ip_prefix: %v", err)
		}
	}

	if err = d.Set("dst_mac", flattenObjectSwitchControllerAclIngressClassifierDstMac2edl(o["dst-mac"], d, "dst_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-mac"], "ObjectSwitchControllerAclIngressClassifier-DstMac"); ok {
			if err = d.Set("dst_mac", vv); err != nil {
				return fmt.Errorf("Error reading dst_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_mac: %v", err)
		}
	}

	if err = d.Set("src_ip_prefix", flattenObjectSwitchControllerAclIngressClassifierSrcIpPrefix2edl(o["src-ip-prefix"], d, "src_ip_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip-prefix"], "ObjectSwitchControllerAclIngressClassifier-SrcIpPrefix"); ok {
			if err = d.Set("src_ip_prefix", vv); err != nil {
				return fmt.Errorf("Error reading src_ip_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip_prefix: %v", err)
		}
	}

	if err = d.Set("src_mac", flattenObjectSwitchControllerAclIngressClassifierSrcMac2edl(o["src-mac"], d, "src_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-mac"], "ObjectSwitchControllerAclIngressClassifier-SrcMac"); ok {
			if err = d.Set("src_mac", vv); err != nil {
				return fmt.Errorf("Error reading src_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_mac: %v", err)
		}
	}

	if err = d.Set("vlan", flattenObjectSwitchControllerAclIngressClassifierVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "ObjectSwitchControllerAclIngressClassifier-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerAclIngressClassifierFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerAclIngressClassifierDstIpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierDstMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierSrcIpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierSrcMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerAclIngressClassifier(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_ip_prefix"); ok || d.HasChange("dst_ip_prefix") {
		t, err := expandObjectSwitchControllerAclIngressClassifierDstIpPrefix2edl(d, v, "dst_ip_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-ip-prefix"] = t
		}
	}

	if v, ok := d.GetOk("dst_mac"); ok || d.HasChange("dst_mac") {
		t, err := expandObjectSwitchControllerAclIngressClassifierDstMac2edl(d, v, "dst_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-mac"] = t
		}
	}

	if v, ok := d.GetOk("src_ip_prefix"); ok || d.HasChange("src_ip_prefix") {
		t, err := expandObjectSwitchControllerAclIngressClassifierSrcIpPrefix2edl(d, v, "src_ip_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip-prefix"] = t
		}
	}

	if v, ok := d.GetOk("src_mac"); ok || d.HasChange("src_mac") {
		t, err := expandObjectSwitchControllerAclIngressClassifierSrcMac2edl(d, v, "src_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-mac"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandObjectSwitchControllerAclIngressClassifierVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
