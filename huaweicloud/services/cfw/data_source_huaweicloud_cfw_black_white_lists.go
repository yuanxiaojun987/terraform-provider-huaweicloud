// Generated by PMS #144
package cfw

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/filters"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceCfwBlackWhiteLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCfwBlackWhiteListsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"object_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the protected object ID.`,
			},
			"list_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the blacklist/whitelist type.`,
			},
			"address_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the IP address type.`,
			},
			"list_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the blacklist/whitelist ID.`,
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the IP address.`,
			},
			"direction": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the direction of a black or white address.`,
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the port.`,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Specifies The protocol type.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the description.`,
			},
			"fw_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the firewall instance ID.`,
			},
			"records": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The blacklist and whitelist records.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"list_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The blacklist/whitelist ID.`,
						},
						"direction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The direction of a black or white address.`,
						},
						"address_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IP address type.`,
						},
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IP address.`,
						},
						"protocol": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The protocol type.`,
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The port.`,
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The description.`,
						},
					},
				},
			},
		},
	}
}

type BlackWhiteListsDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newBlackWhiteListsDSWrapper(d *schema.ResourceData, meta interface{}) *BlackWhiteListsDSWrapper {
	return &BlackWhiteListsDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceCfwBlackWhiteListsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newBlackWhiteListsDSWrapper(d, meta)
	lisBlaWhiLisRst, err := wrapper.ListBlackWhiteLists()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listBlackWhiteListsToSchema(lisBlaWhiLisRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API CFW GET /v1/{project_id}/black-white-lists
func (w *BlackWhiteListsDSWrapper) ListBlackWhiteLists() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "cfw")
	if err != nil {
		return nil, err
	}

	uri := "/v1/{project_id}/black-white-lists"
	params := map[string]any{
		"object_id":      w.Get("object_id"),
		"list_type":      w.Get("list_type"),
		"address_type":   w.Get("address_type"),
		"address":        w.Get("address"),
		"port":           w.Get("port"),
		"fw_instance_id": w.Get("fw_instance_id"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		OffsetPager("data.records", "offset", "limit", 1024).
		Filter(
			filters.New().From("data.records").
				Where("list_id", "=", w.Get("list_id")).
				Where("description", "contains", w.Get("description")).
				Where("protocol", "=", w.Get("protocol")).
				Where("direction", "=", w.GetToInt("direction")),
		).
		Request().
		Result()
}

func (w *BlackWhiteListsDSWrapper) listBlackWhiteListsToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("records", schemas.SliceToList(body.Get("data.records"),
			func(record gjson.Result) any {
				return map[string]any{
					"list_id":      record.Get("list_id").Value(),
					"direction":    record.Get("direction").String(),
					"address_type": record.Get("address_type").String(),
					"address":      record.Get("address").Value(),
					"protocol":     record.Get("protocol").Value(),
					"port":         record.Get("port").Value(),
					"description":  record.Get("description").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
