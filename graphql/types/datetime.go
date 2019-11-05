package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// DateTime is a custom GraphQL type to represent an instant in time. It has to be added to a schema
// via "scalar DateTime" since it is not a predeclared GraphQL type like "ID".
type DateTime struct {
	time.Time
}

const (
	DatetimeFormat = "Jan 02, 2006 15:04:05 AM"
)

func NewDateTime(t time.Time) DateTime {
	return DateTime{
		Time: t,
	}
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

// UnmarshalGraphQL is a custom unmarshaler for DateTime
//
// This function will be called whenever you use the
// time scalar as an input
func (t *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		t.Time = input
		return nil
	case string:
		var err error
		t.Time, err = time.Parse(DatetimeFormat, input)
		return err
	case int:
		t.Time = time.Unix(int64(input), 0)
		return nil
	case float64:
		t.Time = time.Unix(int64(input), 0)
		return nil
	default:
		return fmt.Errorf("wrong type: %T", input)
	}
}

// MarshalJSON is a custom marshaler for DateTime
//
// This function will be called whenever you
// query for fields that use the DateTime type
func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(DatetimeFormat))
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	var timeStr string
	err := json.Unmarshal(b, &timeStr)
	if err != nil {
		return err
	}

	tm, err := time.Parse(DatetimeFormat, timeStr)
	if err != nil {
		return err
	}

	t.Time = tm

	return nil
}
