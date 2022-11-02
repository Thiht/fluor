package fluor

import (
	"errors"
	"testing"
)

func TestStdErrors_Is(t *testing.T) {
	exampleError := errors.New("example")

	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"New error is not an exampleError",
			args{
				err:    New(),
				target: exampleError,
			},
			false,
		},
		{
			"New error is a FluentErrorKindDefault",
			args{
				err:    New(),
				target: KindDefault,
			},
			true,
		},
		{
			"New error wrapping exampleError is an exampleError",
			args{
				err:    New().WithError(exampleError),
				target: exampleError,
			},
			true,
		},
		{
			"New error wrapping exampleError is a FluentErrorKindDefault",
			args{
				err:    New().WithError(exampleError),
				target: KindDefault,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("FluentError.WithHTTPStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
