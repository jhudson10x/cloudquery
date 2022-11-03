// Code generated by codegen; DO NOT EDIT.

package mwaa

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Environments() *schema.Table {
	return &schema.Table{
		Name:                "aws_mwaa_environments",
		Description:         "https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html",
		Resolver:            fetchMwaaEnvironments,
		PreResourceResolver: getEnvironment,
		Multiplex:           client.ServiceAccountRegionMultiplexer("airflow"),
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
				Name:     "airflow_configuration_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AirflowConfigurationOptions"),
			},
			{
				Name:     "airflow_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AirflowVersion"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "dag_s3_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DagS3Path"),
			},
			{
				Name:     "environment_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnvironmentClass"),
			},
			{
				Name:     "execution_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExecutionRoleArn"),
			},
			{
				Name:     "kms_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKey"),
			},
			{
				Name:     "last_update",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastUpdate"),
			},
			{
				Name:     "logging_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoggingConfiguration"),
			},
			{
				Name:     "max_workers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxWorkers"),
			},
			{
				Name:     "min_workers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinWorkers"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkConfiguration"),
			},
			{
				Name:     "plugins_s3_object_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PluginsS3ObjectVersion"),
			},
			{
				Name:     "plugins_s3_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PluginsS3Path"),
			},
			{
				Name:     "requirements_s3_object_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequirementsS3ObjectVersion"),
			},
			{
				Name:     "requirements_s3_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequirementsS3Path"),
			},
			{
				Name:     "schedulers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Schedulers"),
			},
			{
				Name:     "service_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceRoleArn"),
			},
			{
				Name:     "source_bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceBucketArn"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "webserver_access_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebserverAccessMode"),
			},
			{
				Name:     "webserver_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebserverUrl"),
			},
			{
				Name:     "weekly_maintenance_window_start",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WeeklyMaintenanceWindowStart"),
			},
		},
	}
}