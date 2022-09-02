package factory

// 规则配置解析
type IRuleConfigParser_factorySimple interface {
	Parse(data []byte)
}

// json 规则配置
type JsonRuleConfigParser_factorySimple struct {
}

// json 参数
func (J JsonRuleConfigParser_factorySimple) Parse(data []byte) {
	panic("implement me")
}

// yaml 规则配置
type YamlRuleConfigParser_factorySimple struct {
}

// yaml 参数
func (Y YamlRuleConfigParser_factorySimple) Parse(data []byte) {
	panic("implement me")
}

// 创建规则配置解析
func NewIRuleConfigParser_factorySimple(t string) IRuleConfigParser_factorySimple {
	switch t {
	case "json":
		return JsonRuleConfigParser_factorySimple{}
	case "yaml":
		return YamlRuleConfigParser_factorySimple{}
	}
	return nil
}
