// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Authentication Rules.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"pkg_folder_path": &schema.Schema{
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
			},
			"cert_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cors_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cors_stateful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sso_auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transaction_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePackagesAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesAuthenticationRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreatePackagesAuthenticationRule(obj, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesAuthenticationRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesAuthenticationRule(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	wsParams["adom"] = adomv

	err = c.DeletePackagesAuthenticationRule(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
	}
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if pkg == "" {
			return fmt.Errorf("Parameter pkg is missing")
		}
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	o, err := c.ReadPackagesAuthenticationRule(mkey, paradict)
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
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesAuthenticationRuleCertAuthCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleCorsDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleCorsStateful(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesAuthenticationRuleDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesAuthenticationRuleIpBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesAuthenticationRuleSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesAuthenticationRuleSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesAuthenticationRuleSsoAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesAuthenticationRuleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleTransactionBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleWebAuthCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationRuleWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesAuthenticationRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("active_auth_method", flattenPackagesAuthenticationRuleActiveAuthMethod(o["active-auth-method"], d, "active_auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-auth-method"], "PackagesAuthenticationRule-ActiveAuthMethod"); ok {
			if err = d.Set("active_auth_method", vv); err != nil {
				return fmt.Errorf("Error reading active_auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_auth_method: %v", err)
		}
	}

	if err = d.Set("cert_auth_cookie", flattenPackagesAuthenticationRuleCertAuthCookie(o["cert-auth-cookie"], d, "cert_auth_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-auth-cookie"], "PackagesAuthenticationRule-CertAuthCookie"); ok {
			if err = d.Set("cert_auth_cookie", vv); err != nil {
				return fmt.Errorf("Error reading cert_auth_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_auth_cookie: %v", err)
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

	if err = d.Set("cors_depth", flattenPackagesAuthenticationRuleCorsDepth(o["cors-depth"], d, "cors_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["cors-depth"], "PackagesAuthenticationRule-CorsDepth"); ok {
			if err = d.Set("cors_depth", vv); err != nil {
				return fmt.Errorf("Error reading cors_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cors_depth: %v", err)
		}
	}

	if err = d.Set("cors_stateful", flattenPackagesAuthenticationRuleCorsStateful(o["cors-stateful"], d, "cors_stateful")); err != nil {
		if vv, ok := fortiAPIPatch(o["cors-stateful"], "PackagesAuthenticationRule-CorsStateful"); ok {
			if err = d.Set("cors_stateful", vv); err != nil {
				return fmt.Errorf("Error reading cors_stateful: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cors_stateful: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesAuthenticationRuleDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesAuthenticationRule-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenPackagesAuthenticationRuleDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "PackagesAuthenticationRule-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
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

	if err = d.Set("srcintf", flattenPackagesAuthenticationRuleSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesAuthenticationRule-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
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
	return convstr2list(v, nil), nil
}

func expandPackagesAuthenticationRuleCertAuthCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleCorsDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleCorsStateful(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationRuleDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesAuthenticationRuleDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
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
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesAuthenticationRuleSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesAuthenticationRuleSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesAuthenticationRuleSsoAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
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

	if v, ok := d.GetOk("active_auth_method"); ok || d.HasChange("active_auth_method") {
		t, err := expandPackagesAuthenticationRuleActiveAuthMethod(d, v, "active_auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-method"] = t
		}
	}

	if v, ok := d.GetOk("cert_auth_cookie"); ok || d.HasChange("cert_auth_cookie") {
		t, err := expandPackagesAuthenticationRuleCertAuthCookie(d, v, "cert_auth_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesAuthenticationRuleComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("cors_depth"); ok || d.HasChange("cors_depth") {
		t, err := expandPackagesAuthenticationRuleCorsDepth(d, v, "cors_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-depth"] = t
		}
	}

	if v, ok := d.GetOk("cors_stateful"); ok || d.HasChange("cors_stateful") {
		t, err := expandPackagesAuthenticationRuleCorsStateful(d, v, "cors_stateful")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-stateful"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandPackagesAuthenticationRuleDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandPackagesAuthenticationRuleDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("ip_based"); ok || d.HasChange("ip_based") {
		t, err := expandPackagesAuthenticationRuleIpBased(d, v, "ip_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-based"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesAuthenticationRuleName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandPackagesAuthenticationRuleProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandPackagesAuthenticationRuleSrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandPackagesAuthenticationRuleSrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandPackagesAuthenticationRuleSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("sso_auth_method"); ok || d.HasChange("sso_auth_method") {
		t, err := expandPackagesAuthenticationRuleSsoAuthMethod(d, v, "sso_auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-method"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesAuthenticationRuleStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transaction_based"); ok || d.HasChange("transaction_based") {
		t, err := expandPackagesAuthenticationRuleTransactionBased(d, v, "transaction_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transaction-based"] = t
		}
	}

	if v, ok := d.GetOk("web_auth_cookie"); ok || d.HasChange("web_auth_cookie") {
		t, err := expandPackagesAuthenticationRuleWebAuthCookie(d, v, "web_auth_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("web_portal"); ok || d.HasChange("web_portal") {
		t, err := expandPackagesAuthenticationRuleWebPortal(d, v, "web_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-portal"] = t
		}
	}

	return &obj, nil
}
