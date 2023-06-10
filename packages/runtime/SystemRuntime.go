package runtime
import (
	"strings"
	"path/filepath"
	"os"
)

func GetModulePath() (string, error) {
	var path, err = filepath.Abs(filepath.Dir(os.Args[0]));
	if err != nil {
		return "", err;
	}
	return strings.Replace(path, "\\", "/", -1), nil;
}

