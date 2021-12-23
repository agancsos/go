package common
import (
	"strings"
)

type SystemService struct {
	ModulePath     string
}
func (x *SystemService) BuildModuleContainerPath() string {
	var result  = strings.Replace(x.ModulePath, "\\", "/", -1);
	return result;
}
