package monitors

// Interface
type IMonitorResponse interface {
    DataType()          string
    Data()              interface{}
    SetData(a interface{})
}
/*****************************************************************************/

// Basic response
type BasicMonitorReponse struct {
    data         string   `json:data`
}
func (x *BasicMonitorReponse) GetDataType() string { return "String"; }
func (x *BasicMonitorReponse) GetData() interface{} { return x.data; }
func (x *BasicMonitorReponse) SetData(a interface{}) { x.data = a.(string); }
/*****************************************************************************/
