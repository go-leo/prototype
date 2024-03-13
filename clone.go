package prototype

import (
	"context"
	"reflect"
	"strings"
	"time"
)

// Cloner is a clone hook.
type Cloner func(labels []string, tgtVal reflect.Value, srcVal reflect.Value) (bool, error)

type options struct {
	TagKey           string
	DeepClone        bool
	EqualFold        func(t, s string) bool
	IntToTime        func(i int64) time.Time
	StringToTime     func(s string) (time.Time, error)
	TimeToInt        func(t time.Time) int64
	TimeToString     func(t time.Time) string
	GetterPrefix     string
	SetterPrefix     string
	Context          context.Context
	Cloners          []Cloner
	InterruptOnError bool
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) correct() *options {
	if o.EqualFold == nil {
		o.EqualFold = strings.EqualFold
	}
	if o.IntToTime == nil {
		o.IntToTime = func(i int64) time.Time { return time.Unix(i, 0) }
	}
	if o.StringToTime == nil {
		o.StringToTime = func(s string) (time.Time, error) { return time.ParseInLocation(time.DateTime, s, time.Local) }
	}
	if o.TimeToInt == nil {
		o.TimeToInt = func(t time.Time) int64 { return t.Unix() }
	}
	if o.TimeToString == nil {
		o.TimeToString = func(t time.Time) string { return t.Local().Format(time.DateTime) }
	}
	if o.Context == nil {
		o.Context = context.TODO()
	}
	return o
}

type Option func(o *options)

func TagKey(key string) Option {
	return func(o *options) {
		o.TagKey = key
	}
}

func DisableDeepClone() Option {
	return func(o *options) {
		o.DeepClone = false
	}
}

func EqualFold(f func(t, s string) bool) Option {
	return func(o *options) {
		o.EqualFold = f
	}
}

func IntToTime(f func(i int64) time.Time) Option {
	return func(o *options) {
		o.IntToTime = f
	}
}

func StringToTime(f func(s string) (time.Time, error)) Option {
	return func(o *options) {
		o.StringToTime = f
	}
}

func TimeToInt(f func(t time.Time) int64) Option {
	return func(o *options) {
		o.TimeToInt = f
	}
}

func TimeToString(f func(t time.Time) string) Option {
	return func(o *options) {
		o.TimeToString = f
	}
}

func GetterPrefix(p string) Option {
	return func(o *options) {
		o.GetterPrefix = p
	}
}

func SetterPrefix(p string) Option {
	return func(o *options) {
		o.SetterPrefix = p
	}
}

func Context(ctx context.Context) Option {
	return func(o *options) {
		o.Context = ctx
	}
}

func Cloners(f ...Cloner) Option {
	return func(o *options) {
		o.Cloners = append(o.Cloners, f...)
	}
}

func InterruptOnError() Option {
	return func(o *options) {
		o.InterruptOnError = true
	}
}

// Clone 将值从 src 克隆到 tgt
func Clone(tgt any, src any, opts ...Option) error {
	// 获取目标值的反射值
	tgtVal := reflect.ValueOf(tgt)
	// 如果目标不是一个指针或者为nil，返回错误
	if tgtVal.Kind() != reflect.Pointer {
		return newNonPointerError(reflect.TypeOf(tgt), reflect.TypeOf(src))
	}
	if tgtVal.IsNil() {
		return newNilError(reflect.TypeOf(tgt), reflect.TypeOf(src))
	}

	// 初始化options
	o := new(options).apply(opts...).correct()

	// 从对象池中获取克隆状态上下文
	ctx := newCloneContext(o)
	// 克隆结束后，将克隆状态上下文放入对象池中
	defer freeCloneContext(ctx)

	// 处理对象克隆逻辑
	return clone(ctx, tgtVal, reflect.ValueOf(src), o)
}

func clone(g *cloneContext, tgtVal, srcVal reflect.Value, opts *options) error {
	return valueCloner(srcVal, opts)(g, []string{}, tgtVal, srcVal, opts)
}
