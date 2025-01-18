package util

import (
	"context"

	"github.com/goccy/go-json"
)

// ToByte converts any type to a byte slice.
func ToByte(i any) []byte {
	bt, _ := json.Marshal(i)
	return bt
}

// Dump to json using json marshal
func Dump(i any) string {
	return string(ToByte(i))
}

// bindingCtxValidator if you want to validate struct after using BindingFromContext
// you can make your own validator with this interface
type bindingCtxValidator[T any] func(value T) error

// BindingFromContext for binding any value from context to destination struct
func BindingFromContext[T any](ctx context.Context, keys []string, validator bindingCtxValidator[T]) (*T, error) {
	var (
		dataMap = make(map[string]any)
	)

	for _, key := range keys {
		val := ctx.Value(key)
		if val != nil {
			dataMap[key] = val
		}
	}

	dataByte, err := json.Marshal(dataMap)
	if err != nil {
		return nil, err
	}

	var out T
	err = json.Unmarshal(dataByte, &out)
	if err != nil {
		return nil, err
	}

	if validator != nil {
		err := validator(out)
		if err != nil {
			return nil, err
		}
	}

	return &out, nil
}

// DumpIncomingContext converts the metadata from the incoming context to a string representation using json marshal.
func DumpIncomingContext[T any](ctx context.Context, keys []string) string {
	md, _ := BindingFromContext[T](ctx, keys, nil)
	return Dump(md)
}
