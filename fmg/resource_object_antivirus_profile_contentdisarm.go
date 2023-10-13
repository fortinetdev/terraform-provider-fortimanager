// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AV Content Disarm and Reconstruction settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectAntivirusProfileContentDisarm() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectAntivirusProfileContentDisarmUpdate,
		Read:   resourceObjectAntivirusProfileContentDisarmRead,
		Update: resourceObjectAntivirusProfileContentDisarmUpdate,
		Delete: resourceObjectAntivirusProfileContentDisarmDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cover_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"office_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"office_dde": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"office_embed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"office_hylink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"office_linked": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"office_macro": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"original_file_destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_act_form": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_act_gotor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_act_java": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_act_launch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_act_movie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_act_sound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_embedfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_hyperlink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_javacode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectAntivirusProfileContentDisarmUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectAntivirusProfileContentDisarm(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfileContentDisarm resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectAntivirusProfileContentDisarm(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfileContentDisarm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectAntivirusProfileContentDisarm")

	return resourceObjectAntivirusProfileContentDisarmRead(d, m)
}

func resourceObjectAntivirusProfileContentDisarmDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectAntivirusProfileContentDisarm(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectAntivirusProfileContentDisarm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectAntivirusProfileContentDisarmRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectAntivirusProfileContentDisarm(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfileContentDisarm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectAntivirusProfileContentDisarm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfileContentDisarm resource from API: %v", err)
	}
	return nil
}

func flattenObjectAntivirusProfileContentDisarmCoverPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmDetectOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmErrorAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeDde2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeEmbed2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeHylink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeLinked2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeMacro2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOriginalFileDestination2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActForm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActGotor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActJava2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActLaunch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActMovie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActSound2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfEmbedfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfHyperlink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfJavacode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectAntivirusProfileContentDisarm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cover_page", flattenObjectAntivirusProfileContentDisarmCoverPage2edl(o["cover-page"], d, "cover_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["cover-page"], "ObjectAntivirusProfileContentDisarm-CoverPage"); ok {
			if err = d.Set("cover_page", vv); err != nil {
				return fmt.Errorf("Error reading cover_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cover_page: %v", err)
		}
	}

	if err = d.Set("detect_only", flattenObjectAntivirusProfileContentDisarmDetectOnly2edl(o["detect-only"], d, "detect_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-only"], "ObjectAntivirusProfileContentDisarm-DetectOnly"); ok {
			if err = d.Set("detect_only", vv); err != nil {
				return fmt.Errorf("Error reading detect_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_only: %v", err)
		}
	}

	if err = d.Set("error_action", flattenObjectAntivirusProfileContentDisarmErrorAction2edl(o["error-action"], d, "error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["error-action"], "ObjectAntivirusProfileContentDisarm-ErrorAction"); ok {
			if err = d.Set("error_action", vv); err != nil {
				return fmt.Errorf("Error reading error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading error_action: %v", err)
		}
	}

	if err = d.Set("office_action", flattenObjectAntivirusProfileContentDisarmOfficeAction2edl(o["office-action"], d, "office_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["office-action"], "ObjectAntivirusProfileContentDisarm-OfficeAction"); ok {
			if err = d.Set("office_action", vv); err != nil {
				return fmt.Errorf("Error reading office_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading office_action: %v", err)
		}
	}

	if err = d.Set("office_dde", flattenObjectAntivirusProfileContentDisarmOfficeDde2edl(o["office-dde"], d, "office_dde")); err != nil {
		if vv, ok := fortiAPIPatch(o["office-dde"], "ObjectAntivirusProfileContentDisarm-OfficeDde"); ok {
			if err = d.Set("office_dde", vv); err != nil {
				return fmt.Errorf("Error reading office_dde: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading office_dde: %v", err)
		}
	}

	if err = d.Set("office_embed", flattenObjectAntivirusProfileContentDisarmOfficeEmbed2edl(o["office-embed"], d, "office_embed")); err != nil {
		if vv, ok := fortiAPIPatch(o["office-embed"], "ObjectAntivirusProfileContentDisarm-OfficeEmbed"); ok {
			if err = d.Set("office_embed", vv); err != nil {
				return fmt.Errorf("Error reading office_embed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading office_embed: %v", err)
		}
	}

	if err = d.Set("office_hylink", flattenObjectAntivirusProfileContentDisarmOfficeHylink2edl(o["office-hylink"], d, "office_hylink")); err != nil {
		if vv, ok := fortiAPIPatch(o["office-hylink"], "ObjectAntivirusProfileContentDisarm-OfficeHylink"); ok {
			if err = d.Set("office_hylink", vv); err != nil {
				return fmt.Errorf("Error reading office_hylink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading office_hylink: %v", err)
		}
	}

	if err = d.Set("office_linked", flattenObjectAntivirusProfileContentDisarmOfficeLinked2edl(o["office-linked"], d, "office_linked")); err != nil {
		if vv, ok := fortiAPIPatch(o["office-linked"], "ObjectAntivirusProfileContentDisarm-OfficeLinked"); ok {
			if err = d.Set("office_linked", vv); err != nil {
				return fmt.Errorf("Error reading office_linked: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading office_linked: %v", err)
		}
	}

	if err = d.Set("office_macro", flattenObjectAntivirusProfileContentDisarmOfficeMacro2edl(o["office-macro"], d, "office_macro")); err != nil {
		if vv, ok := fortiAPIPatch(o["office-macro"], "ObjectAntivirusProfileContentDisarm-OfficeMacro"); ok {
			if err = d.Set("office_macro", vv); err != nil {
				return fmt.Errorf("Error reading office_macro: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading office_macro: %v", err)
		}
	}

	if err = d.Set("original_file_destination", flattenObjectAntivirusProfileContentDisarmOriginalFileDestination2edl(o["original-file-destination"], d, "original_file_destination")); err != nil {
		if vv, ok := fortiAPIPatch(o["original-file-destination"], "ObjectAntivirusProfileContentDisarm-OriginalFileDestination"); ok {
			if err = d.Set("original_file_destination", vv); err != nil {
				return fmt.Errorf("Error reading original_file_destination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading original_file_destination: %v", err)
		}
	}

	if err = d.Set("pdf_act_form", flattenObjectAntivirusProfileContentDisarmPdfActForm2edl(o["pdf-act-form"], d, "pdf_act_form")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-act-form"], "ObjectAntivirusProfileContentDisarm-PdfActForm"); ok {
			if err = d.Set("pdf_act_form", vv); err != nil {
				return fmt.Errorf("Error reading pdf_act_form: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_act_form: %v", err)
		}
	}

	if err = d.Set("pdf_act_gotor", flattenObjectAntivirusProfileContentDisarmPdfActGotor2edl(o["pdf-act-gotor"], d, "pdf_act_gotor")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-act-gotor"], "ObjectAntivirusProfileContentDisarm-PdfActGotor"); ok {
			if err = d.Set("pdf_act_gotor", vv); err != nil {
				return fmt.Errorf("Error reading pdf_act_gotor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_act_gotor: %v", err)
		}
	}

	if err = d.Set("pdf_act_java", flattenObjectAntivirusProfileContentDisarmPdfActJava2edl(o["pdf-act-java"], d, "pdf_act_java")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-act-java"], "ObjectAntivirusProfileContentDisarm-PdfActJava"); ok {
			if err = d.Set("pdf_act_java", vv); err != nil {
				return fmt.Errorf("Error reading pdf_act_java: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_act_java: %v", err)
		}
	}

	if err = d.Set("pdf_act_launch", flattenObjectAntivirusProfileContentDisarmPdfActLaunch2edl(o["pdf-act-launch"], d, "pdf_act_launch")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-act-launch"], "ObjectAntivirusProfileContentDisarm-PdfActLaunch"); ok {
			if err = d.Set("pdf_act_launch", vv); err != nil {
				return fmt.Errorf("Error reading pdf_act_launch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_act_launch: %v", err)
		}
	}

	if err = d.Set("pdf_act_movie", flattenObjectAntivirusProfileContentDisarmPdfActMovie2edl(o["pdf-act-movie"], d, "pdf_act_movie")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-act-movie"], "ObjectAntivirusProfileContentDisarm-PdfActMovie"); ok {
			if err = d.Set("pdf_act_movie", vv); err != nil {
				return fmt.Errorf("Error reading pdf_act_movie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_act_movie: %v", err)
		}
	}

	if err = d.Set("pdf_act_sound", flattenObjectAntivirusProfileContentDisarmPdfActSound2edl(o["pdf-act-sound"], d, "pdf_act_sound")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-act-sound"], "ObjectAntivirusProfileContentDisarm-PdfActSound"); ok {
			if err = d.Set("pdf_act_sound", vv); err != nil {
				return fmt.Errorf("Error reading pdf_act_sound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_act_sound: %v", err)
		}
	}

	if err = d.Set("pdf_embedfile", flattenObjectAntivirusProfileContentDisarmPdfEmbedfile2edl(o["pdf-embedfile"], d, "pdf_embedfile")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-embedfile"], "ObjectAntivirusProfileContentDisarm-PdfEmbedfile"); ok {
			if err = d.Set("pdf_embedfile", vv); err != nil {
				return fmt.Errorf("Error reading pdf_embedfile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_embedfile: %v", err)
		}
	}

	if err = d.Set("pdf_hyperlink", flattenObjectAntivirusProfileContentDisarmPdfHyperlink2edl(o["pdf-hyperlink"], d, "pdf_hyperlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-hyperlink"], "ObjectAntivirusProfileContentDisarm-PdfHyperlink"); ok {
			if err = d.Set("pdf_hyperlink", vv); err != nil {
				return fmt.Errorf("Error reading pdf_hyperlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_hyperlink: %v", err)
		}
	}

	if err = d.Set("pdf_javacode", flattenObjectAntivirusProfileContentDisarmPdfJavacode2edl(o["pdf-javacode"], d, "pdf_javacode")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-javacode"], "ObjectAntivirusProfileContentDisarm-PdfJavacode"); ok {
			if err = d.Set("pdf_javacode", vv); err != nil {
				return fmt.Errorf("Error reading pdf_javacode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_javacode: %v", err)
		}
	}

	return nil
}

func flattenObjectAntivirusProfileContentDisarmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAntivirusProfileContentDisarmCoverPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmDetectOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmErrorAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeDde2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeEmbed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeHylink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeLinked2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeMacro2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOriginalFileDestination2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActForm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActGotor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActJava2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActLaunch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActMovie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActSound2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfEmbedfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfHyperlink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfJavacode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectAntivirusProfileContentDisarm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cover_page"); ok || d.HasChange("cover_page") {
		t, err := expandObjectAntivirusProfileContentDisarmCoverPage2edl(d, v, "cover_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cover-page"] = t
		}
	}

	if v, ok := d.GetOk("detect_only"); ok || d.HasChange("detect_only") {
		t, err := expandObjectAntivirusProfileContentDisarmDetectOnly2edl(d, v, "detect_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-only"] = t
		}
	}

	if v, ok := d.GetOk("error_action"); ok || d.HasChange("error_action") {
		t, err := expandObjectAntivirusProfileContentDisarmErrorAction2edl(d, v, "error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["error-action"] = t
		}
	}

	if v, ok := d.GetOk("office_action"); ok || d.HasChange("office_action") {
		t, err := expandObjectAntivirusProfileContentDisarmOfficeAction2edl(d, v, "office_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["office-action"] = t
		}
	}

	if v, ok := d.GetOk("office_dde"); ok || d.HasChange("office_dde") {
		t, err := expandObjectAntivirusProfileContentDisarmOfficeDde2edl(d, v, "office_dde")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["office-dde"] = t
		}
	}

	if v, ok := d.GetOk("office_embed"); ok || d.HasChange("office_embed") {
		t, err := expandObjectAntivirusProfileContentDisarmOfficeEmbed2edl(d, v, "office_embed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["office-embed"] = t
		}
	}

	if v, ok := d.GetOk("office_hylink"); ok || d.HasChange("office_hylink") {
		t, err := expandObjectAntivirusProfileContentDisarmOfficeHylink2edl(d, v, "office_hylink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["office-hylink"] = t
		}
	}

	if v, ok := d.GetOk("office_linked"); ok || d.HasChange("office_linked") {
		t, err := expandObjectAntivirusProfileContentDisarmOfficeLinked2edl(d, v, "office_linked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["office-linked"] = t
		}
	}

	if v, ok := d.GetOk("office_macro"); ok || d.HasChange("office_macro") {
		t, err := expandObjectAntivirusProfileContentDisarmOfficeMacro2edl(d, v, "office_macro")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["office-macro"] = t
		}
	}

	if v, ok := d.GetOk("original_file_destination"); ok || d.HasChange("original_file_destination") {
		t, err := expandObjectAntivirusProfileContentDisarmOriginalFileDestination2edl(d, v, "original_file_destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["original-file-destination"] = t
		}
	}

	if v, ok := d.GetOk("pdf_act_form"); ok || d.HasChange("pdf_act_form") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfActForm2edl(d, v, "pdf_act_form")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-act-form"] = t
		}
	}

	if v, ok := d.GetOk("pdf_act_gotor"); ok || d.HasChange("pdf_act_gotor") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfActGotor2edl(d, v, "pdf_act_gotor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-act-gotor"] = t
		}
	}

	if v, ok := d.GetOk("pdf_act_java"); ok || d.HasChange("pdf_act_java") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfActJava2edl(d, v, "pdf_act_java")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-act-java"] = t
		}
	}

	if v, ok := d.GetOk("pdf_act_launch"); ok || d.HasChange("pdf_act_launch") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfActLaunch2edl(d, v, "pdf_act_launch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-act-launch"] = t
		}
	}

	if v, ok := d.GetOk("pdf_act_movie"); ok || d.HasChange("pdf_act_movie") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfActMovie2edl(d, v, "pdf_act_movie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-act-movie"] = t
		}
	}

	if v, ok := d.GetOk("pdf_act_sound"); ok || d.HasChange("pdf_act_sound") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfActSound2edl(d, v, "pdf_act_sound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-act-sound"] = t
		}
	}

	if v, ok := d.GetOk("pdf_embedfile"); ok || d.HasChange("pdf_embedfile") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfEmbedfile2edl(d, v, "pdf_embedfile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-embedfile"] = t
		}
	}

	if v, ok := d.GetOk("pdf_hyperlink"); ok || d.HasChange("pdf_hyperlink") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfHyperlink2edl(d, v, "pdf_hyperlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-hyperlink"] = t
		}
	}

	if v, ok := d.GetOk("pdf_javacode"); ok || d.HasChange("pdf_javacode") {
		t, err := expandObjectAntivirusProfileContentDisarmPdfJavacode2edl(d, v, "pdf_javacode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-javacode"] = t
		}
	}

	return &obj, nil
}
