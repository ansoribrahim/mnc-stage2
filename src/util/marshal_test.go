package util

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type testJSONStruct struct {
	Test string `json:"test"`
}

func TestToByte(t *testing.T) {
	respSuccess, _ := json.Marshal(testJSONStruct{})
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "normal test",
			args: args{i: testJSONStruct{}},
			want: respSuccess,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToByte(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDump(t *testing.T) {
	respSuccess, _ := json.Marshal(testJSONStruct{})
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "normal test",
			args: args{i: testJSONStruct{}},
			want: respSuccess,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Dump(tt.args.i)
		})
	}
}

func TestBindingFromContext(t *testing.T) {
	var (
		ctx      = context.Background()
		mockUUID = uuid.New()

		patcher = gomonkey.NewPatches()
	)

	ctx = context.WithValue(ctx, "id", mockUUID)              //nolint
	ctx = context.WithValue(ctx, "name", "bagas kertenagera") //nolint

	type testStruct struct {
		ID   uuid.UUID
		Name string
	}

	type args[T any] struct {
		ctx       context.Context
		keys      []string
		validator bindingCtxValidator[T]
	}
	type testCase[T any] struct {
		name        string
		args        args[T]
		want        *T
		wantErr     error
		prepareMock func()
	}
	tests := []testCase[testStruct]{
		{
			name: "normal",
			args: args[testStruct]{
				ctx:  ctx,
				keys: []string{"id", "name"},
				validator: func(value testStruct) error {
					if value.Name != "" {
						return nil
					}

					return errors.New("err")
				},
			},
			want: &testStruct{
				ID:   mockUUID,
				Name: "bagas kertenagera",
			},
			wantErr: nil,
		},
		{
			name: "validator err",
			args: args[testStruct]{
				ctx:  ctx,
				keys: []string{"id", "name"},
				validator: func(value testStruct) error {
					if value.Name != "" {
						return errors.New("err")
					}

					return nil
				},
			},
			want:    nil,
			wantErr: errors.New("err"),
		},
		{
			name: "json marshal err",
			args: args[testStruct]{
				ctx:  ctx,
				keys: []string{"id", "name"},
				validator: func(value testStruct) error {
					if value.Name != "" {
						return nil
					}

					return nil
				},
			},
			want:    nil,
			wantErr: errors.New("err"),
			prepareMock: func() {
				patcher.ApplyFunc(json.Marshal, func(v interface{}) ([]byte, error) {
					return nil, errors.New("err")
				})
			},
		},
		{
			name: "json unmarshal err",
			args: args[testStruct]{
				ctx:  ctx,
				keys: []string{"id", "name"},
				validator: func(value testStruct) error {
					if value.Name != "" {
						return nil
					}

					return nil
				},
			},
			want:    nil,
			wantErr: errors.New("err"),
			prepareMock: func() {
				patcher.ApplyFunc(json.Unmarshal, func(data []byte, v interface{}) error {
					return errors.New("err")
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer patcher.Reset()
			if tt.prepareMock != nil {
				tt.prepareMock()
			}

			got, err := BindingFromContext(tt.args.ctx, tt.args.keys, tt.args.validator)
			assert.Equalf(t, tt.wantErr, err, "BindingFromContext(%v, %v, %v)", tt.args.ctx, tt.args.keys, tt.args.validator)
			assert.Equalf(t, tt.want, got, "BindingFromContext(%v, %v, %v)", tt.args.ctx, tt.args.keys, tt.args.validator)
		})
	}
}

func TestDumpIncomingContext(t *testing.T) {
	var (
		ctx      = context.Background()
		mockUUID = uuid.New()
	)

	ctx = context.WithValue(ctx, "id", mockUUID)              //nolint
	ctx = context.WithValue(ctx, "name", "bagas kertenagera") //nolint

	type testStruct struct {
		ID   uuid.UUID
		Name string
	}

	type args struct {
		ctx  context.Context
		keys []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				ctx:  ctx,
				keys: []string{"id", "name"},
			},
			want: fmt.Sprintf(`{"ID":"%s","Name":"bagas kertenagera"}`, mockUUID.String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DumpIncomingContext[testStruct](tt.args.ctx, tt.args.keys), "DumpIncomingContext(%v, %v)", tt.args.ctx, tt.args.keys)
		})
	}
}
