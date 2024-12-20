package resources

import (
	"context"
	"fmt"
	"log"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/internal/provider"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var unsafeExecuteSchema = map[string]*schema.Schema{
	"execute": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "SQL statement to execute. Forces recreation of resource on change.",
	},
	"revert": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "SQL statement to revert the execute statement. Invoked when resource is being destroyed.",
	},
	"query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Optional SQL statement to do a read. Invoked after creation and every time it is changed.",
	},
	"query_results": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "List of key-value maps (text to text) retrieved after executing read query. Will be empty if the query results in an error.",
		Elem: &schema.Schema{
			Type: schema.TypeMap,
			Elem: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

func UnsafeExecute() *schema.Resource {
	return &schema.Resource{
		Create: CreateUnsafeExecute,
		Read:   ReadUnsafeExecute,
		Delete: DeleteUnsafeExecute,
		Update: UpdateUnsafeExecute,

		Schema: unsafeExecuteSchema,

		Description: "Experimental resource allowing execution of ANY SQL statement. It may destroy resources if used incorrectly. It may behave incorrectly combined with other resources. Use at your own risk.",

		CustomizeDiff: TrackingCustomDiffWrapper(resources.UnsafeExecute, func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
			if diff.HasChange("query") {
				err := diff.SetNewComputed("query_results")
				if err != nil {
					return err
				}
			}
			return nil
		}),
	}
}

func ReadUnsafeExecute(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*provider.Context).Client
	ctx := context.Background()

	readStatement := d.Get("query").(string)

	setNilResults := func() error {
		log.Printf(`[DEBUG] Clearing query_results`)
		err := d.Set("query_results", nil)
		if err != nil {
			return err
		}
		return nil
	}

	if readStatement == "" {
		return setNilResults()
	} else {
		rows, err := client.QueryUnsafe(ctx, readStatement)
		if err != nil {
			log.Printf(`[WARN] SQL query "%s" failed with err %v`, readStatement, err)
			return setNilResults()
		}
		log.Printf(`[INFO] SQL query "%s" executed successfully, returned rows count: %d`, readStatement, len(rows))
		rowsTransformed := make([]map[string]any, len(rows))
		for i, row := range rows {
			t := make(map[string]any)
			for k, v := range row {
				if *v == nil {
					t[k] = nil
				} else {
					switch (*v).(type) {
					case fmt.Stringer:
						t[k] = fmt.Sprintf("%v", *v)
					case string:
						t[k] = *v
					default:
						return fmt.Errorf("currently only objects convertible to String are supported by query; got %v", *v)
					}
				}
			}
			rowsTransformed[i] = t
		}
		err = d.Set("query_results", rowsTransformed)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateUnsafeExecute(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*provider.Context).Client
	ctx := context.Background()

	id, err := uuid.GenerateUUID()
	if err != nil {
		return err
	}

	executeStatement := d.Get("execute").(string)
	_, err = client.ExecUnsafe(ctx, executeStatement)
	if err != nil {
		return err
	}

	d.SetId(id)
	log.Printf(`[INFO] SQL "%s" applied successfully\n`, executeStatement)

	return ReadUnsafeExecute(d, meta)
}

func DeleteUnsafeExecute(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*provider.Context).Client
	ctx := context.Background()

	revertStatement := d.Get("revert").(string)
	_, err := client.ExecUnsafe(ctx, revertStatement)
	if err != nil {
		return err
	}

	d.SetId("")
	log.Printf(`[INFO] SQL "%s" applied successfully\n`, revertStatement)

	return nil
}

func UpdateUnsafeExecute(d *schema.ResourceData, meta interface{}) error {
	if d.HasChange("query") {
		return ReadUnsafeExecute(d, meta)
	}
	return nil
}
