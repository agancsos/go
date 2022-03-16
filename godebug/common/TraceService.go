package common
import (
	"os"
	"fmt"
	"time"
)

// TraceLevel
type TraceLevel int;
const (
	TL_NONE TraceLevel = iota
	TL_ERROR
	TL_WARNING
	TL_INFORMATIONAL
	TL_VERBOSE
	TL_DEBUG
)

func GetTraceLevelName(x TraceLevel) string {
	switch (x) {
		case TL_ERROR:
			return "ERR ";
		case TL_WARNING:
			return "WARN";
		case TL_INFORMATIONAL:
			return "INFO";
		case TL_DEBUG:
			return "DEGB";
		case TL_VERBOSE:
			return "VERB";
		default:
			return "";
	}
}
/*****************************************************************************/

// TraceCategory
type TraceCategory int;
const (
	TC_NONE TraceCategory = iota
)
/*****************************************************************************/

// TraceService
type TraceService struct {
	FilePath    string
	TraceLevel  int
}
func (x *TraceService) traceMessage(message string, level int, category int) {
    if x.TraceLevel >= level {
		var file, _ = os.OpenFile(x.FilePath, os.O_APPEND|os.O_WRONLY, 0644);
		file.WriteString(fmt.Sprintf("\n%v | %s | %s",
			time.Now(),
			GetTraceLevelName((TraceLevel)(level)),
			message));
		file.Close();
    }
}

func (x *TraceService) TraceError(message string, category int) {
    x.traceMessage(message, int(TL_ERROR), category);
}

func (x *TraceService) TraceWarning(message string, category int) {
    x.traceMessage(message, int(TL_WARNING), category);
}

func (x *TraceService) TraceInformational(message string, category int) {
    x.traceMessage(message, int(TL_INFORMATIONAL), category);
}

func (x *TraceService) TraceVerbose(message string, category int) {
	x.traceMessage(message, int(TL_VERBOSE), category);
}

func (x *TraceService) TraceDebug(message string, category int) {
    x.traceMessage(message, int(TL_DEBUG), category);
}
/*****************************************************************************/
