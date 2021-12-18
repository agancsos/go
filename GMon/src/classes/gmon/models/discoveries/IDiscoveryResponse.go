package discoveries

// Interface
type IDiscoveryResponse interface {
	DataType()          string
	Data()              interface{}
	SetData(a interface{})
}
/*****************************************************************************/

// Basic response
type BasicDiscoveryReponse struct {
	data         string   `json:data`
}
func (x BasicDiscoveryReponse) DataType() string { return "String"; }
func (x BasicDiscoveryReponse) Data() interface{} { return x.data; }
func (x *BasicDiscoveryReponse) SetData(a interface{}) { x.data = a.(string); }
/*****************************************************************************/

