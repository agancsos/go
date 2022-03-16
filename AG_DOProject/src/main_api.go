package main
import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"net/http"
	"./classes/sr"
	"./classes/helpers"
	"os"
	"./classes/common"
	"path/filepath"
)

// RestService
type RestService struct {
	servicePort	   int
	serviceName	   string
	basePath       string
	packages	   map[string]*helpers.Package
}
var rs *RestService;
func GetRestServiceInstance() *RestService {
	if rs == nil {
		rs = &RestService{}
		var cs = helpers.GetConfigurationServiceInstance();
		var err error;
		rs.basePath = cs.GetProperty("serverBasePath", fmt.Sprintf("%s/agdo", (sr.GetSRInstance()).SS.BuildModuleContainerPath())).(string);
		rs.servicePort, err = strconv.Atoi(cs.GetProperty("apiPort", "4455").(string));
		if err != nil { panic(err); }
		rs.serviceName = "Abel Gancsos DigitalOcean Project REST Service";
		rs.packages = map[string]*helpers.Package{};
		rs.loadPackages();
	}
	return rs;
}

func (x *RestService) loadPackages() {
	var _, err = os.Stat(x.basePath);
	if err != nil { fmt.Printf("Repository not found.  %v\n", err); }
	containers, err := ioutil.ReadDir(x.basePath);
	if err != nil { fmt.Printf("Failed to read repository.  %v\n", err); }
	for _, p := range containers {
		if strings.Contains(p.Name(), ".agdo") {
			rawText := common.RunCmd(fmt.Sprintf("tar xfO %s/%s %s/package.json", x.basePath, p.Name(), strings.Replace(p.Name(), ".agdo", "", -1)));
			if rawText != "" {
				var tempPackage *helpers.Package;
				json.Unmarshal([]byte(rawText), &tempPackage);
				if tempPackage != nil {
					x.packages[tempPackage.Name] = tempPackage;
				}
			}
		}
	}
}

func (x *RestService) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("{\"version\":\"%s\"}", sr.APPLICATION_VERSION)));
}

func (x *RestService) GetHandler(w http.ResponseWriter, r *http.Request) {
	if !common.EnsureRestMethod(r, "GET") { return; }
	var rawContent, err = ioutil.ReadFile(fmt.Sprintf("%s/../index.htm", x.basePath));
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)));
	} else {
		w.Write([]byte(rawContent));
	}
}

func (x *RestService) ListHandler(w http.ResponseWriter, r *http.Request) {
	if !common.EnsureRestMethod(r, "GET") { return; }
	var p = "";
	var comps = strings.Split(r.URL.Path, "/");
	if len(comps) > 3 {
		p = comps[3];
	}
	var result = "{\"packages\":[";
	var i = 0;
	for _, currentPackage := range x.packages {
		if i > 0 { result += ","; }
		if p == "" || p == currentPackage.Name {
			var serial, err = json.Marshal(currentPackage);
			if err != nil { continue; }
			result += string(serial);
		}
		i++;
	}
	result += "]}";
	w.Write([]byte(result));
}

func (x *RestService) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	var p = "";
	var comps = strings.Split(r.URL.Path, "/");
    if len(comps) > 3 {
        p = comps[3];
    }
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", p));
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	if !strings.Contains(p, ".agdo") { p += ".agdo"; }
	http.ServeFile(w, r, fmt.Sprintf("%s/%s", x.basePath, p));
}

func (x *RestService) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Handle multi-part file upload
	// https://tutorialedge.net/golang/go-file-upload-tutorial/
	r.ParseMultipartForm(10 << 10)
	file, handle, err := r.FormFile("package");
	if err != nil && handle == nil {
		w.Write([]byte(fmt.Sprintf("{\"result\":\"Failed to handle file upload...%v\"}", err)));
		return;
	}
	_, err = os.Stat(fmt.Sprintf("%s/%s.lock", x.basePath, handle.Filename));
	if !os.IsNotExist(err) {
		w.Write([]byte("{\"result\":\"File already being modified...\"}"));
        return;
	}
	if !strings.Contains(handle.Filename, ".agdo") {
		w.Write([]byte("{\"result\":\"File is not a package file...\"}"));
        return;
	}
	defer file.Close();
	ioutil.WriteFile(fmt.Sprintf("%s/%s.lock", x.basePath, handle.Filename), []byte(""), 0755);
	tempFile, err := ioutil.TempFile(x.basePath, "agdo_*.agdo")
	if err != nil {
		w.Write([]byte("{\"result\":\"Failed to create temporary file...\"}"));
		os.Remove(fmt.Sprintf("%s/%s.lock", x.basePath, handle.Filename));
        return;
	}
	defer tempFile.Close();
	data, err := ioutil.ReadAll(file);
	if err != nil {
        w.Write([]byte("{\"result\":\"Failed to read temporary file...\"}"));
		os.Remove(fmt.Sprintf("%s/%s.lock", x.basePath, handle.Filename));
        return;
    }
	tempFile.Write(data);

	// Move package into place and add to cache
	rawText := common.RunCmd(fmt.Sprintf("tar xfO %s/%s %s/package.json", x.basePath, handle.Filename, strings.Replace(handle.Filename, ".agdo", "", -1)));
	if rawText != "" {
		var tempPackage *helpers.Package;
		json.Unmarshal([]byte(rawText), &tempPackage);
		if tempPackage != nil {
			x.packages[tempPackage.Name] = tempPackage;
		}
	}
	tempFile.Close();
	os.Rename(tempFile.Name(), fmt.Sprintf("%s/%s", x.basePath, handle.Filename));
	os.Remove(fmt.Sprintf("%s/%s.lock", x.basePath, handle.Filename));
	w.Write([]byte("{\"result\":\"1\"}"));
}

func (x *RestService) Initialize() {
	// Set handlers
	http.HandleFunc("/api/version/", x.GetVersion);
	http.HandleFunc("/api/", x.GetHandler);
	http.HandleFunc("/api/upgrade/", x.ListHandler);
	http.HandleFunc("/api/update/", x.ListHandler);
	http.HandleFunc("/api/upload/", x.UploadHandler);
	http.HandleFunc("/api/download/", x.DownloadHandler);

	// Start listener
	http.ListenAndServe(fmt.Sprintf(":%d", x.servicePort), nil);
}
/*****************************************************************************/

func main() {
	var SRI			  = sr.GetSRInstance();
	var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	var ss			  = &common.SystemService{};
	ss.ModulePath	  = binaryPath;
	SRI.SS			  = ss;

	var rest = GetRestServiceInstance();
	rest.Initialize();
	os.Exit(0);
}

