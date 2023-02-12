// Package types UnixTime类型
// @author: kbj
// @date: 2023/2/7
package types

//import (
//	"database/sql/driver"
//	"fmt"
//	"github.com/gookit/goutil/strutil"
//	"time"
//)
//
//type UnixTime time.Time
//
//// MarshalJSON 重写JSON序列化方法
//func (u *UnixTime) MarshalJSON() ([]byte, error) {
//	t := time.Time(*u)
//	if t.IsZero() {
//		// 零值处理
//		return nil, nil
//	}
//	return []byte(fmt.Sprintf("%d", t.UnixMilli())), nil
//}
//
//// UnmarshalJSON 重写JSON反序列化方法
//func (u *UnixTime) UnmarshalJSON(b []byte) error {
//	timestamp, err := strutil.ToInt64(string(b))
//	if err != nil {
//		return err
//	}
//
//	*u = UnixTime(time.UnixMilli(timestamp))
//	return nil
//}
//
//// Scan Ent使用自定义类型所需要的方法
//func (u *UnixTime) Scan(value interface{}) error {
//	switch v := value.(type) {
//	case nil:
//	case time.Time:
//		*u = UnixTime(v)
//	default:
//		return fmt.Errorf("unexpected type %T", v)
//	}
//	return nil
//}
//
//// Value Ent使用自定义类型所需要的方法
//func (u *UnixTime) Value() (driver.Value, error) {
//	if u == nil {
//		return nil, nil
//	}
//
//	return time.Time(*u), nil
//}
//
//// UnixTimeNow 获得UnixTime类型的当前时间
//func UnixTimeNow() UnixTime {
//	return UnixTime(time.Now())
//}
