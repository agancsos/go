package services
import (
	"../../common"
    "fmt"
	"strconv"
    "../models"
)

type DbTraceService struct {
    ds          *DataService
	traceLevel  int
}

var __db_trace_service__ *DbTraceService;

func (x *DbTraceService) Initialize() {
}

func GetDbTraceServiceInstance() IService {
    if __db_trace_service__ == nil {
        __db_trace_service__ = &DbTraceService{};
		var config = (GetConfigurationServiceInstance()).GetProperty("traceLevel", "3").(string);
		__db_trace_service__.ds = GetDataServiceInstance().(*DataService);
		if config != "" {
			__db_trace_service__.traceLevel, _ = strconv.Atoi(config);
		} else {
			__db_trace_service__.traceLevel = 3;
		}
    }
    return __db_trace_service__;
}

func (x *DbTraceService) traceMessage(message string, level int, category int) {
	if x.traceLevel >= level {
        x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO MESSAGES (MESSAGE_TEXT, MESSAGE_LEVEL, MESSAGE_CATEGORY) values ('%s', '%d', '%d')", message, level, category));
    }
}

func (x *DbTraceService) TraceError(message string, category int) {
    x.traceMessage(message, int(common.TL_ERROR), category);
}

func (x *DbTraceService) TraceWarning(message string, category int) {
    x.traceMessage(message, int(common.TL_WARNING), category);
}

func (x *DbTraceService) TraceInformational(message string, category int) {
    x.traceMessage(message, int(common.TL_INFORMATIONAL), category);
}

func (x *DbTraceService) TraceDebug(message string, category int) {
    x.traceMessage(message, int(common.TL_DEBUG), category);
}

func (x *DbTraceService) TraceVerbose(message string, category int) {
    x.traceMessage(message, int(common.TL_VERBOSE), category);
}

func (x *DbTraceService) GetMessages() []models.Message {
    var result []models.Message;
    var results = x.ds.ServiceQuery("SELECT * FROM MESSAGES ORDER BY LAST_UPDATED_DATE DESC");
    for _, row := range results.Rows() {
		var id, _ = strconv.Atoi(row.Column("MESSAGE_ID").Value());
        var message = models.Message{};
		message.SetID(id);
        message.SetText(row.Column("MESSAGE_TEXT").Value());
        message.SetLastUpdatedDate(row.Column("LAST_UPDATED_DATE").Value());
        var level, _ = strconv.Atoi(row.Column("MESSAGE_LEVEL").Value());
		message.SetLevel(level);
        var category, _ = strconv.Atoi(row.Column("MESSAGE_CATEGORY").Value());
		message.SetCategory(category);
        result = append(result, message);
    }
    return result;
}

func (x *DbTraceService) PurgeAudits(days int) {
    if days < 1 {
        return;
	}
    x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM MESSAGES WHERE LAST_UPDATED_DATE < CAST((CURRENT_TIMESTAMP - %d) AS DATETIME)", days));
}

