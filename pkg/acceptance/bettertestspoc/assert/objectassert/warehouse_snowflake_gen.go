// Code generated by sdk-to-schema generator; DO NOT EDIT.

package objectassert

import (
	"fmt"
	"testing"
	"time"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type WarehouseAssert struct {
	*assert.SnowflakeObjectAssert[sdk.Warehouse, sdk.AccountObjectIdentifier]
}

func Warehouse(t *testing.T, id sdk.AccountObjectIdentifier) *WarehouseAssert {
	t.Helper()
	return &WarehouseAssert{
		assert.NewSnowflakeObjectAssertWithProvider(sdk.ObjectTypeWarehouse, id, acc.TestClient().Warehouse.Show),
	}
}

func WarehouseFromObject(t *testing.T, warehouse *sdk.Warehouse) *WarehouseAssert {
	t.Helper()
	return &WarehouseAssert{
		assert.NewSnowflakeObjectAssertWithObject(sdk.ObjectTypeWarehouse, warehouse.ID(), warehouse),
	}
}

func (w *WarehouseAssert) HasName(expected string) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Name != expected {
			return fmt.Errorf("expected name: %v; got: %v", expected, o.Name)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasState(expected sdk.WarehouseState) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.State != expected {
			return fmt.Errorf("expected state: %v; got: %v", expected, o.State)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasType(expected sdk.WarehouseType) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Type != expected {
			return fmt.Errorf("expected type: %v; got: %v", expected, o.Type)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasSize(expected sdk.WarehouseSize) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Size != expected {
			return fmt.Errorf("expected size: %v; got: %v", expected, o.Size)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasMinClusterCount(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.MinClusterCount != expected {
			return fmt.Errorf("expected min cluster count: %v; got: %v", expected, o.MinClusterCount)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasMaxClusterCount(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.MaxClusterCount != expected {
			return fmt.Errorf("expected max cluster count: %v; got: %v", expected, o.MaxClusterCount)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasStartedClusters(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.StartedClusters != expected {
			return fmt.Errorf("expected started clusters: %v; got: %v", expected, o.StartedClusters)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasRunning(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Running != expected {
			return fmt.Errorf("expected running: %v; got: %v", expected, o.Running)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasQueued(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Queued != expected {
			return fmt.Errorf("expected queued: %v; got: %v", expected, o.Queued)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasIsDefault(expected bool) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.IsDefault != expected {
			return fmt.Errorf("expected is default: %v; got: %v", expected, o.IsDefault)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasIsCurrent(expected bool) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.IsCurrent != expected {
			return fmt.Errorf("expected is current: %v; got: %v", expected, o.IsCurrent)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasAutoSuspend(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.AutoSuspend != expected {
			return fmt.Errorf("expected auto suspend: %v; got: %v", expected, o.AutoSuspend)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasAutoResume(expected bool) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.AutoResume != expected {
			return fmt.Errorf("expected auto resume: %v; got: %v", expected, o.AutoResume)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasAvailable(expected float64) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Available != expected {
			return fmt.Errorf("expected available: %v; got: %v", expected, o.Available)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasProvisioning(expected float64) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Provisioning != expected {
			return fmt.Errorf("expected provisioning: %v; got: %v", expected, o.Provisioning)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasQuiescing(expected float64) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Quiescing != expected {
			return fmt.Errorf("expected quiescing: %v; got: %v", expected, o.Quiescing)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasOther(expected float64) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Other != expected {
			return fmt.Errorf("expected other: %v; got: %v", expected, o.Other)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasCreatedOn(expected time.Time) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.CreatedOn != expected {
			return fmt.Errorf("expected created on: %v; got: %v", expected, o.CreatedOn)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasResumedOn(expected time.Time) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.ResumedOn != expected {
			return fmt.Errorf("expected resumed on: %v; got: %v", expected, o.ResumedOn)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasUpdatedOn(expected time.Time) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.UpdatedOn != expected {
			return fmt.Errorf("expected updated on: %v; got: %v", expected, o.UpdatedOn)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasOwner(expected string) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Owner != expected {
			return fmt.Errorf("expected owner: %v; got: %v", expected, o.Owner)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasComment(expected string) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.Comment != expected {
			return fmt.Errorf("expected comment: %v; got: %v", expected, o.Comment)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasEnableQueryAcceleration(expected bool) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.EnableQueryAcceleration != expected {
			return fmt.Errorf("expected enable query acceleration: %v; got: %v", expected, o.EnableQueryAcceleration)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasQueryAccelerationMaxScaleFactor(expected int) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.QueryAccelerationMaxScaleFactor != expected {
			return fmt.Errorf("expected query acceleration max scale factor: %v; got: %v", expected, o.QueryAccelerationMaxScaleFactor)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasResourceMonitor(expected sdk.AccountObjectIdentifier) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.ResourceMonitor.Name() != expected.Name() {
			return fmt.Errorf("expected resource monitor: %v; got: %v", expected.Name(), o.ResourceMonitor.Name())
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasScalingPolicy(expected sdk.ScalingPolicy) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.ScalingPolicy != expected {
			return fmt.Errorf("expected scaling policy: %v; got: %v", expected, o.ScalingPolicy)
		}
		return nil
	})
	return w
}

func (w *WarehouseAssert) HasOwnerRoleType(expected string) *WarehouseAssert {
	w.AddAssertion(func(t *testing.T, o *sdk.Warehouse) error {
		t.Helper()
		if o.OwnerRoleType != expected {
			return fmt.Errorf("expected owner role type: %v; got: %v", expected, o.OwnerRoleType)
		}
		return nil
	})
	return w
}