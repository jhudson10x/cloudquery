// Code generated by codegen; DO NOT EDIT.

package billing

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Package() *schema.Table {
	return &schema.Table{
		Name:      "github_billing_package",
		Resolver:  fetchPackage,
		Multiplex: client.OrgMultiplex,
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "total_gigabytes_bandwidth_used",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalGigabytesBandwidthUsed"),
			},
			{
				Name:     "total_paid_gigabytes_bandwidth_used",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalPaidGigabytesBandwidthUsed"),
			},
			{
				Name:     "included_gigabytes_bandwidth",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IncludedGigabytesBandwidth"),
			},
		},
	}
}
