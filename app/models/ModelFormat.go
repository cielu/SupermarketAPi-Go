package models

import (
	"fmt"
	"time"
)

// Custom format time
type FormatTime struct {
	time.Time
}

// MarshalJSON on FormatTime format Time field with %Y-%m-%d %H:%M:%S
func (t FormatTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04"))), nil
}

// Scan valueof time.Time
func (t *FormatTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = FormatTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Value insert timestamp into mysql need this function.
// func (t FormatTime) Value() (driver.Value, error) {
// 	var zeroTime time.Time
// 	if t.Time.UnixNano() == zeroTime.UnixNano() {
// 		return nil, nil
// 	}
// 	return t.Time, nil
// }