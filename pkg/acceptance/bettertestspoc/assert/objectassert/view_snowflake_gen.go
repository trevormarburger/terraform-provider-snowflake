// Code generated by assertions generator; DO NOT EDIT.

package objectassert

import (
	"fmt"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type ViewAssert struct {
	*assert.SnowflakeObjectAssert[sdk.View, sdk.SchemaObjectIdentifier]
}

func View(t *testing.T, id sdk.SchemaObjectIdentifier) *ViewAssert {
	t.Helper()
	return &ViewAssert{
		assert.NewSnowflakeObjectAssertWithProvider(sdk.ObjectTypeView, id, acc.TestClient().View.Show),
	}
}

func ViewFromObject(t *testing.T, view *sdk.View) *ViewAssert {
	t.Helper()
	return &ViewAssert{
		assert.NewSnowflakeObjectAssertWithObject(sdk.ObjectTypeView, view.ID(), view),
	}
}

func (v *ViewAssert) HasCreatedOn(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.CreatedOn != expected {
			return fmt.Errorf("expected created on: %v; got: %v", expected, o.CreatedOn)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasName(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.Name != expected {
			return fmt.Errorf("expected name: %v; got: %v", expected, o.Name)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasKind(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.Kind != expected {
			return fmt.Errorf("expected kind: %v; got: %v", expected, o.Kind)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasReserved(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.Reserved != expected {
			return fmt.Errorf("expected reserved: %v; got: %v", expected, o.Reserved)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasDatabaseName(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.DatabaseName != expected {
			return fmt.Errorf("expected database name: %v; got: %v", expected, o.DatabaseName)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasSchemaName(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.SchemaName != expected {
			return fmt.Errorf("expected schema name: %v; got: %v", expected, o.SchemaName)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasOwner(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.Owner != expected {
			return fmt.Errorf("expected owner: %v; got: %v", expected, o.Owner)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasComment(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.Comment != expected {
			return fmt.Errorf("expected comment: %v; got: %v", expected, o.Comment)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasText(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.Text != expected {
			return fmt.Errorf("expected text: %v; got: %v", expected, o.Text)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasIsSecure(expected bool) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.IsSecure != expected {
			return fmt.Errorf("expected is secure: %v; got: %v", expected, o.IsSecure)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasIsMaterialized(expected bool) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.IsMaterialized != expected {
			return fmt.Errorf("expected is materialized: %v; got: %v", expected, o.IsMaterialized)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasOwnerRoleType(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.OwnerRoleType != expected {
			return fmt.Errorf("expected owner role type: %v; got: %v", expected, o.OwnerRoleType)
		}
		return nil
	})
	return v
}

func (v *ViewAssert) HasChangeTracking(expected string) *ViewAssert {
	v.AddAssertion(func(t *testing.T, o *sdk.View) error {
		t.Helper()
		if o.ChangeTracking != expected {
			return fmt.Errorf("expected change tracking: %v; got: %v", expected, o.ChangeTracking)
		}
		return nil
	})
	return v
}