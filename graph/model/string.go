package model

import (
	"fmt"
	"io"

	"gopkg.in/guregu/null.v4"
)

type NullString struct {
	null.String
}

func (s NullString) MarshalGQL(w io.Writer) {
	if s.Valid {
		_, _ = io.WriteString(w, s.String.String)
	} else {
		_, _ = io.WriteString(w, "null")
	}
}

func (s *NullString) UnmarshalGQL(v interface{}) (err error) {
	if str, ok := v.(string); ok {
		*s = NullString{null.StringFrom(str)}
		return
	}

	err = fmt.Errorf("time should be a string")
	return
}
