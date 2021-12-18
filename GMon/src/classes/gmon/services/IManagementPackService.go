package services
import (
	"../models"
	"fmt"
	"net/http"
	"../../common"
	"strconv"
	"io/ioutil"
)

// Interface
type IManagementPackService interface{
	ImportMP(a string)                 bool
	ExportMP(a string)                 string
	AddMP(a *models.ManagementPack)    bool
	DeleteMP(a *models.ManagementPack) bool
	GetManagementPacks()               []*models.ManagementPack
	GetManagementPack(a int)           *models.ManagementPack
}
/*****************************************************************************/

// Local service
type LocalManagementPackService struct {
	ds     *DataService
}
var __local_mp_service__ *LocalManagementPackService;
func GetLocalMPServiceInstance() IService {
	if __local_mp_service__ == nil {
		__local_mp_service__ = &LocalManagementPackService{};
		__local_mp_service__.ds = GetDataServiceInstance().(*DataService);
	}
	return __local_mp_service__;
}

// wip - Pending database design
func (x *LocalManagementPackService) ImportMP(a string) bool {
	var dict = common.StrToDictionary([]byte(a));
	if x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO MANAGEMENT_PACKS (MANAGEMENT_PACK_VERSION, MANAGEMENT_PACK_NAMGE) VALUES ('%S', '%S')", dict["version"].(string), dict["name"].(string))) {
		var rows = x.ds.ServiceQuery(fmt.Sprintf("SELECT * FROM MANAGEMENT_PACKS WHERE MANAGEMENT_PACK_NAME = '%s'", dict["name"].(string))).Rows();
		var row = rows[0];
		id, _ := strconv.Atoi(row.Column("MANAGEMENT_PACK_ID").Value());
		var mp = &models.ManagementPack{};
		mp.ReloadFromJson(fmt.Sprintf("%v", dict["manifest"]));
		mp.SetID(id);
		return x.AddMP(mp);
	} else {
		return false;
	}
	return false;
}
func (x *LocalManagementPackService) ExportMP(a string) string {
	id, _ := strconv.Atoi(a);
	var mp = x.GetManagementPack(id);
	return mp.ToJsonString();
}
func (x *LocalManagementPackService) AddMP(a *models.ManagementPack) bool {
	return x.ds.RunServiceQuery(fmt.Sprintf("UPDATE MANAGEMENT_PACKS SET MANAGEMENT_PACK_MANIFEST = '%s' WHERE MANAGEMENT_PACK_ID = '%d'", a.ToJsonString(), a.ID()));
}
func (x *LocalManagementPackService) DeleteMP(a *models.ManagementPack) bool {
	return x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM MANAGEMENT_PACKS WHERE MANAGEMENT_PACK_ID = '%d'", a.ID()));
}
func (x *LocalManagementPackService) GetManagementPacks() []*models.ManagementPack {
	var result = []*models.ManagementPack{};
	var rows = x.ds.ServiceQuery(fmt.Sprintf("SELECT MANAGEMENT_PACK_ID FROM MANAGEMENT_PACKS")).Rows();
	for _, row := range rows {
		id, _ := strconv.Atoi(row.Column("MANAGEMENT_PACK_ID").Value());
		result = append(result, x.GetManagementPack(id));
	}
	return result;
}
func (x *LocalManagementPackService) GetManagementPack(a int) *models.ManagementPack {
	var result = &models.ManagementPack{};
	var rows = x.ds.ServiceQuery(fmt.Sprintf("SELECT MANAGEMENT_PACK_ID FROM MANAGEMENT_PACKS WHERE MANAGEMENT_PACK_ID = '%d'", a)).Rows();
	if len(rows) == 1 {
		var row = rows[0];
		result.ReloadFromJson(row.Column("MANAGEMENT_PACK_MANIFEST").Value());
	}
	return result;
}
func (x *LocalManagementPackService) Initialize() {
}
/*****************************************************************************/

// Rest service
type RestManagementPackService struct {
    mps     *LocalManagementPackService
}
var __rest_mp_service__ *RestManagementPackService;
func GetRestMPServiceInstance() *RestManagementPackService {
    if __rest_mp_service__ == nil {
        __rest_mp_service__ = &RestManagementPackService{};
        __rest_mp_service__.mps = GetLocalMPServiceInstance().(*LocalManagementPackService);
    }
    return __rest_mp_service__;
}
func (x *RestManagementPackService) ImportMP(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
		return;
	}
	data, _ := ioutil.ReadAll(r.Body);
	w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", x.mps.ImportMP(string(data)))));
}
func (x *RestManagementPackService) ExportMP(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
        return;
    }
	data, _ := ioutil.ReadAll(r.Body);
	var pack = x.mps.ExportMP(string(data));
	w.Write([]byte(pack));
}
func (x *RestManagementPackService) AddMP(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
        return;
    }
	data, _ := ioutil.ReadAll(r.Body);
	var mp = &models.ManagementPack{};
	mp.ReloadFromJson(string(data));
	w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", x.mps.AddMP(mp))));
}
func (x *RestManagementPackService) DeleteMP(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
        return;
    }
	data, _ := ioutil.ReadAll(r.Body);
    var mp = &models.ManagementPack{};
    mp.ReloadFromJson(string(data));
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", x.mps.DeleteMP(mp))));
}
func (x *RestManagementPackService) GetManagementPacks(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
        return;
    }
	var rsp = "{\"managementPacks\":[";
	var packs = x.mps.GetManagementPacks();
	for i, pack := range packs {
		if i > 0 {
			rsp += ",";
		}
		rsp += pack.ToJsonString();
	}
	rsp += "]";
	w.Write([]byte(rsp));
}
func (x *RestManagementPackService) Initialize() {
	http.HandleFunc("/mp/import/", x.ImportMP);
	http.HandleFunc("/mp/export/", x.ExportMP);
	http.HandleFunc("/mp/add/", x.AddMP);
	http.HandleFunc("/mp/delete/", x.DeleteMP);
	http.HandleFunc("/mp/list/", x.GetManagementPacks);
}
/*****************************************************************************/
