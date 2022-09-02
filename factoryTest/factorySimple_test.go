package factory

import (
	factory "design/factory"
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser_factorySimple(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want factory.IRuleConfigParser_factorySimple
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: factory.JsonRuleConfigParser_factorySimple{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: factory.YamlRuleConfigParser_factorySimple{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.NewIRuleConfigParser_factorySimple(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
