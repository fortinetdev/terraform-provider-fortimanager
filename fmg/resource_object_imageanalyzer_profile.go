// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectImageAnalyzer Profile

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectImageAnalyzerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectImageAnalyzerProfileCreate,
		Read:   resourceObjectImageAnalyzerProfileRead,
		Update: resourceObjectImageAnalyzerProfileUpdate,
		Delete: resourceObjectImageAnalyzerProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"alcohol_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"alcohol_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"blocked_img_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drugs_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"drugs_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extremism_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"extremism_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gambling_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gambling_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gore_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gore_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_skip_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"image_skip_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"image_skip_width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ocr_activation_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"optical_character_recognition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"porn_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"porn_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rating_err_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replace_image": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"swim_underwear_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"swim_underwear_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weapons_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weapons_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectImageAnalyzerProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectImageAnalyzerProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectImageAnalyzerProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectImageAnalyzerProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectImageAnalyzerProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectImageAnalyzerProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectImageAnalyzerProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectImageAnalyzerProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectImageAnalyzerProfileRead(d, m)
}

func resourceObjectImageAnalyzerProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectImageAnalyzerProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectImageAnalyzerProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectImageAnalyzerProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectImageAnalyzerProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectImageAnalyzerProfileRead(d, m)
}

func resourceObjectImageAnalyzerProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectImageAnalyzerProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectImageAnalyzerProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectImageAnalyzerProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectImageAnalyzerProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectImageAnalyzerProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectImageAnalyzerProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectImageAnalyzerProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectImageAnalyzerProfileAlcoholBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileAlcoholStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileBlockedImgCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileDrugsBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileDrugsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileExtremismBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileExtremismStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileGamblingBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileGamblingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileGoreBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileGoreStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileImageSkipHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileImageSkipSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileImageSkipWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileLogOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileOcrActivationThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileOpticalCharacterRecognition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfilePornBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfilePornStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileRatingErrAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileReplaceImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectImageAnalyzerProfileSourceUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileSwimUnderwearStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileWeaponsBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectImageAnalyzerProfileWeaponsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectImageAnalyzerProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("alcohol_block_strictness_level", flattenObjectImageAnalyzerProfileAlcoholBlockStrictnessLevel(o["alcohol-block-strictness-level"], d, "alcohol_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["alcohol-block-strictness-level"], "ObjectImageAnalyzerProfile-AlcoholBlockStrictnessLevel"); ok {
			if err = d.Set("alcohol_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading alcohol_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alcohol_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("alcohol_status", flattenObjectImageAnalyzerProfileAlcoholStatus(o["alcohol-status"], d, "alcohol_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["alcohol-status"], "ObjectImageAnalyzerProfile-AlcoholStatus"); ok {
			if err = d.Set("alcohol_status", vv); err != nil {
				return fmt.Errorf("Error reading alcohol_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alcohol_status: %v", err)
		}
	}

	if err = d.Set("blocked_img_cache", flattenObjectImageAnalyzerProfileBlockedImgCache(o["blocked-img-cache"], d, "blocked_img_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-img-cache"], "ObjectImageAnalyzerProfile-BlockedImgCache"); ok {
			if err = d.Set("blocked_img_cache", vv); err != nil {
				return fmt.Errorf("Error reading blocked_img_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_img_cache: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectImageAnalyzerProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectImageAnalyzerProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("drugs_block_strictness_level", flattenObjectImageAnalyzerProfileDrugsBlockStrictnessLevel(o["drugs-block-strictness-level"], d, "drugs_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["drugs-block-strictness-level"], "ObjectImageAnalyzerProfile-DrugsBlockStrictnessLevel"); ok {
			if err = d.Set("drugs_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading drugs_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drugs_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("drugs_status", flattenObjectImageAnalyzerProfileDrugsStatus(o["drugs-status"], d, "drugs_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["drugs-status"], "ObjectImageAnalyzerProfile-DrugsStatus"); ok {
			if err = d.Set("drugs_status", vv); err != nil {
				return fmt.Errorf("Error reading drugs_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drugs_status: %v", err)
		}
	}

	if err = d.Set("extremism_block_strictness_level", flattenObjectImageAnalyzerProfileExtremismBlockStrictnessLevel(o["extremism-block-strictness-level"], d, "extremism_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["extremism-block-strictness-level"], "ObjectImageAnalyzerProfile-ExtremismBlockStrictnessLevel"); ok {
			if err = d.Set("extremism_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading extremism_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extremism_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("extremism_status", flattenObjectImageAnalyzerProfileExtremismStatus(o["extremism-status"], d, "extremism_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["extremism-status"], "ObjectImageAnalyzerProfile-ExtremismStatus"); ok {
			if err = d.Set("extremism_status", vv); err != nil {
				return fmt.Errorf("Error reading extremism_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extremism_status: %v", err)
		}
	}

	if err = d.Set("gambling_block_strictness_level", flattenObjectImageAnalyzerProfileGamblingBlockStrictnessLevel(o["gambling-block-strictness-level"], d, "gambling_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["gambling-block-strictness-level"], "ObjectImageAnalyzerProfile-GamblingBlockStrictnessLevel"); ok {
			if err = d.Set("gambling_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading gambling_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gambling_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("gambling_status", flattenObjectImageAnalyzerProfileGamblingStatus(o["gambling-status"], d, "gambling_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["gambling-status"], "ObjectImageAnalyzerProfile-GamblingStatus"); ok {
			if err = d.Set("gambling_status", vv); err != nil {
				return fmt.Errorf("Error reading gambling_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gambling_status: %v", err)
		}
	}

	if err = d.Set("gore_block_strictness_level", flattenObjectImageAnalyzerProfileGoreBlockStrictnessLevel(o["gore-block-strictness-level"], d, "gore_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["gore-block-strictness-level"], "ObjectImageAnalyzerProfile-GoreBlockStrictnessLevel"); ok {
			if err = d.Set("gore_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading gore_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gore_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("gore_status", flattenObjectImageAnalyzerProfileGoreStatus(o["gore-status"], d, "gore_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["gore-status"], "ObjectImageAnalyzerProfile-GoreStatus"); ok {
			if err = d.Set("gore_status", vv); err != nil {
				return fmt.Errorf("Error reading gore_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gore_status: %v", err)
		}
	}

	if err = d.Set("image_skip_height", flattenObjectImageAnalyzerProfileImageSkipHeight(o["image-skip-height"], d, "image_skip_height")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-skip-height"], "ObjectImageAnalyzerProfile-ImageSkipHeight"); ok {
			if err = d.Set("image_skip_height", vv); err != nil {
				return fmt.Errorf("Error reading image_skip_height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_skip_height: %v", err)
		}
	}

	if err = d.Set("image_skip_size", flattenObjectImageAnalyzerProfileImageSkipSize(o["image-skip-size"], d, "image_skip_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-skip-size"], "ObjectImageAnalyzerProfile-ImageSkipSize"); ok {
			if err = d.Set("image_skip_size", vv); err != nil {
				return fmt.Errorf("Error reading image_skip_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_skip_size: %v", err)
		}
	}

	if err = d.Set("image_skip_width", flattenObjectImageAnalyzerProfileImageSkipWidth(o["image-skip-width"], d, "image_skip_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-skip-width"], "ObjectImageAnalyzerProfile-ImageSkipWidth"); ok {
			if err = d.Set("image_skip_width", vv); err != nil {
				return fmt.Errorf("Error reading image_skip_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_skip_width: %v", err)
		}
	}

	if err = d.Set("log_option", flattenObjectImageAnalyzerProfileLogOption(o["log-option"], d, "log_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-option"], "ObjectImageAnalyzerProfile-LogOption"); ok {
			if err = d.Set("log_option", vv); err != nil {
				return fmt.Errorf("Error reading log_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_option: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectImageAnalyzerProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectImageAnalyzerProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ocr_activation_threshold", flattenObjectImageAnalyzerProfileOcrActivationThreshold(o["ocr-activation-threshold"], d, "ocr_activation_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocr-activation-threshold"], "ObjectImageAnalyzerProfile-OcrActivationThreshold"); ok {
			if err = d.Set("ocr_activation_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ocr_activation_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocr_activation_threshold: %v", err)
		}
	}

	if err = d.Set("optical_character_recognition", flattenObjectImageAnalyzerProfileOpticalCharacterRecognition(o["optical-character-recognition"], d, "optical_character_recognition")); err != nil {
		if vv, ok := fortiAPIPatch(o["optical-character-recognition"], "ObjectImageAnalyzerProfile-OpticalCharacterRecognition"); ok {
			if err = d.Set("optical_character_recognition", vv); err != nil {
				return fmt.Errorf("Error reading optical_character_recognition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optical_character_recognition: %v", err)
		}
	}

	if err = d.Set("porn_block_strictness_level", flattenObjectImageAnalyzerProfilePornBlockStrictnessLevel(o["porn-block-strictness-level"], d, "porn_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["porn-block-strictness-level"], "ObjectImageAnalyzerProfile-PornBlockStrictnessLevel"); ok {
			if err = d.Set("porn_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading porn_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading porn_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("porn_status", flattenObjectImageAnalyzerProfilePornStatus(o["porn-status"], d, "porn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["porn-status"], "ObjectImageAnalyzerProfile-PornStatus"); ok {
			if err = d.Set("porn_status", vv); err != nil {
				return fmt.Errorf("Error reading porn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading porn_status: %v", err)
		}
	}

	if err = d.Set("rating_err_action", flattenObjectImageAnalyzerProfileRatingErrAction(o["rating-err-action"], d, "rating_err_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["rating-err-action"], "ObjectImageAnalyzerProfile-RatingErrAction"); ok {
			if err = d.Set("rating_err_action", vv); err != nil {
				return fmt.Errorf("Error reading rating_err_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rating_err_action: %v", err)
		}
	}

	if err = d.Set("replace_image", flattenObjectImageAnalyzerProfileReplaceImage(o["replace-image"], d, "replace_image")); err != nil {
		if vv, ok := fortiAPIPatch(o["replace-image"], "ObjectImageAnalyzerProfile-ReplaceImage"); ok {
			if err = d.Set("replace_image", vv); err != nil {
				return fmt.Errorf("Error reading replace_image: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replace_image: %v", err)
		}
	}

	if err = d.Set("source_url", flattenObjectImageAnalyzerProfileSourceUrl(o["source-url"], d, "source_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-url"], "ObjectImageAnalyzerProfile-SourceUrl"); ok {
			if err = d.Set("source_url", vv); err != nil {
				return fmt.Errorf("Error reading source_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_url: %v", err)
		}
	}

	if err = d.Set("swim_underwear_block_strictness_level", flattenObjectImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(o["swim_underwear-block-strictness-level"], d, "swim_underwear_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["swim_underwear-block-strictness-level"], "ObjectImageAnalyzerProfile-SwimUnderwearBlockStrictnessLevel"); ok {
			if err = d.Set("swim_underwear_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading swim_underwear_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading swim_underwear_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("swim_underwear_status", flattenObjectImageAnalyzerProfileSwimUnderwearStatus(o["swim_underwear-status"], d, "swim_underwear_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["swim_underwear-status"], "ObjectImageAnalyzerProfile-SwimUnderwearStatus"); ok {
			if err = d.Set("swim_underwear_status", vv); err != nil {
				return fmt.Errorf("Error reading swim_underwear_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading swim_underwear_status: %v", err)
		}
	}

	if err = d.Set("weapons_block_strictness_level", flattenObjectImageAnalyzerProfileWeaponsBlockStrictnessLevel(o["weapons-block-strictness-level"], d, "weapons_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["weapons-block-strictness-level"], "ObjectImageAnalyzerProfile-WeaponsBlockStrictnessLevel"); ok {
			if err = d.Set("weapons_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading weapons_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weapons_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("weapons_status", flattenObjectImageAnalyzerProfileWeaponsStatus(o["weapons-status"], d, "weapons_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["weapons-status"], "ObjectImageAnalyzerProfile-WeaponsStatus"); ok {
			if err = d.Set("weapons_status", vv); err != nil {
				return fmt.Errorf("Error reading weapons_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weapons_status: %v", err)
		}
	}

	return nil
}

func flattenObjectImageAnalyzerProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectImageAnalyzerProfileAlcoholBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileAlcoholStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileBlockedImgCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileDrugsBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileDrugsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileExtremismBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileExtremismStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileGamblingBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileGamblingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileGoreBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileGoreStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileImageSkipHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileImageSkipSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileImageSkipWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileLogOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileOcrActivationThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileOpticalCharacterRecognition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfilePornBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfilePornStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileRatingErrAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileReplaceImage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectImageAnalyzerProfileSourceUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileSwimUnderwearStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileWeaponsBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectImageAnalyzerProfileWeaponsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectImageAnalyzerProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alcohol_block_strictness_level"); ok || d.HasChange("alcohol_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileAlcoholBlockStrictnessLevel(d, v, "alcohol_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alcohol-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("alcohol_status"); ok || d.HasChange("alcohol_status") {
		t, err := expandObjectImageAnalyzerProfileAlcoholStatus(d, v, "alcohol_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alcohol-status"] = t
		}
	}

	if v, ok := d.GetOk("blocked_img_cache"); ok || d.HasChange("blocked_img_cache") {
		t, err := expandObjectImageAnalyzerProfileBlockedImgCache(d, v, "blocked_img_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-img-cache"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectImageAnalyzerProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("drugs_block_strictness_level"); ok || d.HasChange("drugs_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileDrugsBlockStrictnessLevel(d, v, "drugs_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drugs-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("drugs_status"); ok || d.HasChange("drugs_status") {
		t, err := expandObjectImageAnalyzerProfileDrugsStatus(d, v, "drugs_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drugs-status"] = t
		}
	}

	if v, ok := d.GetOk("extremism_block_strictness_level"); ok || d.HasChange("extremism_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileExtremismBlockStrictnessLevel(d, v, "extremism_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extremism-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("extremism_status"); ok || d.HasChange("extremism_status") {
		t, err := expandObjectImageAnalyzerProfileExtremismStatus(d, v, "extremism_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extremism-status"] = t
		}
	}

	if v, ok := d.GetOk("gambling_block_strictness_level"); ok || d.HasChange("gambling_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileGamblingBlockStrictnessLevel(d, v, "gambling_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gambling-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("gambling_status"); ok || d.HasChange("gambling_status") {
		t, err := expandObjectImageAnalyzerProfileGamblingStatus(d, v, "gambling_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gambling-status"] = t
		}
	}

	if v, ok := d.GetOk("gore_block_strictness_level"); ok || d.HasChange("gore_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileGoreBlockStrictnessLevel(d, v, "gore_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gore-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("gore_status"); ok || d.HasChange("gore_status") {
		t, err := expandObjectImageAnalyzerProfileGoreStatus(d, v, "gore_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gore-status"] = t
		}
	}

	if v, ok := d.GetOk("image_skip_height"); ok || d.HasChange("image_skip_height") {
		t, err := expandObjectImageAnalyzerProfileImageSkipHeight(d, v, "image_skip_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-skip-height"] = t
		}
	}

	if v, ok := d.GetOk("image_skip_size"); ok || d.HasChange("image_skip_size") {
		t, err := expandObjectImageAnalyzerProfileImageSkipSize(d, v, "image_skip_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-skip-size"] = t
		}
	}

	if v, ok := d.GetOk("image_skip_width"); ok || d.HasChange("image_skip_width") {
		t, err := expandObjectImageAnalyzerProfileImageSkipWidth(d, v, "image_skip_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-skip-width"] = t
		}
	}

	if v, ok := d.GetOk("log_option"); ok || d.HasChange("log_option") {
		t, err := expandObjectImageAnalyzerProfileLogOption(d, v, "log_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-option"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectImageAnalyzerProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ocr_activation_threshold"); ok || d.HasChange("ocr_activation_threshold") {
		t, err := expandObjectImageAnalyzerProfileOcrActivationThreshold(d, v, "ocr_activation_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocr-activation-threshold"] = t
		}
	}

	if v, ok := d.GetOk("optical_character_recognition"); ok || d.HasChange("optical_character_recognition") {
		t, err := expandObjectImageAnalyzerProfileOpticalCharacterRecognition(d, v, "optical_character_recognition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optical-character-recognition"] = t
		}
	}

	if v, ok := d.GetOk("porn_block_strictness_level"); ok || d.HasChange("porn_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfilePornBlockStrictnessLevel(d, v, "porn_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["porn-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("porn_status"); ok || d.HasChange("porn_status") {
		t, err := expandObjectImageAnalyzerProfilePornStatus(d, v, "porn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["porn-status"] = t
		}
	}

	if v, ok := d.GetOk("rating_err_action"); ok || d.HasChange("rating_err_action") {
		t, err := expandObjectImageAnalyzerProfileRatingErrAction(d, v, "rating_err_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rating-err-action"] = t
		}
	}

	if v, ok := d.GetOk("replace_image"); ok || d.HasChange("replace_image") {
		t, err := expandObjectImageAnalyzerProfileReplaceImage(d, v, "replace_image")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replace-image"] = t
		}
	}

	if v, ok := d.GetOk("source_url"); ok || d.HasChange("source_url") {
		t, err := expandObjectImageAnalyzerProfileSourceUrl(d, v, "source_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-url"] = t
		}
	}

	if v, ok := d.GetOk("swim_underwear_block_strictness_level"); ok || d.HasChange("swim_underwear_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(d, v, "swim_underwear_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swim_underwear-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("swim_underwear_status"); ok || d.HasChange("swim_underwear_status") {
		t, err := expandObjectImageAnalyzerProfileSwimUnderwearStatus(d, v, "swim_underwear_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swim_underwear-status"] = t
		}
	}

	if v, ok := d.GetOk("weapons_block_strictness_level"); ok || d.HasChange("weapons_block_strictness_level") {
		t, err := expandObjectImageAnalyzerProfileWeaponsBlockStrictnessLevel(d, v, "weapons_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weapons-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("weapons_status"); ok || d.HasChange("weapons_status") {
		t, err := expandObjectImageAnalyzerProfileWeaponsStatus(d, v, "weapons_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weapons-status"] = t
		}
	}

	return &obj, nil
}
