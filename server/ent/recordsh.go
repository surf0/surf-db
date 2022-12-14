// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/ent/recordsh"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// RecordSh is the model entity for the RecordSh schema.
type RecordSh struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// PlayerName holds the value of the "player_name" field.
	PlayerName string `json:"player_name,omitempty"`
	// PlayerID holds the value of the "player_id" field.
	PlayerID string `json:"player_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Track holds the value of the "track" field.
	Track int `json:"track,omitempty"`
	// MapName holds the value of the "map_name" field.
	MapName string `json:"map_name,omitempty"`
	// Time holds the value of the "time" field.
	Time string `json:"time,omitempty"`
	// Improvement holds the value of the "improvement" field.
	Improvement string `json:"improvement,omitempty"`
	// Server holds the value of the "server" field.
	Server string `json:"server,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RecordSh) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case recordsh.FieldTrack:
			values[i] = new(sql.NullInt64)
		case recordsh.FieldID, recordsh.FieldPlayerName, recordsh.FieldPlayerID, recordsh.FieldType, recordsh.FieldMapName, recordsh.FieldTime, recordsh.FieldImprovement, recordsh.FieldServer:
			values[i] = new(sql.NullString)
		case recordsh.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RecordSh", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RecordSh fields.
func (rs *RecordSh) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recordsh.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rs.ID = value.String
			}
		case recordsh.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				rs.Timestamp = value.Time
			}
		case recordsh.FieldPlayerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field player_name", values[i])
			} else if value.Valid {
				rs.PlayerName = value.String
			}
		case recordsh.FieldPlayerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field player_id", values[i])
			} else if value.Valid {
				rs.PlayerID = value.String
			}
		case recordsh.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				rs.Type = value.String
			}
		case recordsh.FieldTrack:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field track", values[i])
			} else if value.Valid {
				rs.Track = int(value.Int64)
			}
		case recordsh.FieldMapName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field map_name", values[i])
			} else if value.Valid {
				rs.MapName = value.String
			}
		case recordsh.FieldTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				rs.Time = value.String
			}
		case recordsh.FieldImprovement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field improvement", values[i])
			} else if value.Valid {
				rs.Improvement = value.String
			}
		case recordsh.FieldServer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field server", values[i])
			} else if value.Valid {
				rs.Server = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RecordSh.
// Note that you need to call RecordSh.Unwrap() before calling this method if this RecordSh
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *RecordSh) Update() *RecordShUpdateOne {
	return (&RecordShClient{config: rs.config}).UpdateOne(rs)
}

// Unwrap unwraps the RecordSh entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rs *RecordSh) Unwrap() *RecordSh {
	_tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: RecordSh is not a transactional entity")
	}
	rs.config.driver = _tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *RecordSh) String() string {
	var builder strings.Builder
	builder.WriteString("RecordSh(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rs.ID))
	builder.WriteString("timestamp=")
	builder.WriteString(rs.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("player_name=")
	builder.WriteString(rs.PlayerName)
	builder.WriteString(", ")
	builder.WriteString("player_id=")
	builder.WriteString(rs.PlayerID)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(rs.Type)
	builder.WriteString(", ")
	builder.WriteString("track=")
	builder.WriteString(fmt.Sprintf("%v", rs.Track))
	builder.WriteString(", ")
	builder.WriteString("map_name=")
	builder.WriteString(rs.MapName)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(rs.Time)
	builder.WriteString(", ")
	builder.WriteString("improvement=")
	builder.WriteString(rs.Improvement)
	builder.WriteString(", ")
	builder.WriteString("server=")
	builder.WriteString(rs.Server)
	builder.WriteByte(')')
	return builder.String()
}

// RecordShs is a parsable slice of RecordSh.
type RecordShs []*RecordSh

func (rs RecordShs) config(cfg config) {
	for _i := range rs {
		rs[_i].config = cfg
	}
}
