package helpers
import (
	"bytes"
	"../common"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/user"
	"strings"
	"strconv"
)

// PackageDependency
type PackageDependency struct {
	Name	string
	Version string
}
func (x PackageDependency) ToJsonString() string {
	rawJson, _ := json.Marshal(x);
	return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
/*****************************************************************************/

// Package
type Package struct {
	Name		 string
	Version	     string
	Author	     string
	Dependencies []*PackageDependency
	repository   *Repository
}
func (x *Package) AddDependency(a *PackageDependency) {
	if x.Dependencies == nil {
		x.Dependencies = []*PackageDependency{};
	}
	x.Dependencies = append(x.Dependencies, a);
}
func (x Package) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
func (x Package) Repository() *Repository { return x.repository; }
func (x *Package) SetRepository(repo *Repository) { x.repository = repo; }
/*****************************************************************************/

// Repository
type Repository struct {
	Name        string
    Url         string
	basePath    string
	Packages    map[string]*Package
}
func (x *Repository) loadPackages(basePath string) {
    x.Packages = map[string]*Package{};
	x.basePath = basePath;
    var _, err = os.Stat(basePath);
    if err != nil { fmt.Printf("Local cache not found.  %v\n", err); }
    containers, err := ioutil.ReadDir(basePath);
    if err != nil { fmt.Printf("Failed to read local cache.  %v\n", err); }
    for _, container := range containers {
        _, err = os.Stat(fmt.Sprintf("%s/%s/package.json", x.basePath, container.Name()));
        if !os.IsNotExist(err) {
            var tempPackage *Package;
            rawText, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/package.json", x.basePath, container.Name()));
            if err == nil {
                json.Unmarshal(rawText, &tempPackage);
                if tempPackage != nil {
                    x.Packages[tempPackage.Name] = tempPackage;
                }
            }
        }
    }
}
func(x Repository) BasePath() string { return x.basePath; }
func(x *Repository) SetBasePath(path string) { x.basePath = path; }
/*****************************************************************************/

// RegistryCache
type RegistryCache struct {
	basePath	    string
	repositories    map[string]*Repository
	cs		        *ConfigurationService
	url		        string
	packagePaths    map[string]string
	packageMap      map[string]*Package
}

func (x *RegistryCache) ensureLocalCache() {
	var err error;
	_, err = os.Stat(x.basePath);
	if os.IsNotExist(err) {
		err = os.Mkdir(x.basePath, 0755);
		if err != nil { fmt.Printf("%s\n", err); }
	}
}

func (x Repository) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}

func (x *RegistryCache) ContainsRepository(reg *Repository) bool {
	for _, r := range x.repositories {
		if r.Url == reg.Url {
			return true;
		}
	}
	return false;
}

func (x *RegistryCache) PackageExists(packages map[string]*Package, p *Package) bool {
	for _, check := range packages {
		if check.Name == p.Name {
			return true;
		}
	}
	return false;
}

func (x *RegistryCache) loadRepositories() {
	x.repositories = map[string]*Repository{};
	x.packagePaths = map[string]string{};
	x.packageMap = map[string]*Package{};
	var _, err = os.Stat(x.basePath);
	if err != nil { fmt.Printf("Local cache not found.  %v\n", err); }
	containers, err := ioutil.ReadDir(x.basePath);
	if err != nil { fmt.Printf("Failed to read local cache.  %v\n", err); }
	for _, container := range containers {
		_, err = os.Stat(fmt.Sprintf("%s/%s/repo.json", x.basePath, container.Name()));
		if !os.IsNotExist(err) {
			var tempRepo *Repository;
			rawText, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/repo.json", x.basePath, container.Name()));
			if err == nil {
				json.Unmarshal(rawText, &tempRepo);
				if tempRepo != nil && !x.ContainsRepository(tempRepo) {
					tempRepo.SetBasePath(fmt.Sprintf("%s/%s", x.basePath, container.Name()));
					//tempRepo.loadPackages(fmt.Sprintf("%s/%s", x.basePath, container.Name()));
					for _, p := range tempRepo.Packages {
						p.SetRepository(tempRepo);
						x.packagePaths[p.Name] = fmt.Sprintf("%s/%s/%s", x.basePath, container.Name(), p.Name);
						x.packageMap[p.Name] = p;
					}
					x.repositories[tempRepo.Name] = tempRepo;
				}
			}
		}
	}
}

func (x *RegistryCache) SetRegistryUrl(url string) {
	if url != "" {
		x.url = url;
	}
}

func (x *RegistryCache) comparePackageVersions(p1 *Package, p2 *Package) bool {
	if p1 == nil || p2 == nil {
		fmt.Printf("Packages cannot be nil...\n");
		return false;
	}
	var comps1 = strings.Split(p1.Version, ".");
	var comps2 = strings.Split(p2.Version, ".");
	if len(comps1) != len(comps2) {
		fmt.Printf("Version format mismatch...\n");
		return false;
	}
	for i := 0; i < len(comps1); i++ {
		temp1, err := strconv.Atoi(comps1[i]);
		if err != nil { return false; }
		temp2, err := strconv.Atoi(comps2[i]);
		if err != nil { return false; }
		if temp2 > temp1 { return true; }
	}
	return false;
}

func (x *RegistryCache) downloadPackage(src string, tar string) {
	_, err := os.Stat(fmt.Sprintf("%s.lock", tar));
	if !os.IsNotExist(err) {
		fmt.Printf("Package already being modified...\n");
		return;
	}
	ioutil.WriteFile(fmt.Sprintf("%s.lock", tar), []byte(""), 0755);
	var client = http.Client{};
    req, err := http.NewRequest("GET", src, nil);
    rsp, err := client.Do(req);
    if err == nil {
        rspData, _ := ioutil.ReadAll(rsp.Body);
		ioutil.WriteFile(tar, rspData, 0755);
    }
	os.Remove(fmt.Sprintf("%s.lock", tar));
}

// https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
func (x *RegistryCache) uploadPackage(src string) {
	info, err := os.Stat(src);
	if os.IsNotExist(err) {
		fmt.Printf("Unable to find package.  Please provide the full path to a valid package...\n");
		return;
	}
	fh, err := os.Open(src);
	if err != nil {
		fmt.Printf("Failed to open: %s\n", src);
		return;
	}
	data, err := ioutil.ReadAll(fh);
	if err != nil {
		fmt.Printf("Failed to read: %s\n", src);
		return;
	}
	fh.Close();
	var body = &bytes.Buffer{};
	var w = multipart.NewWriter(body);
	part, err := w.CreateFormFile("package", info.Name());
	if err != nil {
		fmt.Printf("Failed to create form data...\n");
		return;
	}
	part.Write(data);
	w.Close();
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/upload/", x.url), body);
	req.Header.Add("Content-Type", w.FormDataContentType())
	if err != nil {
		fmt.Printf("Failed to create request... %s\n", err);
		return;
	}
	client := &http.Client{};
	_, err = client.Do(req);
	if err != nil {
		fmt.Sprintf("Failed to upload package (%s): %s\n", src, err);
		return;
	}
}

func (x *RegistryCache) Invoke(operation string, packageName string, repoName string, repoUrl string) {
	if (operation == "upload" || operation == "purge" || operation == "install") && packageName == "" {
		fmt.Printf("Invalid package (%s) for operation (%s)\n", packageName, operation);
		os.Exit(3);
	}
	switch (operation) {
		case "get":
			for _, p := range x.Packages() {
				if packageName == "" || p.Name == packageName {
					fmt.Printf("%s: %s\n", p.Name, p.Version);
				}
			}
			break;
		case "upload":
			_, err := os.Stat(packageName);
			if os.IsNotExist(err) {
				fmt.Printf("Unable to find package.  Please provide the full path to a valid package...\n");
				return;
			}
			if packageName[len(packageName) - 5:] != ".agdo" {
				var comps = strings.Split(packageName, "/");
				var pName = comps[len(comps) - 1];
				// Create the package file
				var cmd = fmt.Sprintf("tar -cf %s.agdo -C %s/../ ./%s", packageName, packageName, pName);
				common.RunCmd(cmd);
				_, err := os.Stat(fmt.Sprintf("%s.agdo", packageName));
				if os.IsNotExist(err) {
					fmt.Printf("Failed to create package file.  Please report this bug to the vendor...\n");
					return;
				}
				packageName += ".agdo";
			}
			x.uploadPackage(packageName);
			break;
		case "update":
//debug
			for _, repo := range x.repositories {
				repo.Packages = map[string]*Package{};
				//go func(repo *Repository) {
					repo.loadPackages(repo.BasePath());
					var packages = (&common.RestHelper{}).InvokeGet(fmt.Sprintf("%s/api/update/", repo.Url), map[string]string{});
					if packages != nil {
						var packages2 = packages["packages"];
						for _, p := range packages2.([]interface{}) {
							var rawPackage = p.(map[string]interface{});
							var tempPackage *Package;
							jsonString, _ := json.Marshal(rawPackage);
							json.Unmarshal([]byte(jsonString), &tempPackage);
							if tempPackage != nil {
								repo.Packages[tempPackage.Name] = tempPackage;
							}
						}
						jsonString, _ := json.Marshal(repo);
						ioutil.WriteFile(fmt.Sprintf("%s/repo.json", repo.BasePath()), []byte(jsonString), 0755);
					}
				//}(repo);
			}
			break;
		case "upgrade":
			for _, repo := range x.repositories {
				go func(repo *Repository) {
					var packages = (&common.RestHelper{}).InvokeGet(fmt.Sprintf("%s/api/upgrade/", repo.Url), map[string]string{})["packages"].([]*Package);
					for _, p := range packages {
						var pName = strings.Replace(packageName, ".agdo", "", -1);
						if (packageName == "" || pName  == p.Name) &&  x.comparePackageVersions(repo.Packages[p.Name], p) {
							// Download latest
							x.downloadPackage(fmt.Sprintf("%s/api/download/%s", p.Repository().Url, pName), fmt.Sprintf("%s/%s.agdo", p.Repository().BasePath(), pName));
						}
					}
				}(repo);
			}
			break;
		case "install":
			var pName = strings.Replace(packageName, ".agdo", "", -1);
			var p = x.Packages()[pName];
			if p == nil {
				fmt.Printf("Failed to lookup package (%s)...\n", pName);
				return;
			}
			x.downloadPackage(fmt.Sprintf("%s/api/download/%s", p.Repository().Url, pName), fmt.Sprintf("%s/%s.agdo", p.Repository().BasePath(), pName));
			break;
		case "purge":
			os.RemoveAll(x.packagePaths[packageName]);
			os.Remove(fmt.Sprintf("%s.agdo", x.packagePaths[packageName]));
			break;
		default:
			fmt.Printf("Invalid operation (%s)\n", operation);
			os.Exit(1);
	}
}

func (x RegistryCache) Packages() map[string]*Package {
	var packages = map[string]*Package{};
	for _, repo := range x.repositories {
		for _, p := range repo.Packages {
			if !x.PackageExists(packages, p) {
				packages[p.Name] = p;
			}
		}
	}
	return packages;
}

func (x RegistryCache) Repositories() map[string]*Repository { return x.repositories; }
func (x RegistryCache) BasePath() string { return x.basePath; }
func (x *RegistryCache) SetBasePath(p string) { x.basePath = p; }
/*****************************************************************************/

// RegistryCache Instance
var rc *RegistryCache;
func GetRegistryCacheInstance() *RegistryCache {
	if rc == nil {
		rc = &RegistryCache{};
		rc.cs = GetConfigurationServiceInstance();
		rc.url = rc.cs.GetProperty("registryUrl", "http://localhost:4455").(string);
		var currentUser, _  = user.Current();
		var homeDirectory = currentUser.HomeDir;
		rc.SetBasePath(rc.cs.GetProperty("basePath", fmt.Sprintf("%s/agdo", homeDirectory)).(string));
		rc.ensureLocalCache();
		rc.loadRepositories();
	}
	return rc;
}
/*****************************************************************************/

