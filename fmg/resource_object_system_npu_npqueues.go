// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure queue assignment on NP7.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpQueues() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpQueuesUpdate,
		Read:   resourceObjectSystemNpuNpQueuesRead,
		Update: resourceObjectSystemNpuNpQueuesUpdate,
		Delete: resourceObjectSystemNpuNpQueuesDelete,

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
			"ethernet_type": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip_protocol": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp10": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp11": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp12": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp13": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp14": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp15": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp17": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp18": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp19": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp20": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp21": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp22": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp23": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp24": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp25": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp26": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp27": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp28": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp29": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp30": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp31": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp32": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp33": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp34": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp35": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp36": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp37": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp38": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp39": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp40": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp41": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp42": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp43": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp44": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp45": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp47": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp48": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp49": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp50": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp51": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp52": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp53": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp54": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp55": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp56": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp57": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp58": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp59": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp60": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp61": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp62": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp63": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp9": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"scheduler": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectSystemNpuNpQueuesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpQueues(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueues resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpQueues(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueues resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpQueues")

	return resourceObjectSystemNpuNpQueuesRead(d, m)
}

func resourceObjectSystemNpuNpQueuesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpQueues(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpQueues resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpQueuesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpQueues(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueues resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpQueues(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueues resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpQueuesEthernetTypeOsnna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeNameOsnna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeQueueOsnna(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeTypeOsnna(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeWeightOsnna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesEthernetTypeNameOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeQueueOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeTypeOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeWeightOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolOsnna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolNameOsnna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolProtocolOsnna(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolQueueOsnna(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolWeightOsnna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpProtocolNameOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolProtocolOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolQueueOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolWeightOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceOsnna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := i["dport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceDportOsnna(i["dport"], d, pre_append)
			tmp["dport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Dport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceNameOsnna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceProtocolOsnna(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceQueueOsnna(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := i["sport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceSportOsnna(i["sport"], d, pre_append)
			tmp["sport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Sport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceWeightOsnna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpServiceDportOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceNameOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceProtocolOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceQueueOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceSportOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceWeightOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileOsnna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := i["cos0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos0Osnna(i["cos0"], d, pre_append)
			tmp["cos0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := i["cos1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos1Osnna(i["cos1"], d, pre_append)
			tmp["cos1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := i["cos2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos2Osnna(i["cos2"], d, pre_append)
			tmp["cos2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := i["cos3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos3Osnna(i["cos3"], d, pre_append)
			tmp["cos3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := i["cos4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos4Osnna(i["cos4"], d, pre_append)
			tmp["cos4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := i["cos5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos5Osnna(i["cos5"], d, pre_append)
			tmp["cos5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := i["cos6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos6Osnna(i["cos6"], d, pre_append)
			tmp["cos6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := i["cos7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos7Osnna(i["cos7"], d, pre_append)
			tmp["cos7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := i["dscp0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp0Osnna(i["dscp0"], d, pre_append)
			tmp["dscp0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := i["dscp1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp1Osnna(i["dscp1"], d, pre_append)
			tmp["dscp1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := i["dscp10"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp10Osnna(i["dscp10"], d, pre_append)
			tmp["dscp10"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp10")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := i["dscp11"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp11Osnna(i["dscp11"], d, pre_append)
			tmp["dscp11"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp11")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := i["dscp12"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp12Osnna(i["dscp12"], d, pre_append)
			tmp["dscp12"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := i["dscp13"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp13Osnna(i["dscp13"], d, pre_append)
			tmp["dscp13"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp13")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := i["dscp14"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp14Osnna(i["dscp14"], d, pre_append)
			tmp["dscp14"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp14")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := i["dscp15"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp15Osnna(i["dscp15"], d, pre_append)
			tmp["dscp15"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp15")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := i["dscp16"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp16Osnna(i["dscp16"], d, pre_append)
			tmp["dscp16"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp16")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := i["dscp17"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp17Osnna(i["dscp17"], d, pre_append)
			tmp["dscp17"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp17")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := i["dscp18"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp18Osnna(i["dscp18"], d, pre_append)
			tmp["dscp18"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp18")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := i["dscp19"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp19Osnna(i["dscp19"], d, pre_append)
			tmp["dscp19"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp19")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := i["dscp2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp2Osnna(i["dscp2"], d, pre_append)
			tmp["dscp2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := i["dscp20"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp20Osnna(i["dscp20"], d, pre_append)
			tmp["dscp20"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp20")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := i["dscp21"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp21Osnna(i["dscp21"], d, pre_append)
			tmp["dscp21"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp21")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := i["dscp22"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp22Osnna(i["dscp22"], d, pre_append)
			tmp["dscp22"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp22")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := i["dscp23"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp23Osnna(i["dscp23"], d, pre_append)
			tmp["dscp23"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp23")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := i["dscp24"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp24Osnna(i["dscp24"], d, pre_append)
			tmp["dscp24"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp24")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := i["dscp25"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp25Osnna(i["dscp25"], d, pre_append)
			tmp["dscp25"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp25")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := i["dscp26"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp26Osnna(i["dscp26"], d, pre_append)
			tmp["dscp26"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp26")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := i["dscp27"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp27Osnna(i["dscp27"], d, pre_append)
			tmp["dscp27"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp27")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := i["dscp28"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp28Osnna(i["dscp28"], d, pre_append)
			tmp["dscp28"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp28")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := i["dscp29"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp29Osnna(i["dscp29"], d, pre_append)
			tmp["dscp29"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp29")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := i["dscp3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp3Osnna(i["dscp3"], d, pre_append)
			tmp["dscp3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := i["dscp30"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp30Osnna(i["dscp30"], d, pre_append)
			tmp["dscp30"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp30")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := i["dscp31"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp31Osnna(i["dscp31"], d, pre_append)
			tmp["dscp31"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp31")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := i["dscp32"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp32Osnna(i["dscp32"], d, pre_append)
			tmp["dscp32"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp32")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := i["dscp33"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp33Osnna(i["dscp33"], d, pre_append)
			tmp["dscp33"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp33")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := i["dscp34"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp34Osnna(i["dscp34"], d, pre_append)
			tmp["dscp34"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := i["dscp35"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp35Osnna(i["dscp35"], d, pre_append)
			tmp["dscp35"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp35")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := i["dscp36"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp36Osnna(i["dscp36"], d, pre_append)
			tmp["dscp36"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp36")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := i["dscp37"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp37Osnna(i["dscp37"], d, pre_append)
			tmp["dscp37"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp37")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := i["dscp38"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp38Osnna(i["dscp38"], d, pre_append)
			tmp["dscp38"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp38")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := i["dscp39"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp39Osnna(i["dscp39"], d, pre_append)
			tmp["dscp39"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp39")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := i["dscp4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp4Osnna(i["dscp4"], d, pre_append)
			tmp["dscp4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := i["dscp40"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp40Osnna(i["dscp40"], d, pre_append)
			tmp["dscp40"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp40")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := i["dscp41"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp41Osnna(i["dscp41"], d, pre_append)
			tmp["dscp41"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp41")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := i["dscp42"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp42Osnna(i["dscp42"], d, pre_append)
			tmp["dscp42"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp42")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := i["dscp43"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp43Osnna(i["dscp43"], d, pre_append)
			tmp["dscp43"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp43")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := i["dscp44"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp44Osnna(i["dscp44"], d, pre_append)
			tmp["dscp44"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := i["dscp45"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp45Osnna(i["dscp45"], d, pre_append)
			tmp["dscp45"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp45")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := i["dscp46"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp46Osnna(i["dscp46"], d, pre_append)
			tmp["dscp46"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := i["dscp47"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp47Osnna(i["dscp47"], d, pre_append)
			tmp["dscp47"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp47")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := i["dscp48"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp48Osnna(i["dscp48"], d, pre_append)
			tmp["dscp48"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp48")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := i["dscp49"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp49Osnna(i["dscp49"], d, pre_append)
			tmp["dscp49"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp49")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := i["dscp5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp5Osnna(i["dscp5"], d, pre_append)
			tmp["dscp5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := i["dscp50"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp50Osnna(i["dscp50"], d, pre_append)
			tmp["dscp50"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp50")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := i["dscp51"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp51Osnna(i["dscp51"], d, pre_append)
			tmp["dscp51"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp51")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := i["dscp52"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp52Osnna(i["dscp52"], d, pre_append)
			tmp["dscp52"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp52")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := i["dscp53"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp53Osnna(i["dscp53"], d, pre_append)
			tmp["dscp53"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp53")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := i["dscp54"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp54Osnna(i["dscp54"], d, pre_append)
			tmp["dscp54"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp54")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := i["dscp55"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp55Osnna(i["dscp55"], d, pre_append)
			tmp["dscp55"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp55")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := i["dscp56"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp56Osnna(i["dscp56"], d, pre_append)
			tmp["dscp56"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp56")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := i["dscp57"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp57Osnna(i["dscp57"], d, pre_append)
			tmp["dscp57"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp57")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := i["dscp58"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp58Osnna(i["dscp58"], d, pre_append)
			tmp["dscp58"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp58")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := i["dscp59"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp59Osnna(i["dscp59"], d, pre_append)
			tmp["dscp59"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp59")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := i["dscp6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp6Osnna(i["dscp6"], d, pre_append)
			tmp["dscp6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := i["dscp60"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp60Osnna(i["dscp60"], d, pre_append)
			tmp["dscp60"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp60")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := i["dscp61"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp61Osnna(i["dscp61"], d, pre_append)
			tmp["dscp61"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp61")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := i["dscp62"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp62Osnna(i["dscp62"], d, pre_append)
			tmp["dscp62"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp62")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := i["dscp63"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp63Osnna(i["dscp63"], d, pre_append)
			tmp["dscp63"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp63")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := i["dscp7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp7Osnna(i["dscp7"], d, pre_append)
			tmp["dscp7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := i["dscp8"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp8Osnna(i["dscp8"], d, pre_append)
			tmp["dscp8"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp8")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := i["dscp9"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp9Osnna(i["dscp9"], d, pre_append)
			tmp["dscp9"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp9")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileIdOsnna(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileTypeOsnna(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileWeightOsnna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesProfileCos0Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos1Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos2Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos3Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos4Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos5Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos6Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos7Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp0Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp1Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp10Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp11Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp12Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp13Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp14Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp15Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp16Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp17Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp18Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp19Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp2Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp20Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp21Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp22Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp23Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp24Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp25Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp26Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp27Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp28Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp29Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp3Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp30Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp31Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp32Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp33Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp34Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp35Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp36Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp37Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp38Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp39Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp4Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp40Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp41Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp42Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp43Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp44Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp45Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp46Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp47Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp48Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp49Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp5Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp50Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp51Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp52Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp53Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp54Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp55Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp56Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp57Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp58Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp59Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp6Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp60Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp61Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp62Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp63Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp7Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp8Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp9Osnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileIdOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileTypeOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileWeightOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesSchedulerOsnna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerModeOsnna(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerNameOsnna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesSchedulerModeOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesSchedulerNameOsnna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpQueues(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("ethernet_type", flattenObjectSystemNpuNpQueuesEthernetTypeOsnna(o["ethernet-type"], d, "ethernet_type")); err != nil {
			if vv, ok := fortiAPIPatch(o["ethernet-type"], "ObjectSystemNpuNpQueues-EthernetType"); ok {
				if err = d.Set("ethernet_type", vv); err != nil {
					return fmt.Errorf("Error reading ethernet_type: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ethernet_type: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ethernet_type"); ok {
			if err = d.Set("ethernet_type", flattenObjectSystemNpuNpQueuesEthernetTypeOsnna(o["ethernet-type"], d, "ethernet_type")); err != nil {
				if vv, ok := fortiAPIPatch(o["ethernet-type"], "ObjectSystemNpuNpQueues-EthernetType"); ok {
					if err = d.Set("ethernet_type", vv); err != nil {
						return fmt.Errorf("Error reading ethernet_type: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ethernet_type: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_protocol", flattenObjectSystemNpuNpQueuesIpProtocolOsnna(o["ip-protocol"], d, "ip_protocol")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-protocol"], "ObjectSystemNpuNpQueues-IpProtocol"); ok {
				if err = d.Set("ip_protocol", vv); err != nil {
					return fmt.Errorf("Error reading ip_protocol: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_protocol: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_protocol"); ok {
			if err = d.Set("ip_protocol", flattenObjectSystemNpuNpQueuesIpProtocolOsnna(o["ip-protocol"], d, "ip_protocol")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-protocol"], "ObjectSystemNpuNpQueues-IpProtocol"); ok {
					if err = d.Set("ip_protocol", vv); err != nil {
						return fmt.Errorf("Error reading ip_protocol: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_protocol: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_service", flattenObjectSystemNpuNpQueuesIpServiceOsnna(o["ip-service"], d, "ip_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-service"], "ObjectSystemNpuNpQueues-IpService"); ok {
				if err = d.Set("ip_service", vv); err != nil {
					return fmt.Errorf("Error reading ip_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_service"); ok {
			if err = d.Set("ip_service", flattenObjectSystemNpuNpQueuesIpServiceOsnna(o["ip-service"], d, "ip_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-service"], "ObjectSystemNpuNpQueues-IpService"); ok {
					if err = d.Set("ip_service", vv); err != nil {
						return fmt.Errorf("Error reading ip_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("profile", flattenObjectSystemNpuNpQueuesProfileOsnna(o["profile"], d, "profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["profile"], "ObjectSystemNpuNpQueues-Profile"); ok {
				if err = d.Set("profile", vv); err != nil {
					return fmt.Errorf("Error reading profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("profile"); ok {
			if err = d.Set("profile", flattenObjectSystemNpuNpQueuesProfileOsnna(o["profile"], d, "profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["profile"], "ObjectSystemNpuNpQueues-Profile"); ok {
					if err = d.Set("profile", vv); err != nil {
						return fmt.Errorf("Error reading profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading profile: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("scheduler", flattenObjectSystemNpuNpQueuesSchedulerOsnna(o["scheduler"], d, "scheduler")); err != nil {
			if vv, ok := fortiAPIPatch(o["scheduler"], "ObjectSystemNpuNpQueues-Scheduler"); ok {
				if err = d.Set("scheduler", vv); err != nil {
					return fmt.Errorf("Error reading scheduler: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scheduler: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scheduler"); ok {
			if err = d.Set("scheduler", flattenObjectSystemNpuNpQueuesSchedulerOsnna(o["scheduler"], d, "scheduler")); err != nil {
				if vv, ok := fortiAPIPatch(o["scheduler"], "ObjectSystemNpuNpQueues-Scheduler"); ok {
					if err = d.Set("scheduler", vv); err != nil {
						return fmt.Errorf("Error reading scheduler: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scheduler: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectSystemNpuNpQueuesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpQueuesEthernetTypeOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesEthernetTypeNameOsnna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesEthernetTypeQueueOsnna(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesEthernetTypeTypeOsnna(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesEthernetTypeWeightOsnna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeNameOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeQueueOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeTypeOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeWeightOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpProtocolNameOsnna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpProtocolProtocolOsnna(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpProtocolQueueOsnna(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpProtocolWeightOsnna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolNameOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolProtocolOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolQueueOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolWeightOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dport"], _ = expandObjectSystemNpuNpQueuesIpServiceDportOsnna(d, i["dport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpServiceNameOsnna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpServiceProtocolOsnna(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpServiceQueueOsnna(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sport"], _ = expandObjectSystemNpuNpQueuesIpServiceSportOsnna(d, i["sport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpServiceWeightOsnna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpServiceDportOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceNameOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceProtocolOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceQueueOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceSportOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceWeightOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos0"], _ = expandObjectSystemNpuNpQueuesProfileCos0Osnna(d, i["cos0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos1"], _ = expandObjectSystemNpuNpQueuesProfileCos1Osnna(d, i["cos1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos2"], _ = expandObjectSystemNpuNpQueuesProfileCos2Osnna(d, i["cos2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos3"], _ = expandObjectSystemNpuNpQueuesProfileCos3Osnna(d, i["cos3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos4"], _ = expandObjectSystemNpuNpQueuesProfileCos4Osnna(d, i["cos4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos5"], _ = expandObjectSystemNpuNpQueuesProfileCos5Osnna(d, i["cos5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos6"], _ = expandObjectSystemNpuNpQueuesProfileCos6Osnna(d, i["cos6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos7"], _ = expandObjectSystemNpuNpQueuesProfileCos7Osnna(d, i["cos7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp0"], _ = expandObjectSystemNpuNpQueuesProfileDscp0Osnna(d, i["dscp0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp1"], _ = expandObjectSystemNpuNpQueuesProfileDscp1Osnna(d, i["dscp1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp10"], _ = expandObjectSystemNpuNpQueuesProfileDscp10Osnna(d, i["dscp10"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp11"], _ = expandObjectSystemNpuNpQueuesProfileDscp11Osnna(d, i["dscp11"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp12"], _ = expandObjectSystemNpuNpQueuesProfileDscp12Osnna(d, i["dscp12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp13"], _ = expandObjectSystemNpuNpQueuesProfileDscp13Osnna(d, i["dscp13"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp14"], _ = expandObjectSystemNpuNpQueuesProfileDscp14Osnna(d, i["dscp14"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp15"], _ = expandObjectSystemNpuNpQueuesProfileDscp15Osnna(d, i["dscp15"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp16"], _ = expandObjectSystemNpuNpQueuesProfileDscp16Osnna(d, i["dscp16"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp17"], _ = expandObjectSystemNpuNpQueuesProfileDscp17Osnna(d, i["dscp17"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp18"], _ = expandObjectSystemNpuNpQueuesProfileDscp18Osnna(d, i["dscp18"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp19"], _ = expandObjectSystemNpuNpQueuesProfileDscp19Osnna(d, i["dscp19"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp2"], _ = expandObjectSystemNpuNpQueuesProfileDscp2Osnna(d, i["dscp2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp20"], _ = expandObjectSystemNpuNpQueuesProfileDscp20Osnna(d, i["dscp20"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp21"], _ = expandObjectSystemNpuNpQueuesProfileDscp21Osnna(d, i["dscp21"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp22"], _ = expandObjectSystemNpuNpQueuesProfileDscp22Osnna(d, i["dscp22"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp23"], _ = expandObjectSystemNpuNpQueuesProfileDscp23Osnna(d, i["dscp23"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp24"], _ = expandObjectSystemNpuNpQueuesProfileDscp24Osnna(d, i["dscp24"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp25"], _ = expandObjectSystemNpuNpQueuesProfileDscp25Osnna(d, i["dscp25"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp26"], _ = expandObjectSystemNpuNpQueuesProfileDscp26Osnna(d, i["dscp26"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp27"], _ = expandObjectSystemNpuNpQueuesProfileDscp27Osnna(d, i["dscp27"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp28"], _ = expandObjectSystemNpuNpQueuesProfileDscp28Osnna(d, i["dscp28"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp29"], _ = expandObjectSystemNpuNpQueuesProfileDscp29Osnna(d, i["dscp29"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp3"], _ = expandObjectSystemNpuNpQueuesProfileDscp3Osnna(d, i["dscp3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp30"], _ = expandObjectSystemNpuNpQueuesProfileDscp30Osnna(d, i["dscp30"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp31"], _ = expandObjectSystemNpuNpQueuesProfileDscp31Osnna(d, i["dscp31"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp32"], _ = expandObjectSystemNpuNpQueuesProfileDscp32Osnna(d, i["dscp32"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp33"], _ = expandObjectSystemNpuNpQueuesProfileDscp33Osnna(d, i["dscp33"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp34"], _ = expandObjectSystemNpuNpQueuesProfileDscp34Osnna(d, i["dscp34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp35"], _ = expandObjectSystemNpuNpQueuesProfileDscp35Osnna(d, i["dscp35"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp36"], _ = expandObjectSystemNpuNpQueuesProfileDscp36Osnna(d, i["dscp36"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp37"], _ = expandObjectSystemNpuNpQueuesProfileDscp37Osnna(d, i["dscp37"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp38"], _ = expandObjectSystemNpuNpQueuesProfileDscp38Osnna(d, i["dscp38"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp39"], _ = expandObjectSystemNpuNpQueuesProfileDscp39Osnna(d, i["dscp39"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp4"], _ = expandObjectSystemNpuNpQueuesProfileDscp4Osnna(d, i["dscp4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp40"], _ = expandObjectSystemNpuNpQueuesProfileDscp40Osnna(d, i["dscp40"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp41"], _ = expandObjectSystemNpuNpQueuesProfileDscp41Osnna(d, i["dscp41"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp42"], _ = expandObjectSystemNpuNpQueuesProfileDscp42Osnna(d, i["dscp42"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp43"], _ = expandObjectSystemNpuNpQueuesProfileDscp43Osnna(d, i["dscp43"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp44"], _ = expandObjectSystemNpuNpQueuesProfileDscp44Osnna(d, i["dscp44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp45"], _ = expandObjectSystemNpuNpQueuesProfileDscp45Osnna(d, i["dscp45"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp46"], _ = expandObjectSystemNpuNpQueuesProfileDscp46Osnna(d, i["dscp46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp47"], _ = expandObjectSystemNpuNpQueuesProfileDscp47Osnna(d, i["dscp47"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp48"], _ = expandObjectSystemNpuNpQueuesProfileDscp48Osnna(d, i["dscp48"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp49"], _ = expandObjectSystemNpuNpQueuesProfileDscp49Osnna(d, i["dscp49"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp5"], _ = expandObjectSystemNpuNpQueuesProfileDscp5Osnna(d, i["dscp5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp50"], _ = expandObjectSystemNpuNpQueuesProfileDscp50Osnna(d, i["dscp50"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp51"], _ = expandObjectSystemNpuNpQueuesProfileDscp51Osnna(d, i["dscp51"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp52"], _ = expandObjectSystemNpuNpQueuesProfileDscp52Osnna(d, i["dscp52"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp53"], _ = expandObjectSystemNpuNpQueuesProfileDscp53Osnna(d, i["dscp53"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp54"], _ = expandObjectSystemNpuNpQueuesProfileDscp54Osnna(d, i["dscp54"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp55"], _ = expandObjectSystemNpuNpQueuesProfileDscp55Osnna(d, i["dscp55"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp56"], _ = expandObjectSystemNpuNpQueuesProfileDscp56Osnna(d, i["dscp56"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp57"], _ = expandObjectSystemNpuNpQueuesProfileDscp57Osnna(d, i["dscp57"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp58"], _ = expandObjectSystemNpuNpQueuesProfileDscp58Osnna(d, i["dscp58"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp59"], _ = expandObjectSystemNpuNpQueuesProfileDscp59Osnna(d, i["dscp59"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp6"], _ = expandObjectSystemNpuNpQueuesProfileDscp6Osnna(d, i["dscp6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp60"], _ = expandObjectSystemNpuNpQueuesProfileDscp60Osnna(d, i["dscp60"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp61"], _ = expandObjectSystemNpuNpQueuesProfileDscp61Osnna(d, i["dscp61"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp62"], _ = expandObjectSystemNpuNpQueuesProfileDscp62Osnna(d, i["dscp62"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp63"], _ = expandObjectSystemNpuNpQueuesProfileDscp63Osnna(d, i["dscp63"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp7"], _ = expandObjectSystemNpuNpQueuesProfileDscp7Osnna(d, i["dscp7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp8"], _ = expandObjectSystemNpuNpQueuesProfileDscp8Osnna(d, i["dscp8"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp9"], _ = expandObjectSystemNpuNpQueuesProfileDscp9Osnna(d, i["dscp9"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemNpuNpQueuesProfileIdOsnna(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesProfileTypeOsnna(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesProfileWeightOsnna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesProfileCos0Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos1Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos2Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos3Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos4Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos5Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos6Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos7Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp0Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp1Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp10Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp11Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp12Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp13Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp14Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp15Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp16Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp17Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp18Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp19Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp2Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp20Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp21Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp22Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp23Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp24Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp25Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp26Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp27Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp28Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp29Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp3Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp30Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp31Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp32Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp33Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp34Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp35Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp36Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp37Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp38Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp39Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp4Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp40Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp41Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp42Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp43Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp44Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp45Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp46Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp47Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp48Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp49Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp5Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp50Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp51Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp52Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp53Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp54Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp55Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp56Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp57Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp58Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp59Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp6Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp60Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp61Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp62Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp63Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp7Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp8Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp9Osnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileIdOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileTypeOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileWeightOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesSchedulerOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandObjectSystemNpuNpQueuesSchedulerModeOsnna(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesSchedulerNameOsnna(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesSchedulerModeOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesSchedulerNameOsnna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpQueues(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ethernet_type"); ok || d.HasChange("ethernet_type") {
		t, err := expandObjectSystemNpuNpQueuesEthernetTypeOsnna(d, v, "ethernet_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ethernet-type"] = t
		}
	}

	if v, ok := d.GetOk("ip_protocol"); ok || d.HasChange("ip_protocol") {
		t, err := expandObjectSystemNpuNpQueuesIpProtocolOsnna(d, v, "ip_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ip_service"); ok || d.HasChange("ip_service") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceOsnna(d, v, "ip_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-service"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandObjectSystemNpuNpQueuesProfileOsnna(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("scheduler"); ok || d.HasChange("scheduler") {
		t, err := expandObjectSystemNpuNpQueuesSchedulerOsnna(d, v, "scheduler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scheduler"] = t
		}
	}

	return &obj, nil
}
