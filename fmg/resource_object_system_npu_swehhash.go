// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure switch enhanced hashing.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuSwEhHash() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuSwEhHashUpdate,
		Read:   resourceObjectSystemNpuSwEhHashRead,
		Update: resourceObjectSystemNpuSwEhHashUpdate,
		Delete: resourceObjectSystemNpuSwEhHashDelete,

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
			"computation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_ip_lower_16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_ip_upper_16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip_lower_16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_upper_16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuSwEhHashUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuSwEhHash(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuSwEhHash resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuSwEhHash(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuSwEhHash resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuSwEhHash")

	return resourceObjectSystemNpuSwEhHashRead(d, m)
}

func resourceObjectSystemNpuSwEhHashDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuSwEhHash(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuSwEhHash resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuSwEhHashRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuSwEhHash(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuSwEhHash resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuSwEhHash(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuSwEhHash resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuSwEhHashComputation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationIpLower162edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationIpUpper162edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashIpProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashNetmaskLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourceIpLower162edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourceIpUpper162edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourcePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuSwEhHash(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("computation", flattenObjectSystemNpuSwEhHashComputation2edl(o["computation"], d, "computation")); err != nil {
		if vv, ok := fortiAPIPatch(o["computation"], "ObjectSystemNpuSwEhHash-Computation"); ok {
			if err = d.Set("computation", vv); err != nil {
				return fmt.Errorf("Error reading computation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading computation: %v", err)
		}
	}

	if err = d.Set("destination_ip_lower_16", flattenObjectSystemNpuSwEhHashDestinationIpLower162edl(o["destination-ip-lower-16"], d, "destination_ip_lower_16")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-ip-lower-16"], "ObjectSystemNpuSwEhHash-DestinationIpLower16"); ok {
			if err = d.Set("destination_ip_lower_16", vv); err != nil {
				return fmt.Errorf("Error reading destination_ip_lower_16: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_ip_lower_16: %v", err)
		}
	}

	if err = d.Set("destination_ip_upper_16", flattenObjectSystemNpuSwEhHashDestinationIpUpper162edl(o["destination-ip-upper-16"], d, "destination_ip_upper_16")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-ip-upper-16"], "ObjectSystemNpuSwEhHash-DestinationIpUpper16"); ok {
			if err = d.Set("destination_ip_upper_16", vv); err != nil {
				return fmt.Errorf("Error reading destination_ip_upper_16: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_ip_upper_16: %v", err)
		}
	}

	if err = d.Set("destination_port", flattenObjectSystemNpuSwEhHashDestinationPort2edl(o["destination-port"], d, "destination_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-port"], "ObjectSystemNpuSwEhHash-DestinationPort"); ok {
			if err = d.Set("destination_port", vv); err != nil {
				return fmt.Errorf("Error reading destination_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_port: %v", err)
		}
	}

	if err = d.Set("ip_protocol", flattenObjectSystemNpuSwEhHashIpProtocol2edl(o["ip-protocol"], d, "ip_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-protocol"], "ObjectSystemNpuSwEhHash-IpProtocol"); ok {
			if err = d.Set("ip_protocol", vv); err != nil {
				return fmt.Errorf("Error reading ip_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_protocol: %v", err)
		}
	}

	if err = d.Set("netmask_length", flattenObjectSystemNpuSwEhHashNetmaskLength2edl(o["netmask-length"], d, "netmask_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask-length"], "ObjectSystemNpuSwEhHash-NetmaskLength"); ok {
			if err = d.Set("netmask_length", vv); err != nil {
				return fmt.Errorf("Error reading netmask_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask_length: %v", err)
		}
	}

	if err = d.Set("source_ip_lower_16", flattenObjectSystemNpuSwEhHashSourceIpLower162edl(o["source-ip-lower-16"], d, "source_ip_lower_16")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-lower-16"], "ObjectSystemNpuSwEhHash-SourceIpLower16"); ok {
			if err = d.Set("source_ip_lower_16", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_lower_16: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_lower_16: %v", err)
		}
	}

	if err = d.Set("source_ip_upper_16", flattenObjectSystemNpuSwEhHashSourceIpUpper162edl(o["source-ip-upper-16"], d, "source_ip_upper_16")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-upper-16"], "ObjectSystemNpuSwEhHash-SourceIpUpper16"); ok {
			if err = d.Set("source_ip_upper_16", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_upper_16: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_upper_16: %v", err)
		}
	}

	if err = d.Set("source_port", flattenObjectSystemNpuSwEhHashSourcePort2edl(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "ObjectSystemNpuSwEhHash-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuSwEhHashFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuSwEhHashComputation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationIpLower162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationIpUpper162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashIpProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashNetmaskLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourceIpLower162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourceIpUpper162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourcePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuSwEhHash(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("computation"); ok || d.HasChange("computation") {
		t, err := expandObjectSystemNpuSwEhHashComputation2edl(d, v, "computation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["computation"] = t
		}
	}

	if v, ok := d.GetOk("destination_ip_lower_16"); ok || d.HasChange("destination_ip_lower_16") {
		t, err := expandObjectSystemNpuSwEhHashDestinationIpLower162edl(d, v, "destination_ip_lower_16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-ip-lower-16"] = t
		}
	}

	if v, ok := d.GetOk("destination_ip_upper_16"); ok || d.HasChange("destination_ip_upper_16") {
		t, err := expandObjectSystemNpuSwEhHashDestinationIpUpper162edl(d, v, "destination_ip_upper_16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-ip-upper-16"] = t
		}
	}

	if v, ok := d.GetOk("destination_port"); ok || d.HasChange("destination_port") {
		t, err := expandObjectSystemNpuSwEhHashDestinationPort2edl(d, v, "destination_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-port"] = t
		}
	}

	if v, ok := d.GetOk("ip_protocol"); ok || d.HasChange("ip_protocol") {
		t, err := expandObjectSystemNpuSwEhHashIpProtocol2edl(d, v, "ip_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-protocol"] = t
		}
	}

	if v, ok := d.GetOk("netmask_length"); ok || d.HasChange("netmask_length") {
		t, err := expandObjectSystemNpuSwEhHashNetmaskLength2edl(d, v, "netmask_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask-length"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_lower_16"); ok || d.HasChange("source_ip_lower_16") {
		t, err := expandObjectSystemNpuSwEhHashSourceIpLower162edl(d, v, "source_ip_lower_16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-lower-16"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_upper_16"); ok || d.HasChange("source_ip_upper_16") {
		t, err := expandObjectSystemNpuSwEhHashSourceIpUpper162edl(d, v, "source_ip_upper_16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-upper-16"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandObjectSystemNpuSwEhHashSourcePort2edl(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	return &obj, nil
}
