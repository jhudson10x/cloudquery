// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Archives() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_archives",
		Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Archive.html",
		Resolver:    fetchEventbridgeArchives,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
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
				Resolver: resolveArchiveArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "archive_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ArchiveName"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "event_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EventCount"),
			},
			{
				Name:     "event_source_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventSourceArn"),
			},
			{
				Name:     "retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RetentionDays"),
			},
			{
				Name:     "size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeBytes"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason"),
			},
		},
	}
}
