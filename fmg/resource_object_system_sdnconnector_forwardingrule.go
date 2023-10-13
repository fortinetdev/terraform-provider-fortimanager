// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure GCP forwarding rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorForwardingRuleCreate,
		Read:   resourceObjectSystemSdnConnectorForwardingRuleRead,
		Update: resourceObjectSystemSdnConnectorForwardingRuleUpdate,
		Delete: resourceObjectSystemSdnConnectorForwardingRuleDelete,

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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorForwardingRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectObjectSystemSdnConnectorForwardingRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorForwardingRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemSdnConnectorForwardingRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorForwardingRule resource: %v", err)
	}

	d.SetId(getStringKey(d, "rule_name"))

	return resourceObjectSystemSdnConnectorForwardingRuleRead(d, m)
}

func resourceObjectSystemSdnConnectorForwardingRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectObjectSystemSdnConnectorForwardingRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorForwardingRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemSdnConnectorForwardingRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorForwardingRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "rule_name"))

	return resourceObjectSystemSdnConnectorForwardingRuleRead(d, m)
}

func resourceObjectSystemSdnConnectorForwardingRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	err = c.DeleteObjectSystemSdnConnectorForwardingRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorForwardingRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorForwardingRuleRead(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	if sdn_connector == "" {
		sdn_connector = importOptionChecking(m.(*FortiClient).Cfg, "sdn_connector")
		if sdn_connector == "" {
			return fmt.Errorf("Parameter sdn_connector is missing")
		}
		if err = d.Set("sdn_connector", sdn_connector); err != nil {
			return fmt.Errorf("Error set params sdn_connector: %v", err)
		}
	}
	paradict["sdn_connector"] = sdn_connector

	o, err := c.ReadObjectSystemSdnConnectorForwardingRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorForwardingRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorForwardingRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorForwardingRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorForwardingRuleRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorForwardingRuleTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorForwardingRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("rule_name", flattenObjectSystemSdnConnectorForwardingRuleRuleName2edl(o["rule-name"], d, "rule_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule-name"], "ObjectSystemSdnConnectorForwardingRule-RuleName"); ok {
			if err = d.Set("rule_name", vv); err != nil {
				return fmt.Errorf("Error reading rule_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule_name: %v", err)
		}
	}

	if err = d.Set("target", flattenObjectSystemSdnConnectorForwardingRuleTarget2edl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "ObjectSystemSdnConnectorForwardingRule-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorForwardingRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorForwardingRuleRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorForwardingRuleTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorForwardingRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rule_name"); ok || d.HasChange("rule_name") {
		t, err := expandObjectSystemSdnConnectorForwardingRuleRuleName2edl(d, v, "rule_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-name"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectSystemSdnConnectorForwardingRuleTarget2edl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
