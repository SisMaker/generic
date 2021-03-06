package generic

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalTimestampMS(t *testing.T) {
	v := time.Now()
	expected := v.UnixNano() / 1000000
	ts, err := MarshalTimestampMS(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if ts.Weak() != expected {
		t.Errorf("actual:%[1]v(%[1]T), expected:%[2]v(%[2]T)", ts.Weak(), expected)
	}
}

func TestMustTimestampMS(t *testing.T) {
	v := time.Now()
	expected := v.UnixNano() / 1000000
	tests := []struct {
		name      string
		args      interface{}
		want      TimestampMS
		wantPanic bool
	}{
		{
			name: "valid",
			args: v,
			want: TimestampMS{
				ValidFlag: true,
				time:      v,
			},
			wantPanic: false,
		},
		{
			name: "panic",
			args: "valid paramenter",
			want: TimestampMS{
				ValidFlag: false,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				p := assert.Panics(t, func() {
					MustTimestampMS(tt.args)
				})
				if !p {
					t.Errorf("MustTimestampMS() panic = %v, want panic %v", p, tt.wantPanic)
				}
				return
			}
			if got := MustTimestampMS(tt.args); got.Weak() != expected {
				t.Errorf("MustTimestampMS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestampMSJsonMarshal(t *testing.T) {
	v := time.Now()
	tm := TimestampMS{
		ValidFlag: true,
		time:      v,
	}
	expected := strconv.FormatInt(v.UnixNano()/1000000, 10)
	actual, err := json.Marshal(tm)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestTimestampMSJsonMarshalValidFalse(t *testing.T) {
	tm := TimestampMS{
		ValidFlag: false,
		time:      time.Now(),
	}
	expected := []byte("null")
	actual, err := json.Marshal(tm)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != string(expected) {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestTimestampMSJsonUnmarshal(t *testing.T) {
	v := time.Now()
	in, _ := v.MarshalJSON()
	tm := TimestampMS{}
	if err := tm.UnmarshalJSON(in); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if !tm.Valid() {
		t.Error("ValidFlag should be TRUE")
	}
	if tm.Int64() != v.UnixNano()/1000000 {
		t.Errorf("actual:%d, expected:%d", tm.Int64(), v.UnixNano()/1000000)
	}
}

func TestTimestampMSJsonUnmarshalNil(t *testing.T) {
	tm := TimestampMS{}
	if err := tm.UnmarshalJSON(nil); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if tm.Valid() {
		t.Error("ValidFlag should be FALSE")
	}
	if tm.Int64() != 0 {
		t.Errorf("actual:%d, expected:%d", tm.Int64(), 0)
	}
}

func TestTimestampMSJsonUnmarshalInvalid(t *testing.T) {
	tm := TimestampMS{}
	if err := tm.UnmarshalJSON([]byte(`"a`)); err == nil {
		t.Errorf("Expected error when json.Unmarshal, but not; %#v", tm)
	}
}

func TestTimestampMSSetNil(t *testing.T) {
	tm := TimestampMS{}
	err := tm.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if _, err = tm.Value(); err != nil {
		t.Errorf("This value should return nil. error:%s", err.Error())
	}
}

func TestTimestampMSSetTime(t *testing.T) {
	v := time.Now()
	expected := v
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tm.Weak() != expected.UnixNano()/1000000 {
		t.Errorf("actual:%v, expected:%v", tm.Weak(), expected)
	}
}

func TestTimestampMSSetInt64(t *testing.T) {
	var v int64 = 1367059792
	expected := time.Unix(0, v*1000000)
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tm.Weak() != expected.UnixNano()/1000000 {
		t.Errorf("actual:%v, expected:%v", tm.Weak(), expected)
	}
}

func TestTimestampMSSetNumericString(t *testing.T) {
	v := "1467059792"
	tm := TimestampMS{}
	err := tm.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tm.Weak() != nil {
		t.Errorf("This value should return nil. value:%#v", tm.Weak())
	}
}

func TestTimestampMSSetNonNumericString(t *testing.T) {
	v := "a"
	tm := TimestampMS{}
	err := tm.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tm.Weak() != nil {
		t.Errorf("This value should return nil. value:%#v", tm.Weak())
	}
}

func TestTimestampMSSetBool(t *testing.T) {
	v := true
	tm := TimestampMS{}
	err := tm.Set(v)
	if err == nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tm.Weak() != nil {
		t.Errorf("This value should return nil. value:%#v", tm.Weak())
	}
}

func TestTimestampMSInt64(t *testing.T) {
	v := time.Now()
	expected := v.UnixNano() / 1000000
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.Int64() != expected {
		t.Errorf("This value should return %d. value:%d", expected, tm.Int())
	}
}

func TestTimestampMSInt64Zero(t *testing.T) {
	v := time.Unix(0, 0)
	var expected int64
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.Int64() != expected {
		t.Errorf("This value should return %d. value:%d", expected, tm.Int())
	}
}

func TestTimestampMSInt(t *testing.T) {
	v := time.Now()
	expected := int(v.UnixNano() / 1000000)
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.Int() != expected {
		t.Errorf("This value should return %d. value:%d", expected, tm.Int())
	}
}

func TestTimestampMSString(t *testing.T) {
	v := time.Now()
	expected := strconv.FormatInt(v.UnixNano()/1000000, 10)
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.String() != expected {
		t.Errorf("This value should return %s. value:%s", expected, tm.String())
	}
}

func TestTimestampMS_Time(t *testing.T) {
	now := time.Now()

	type fields struct {
		ValidFlag ValidFlag
		time      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "now",
			fields: fields{
				ValidFlag: true,
				time:      now,
			},
			want: now,
		},
		{
			name: "invalid",
			fields: fields{
				ValidFlag: false,
				time:      now,
			},
			want: time.Unix(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := TimestampMS{
				ValidFlag: tt.fields.ValidFlag,
				time:      tt.fields.time,
			}
			if got := v.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimestampMS.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}
