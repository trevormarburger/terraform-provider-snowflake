// Code generated by sdk-to-schema generator; DO NOT EDIT.

package schemas

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ShowFileFormatSchema represents output of SHOW query for the single FileFormat.
var ShowFileFormatSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"created_on": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"owner": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"comment": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"owner_role_type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"options": {
		Type:     schema.TypeInvalid,
		Computed: true,
	},
}

var _ = ShowFileFormatSchema

func FileFormatToSchema(fileFormat *sdk.FileFormat) map[string]any {
	fileFormatSchema := make(map[string]any)
	fileFormatSchema["name"] = fileFormat.Name.FullyQualifiedName()
	fileFormatSchema["created_on"] = fileFormat.CreatedOn.String()
	fileFormatSchema["type"] = string(fileFormat.Type)
	fileFormatSchema["owner"] = fileFormat.Owner
	fileFormatSchema["comment"] = fileFormat.Comment
	fileFormatSchema["owner_role_type"] = fileFormat.OwnerRoleType
	fileFormatSchema["options"] = fileFormat.Options
	return fileFormatSchema
}

var _ = FileFormatToSchema