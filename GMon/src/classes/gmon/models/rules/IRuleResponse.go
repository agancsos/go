package rules

// Interface
type IRuleResponse interface {
    DataType()          string
    Data()              interface{}
    SetData(a interface{})
}
/*****************************************************************************/

// Basic response
type BasicRuleReponse struct {
    data         string   `json:data`
}
func (x BasicRuleReponse) DataType() string { return "String"; }
func (x BasicRuleReponse) Data() interface{} { return x.data; }
func (x *BasicRuleReponse) SetData(a interface{}) { x.data = a.(string); }
/*****************************************************************************/
