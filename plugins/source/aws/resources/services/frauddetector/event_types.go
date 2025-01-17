// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventTypes() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_event_types",
		Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_EventType.html",
		Resolver:    fetchFrauddetectorEventTypes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("frauddetector"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceTags,
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "entity_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EntityTypes"),
			},
			{
				Name:     "event_ingestion",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventIngestion"),
			},
			{
				Name:     "event_variables",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EventVariables"),
			},
			{
				Name:     "ingested_event_statistics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IngestedEventStatistics"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}
