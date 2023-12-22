package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"gopkg.in/guregu/null.v4"
)

type Timestamp struct {
	null.Time
}

func (t Timestamp) MarshalGQL(w io.Writer) {
	if t.Valid {
		_, _ = io.WriteString(w, t.Time.Time.Format(strconv.Quote(time.RFC3339)))
		return
	}

	_, _ = io.WriteString(w, "null")
}

func (t *Timestamp) UnmarshalGQL(v interface{}) (err error) {
	if timeString, ok := v.(string); ok {
		var _t time.Time
		_t, err = time.Parse(time.RFC3339, timeString)
		if err != nil {
			return
		}
		*t = Timestamp{null.TimeFrom(_t)}
		return
	}

	err = fmt.Errorf("time should be a string")
	return
}
