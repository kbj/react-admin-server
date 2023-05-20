package types

import "time"

type DateTime time.Time

const timeFormat = "2006-01-02 15:04:05"

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
