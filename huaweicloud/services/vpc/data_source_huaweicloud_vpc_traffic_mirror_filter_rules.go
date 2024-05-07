// Generated by PMS #47
package vpc

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceVpcTrafficMirrorFilterRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVpcTrafficMirrorFilterRulesRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"traffic_mirror_filter_rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The traffic mirror filter rule ID used as the query filter.`,
			},
			"traffic_mirror_filter_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The traffic mirror filter ID used as the query filter.`,
			},
			"direction": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The direction of the traffic mirror filter rule.`,
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The protocol of the traffic mirror filter rule.`,
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The policy of in the traffic mirror filter rule.`,
			},
			"priority": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The priority number of the traffic mirror filter rule.`,
			},
			"source_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The source IP address of the traffic mirror filter rule.`,
			},
			"source_port_range": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The source port number range of the traffic mirror filter rule.`,
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The destination IP address of the traffic mirror filter rule.`,
			},
			"destination_port_range": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The destination port number range of the traffic mirror filter rule.`,
			},
			"traffic_mirror_filter_rules": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `List of traffic mirror filter rules.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Time when a traffic mirror filter rule is created.`,
						},
						"source_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Source CIDR block of the mirrored traffic.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Project ID.`,
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Traffic mirror filter rule ID.`,
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Whether to accept or reject traffic.`,
						},
						"source_port_range": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Source port range.`,
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Destination CIDR block of the mirrored traffic.`,
						},
						"traffic_mirror_filter_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Traffic mirror filter ID.`,
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Description of a traffic mirror filter rule.`,
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Time when a traffic mirror filter rule is updated.`,
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Mirror filter rule priority.`,
						},
						"ethertype": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `IP address version of the mirrored traffic.`,
						},
						"destination_port_range": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Source port range.`,
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Protocol of the mirrored traffic.`,
						},
						"direction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Traffic direction.`,
						},
					},
				},
			},
		},
	}
}

type TrafficMirrorFilterRulesDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newTrafficMirrorFilterRulesDSWrapper(d *schema.ResourceData, meta interface{}) *TrafficMirrorFilterRulesDSWrapper {
	return &TrafficMirrorFilterRulesDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceVpcTrafficMirrorFilterRulesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newTrafficMirrorFilterRulesDSWrapper(d, meta)
	lisTraMirFilRulRst, err := wrapper.ListTrafficMirrorFilterRules()
	if err != nil {
		return diag.FromErr(err)
	}

	id, _ := uuid.GenerateUUID()
	d.SetId(id)

	err = wrapper.listTrafficMirrorFilterRulesToSchema(lisTraMirFilRulRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API VPC GET /v3/{project_id}/vpc/traffic-mirror-filter-rules
func (w *TrafficMirrorFilterRulesDSWrapper) ListTrafficMirrorFilterRules() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "vpc")
	if err != nil {
		return nil, err
	}

	uri := "/v3/{project_id}/vpc/traffic-mirror-filter-rules"
	params := map[string]any{
		"id":                       w.Get("traffic_mirror_filter_rule_id"),
		"traffic_mirror_filter_id": w.Get("traffic_mirror_filter_id"),
		"direction":                w.Get("direction"),
		"protocol":                 w.Get("protocol"),
		"source_cidr_block":        w.Get("source_cidr_block"),
		"destination_cidr_block":   w.Get("destination_cidr_block"),
		"source_port_range":        w.Get("source_port_range"),
		"destination_port_range":   w.Get("destination_port_range"),
		"action":                   w.Get("action"),
		"priority":                 w.Get("priority"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		MarkerPager("traffic_mirror_filter_rules", "page_info.next_marker", "marker").
		Request().
		Result()
}

func (w *TrafficMirrorFilterRulesDSWrapper) listTrafficMirrorFilterRulesToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("traffic_mirror_filter_rules", schemas.SliceToList(body.Get("traffic_mirror_filter_rules"),
			func(traMirFilRul gjson.Result) any {
				return map[string]any{
					"created_at":               traMirFilRul.Get("created_at").Value(),
					"source_cidr_block":        traMirFilRul.Get("source_cidr_block").Value(),
					"project_id":               traMirFilRul.Get("project_id").Value(),
					"id":                       traMirFilRul.Get("id").Value(),
					"action":                   traMirFilRul.Get("action").Value(),
					"source_port_range":        traMirFilRul.Get("source_port_range").Value(),
					"destination_cidr_block":   traMirFilRul.Get("destination_cidr_block").Value(),
					"traffic_mirror_filter_id": traMirFilRul.Get("traffic_mirror_filter_id").Value(),
					"description":              traMirFilRul.Get("description").Value(),
					"updated_at":               traMirFilRul.Get("updated_at").Value(),
					"priority":                 traMirFilRul.Get("priority").Value(),
					"ethertype":                traMirFilRul.Get("ethertype").Value(),
					"destination_port_range":   traMirFilRul.Get("destination_port_range").Value(),
					"protocol":                 traMirFilRul.Get("protocol").Value(),
					"direction":                traMirFilRul.Get("direction").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
