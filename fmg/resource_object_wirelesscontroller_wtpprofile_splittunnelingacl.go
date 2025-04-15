// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Split tunneling ACL filter list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWtpProfileSplitTunnelingAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWtpProfileSplitTunnelingAclCreate,
		Read:   resourceObjectWirelessControllerWtpProfileSplitTunnelingAclRead,
		Update: resourceObjectWirelessControllerWtpProfileSplitTunnelingAclUpdate,
		Delete: resourceObjectWirelessControllerWtpProfileSplitTunnelingAclDelete,

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
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dest_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWtpProfileSplitTunnelingAclCreate(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectObjectWirelessControllerWtpProfileSplitTunnelingAcl(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerWtpProfileSplitTunnelingAcl(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerWtpProfileSplitTunnelingAclRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileSplitTunnelingAclUpdate(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectObjectWirelessControllerWtpProfileSplitTunnelingAcl(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerWtpProfileSplitTunnelingAcl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerWtpProfileSplitTunnelingAclRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileSplitTunnelingAclDelete(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerWtpProfileSplitTunnelingAcl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWtpProfileSplitTunnelingAclRead(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadObjectWirelessControllerWtpProfileSplitTunnelingAcl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWtpProfileSplitTunnelingAcl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileSplitTunnelingAcl resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dest_ip", flattenObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(o["dest-ip"], d, "dest_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dest-ip"], "ObjectWirelessControllerWtpProfileSplitTunnelingAcl-DestIp"); ok {
			if err = d.Set("dest_ip", vv); err != nil {
				return fmt.Errorf("Error reading dest_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dest_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerWtpProfileSplitTunnelingAclId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerWtpProfileSplitTunnelingAcl-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAclId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dest_ip"); ok || d.HasChange("dest_ip") {
		t, err := expandObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(d, v, "dest_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerWtpProfileSplitTunnelingAclId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
