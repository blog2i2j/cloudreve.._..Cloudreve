// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudreve/Cloudreve/v4/ent/node"
	"github.com/cloudreve/Cloudreve/v4/inventory/types"
	"github.com/cloudreve/Cloudreve/v4/pkg/boolset"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Status holds the value of the "status" field.
	Status node.Status `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type node.Type `json:"type,omitempty"`
	// Server holds the value of the "server" field.
	Server string `json:"server,omitempty"`
	// SlaveKey holds the value of the "slave_key" field.
	SlaveKey string `json:"slave_key,omitempty"`
	// Capabilities holds the value of the "capabilities" field.
	Capabilities *boolset.BooleanSet `json:"capabilities,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings *types.NodeSetting `json:"settings,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges        NodeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// StoragePolicy holds the value of the storage_policy edge.
	StoragePolicy []*StoragePolicy `json:"storage_policy,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StoragePolicyOrErr returns the StoragePolicy value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) StoragePolicyOrErr() ([]*StoragePolicy, error) {
	if e.loadedTypes[0] {
		return e.StoragePolicy, nil
	}
	return nil, &NotLoadedError{edge: "storage_policy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldSettings:
			values[i] = new([]byte)
		case node.FieldCapabilities:
			values[i] = new(boolset.BooleanSet)
		case node.FieldID, node.FieldWeight:
			values[i] = new(sql.NullInt64)
		case node.FieldStatus, node.FieldName, node.FieldType, node.FieldServer, node.FieldSlaveKey:
			values[i] = new(sql.NullString)
		case node.FieldCreatedAt, node.FieldUpdatedAt, node.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case node.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case node.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case node.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				n.DeletedAt = new(time.Time)
				*n.DeletedAt = value.Time
			}
		case node.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				n.Status = node.Status(value.String)
			}
		case node.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case node.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				n.Type = node.Type(value.String)
			}
		case node.FieldServer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field server", values[i])
			} else if value.Valid {
				n.Server = value.String
			}
		case node.FieldSlaveKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slave_key", values[i])
			} else if value.Valid {
				n.SlaveKey = value.String
			}
		case node.FieldCapabilities:
			if value, ok := values[i].(*boolset.BooleanSet); !ok {
				return fmt.Errorf("unexpected type %T for field capabilities", values[i])
			} else if value != nil {
				n.Capabilities = value
			}
		case node.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		case node.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				n.Weight = int(value.Int64)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Node.
// This includes values selected through modifiers, order, etc.
func (n *Node) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryStoragePolicy queries the "storage_policy" edge of the Node entity.
func (n *Node) QueryStoragePolicy() *StoragePolicyQuery {
	return NewNodeClient(n.config).QueryStoragePolicy(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return NewNodeClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := n.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", n.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", n.Type))
	builder.WriteString(", ")
	builder.WriteString("server=")
	builder.WriteString(n.Server)
	builder.WriteString(", ")
	builder.WriteString("slave_key=")
	builder.WriteString(n.SlaveKey)
	builder.WriteString(", ")
	builder.WriteString("capabilities=")
	builder.WriteString(fmt.Sprintf("%v", n.Capabilities))
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(fmt.Sprintf("%v", n.Settings))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", n.Weight))
	builder.WriteByte(')')
	return builder.String()
}

// SetStoragePolicy manually set the edge as loaded state.
func (e *Node) SetStoragePolicy(v []*StoragePolicy) {
	e.Edges.StoragePolicy = v
	e.Edges.loadedTypes[0] = true
}

// Nodes is a parsable slice of Node.
type Nodes []*Node
