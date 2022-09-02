package factory

// 规则配置解析
type IRuleConfigParser_model interface {
	Parse(data []byte)
}

// 工厂方法接口
type IRuleConfigParserFactory_model interface {
	CreateParser() IRuleConfigParser_model
}

// json 规则配置
type jsonRuleConfigParser_model struct {
}

// yaml 规则配置
type yamlRuleConfigParser_model struct {
}

// json 参数
func (J jsonRuleConfigParser_model) Parse(data []byte) {
	panic("implement me")
}

// yaml 参数
func (Y yamlRuleConfigParser_model) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory_model struct {
}

// CreateParser CreateParser
func (y yamlRuleConfigParserFactory_model) CreateParser() IRuleConfigParser_model {
	return yamlRuleConfigParser_model{}
}

// jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory_model struct {
}

// json 创建参数
func (j jsonRuleConfigParserFactory_model) CreateParser() IRuleConfigParser_model {
	return jsonRuleConfigParser_model{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory_model(t string) IRuleConfigParserFactory_model {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory_model{}
	case "yaml":
		return yamlRuleConfigParserFactory_model{}
	}
	return nil
}
