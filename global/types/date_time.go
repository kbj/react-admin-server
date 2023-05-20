package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateTime time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *DateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = DateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Value 注意，这里(t DateTime)不能用指针
func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// UnmarshalJSON 转换成时间戳
func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = DateTime(now)
	return
}

// MarshalJSON 转换成自定义格式
func (t *DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}
