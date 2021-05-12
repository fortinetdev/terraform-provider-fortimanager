// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectVpnmgr Vpntable

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVpnmgrVpntable() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnmgrVpntableCreate,
		Read:   resourceObjectVpnmgrVpntableRead,
		Update: resourceObjectVpnmgrVpntableUpdate,
		Delete: resourceObjectVpnmgrVpntableDelete,

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
			"authmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_zone_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_retrycount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dpd_retryinterval": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"fcc_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hub2spoke_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike1dhgroup": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ike1dpd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike1keylifesec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ike1localid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike1mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike1natkeepalive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ike1nattraversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike1proposal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike2autonego": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike2dhgroup": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ike2keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike2keylifekbs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ike2keylifesec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ike2keylifetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike2proposal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inter_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intf_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localid_type": &schema.Schema{
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
			"negotiate_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"network_overlay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"npu_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psk_auto_generate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsa_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spoke2hub_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"topology": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnmgrVpntableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnmgrVpntable(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrVpntable resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnmgrVpntable(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrVpntable resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnmgrVpntableRead(d, m)
}

func resourceObjectVpnmgrVpntableUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnmgrVpntable(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrVpntable resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnmgrVpntable(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrVpntable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnmgrVpntableRead(d, m)
}

func resourceObjectVpnmgrVpntableDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnmgrVpntable(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnmgrVpntable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnmgrVpntableRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnmgrVpntable(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrVpntable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnmgrVpntable(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrVpntable resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnmgrVpntableAuthmethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "psk",
			3: "signature",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableAutoZonePolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableDpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			2: "on-idle",
			3: "on-demand",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableDpdRetrycount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableDpdRetryinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectVpnmgrVpntableFccEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableHub2SpokeZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableIkeVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "1",
			2: "2",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke1Dhgroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:     "1",
			2:     "2",
			4:     "5",
			8:     "14",
			16:    "15",
			32:    "16",
			64:    "17",
			128:   "18",
			256:   "19",
			512:   "20",
			1024:  "21",
			2048:  "27",
			4096:  "28",
			8192:  "29",
			16384: "30",
			32768: "31",
			65536: "32",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke1Dpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableIke1Keylifesec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableIke1Localid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableIke1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "main",
			2: "aggressive",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke1Natkeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableIke1Nattraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			1:  "enable",
			22: "forced",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke1Proposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			5:  "des-md5",
			6:  "des-sha1",
			7:  "3des-md5",
			8:  "3des-sha1",
			9:  "aes128-md5",
			10: "aes128-sha1",
			11: "aes192-md5",
			12: "aes192-sha1",
			13: "aes256-md5",
			14: "aes256-sha1",
			19: "des-sha256",
			20: "3des-sha256",
			21: "aes128-sha256",
			22: "aes192-sha256",
			23: "aes256-sha256",
			24: "des-sha384",
			25: "des-sha512",
			26: "3des-sha384",
			27: "3des-sha512",
			28: "aes128-sha384",
			29: "aes128-sha512",
			30: "aes192-sha384",
			31: "aes192-sha512",
			32: "aes256-sha384",
			33: "aes256-sha512",
			39: "aria128-md5",
			40: "aria128-sha1",
			41: "aria128-sha256",
			42: "aria128-sha384",
			43: "aria128-sha512",
			45: "aria192-md5",
			46: "aria192-sha1",
			47: "aria192-sha256",
			48: "aria192-sha384",
			49: "aria192-sha512",
			51: "aria256-md5",
			52: "aria256-sha1",
			53: "aria256-sha256",
			54: "aria256-sha384",
			55: "aria256-sha512",
			57: "seed-md5",
			58: "seed-sha1",
			59: "seed-sha256",
			60: "seed-sha384",
			61: "seed-sha512",
			65: "aes128gcm-prfsha1",
			66: "aes128gcm-prfsha256",
			67: "aes128gcm-prfsha384",
			68: "aes128gcm-prfsha512",
			69: "aes256gcm-prfsha1",
			70: "aes256gcm-prfsha256",
			71: "aes256gcm-prfsha384",
			72: "aes256gcm-prfsha512",
			73: "chacha20poly1305-prfsha1",
			74: "chacha20poly1305-prfsha256",
			75: "chacha20poly1305-prfsha384",
			76: "chacha20poly1305-prfsha512",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke2Autonego(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableIke2Dhgroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:     "1",
			2:     "2",
			4:     "5",
			8:     "14",
			16:    "15",
			32:    "16",
			64:    "17",
			128:   "18",
			256:   "19",
			512:   "20",
			1024:  "21",
			2048:  "27",
			4096:  "28",
			8192:  "29",
			16384: "30",
			32768: "31",
			65536: "32",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke2Keepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableIke2Keylifekbs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableIke2Keylifesec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableIke2Keylifetype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "seconds",
			2: "kbs",
			3: "both",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableIke2Proposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "null-md5",
			2:  "null-sha1",
			3:  "des-null",
			4:  "3des-null",
			5:  "des-md5",
			6:  "des-sha1",
			7:  "3des-md5",
			8:  "3des-sha1",
			9:  "aes128-md5",
			10: "aes128-sha1",
			11: "aes192-md5",
			12: "aes192-sha1",
			13: "aes256-md5",
			14: "aes256-sha1",
			15: "aes128-null",
			16: "aes192-null",
			17: "aes256-null",
			18: "null-sha256",
			19: "des-sha256",
			20: "3des-sha256",
			21: "aes128-sha256",
			22: "aes192-sha256",
			23: "aes256-sha256",
			24: "des-sha384",
			25: "des-sha512",
			26: "3des-sha384",
			27: "3des-sha512",
			28: "aes128-sha384",
			29: "aes128-sha512",
			30: "aes192-sha384",
			31: "aes192-sha512",
			32: "aes256-sha384",
			33: "aes256-sha512",
			34: "null-sha384",
			35: "null-sha512",
			38: "aria128-null",
			39: "aria128-md5",
			40: "aria128-sha1",
			41: "aria128-sha256",
			42: "aria128-sha384",
			43: "aria128-sha512",
			44: "aria192-null",
			45: "aria192-md5",
			46: "aria192-sha1",
			47: "aria192-sha256",
			48: "aria192-sha384",
			49: "aria192-sha512",
			50: "aria256-null",
			51: "aria256-md5",
			52: "aria256-sha1",
			53: "aria256-sha256",
			54: "aria256-sha384",
			55: "aria256-sha512",
			56: "seed-null",
			57: "seed-md5",
			58: "seed-sha1",
			59: "seed-sha256",
			60: "seed-sha384",
			61: "seed-sha512",
			62: "aes128gcm",
			63: "aes256gcm",
			64: "chacha20poly1305",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableInterVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableIntfMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "off",
			1: "on",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableLocalidType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "auto",
			1: "fqdn",
			2: "user-fqdn",
			3: "keyid",
			4: "address",
			5: "asn1dn",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableNegotiateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableNetworkOverlay(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableNpuOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntablePfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntablePskAutoGenerate(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntablePsksecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrVpntableReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectVpnmgrVpntableRsaCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableSpoke2HubZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrVpntableTopology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "meshed",
			1: "star",
			2: "dialup",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnmgrVpntableVpnZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnmgrVpntable(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("authmethod", flattenObjectVpnmgrVpntableAuthmethod(o["authmethod"], d, "authmethod")); err != nil {
		if vv, ok := fortiAPIPatch(o["authmethod"], "ObjectVpnmgrVpntable-Authmethod"); ok {
			if err = d.Set("authmethod", vv); err != nil {
				return fmt.Errorf("Error reading authmethod: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authmethod: %v", err)
		}
	}

	if err = d.Set("auto_zone_policy", flattenObjectVpnmgrVpntableAutoZonePolicy(o["auto-zone-policy"], d, "auto_zone_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-zone-policy"], "ObjectVpnmgrVpntable-AutoZonePolicy"); ok {
			if err = d.Set("auto_zone_policy", vv); err != nil {
				return fmt.Errorf("Error reading auto_zone_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_zone_policy: %v", err)
		}
	}

	if err = d.Set("certificate", flattenObjectVpnmgrVpntableCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ObjectVpnmgrVpntable-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectVpnmgrVpntableDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectVpnmgrVpntable-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dpd", flattenObjectVpnmgrVpntableDpd(o["dpd"], d, "dpd")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd"], "ObjectVpnmgrVpntable-Dpd"); ok {
			if err = d.Set("dpd", vv); err != nil {
				return fmt.Errorf("Error reading dpd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd: %v", err)
		}
	}

	if err = d.Set("dpd_retrycount", flattenObjectVpnmgrVpntableDpdRetrycount(o["dpd-retrycount"], d, "dpd_retrycount")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd-retrycount"], "ObjectVpnmgrVpntable-DpdRetrycount"); ok {
			if err = d.Set("dpd_retrycount", vv); err != nil {
				return fmt.Errorf("Error reading dpd_retrycount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd_retrycount: %v", err)
		}
	}

	if err = d.Set("dpd_retryinterval", flattenObjectVpnmgrVpntableDpdRetryinterval(o["dpd-retryinterval"], d, "dpd_retryinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd-retryinterval"], "ObjectVpnmgrVpntable-DpdRetryinterval"); ok {
			if err = d.Set("dpd_retryinterval", vv); err != nil {
				return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
		}
	}

	if err = d.Set("fcc_enforcement", flattenObjectVpnmgrVpntableFccEnforcement(o["fcc-enforcement"], d, "fcc_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["fcc-enforcement"], "ObjectVpnmgrVpntable-FccEnforcement"); ok {
			if err = d.Set("fcc_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading fcc_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fcc_enforcement: %v", err)
		}
	}

	if err = d.Set("hub2spoke_zone", flattenObjectVpnmgrVpntableHub2SpokeZone(o["hub2spoke-zone"], d, "hub2spoke_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["hub2spoke-zone"], "ObjectVpnmgrVpntable-Hub2SpokeZone"); ok {
			if err = d.Set("hub2spoke_zone", vv); err != nil {
				return fmt.Errorf("Error reading hub2spoke_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hub2spoke_zone: %v", err)
		}
	}

	if err = d.Set("ike_version", flattenObjectVpnmgrVpntableIkeVersion(o["ike-version"], d, "ike_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-version"], "ObjectVpnmgrVpntable-IkeVersion"); ok {
			if err = d.Set("ike_version", vv); err != nil {
				return fmt.Errorf("Error reading ike_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_version: %v", err)
		}
	}

	if err = d.Set("ike1dhgroup", flattenObjectVpnmgrVpntableIke1Dhgroup(o["ike1dhgroup"], d, "ike1dhgroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1dhgroup"], "ObjectVpnmgrVpntable-Ike1Dhgroup"); ok {
			if err = d.Set("ike1dhgroup", vv); err != nil {
				return fmt.Errorf("Error reading ike1dhgroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1dhgroup: %v", err)
		}
	}

	if err = d.Set("ike1dpd", flattenObjectVpnmgrVpntableIke1Dpd(o["ike1dpd"], d, "ike1dpd")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1dpd"], "ObjectVpnmgrVpntable-Ike1Dpd"); ok {
			if err = d.Set("ike1dpd", vv); err != nil {
				return fmt.Errorf("Error reading ike1dpd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1dpd: %v", err)
		}
	}

	if err = d.Set("ike1keylifesec", flattenObjectVpnmgrVpntableIke1Keylifesec(o["ike1keylifesec"], d, "ike1keylifesec")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1keylifesec"], "ObjectVpnmgrVpntable-Ike1Keylifesec"); ok {
			if err = d.Set("ike1keylifesec", vv); err != nil {
				return fmt.Errorf("Error reading ike1keylifesec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1keylifesec: %v", err)
		}
	}

	if err = d.Set("ike1localid", flattenObjectVpnmgrVpntableIke1Localid(o["ike1localid"], d, "ike1localid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1localid"], "ObjectVpnmgrVpntable-Ike1Localid"); ok {
			if err = d.Set("ike1localid", vv); err != nil {
				return fmt.Errorf("Error reading ike1localid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1localid: %v", err)
		}
	}

	if err = d.Set("ike1mode", flattenObjectVpnmgrVpntableIke1Mode(o["ike1mode"], d, "ike1mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1mode"], "ObjectVpnmgrVpntable-Ike1Mode"); ok {
			if err = d.Set("ike1mode", vv); err != nil {
				return fmt.Errorf("Error reading ike1mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1mode: %v", err)
		}
	}

	if err = d.Set("ike1natkeepalive", flattenObjectVpnmgrVpntableIke1Natkeepalive(o["ike1natkeepalive"], d, "ike1natkeepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1natkeepalive"], "ObjectVpnmgrVpntable-Ike1Natkeepalive"); ok {
			if err = d.Set("ike1natkeepalive", vv); err != nil {
				return fmt.Errorf("Error reading ike1natkeepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1natkeepalive: %v", err)
		}
	}

	if err = d.Set("ike1nattraversal", flattenObjectVpnmgrVpntableIke1Nattraversal(o["ike1nattraversal"], d, "ike1nattraversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1nattraversal"], "ObjectVpnmgrVpntable-Ike1Nattraversal"); ok {
			if err = d.Set("ike1nattraversal", vv); err != nil {
				return fmt.Errorf("Error reading ike1nattraversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1nattraversal: %v", err)
		}
	}

	if err = d.Set("ike1proposal", flattenObjectVpnmgrVpntableIke1Proposal(o["ike1proposal"], d, "ike1proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike1proposal"], "ObjectVpnmgrVpntable-Ike1Proposal"); ok {
			if err = d.Set("ike1proposal", vv); err != nil {
				return fmt.Errorf("Error reading ike1proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike1proposal: %v", err)
		}
	}

	if err = d.Set("ike2autonego", flattenObjectVpnmgrVpntableIke2Autonego(o["ike2autonego"], d, "ike2autonego")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2autonego"], "ObjectVpnmgrVpntable-Ike2Autonego"); ok {
			if err = d.Set("ike2autonego", vv); err != nil {
				return fmt.Errorf("Error reading ike2autonego: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2autonego: %v", err)
		}
	}

	if err = d.Set("ike2dhgroup", flattenObjectVpnmgrVpntableIke2Dhgroup(o["ike2dhgroup"], d, "ike2dhgroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2dhgroup"], "ObjectVpnmgrVpntable-Ike2Dhgroup"); ok {
			if err = d.Set("ike2dhgroup", vv); err != nil {
				return fmt.Errorf("Error reading ike2dhgroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2dhgroup: %v", err)
		}
	}

	if err = d.Set("ike2keepalive", flattenObjectVpnmgrVpntableIke2Keepalive(o["ike2keepalive"], d, "ike2keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2keepalive"], "ObjectVpnmgrVpntable-Ike2Keepalive"); ok {
			if err = d.Set("ike2keepalive", vv); err != nil {
				return fmt.Errorf("Error reading ike2keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2keepalive: %v", err)
		}
	}

	if err = d.Set("ike2keylifekbs", flattenObjectVpnmgrVpntableIke2Keylifekbs(o["ike2keylifekbs"], d, "ike2keylifekbs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2keylifekbs"], "ObjectVpnmgrVpntable-Ike2Keylifekbs"); ok {
			if err = d.Set("ike2keylifekbs", vv); err != nil {
				return fmt.Errorf("Error reading ike2keylifekbs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2keylifekbs: %v", err)
		}
	}

	if err = d.Set("ike2keylifesec", flattenObjectVpnmgrVpntableIke2Keylifesec(o["ike2keylifesec"], d, "ike2keylifesec")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2keylifesec"], "ObjectVpnmgrVpntable-Ike2Keylifesec"); ok {
			if err = d.Set("ike2keylifesec", vv); err != nil {
				return fmt.Errorf("Error reading ike2keylifesec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2keylifesec: %v", err)
		}
	}

	if err = d.Set("ike2keylifetype", flattenObjectVpnmgrVpntableIke2Keylifetype(o["ike2keylifetype"], d, "ike2keylifetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2keylifetype"], "ObjectVpnmgrVpntable-Ike2Keylifetype"); ok {
			if err = d.Set("ike2keylifetype", vv); err != nil {
				return fmt.Errorf("Error reading ike2keylifetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2keylifetype: %v", err)
		}
	}

	if err = d.Set("ike2proposal", flattenObjectVpnmgrVpntableIke2Proposal(o["ike2proposal"], d, "ike2proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike2proposal"], "ObjectVpnmgrVpntable-Ike2Proposal"); ok {
			if err = d.Set("ike2proposal", vv); err != nil {
				return fmt.Errorf("Error reading ike2proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike2proposal: %v", err)
		}
	}

	if err = d.Set("inter_vdom", flattenObjectVpnmgrVpntableInterVdom(o["inter-vdom"], d, "inter_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["inter-vdom"], "ObjectVpnmgrVpntable-InterVdom"); ok {
			if err = d.Set("inter_vdom", vv); err != nil {
				return fmt.Errorf("Error reading inter_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inter_vdom: %v", err)
		}
	}

	if err = d.Set("intf_mode", flattenObjectVpnmgrVpntableIntfMode(o["intf-mode"], d, "intf_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf-mode"], "ObjectVpnmgrVpntable-IntfMode"); ok {
			if err = d.Set("intf_mode", vv); err != nil {
				return fmt.Errorf("Error reading intf_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf_mode: %v", err)
		}
	}

	if err = d.Set("localid_type", flattenObjectVpnmgrVpntableLocalidType(o["localid-type"], d, "localid_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["localid-type"], "ObjectVpnmgrVpntable-LocalidType"); ok {
			if err = d.Set("localid_type", vv); err != nil {
				return fmt.Errorf("Error reading localid_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localid_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnmgrVpntableName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnmgrVpntable-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("negotiate_timeout", flattenObjectVpnmgrVpntableNegotiateTimeout(o["negotiate-timeout"], d, "negotiate_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["negotiate-timeout"], "ObjectVpnmgrVpntable-NegotiateTimeout"); ok {
			if err = d.Set("negotiate_timeout", vv); err != nil {
				return fmt.Errorf("Error reading negotiate_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negotiate_timeout: %v", err)
		}
	}

	if err = d.Set("network_id", flattenObjectVpnmgrVpntableNetworkId(o["network-id"], d, "network_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-id"], "ObjectVpnmgrVpntable-NetworkId"); ok {
			if err = d.Set("network_id", vv); err != nil {
				return fmt.Errorf("Error reading network_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_id: %v", err)
		}
	}

	if err = d.Set("network_overlay", flattenObjectVpnmgrVpntableNetworkOverlay(o["network-overlay"], d, "network_overlay")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-overlay"], "ObjectVpnmgrVpntable-NetworkOverlay"); ok {
			if err = d.Set("network_overlay", vv); err != nil {
				return fmt.Errorf("Error reading network_overlay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_overlay: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenObjectVpnmgrVpntableNpuOffload(o["npu-offload"], d, "npu_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-offload"], "ObjectVpnmgrVpntable-NpuOffload"); ok {
			if err = d.Set("npu_offload", vv); err != nil {
				return fmt.Errorf("Error reading npu_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	if err = d.Set("pfs", flattenObjectVpnmgrVpntablePfs(o["pfs"], d, "pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfs"], "ObjectVpnmgrVpntable-Pfs"); ok {
			if err = d.Set("pfs", vv); err != nil {
				return fmt.Errorf("Error reading pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfs: %v", err)
		}
	}

	if err = d.Set("psk_auto_generate", flattenObjectVpnmgrVpntablePskAutoGenerate(o["psk-auto-generate"], d, "psk_auto_generate")); err != nil {
		if vv, ok := fortiAPIPatch(o["psk-auto-generate"], "ObjectVpnmgrVpntable-PskAutoGenerate"); ok {
			if err = d.Set("psk_auto_generate", vv); err != nil {
				return fmt.Errorf("Error reading psk_auto_generate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading psk_auto_generate: %v", err)
		}
	}

	if err = d.Set("psksecret", flattenObjectVpnmgrVpntablePsksecret(o["psksecret"], d, "psksecret")); err != nil {
		if vv, ok := fortiAPIPatch(o["psksecret"], "ObjectVpnmgrVpntable-Psksecret"); ok {
			if err = d.Set("psksecret", vv); err != nil {
				return fmt.Errorf("Error reading psksecret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading psksecret: %v", err)
		}
	}

	if err = d.Set("replay", flattenObjectVpnmgrVpntableReplay(o["replay"], d, "replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["replay"], "ObjectVpnmgrVpntable-Replay"); ok {
			if err = d.Set("replay", vv); err != nil {
				return fmt.Errorf("Error reading replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replay: %v", err)
		}
	}

	if err = d.Set("rsa_certificate", flattenObjectVpnmgrVpntableRsaCertificate(o["rsa-certificate"], d, "rsa_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsa-certificate"], "ObjectVpnmgrVpntable-RsaCertificate"); ok {
			if err = d.Set("rsa_certificate", vv); err != nil {
				return fmt.Errorf("Error reading rsa_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsa_certificate: %v", err)
		}
	}

	if err = d.Set("spoke2hub_zone", flattenObjectVpnmgrVpntableSpoke2HubZone(o["spoke2hub-zone"], d, "spoke2hub_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["spoke2hub-zone"], "ObjectVpnmgrVpntable-Spoke2HubZone"); ok {
			if err = d.Set("spoke2hub_zone", vv); err != nil {
				return fmt.Errorf("Error reading spoke2hub_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spoke2hub_zone: %v", err)
		}
	}

	if err = d.Set("topology", flattenObjectVpnmgrVpntableTopology(o["topology"], d, "topology")); err != nil {
		if vv, ok := fortiAPIPatch(o["topology"], "ObjectVpnmgrVpntable-Topology"); ok {
			if err = d.Set("topology", vv); err != nil {
				return fmt.Errorf("Error reading topology: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading topology: %v", err)
		}
	}

	if err = d.Set("vpn_zone", flattenObjectVpnmgrVpntableVpnZone(o["vpn-zone"], d, "vpn_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-zone"], "ObjectVpnmgrVpntable-VpnZone"); ok {
			if err = d.Set("vpn_zone", vv); err != nil {
				return fmt.Errorf("Error reading vpn_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_zone: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnmgrVpntableFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnmgrVpntableAuthmethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableAutoZonePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableDpd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableDpdRetrycount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableDpdRetryinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectVpnmgrVpntableFccEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableHub2SpokeZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIkeVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Dhgroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnmgrVpntableIke1Dpd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Keylifesec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Localid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Natkeepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Nattraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke1Proposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke2Autonego(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke2Dhgroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnmgrVpntableIke2Keepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke2Keylifekbs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke2Keylifesec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke2Keylifetype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIke2Proposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableInterVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableIntfMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableLocalidType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableNegotiateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableNetworkOverlay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableNpuOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntablePfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntablePskAutoGenerate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntablePsksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnmgrVpntableReplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableRsaCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableSpoke2HubZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableTopology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrVpntableVpnZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnmgrVpntable(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authmethod"); ok {
		t, err := expandObjectVpnmgrVpntableAuthmethod(d, v, "authmethod")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod"] = t
		}
	}

	if v, ok := d.GetOk("auto_zone_policy"); ok {
		t, err := expandObjectVpnmgrVpntableAutoZonePolicy(d, v, "auto_zone_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-zone-policy"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandObjectVpnmgrVpntableCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandObjectVpnmgrVpntableDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dpd"); ok {
		t, err := expandObjectVpnmgrVpntableDpd(d, v, "dpd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retrycount"); ok {
		t, err := expandObjectVpnmgrVpntableDpdRetrycount(d, v, "dpd_retrycount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retrycount"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retryinterval"); ok {
		t, err := expandObjectVpnmgrVpntableDpdRetryinterval(d, v, "dpd_retryinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retryinterval"] = t
		}
	}

	if v, ok := d.GetOk("fcc_enforcement"); ok {
		t, err := expandObjectVpnmgrVpntableFccEnforcement(d, v, "fcc_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fcc-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("hub2spoke_zone"); ok {
		t, err := expandObjectVpnmgrVpntableHub2SpokeZone(d, v, "hub2spoke_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hub2spoke-zone"] = t
		}
	}

	if v, ok := d.GetOk("ike_version"); ok {
		t, err := expandObjectVpnmgrVpntableIkeVersion(d, v, "ike_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-version"] = t
		}
	}

	if v, ok := d.GetOk("ike1dhgroup"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Dhgroup(d, v, "ike1dhgroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1dhgroup"] = t
		}
	}

	if v, ok := d.GetOk("ike1dpd"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Dpd(d, v, "ike1dpd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1dpd"] = t
		}
	}

	if v, ok := d.GetOk("ike1keylifesec"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Keylifesec(d, v, "ike1keylifesec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1keylifesec"] = t
		}
	}

	if v, ok := d.GetOk("ike1localid"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Localid(d, v, "ike1localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1localid"] = t
		}
	}

	if v, ok := d.GetOk("ike1mode"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Mode(d, v, "ike1mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1mode"] = t
		}
	}

	if v, ok := d.GetOk("ike1natkeepalive"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Natkeepalive(d, v, "ike1natkeepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1natkeepalive"] = t
		}
	}

	if v, ok := d.GetOk("ike1nattraversal"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Nattraversal(d, v, "ike1nattraversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1nattraversal"] = t
		}
	}

	if v, ok := d.GetOk("ike1proposal"); ok {
		t, err := expandObjectVpnmgrVpntableIke1Proposal(d, v, "ike1proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike1proposal"] = t
		}
	}

	if v, ok := d.GetOk("ike2autonego"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Autonego(d, v, "ike2autonego")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2autonego"] = t
		}
	}

	if v, ok := d.GetOk("ike2dhgroup"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Dhgroup(d, v, "ike2dhgroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2dhgroup"] = t
		}
	}

	if v, ok := d.GetOk("ike2keepalive"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Keepalive(d, v, "ike2keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2keepalive"] = t
		}
	}

	if v, ok := d.GetOk("ike2keylifekbs"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Keylifekbs(d, v, "ike2keylifekbs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2keylifekbs"] = t
		}
	}

	if v, ok := d.GetOk("ike2keylifesec"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Keylifesec(d, v, "ike2keylifesec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2keylifesec"] = t
		}
	}

	if v, ok := d.GetOk("ike2keylifetype"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Keylifetype(d, v, "ike2keylifetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2keylifetype"] = t
		}
	}

	if v, ok := d.GetOk("ike2proposal"); ok {
		t, err := expandObjectVpnmgrVpntableIke2Proposal(d, v, "ike2proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike2proposal"] = t
		}
	}

	if v, ok := d.GetOk("inter_vdom"); ok {
		t, err := expandObjectVpnmgrVpntableInterVdom(d, v, "inter_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-vdom"] = t
		}
	}

	if v, ok := d.GetOk("intf_mode"); ok {
		t, err := expandObjectVpnmgrVpntableIntfMode(d, v, "intf_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf-mode"] = t
		}
	}

	if v, ok := d.GetOk("localid_type"); ok {
		t, err := expandObjectVpnmgrVpntableLocalidType(d, v, "localid_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectVpnmgrVpntableName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_timeout"); ok {
		t, err := expandObjectVpnmgrVpntableNegotiateTimeout(d, v, "negotiate_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-timeout"] = t
		}
	}

	if v, ok := d.GetOk("network_id"); ok {
		t, err := expandObjectVpnmgrVpntableNetworkId(d, v, "network_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-id"] = t
		}
	}

	if v, ok := d.GetOk("network_overlay"); ok {
		t, err := expandObjectVpnmgrVpntableNetworkOverlay(d, v, "network_overlay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-overlay"] = t
		}
	}

	if v, ok := d.GetOk("npu_offload"); ok {
		t, err := expandObjectVpnmgrVpntableNpuOffload(d, v, "npu_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	if v, ok := d.GetOk("pfs"); ok {
		t, err := expandObjectVpnmgrVpntablePfs(d, v, "pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfs"] = t
		}
	}

	if v, ok := d.GetOk("psk_auto_generate"); ok {
		t, err := expandObjectVpnmgrVpntablePskAutoGenerate(d, v, "psk_auto_generate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk-auto-generate"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		t, err := expandObjectVpnmgrVpntablePsksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("replay"); ok {
		t, err := expandObjectVpnmgrVpntableReplay(d, v, "replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay"] = t
		}
	}

	if v, ok := d.GetOk("rsa_certificate"); ok {
		t, err := expandObjectVpnmgrVpntableRsaCertificate(d, v, "rsa_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-certificate"] = t
		}
	}

	if v, ok := d.GetOk("spoke2hub_zone"); ok {
		t, err := expandObjectVpnmgrVpntableSpoke2HubZone(d, v, "spoke2hub_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spoke2hub-zone"] = t
		}
	}

	if v, ok := d.GetOk("topology"); ok {
		t, err := expandObjectVpnmgrVpntableTopology(d, v, "topology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["topology"] = t
		}
	}

	if v, ok := d.GetOk("vpn_zone"); ok {
		t, err := expandObjectVpnmgrVpntableVpnZone(d, v, "vpn_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-zone"] = t
		}
	}

	return &obj, nil
}
