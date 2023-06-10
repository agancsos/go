package runtime
import (
	"fmt"
	"plugin"
	"errors"
	"os"
	"strings"
	"io/ioutil"
)

type IClassLoader interface {
	Plugin(name string) (*plugin.Plugin, error)
	loadPlugins()
}

type ClassLoader struct {
	basePath            string
	plugs               map[string]*plugin.Plugin
}

func (x *ClassLoader) loadPlugins() {
	x.plugs    = map[string]*plugin.Plugin{};
	var _, err = os.Stat(x.basePath);
	if err != nil {
		return;
	}
	var libs, _ = ioutil.ReadDir(x.basePath);
	for _, lib := range libs {
		if !strings.Contains(lib.Name(), ".so") {
			continue;
		}
		var plug, err = plugin.Open(fmt.Sprintf("%s/%s", x.basePath, lib.Name()));
		if err == nil {
			x.plugs[strings.Split(lib.Name(), ".")[0]] = plug;
		}
	}
}

func (x ClassLoader) Plugin(name string) (*plugin.Plugin, error) {
	if x.plugs[name] == nil {
		return nil, errors.New(fmt.Sprintf("No such plugin (%s)", name));
	}
	return x.plugs[name], nil;
}

func (x ClassLoader) Plugins() map[string]*plugin.Plugin {
    return x.plugs;
}

func (x ClassLoader) Lookup(sym string) (plugin.Symbol, error) {
	var err error;
	for _, plug := range x.plugs {	
		var rst, err = plug.Lookup(sym);
		if err == nil {
			return rst, nil;
		}
	}
	return nil, err;
}

func NewClassLoader(path string) *ClassLoader {
	var rst      = &ClassLoader{};
	rst.basePath = path;
	rst.loadPlugins();
	return rst;
}

