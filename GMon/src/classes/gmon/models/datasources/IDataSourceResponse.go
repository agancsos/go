package datasources
import (
	"fmt"
)

// Interface
type IDataSource interface {
	Invoke()               IDataSourceResponse
}

type IDataSourceResponse interface {
    DataType()          string
    Data()              interface{}
    SetData(a interface{})
}
/*****************************************************************************/

// Static data source
type StaticDataSource struct {
    data        []interface{}
	delimeter   string
}
func (x *StaticDataSource) Invoke() IDataSourceResponse {
    var result = &BasicDataSourceReponse{};
    var temp = "";
    for i, val := range x.data {
        if i > 0 {
            temp += x.delimeter;
        }
        temp += fmt.Sprintf("%v", val);
    }
    result.SetData(temp);
    return result;
}
func (x StaticDataSource) Delimeter() string { return x.delimeter; }
func (x *StaticDataSource) SetDelimeter(a string) { x.delimeter = a; }
/*****************************************************************************/

// Basic response
type BasicDataSourceReponse struct {
    data         string   `json:data`
}
func (x BasicDataSourceReponse) DataType() string { return "String"; }
func (x BasicDataSourceReponse) Data() interface{} { return x.data; }
func (x *BasicDataSourceReponse) SetData(a interface{}) { x.data = a.(string); }
/*****************************************************************************/

