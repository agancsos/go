package services
import (
	"../models"
	"../../sr"
	"strconv"
	"fmt"
	"strings"
	"math/rand"
	"../../common"
    "net/http"
    "io/ioutil"
    b64 "encoding/base64"
)

// Interface
type IAuthenticationService interface {
	Initialize();
	IsValid(username string, password string) bool;
	User(username string, password string) *models.User;
	GenerateToken(user *models.User) string;
	Groups(key string) []string;
	AddGroup(name string) bool;
	RemoveGroup(id int) bool;
	IsGroupMember(user *models.User, id int) bool;
	AddGroupMember(user *models.User, id int) bool;
	RemoveGroupMember(user *models.User, id int) bool;
}
/*****************************************************************************/

// Local service (DB handle)
type LocalAuthenticationService struct {
	ds		   *DataService
	dbts	   *DbTraceService
}
var __local_auth_service__ *LocalAuthenticationService;
func GetLocalAuthService() *LocalAuthenticationService {
	if __local_auth_service__ == nil {
		__local_auth_service__ = &LocalAuthenticationService{};
		__local_auth_service__.ds = GetDataServiceInstance().(*DataService);
		__local_auth_service__.dbts = GetDbTraceServiceInstance().(*DbTraceService);
	}
	return __local_auth_service__;
}
func (x *LocalAuthenticationService) IsValid(username string, password string) bool {
	sr.TS.TraceVerbose("AuthenticationService: Checking if valid (" + username + ")", int(sr.TC_SERVICE));
	return x.User(username, password) != nil;
}
func (x *LocalAuthenticationService) User(username string, password string) *models.User {
	sr.TS.TraceVerbose("AuthenticationService: ting user (" + username + ")", int(sr.TC_SERVICE));
	if password != "" && username != "" {
		var rawResult = x.ds.ServiceQuery("SELECT * FROM USERS WHERE USER_USERNAME = '" + username + "' AND USER_PASSWORD = '" + password + "' AND USER_STATE = '1'");
		if len(rawResult.Rows()) > 0 {
			var row = rawResult.Rows()[0];
			var result = &models.User{};
			var id, _ = strconv.Atoi(row.Column("USER_ID").Value());
			var state, _ = strconv.Atoi(row.Column("USER_STATE").Value());
			result.SetID(id);
			result.SetUsername(row.Column("USER_USERNAME").Value());
			result.SetFirstName(row.Column("USER_FIRSTNAME").Value());
			result.SetLastName(row.Column("USER_LASTNAME").Value());
			result.SetPassword(row.Column("USER_PASSWORD").Value());
			result.SetSettings(row.Column("USER_JSON").Value());
			result.SetEmail(row.Column("USER_EMAIL").Value());
			result.SetState(state);
			return result;
		}
	} else {
		var rawResult = x.ds.ServiceQuery("SELECT * FROM USERS WHERE USER_ID = (SELECT USER_ID FROM SESSION WHERE SESSION_TOKEN = '" + username + "') AND USER_STATE = '1'");
		if len(rawResult.Rows()) == 1 {
			var row = rawResult.Rows()[0];
			var result = &models.User{};
			var id, _ = strconv.Atoi(row.Column("USER_ID").Value());
			var state, _ = strconv.Atoi(row.Column("USER_STATE").Value());
			result.SetID(id);
			result.SetUsername(row.Column("USER_USERNAME").Value());
			result.SetFirstName(row.Column("USER_FIRSTNAME").Value());
			result.SetLastName(row.Column("USER_LASTNAME").Value());
			result.SetPassword(row.Column("USER_PASSWORD").Value());
			result.SetSettings(row.Column("USER_JSON").Value());
			result.SetEmail(row.Column("USER_EMAIL").Value());
			result.SetState(state);
			return result;
		}
	}
	sr.TS.TraceVerbose("AuthenticationService: Done fetching user (" + username + ")", int(sr.TC_SERVICE));
	return nil;
}
func (x *LocalAuthenticationService) GenerateToken(user *models.User) string  {
	var isValid = false;
	var maxRetries = 5;
	var current = 0;
	var token = "";
	var chars = "abcdefghijklmnopqrstuvwxyx0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$";
	var chars2 = strings.Split(chars, "");
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: Generating token (%d)", user.ID()), int(sr.TC_SERVICE));
	for i := 0; i < maxRetries; i++ {
		token = "";
		var tokenLength = (30 + rand.Int() % ((120 - 30) + 1));
		for i := 0; i < tokenLength; i++ {
			var charIndex = (rand.Int() % (len(chars2) - 1)) + 0;
			token = token + chars2[charIndex];
		}
		var rawResult = x.ds.ServiceQuery("SELECT 1 FROM SESSIONS WHERE SESSION_TOKEN = '" + token + "'");
		isValid = len(rawResult.Rows()) == 0;
		current++;
		if isValid { break; }
	}

	if x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO SESSIONS (SESSION_TOKEN, USER_ID, CREATED_DATE, LAST_UPDATED_DATE) VALUES ('%s', '%d', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)", token, user.ID())) {
		return token;
	}
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: Done generating token (%d)", user.ID()), int(sr.TC_SERVICE));
	return "";
}
func (x *LocalAuthenticationService) Groups(key string) []string {
	var result = []string{};
	if key == "*" {
		key = "";
	}
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: ting groups (%s)", key), int(sr.TC_SERVICE));
	var table = x.ds.ServiceQuery("SELECT * FROM GROUPS WHERE GROUP_LABEL LIKE '%" + key + "%'");
	for _, row := range table.Rows() {
		result = append(result, fmt.Sprintf("%s:%s", row.Column("GROUP_ID").Value(), row.Column("GROUP_LABEL").Value()));
	}
	return result;
}
func (x *LocalAuthenticationService) AddGroup(name string) bool {
	sr.TS.TraceVerbose("AuthenticationService: Adding group (" + name + ")", int(sr.TC_SERVICE));
	if len(x.ds.ServiceQuery("SELECT 1 FROM GROUPS WHERE GROUP_LABEL = '" + name + "'").Rows()) == 0 {
		return x.ds.RunServiceQuery("INSERT INTO GROUPS (GROUP_LABEL) VALUES ('" + name + "')");
	}
	return true;
}
func (x *LocalAuthenticationService) RemoveGroup(id int) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: Removing group (%d)", id), int(sr.TC_SERVICE));
	x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM GROUP_MEMBERS WHERE GROUP_ID = '%d'", id));
	x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM GROUPS WHERE GROUP_ID = '%d'", id));
	return true;
}
func (x *LocalAuthenticationService) IsGroupMember(user *models.User, id int) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: Checking group membership (%d:%d)", user.ID(), id), int(sr.TC_SERVICE));
	return len(x.ds.ServiceQuery(fmt.Sprintf("SELECT 1 FROM GROUP_MEMBERS WHERE USER_ID = '%s' AND GROUP_ID = '%d'", user.ID(), id)).Rows()) == 1;
}
func (x *LocalAuthenticationService) AddGroupMember(user *models.User, id int) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: Adding group member (%d; %d)", user.ID(), id), int(sr.TC_SERVICE));
	if len(x.ds.ServiceQuery(fmt.Sprintf("SELECT 1 FROM GROUP_MEMBERS WHERE USER_ID = '%s' AND GROUP_ID = '%d'", user.ID(), id)).Rows()) == 0 {
		if !x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO GROUP_MEMBERS (USER_ID, GROUP_ID) VALUES ('%d', '%d')", user.ID(), id)) {
			sr.TS.TraceError(fmt.Sprintf("AuthenticationService: Failed to add group member (%d;%d)", user.ID(), id), int(sr.TC_SERVICE));
		}
	} else {
		sr.TS.TraceWarning(fmt.Sprintf("AuthenticationService: User already in group (%d; %d)", user.ID(), id), int(sr.TC_SERVICE));
	}
	return true;
}
func (x *LocalAuthenticationService) RemoveGroupMember(user *models.User, id int) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AuthenticationService: Removing group member (%d;%d)", user.ID(), id), int(sr.TC_SERVICE));
	if len(x.ds.ServiceQuery(fmt.Sprintf("SELECT 1 FROM GROUP_MEMBERS WHERE USER_ID = '%d' AND GROUP_ID = '%d'", user.ID(), id)).Rows()) == 1 {
		if !x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM GROUP_MEMBERS WHERE USER_ID = '%d' AND GROUP_ID = '%d'", user.ID(), id)) {
			sr.TS.TraceError(fmt.Sprintf("AuthenticationService: Failed to remove group member (%d;%d)", user.ID(), id), int(sr.TC_SERVICE));
		}
	} else {
		sr.TS.TraceError(fmt.Sprintf("AuthenticationService: User not in group (%d;%d)", user.ID(), id), int(sr.TC_SERVICE));
	}
	return true;
}
func (x *LocalAuthenticationService) Initialize() {}
/*****************************************************************************/

// Rest service
type RestAuthenticationService struct {
	las    *LocalAuthenticationService
}
var __rest_auth_service__ *RestAuthenticationService;
func GetRestAuthService() *RestAuthenticationService {
    if __rest_auth_service__ == nil {
        __rest_auth_service__ = &RestAuthenticationService{};
		__rest_auth_service__.las = GetLocalAuthService();
    }
    return __rest_auth_service__;
}
func (x *RestAuthenticationService) IsValid(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
    var body, _ = ioutil.ReadAll(r.Body);
    var json = common.StrToDictionary(body);
    if json["credentials"] == nil {
        w.Write([]byte("{\"message\":\"Credentials missing\"}"));
        return;
    }
    var decoded, _ = b64.StdEncoding.DecodeString(json["credentials"].(string));
    if string(decoded) == "" {
        w.Write([]byte("{\"message\":\"Failed to decode input\"}"));
        return;
    }
    var comps = strings.Split(string(decoded), ":");
    if len(comps) < 2 {
        w.Write([]byte("{\"message\":\"Failed to find components\"}"));
        return;
    }
    var username = comps[0];
    var password = comps[1];
    if !x.las.IsValid(username, password) {
        w.Write([]byte("{\"message\":\"User is invalid\"}"));
        return;
    }
    var user = x.las.User(username, password);
    var token = x.las.GenerateToken(user);
    if token == "" {
        w.Write([]byte("{\"message\":\"Failed to generate token\"}"));
        return;
    }
    w.Write([]byte(fmt.Sprintf("{\"token\":\"%s\"}", token)));
}
func (x *RestAuthenticationService) Groups(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
	if !x.las.IsValid(sr.ExtractApiToken(r), "") {
		w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
		return;
	}
	var search, _ = ioutil.ReadAll(r.Body);
	var s = string(search);
	var groups = x.las.Groups(s);
	var rsp = fmt.Sprintf("{\"groups\":[");
	for i, value := range groups {
		if i > 0 {
			rsp += ",";
		}
		rsp += fmt.Sprintf("\"%s\"", value);
	}
	rsp += "]}";
	w.Write([]byte(rsp));
}
func (x *RestAuthenticationService) AddGroup(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
	if !x.las.IsValid(sr.ExtractApiToken(r), "") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
        return;
    }
	var data, _ = ioutil.ReadAll(r.Body);
	var group = string(data);
	var result = x.las.AddGroup(group);
	w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(result))));
}
func (x *RestAuthenticationService) RemoveGroup(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
	if !x.las.IsValid(sr.ExtractApiToken(r), "") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
        return;
    }
	var data, _ = ioutil.ReadAll(r.Body);
    var id, _ = strconv.Atoi(string(data));
    var result = x.las.RemoveGroup(id);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(result))));
}
func (x *RestAuthenticationService) IsGroupMember(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
	if !x.las.IsValid(sr.ExtractApiToken(r), "") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
        return;
    }
	var data, _ = ioutil.ReadAll(r.Body);
    var obj = common.StrToDictionary(data);
	var user = &models.User{};
	var userId, _ = strconv.Atoi(obj["user"].(string));
	user.SetID(userId);
	var groupId, _ = strconv.Atoi(obj["group"].(string));
    var result = x.las.IsGroupMember(user, groupId);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(result))));
}
func (x *RestAuthenticationService) AddGroupMember(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
	if !x.las.IsValid(sr.ExtractApiToken(r), "") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
        return;
    }
	var data, _ = ioutil.ReadAll(r.Body);
    var obj = common.StrToDictionary(data);
    var user = &models.User{};
    var userId, _ = strconv.Atoi(obj["user"].(string));
    user.SetID(userId);
    var groupId, _ = strconv.Atoi(obj["group"].(string));
    var result = x.las.AddGroupMember(user, groupId);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(result))));
}
func (x *RestAuthenticationService) RemoveGroupMember(w http.ResponseWriter, r *http.Request) {
    if !common.EnsureRestMethod(r, "POST") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return;
    }
	if !x.las.IsValid(sr.ExtractApiToken(r), "") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
        return;
    }
	var data, _ = ioutil.ReadAll(r.Body);
    var obj = common.StrToDictionary(data);
    var user = &models.User{};
    var userId, _ = strconv.Atoi(obj["user"].(string));
    user.SetID(userId);
    var groupId, _ = strconv.Atoi(obj["group"].(string));
    var result = x.las.RemoveGroupMember(user, groupId);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", result)));
}
func (x *RestAuthenticationService) Initialize() {
    http.HandleFunc("/auth/", x.IsValid);
    http.HandleFunc("/group/add/", x.AddGroup);
    http.HandleFunc("/group/remove/", x.RemoveGroup);
    http.HandleFunc("/group/getall/", x.Groups);
    http.HandleFunc("/group/list/", x.Groups);
    http.HandleFunc("/group/ismember/", x.IsGroupMember);
    http.HandleFunc("/group/addmember/", x.AddGroupMember);
    http.HandleFunc("/group/removemember/", x.RemoveGroupMember);
}
func (x *RestAuthenticationService) User(w http.ResponseWriter, r *http.Request) {}
func (x *RestAuthenticationService) GenerateToken(w http.ResponseWriter, r *http.Request)  {}
/*****************************************************************************/

