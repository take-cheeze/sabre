package runtime_test

import (
	"testing"

	"github.com/spy16/sabre/sabre/runtime"
)

func TestList_Eval(t *testing.T) {
	t.Parallel()

	runEvalTests(t, []evalTestCase{
		{
			title: "EmptyList",
			form: &runtime.SliceList{
				Items: []runtime.Value{},
			},
			want: &runtime.SliceList{
				Items: []runtime.Value{},
			},
		},
		{
			title:  "FirstEvalFailure",
			getEnv: func() runtime.Runtime { return runtime.New(nil) },
			form: &runtime.SliceList{Items: []runtime.Value{
				runtime.Symbol{Value: "non-existent"},
			}},
			wantErr: true,
		},
		{
			title:  "NonInvokable",
			getEnv: func() runtime.Runtime { return runtime.New(nil) },
			form: &runtime.SliceList{Items: []runtime.Value{
				runtime.Int64(0),
			}},
			wantErr: true,
		},
		{
			title:  "NonInvokable",
			getEnv: func() runtime.Runtime { return runtime.New(nil) },
			form: &runtime.SliceList{Items: []runtime.Value{
				runtime.GoFunc(func(env runtime.Runtime, args ...runtime.Value) (runtime.Value, error) {
					return runtime.String("called"), nil
				}),
			}},
			want:    runtime.String("called"),
			wantErr: false,
		},
	})
}

func TestList_String(t *testing.T) {
	l := &runtime.SliceList{
		Items: []runtime.Value{
			runtime.Bool(true),
			runtime.Int64(10),
			runtime.Float64(3.1412),
		},
	}
	want := `(true 10 3.141200)`
	got := l.String()

	if want != got {
		t.Errorf("List.String() want=`%s`, got=`%s`",
			want, got)
	}
}