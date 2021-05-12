// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Authentication Rules.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesAuthenticationRuleCreate,
		Read:   resourcePackagesAuthenticationRuleRead,
		Update: resourcePackagesAuthenticationRuleUpdate,
		Delete: resourcePackagesAuthenticationRuleDelete,

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
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"active_auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transaction_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesAuthenticationRule resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesAuthenticationRule(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesAuthenticationRule resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesAuthenticationRuleRead(d, m)
}

func resourcePackagesAuthenticationRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesAuthenticationRule resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesAuthenticationRule(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesAuthenticationRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesAuthenticationRuleRead(d, m)
}

func resourcePackagesAuthenticationRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesAuthenticationRule(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesAuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesAuthenticationRule(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesAuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesAuthenticationRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesAuthenticationRule resource from API: %v", err)
	}
	return nil
}

func flattenPackagesAuthenticationRuleActiveAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleIpBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesAuthenticationRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			7:  "http",
			8:  "ftp",
			10: "socks",
			14: "ssh",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenPackagesAuthenticationRuleSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleSsoAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesAuthenticationRuleTransactionBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesAuthenticationRuleWebAuthCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesAuthenticationRuleWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectPackagesAuthenticationRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("active_auth_method", flattenPackagesAuthenticationRuleActiveAuthMethod(o["active-auth-method"], d, "active_auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-auth-method"], "PackagesAuthenticationRule-ActiveAuthMethod"); ok {
			if err = d.Set("active_auth_method", vv); err != nil {
				return fmt.Errorf("Error reading active_auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_auth_method: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesAuthenticationRuleComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesAuthenticationRule-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("ip_based", flattenPackagesAuthenticationRuleIpBased(o["ip-based"], d, "ip_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-based"], "PackagesAuthenticationRule-IpBased"); ok {
			if err = d.Set("ip_based", vv); err != nil {
				return fmt.Errorf("Error reading ip_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_based: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesAuthenticationRuleName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesAuthenticationRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenPackagesAuthenticationRuleProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "PackagesAuthenticationRule-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesAuthenticationRuleSrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesAuthenticationRule-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesAuthenticationRuleSrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesAuthenticationRule-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("sso_auth_method", flattenPackagesAuthenticationRuleSsoAuthMethod(o["sso-auth-method"], d, "sso_auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-auth-method"], "PackagesAuthenticationRule-SsoAuthMethod"); ok {
			if err = d.Set("sso_auth_method", vv); err != nil {
				return fmt.Errorf("Error reading sso_auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_auth_method: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesAuthenticationRuleStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesAuthenticationRule-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transaction_based", flattenPackagesAuthenticationRuleTransactionBased(o["transaction-based"], d, "transaction_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["transaction-based"], "PackagesAuthenticationRule-TransactionBased"); ok {
			if err = d.Set("transaction_based", vv); err != nil {
				return fmt.Errorf("Error reading transaction_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transaction_based: %v", err)
		}
	}

	if err = d.Set("web_auth_cookie", flattenPackagesAuthenticationRuleWebAuthCookie(o["web-auth-cookie"], d, "web_auth_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-auth-cookie"], "PackagesAuthenticationRule-WebAuthCookie"); ok {
			if err = d.Set("web_auth_cookie", vv); err != nil {
				return fmt.Errorf("Error reading web_auth_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_auth_cookie: %v", err)
		}
	}

	if err = d.Set("web_portal", flattenPackagesAuthenticationRuleWebPortal(o["web-portal"], d, "web_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-portal"], "PackagesAuthenticationRule-WebPortal"); ok {
			if err = d.Set("web_portal", vv); err != nil {
				return fmt.Errorf("Error reading web_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_portal: %v", err)
		}
	}

	return nil
}

func flattenPackagesAuthenticationRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesAuthenticationRuleActiveAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleIpBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleSsoAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleTransactionBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleWebAuthCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleWebPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesAuthenticationRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("active_auth_method"); ok {
		t, err := expandPackagesAuthenticationRuleActiveAuthMethod(d, v, "active_auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-method"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesAuthenticationRuleComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("ip_based"); ok {
		t, err := expandPackagesAuthenticationRuleIpBased(d, v, "ip_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-based"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesAuthenticationRuleName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandPackagesAuthenticationRuleProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesAuthenticationRuleSrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok {
		t, err := expandPackagesAuthenticationRuleSrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("sso_auth_method"); ok {
		t, err := expandPackagesAuthenticationRuleSsoAuthMethod(d, v, "sso_auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-method"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesAuthenticationRuleStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transaction_based"); ok {
		t, err := expandPackagesAuthenticationRuleTransactionBased(d, v, "transaction_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transaction-based"] = t
		}
	}

	if v, ok := d.GetOk("web_auth_cookie"); ok {
		t, err := expandPackagesAuthenticationRuleWebAuthCookie(d, v, "web_auth_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("web_portal"); ok {
		t, err := expandPackagesAuthenticationRuleWebPortal(d, v, "web_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-portal"] = t
		}
	}

	return &obj, nil
}
