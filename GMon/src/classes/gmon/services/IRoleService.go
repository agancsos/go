package services
import (
	"../models"
	"fmt"
    "strconv"
    "io/ioutil"
	"net/http"
	"../../common"
)

// Interface
type IRoleService interface {
	GetRoles()                 []*models.Role
	GetRole(a int)             *models.Role
	AddRole(a *models.Role)    bool
	UpdateRole(a *models.Role) bool
	DeleteRole(a *models.Role) bool
}
/*****************************************************************************/

// Local service
type LocalRoleService struct {
	ds         *DataService
}
var __local_role_service__ *LocalRoleService;
func GetLocalRoleServiceInstance() IService {
	if __local_role_service__ == nil {
		__local_role_service__ = &LocalRoleService{};
		__local_role_service__.ds = GetDataServiceInstance().(*DataService);
	}
	return __local_role_service__;
}
func (x *LocalRoleService) GetRole(a int) *models.Role {
	var result *models.Role;
	var raw = x.ds.ServiceQuery(fmt.Sprintf("SELECT * FROM ROLES WHERE ROLE_ID = '%d'", a)).Rows();
	if len(raw) > 0 {
		result = &models.Role{};
		var row = raw[0];
		id, _ := strconv.Atoi(row.Column("ROLE_ID").Value());
		result.SetID(id);
		result.SetName(row.Column("ROLE_NAME").Value());
	}
	return result;
}
func (x *LocalRoleService) GetRoles() []*models.Role {
	var result = []*models.Role {};
	var rows = x.ds.ServiceQuery("SELECT ROLE_ID FROM ROLES").Rows();
	for _, row := range rows {
		id,_ := strconv.Atoi(row.Column("ROLE_ID").Value());
		result = append(result, x.GetRole(id));
	}
	return result;
}
func (x *LocalRoleService) AddRole(a *models.Role) bool {
	if len(x.ds.ServiceQuery(fmt.Sprintf("SELECT 1 FROM ROLES WHERE ROLE_NAME = '%s'", a.Name())).Rows()) == 0 {
		return x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO ROLES (ROLE_NAME) VALUES ('%s')", a.Name()));
	}
	return true;
}
func (x *LocalRoleService) UpdateRole(a *models.Role) bool {
	if len(x.ds.ServiceQuery(fmt.Sprintf("SELECT 1 FROM ROLES WHERE ROLE_NAME = '%s'", a.Name())).Rows()) != 0 {
        return x.ds.RunServiceQuery(fmt.Sprintf("UPDATE ROLES SET ROLE_NAME = '%s', LAST_UPDATED_DATE = CURRENT_TIMESTAMP WHERE ROLE_ID = '%d'", a.Name(), a.ID()));
    }
    return true;
}
func (x *LocalRoleService) DeleteRole(a *models.Role) bool {
	return x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM ROLES WHERE ROLE_ID = '%d'", a.ID()));
}
func (x *LocalRoleService) Initialize() {}
/*****************************************************************************/

// Rest service
type RestRoleService struct {
	lrs          *LocalRoleService
	baseEndpoint string
}
var __rest_role_service__ *RestRoleService;
func GetRestRoleServiceInstance() *RestRoleService {
	if __rest_role_service__ == nil {
		__rest_role_service__ = &RestRoleService{};
		__rest_role_service__.lrs = GetLocalRoleServiceInstance().(*LocalRoleService);
		__rest_role_service__.baseEndpoint = "role";
	}
	return __rest_role_service__;
}
func (x *RestRoleService) GetRole(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") { return; }
	raw, _ := ioutil.ReadAll(r.Body);
	id,_ := strconv.Atoi(string(raw));
	var role = x.lrs.GetRole(id);
	w.Write([]byte(role.ToJsonString()));
}
func (x *RestRoleService) GetRoles(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "GET") { return; }
	var result = "{\"roles\":[";
	var roles = x.lrs.GetRoles();
	for i, role := range roles {
		if i > 0 {
			result += ",";
		}
		result += role.ToJsonString();
	}
	result += "]}";
	w.Write([]byte(result));
}
func (x *RestRoleService) AddRole(w http.ResponseWriter, r *http.Request){
	if !EnsureAuthenticated(w, r, "POST") { return; }
	raw, _ := ioutil.ReadAll(r.Body);
	var dict = common.StrToDictionary(raw);
	var role = dict["role"].(*models.Role);
	w.Write([]byte(fmt.Sprintf("\"result\":\"%d\"}", common.BoolToInt(x.lrs.AddRole(role)))));
}
func (x *RestRoleService) UpdateRole(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") { return; }
	raw, _ := ioutil.ReadAll(r.Body);
	var dict = common.StrToDictionary(raw);
    var role = dict["role"].(*models.Role);
    w.Write([]byte(fmt.Sprintf("\"result\":\"%d\"}", common.BoolToInt(x.lrs.UpdateRole(role)))));
}
func (x *RestRoleService) DeleteRole(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") { return; }
	raw, _ := ioutil.ReadAll(r.Body);
	var dict = common.StrToDictionary(raw);
    var role = dict["role"].(*models.Role);
    w.Write([]byte(fmt.Sprintf("\"result\":\"%d\"}", common.BoolToInt(x.lrs.DeleteRole(role)))));
}
func (x *RestRoleService) Initialize() {
	http.HandleFunc(fmt.Sprintf("/%s/get/", x.baseEndpoint), x.GetRole);
	http.HandleFunc(fmt.Sprintf("/%s/list/", x.baseEndpoint), x.GetRoles);
	http.HandleFunc(fmt.Sprintf("/%s/add/", x.baseEndpoint), x.AddRole);
	http.HandleFunc(fmt.Sprintf("/%s/update/", x.baseEndpoint), x.UpdateRole);
	http.HandleFunc(fmt.Sprintf("/%s/delete/", x.baseEndpoint), x.DeleteRole);
}
/*****************************************************************************/
