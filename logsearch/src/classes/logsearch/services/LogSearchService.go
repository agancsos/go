package services
import (
	"fmt"
	"../models"
	"../parsers"
	"errors"
	"strings"
)

type LogSearchService struct {
	ds                      *DataService
}

var __log_search_instance__ *LogSearchService;
func GetLogSearchServiceInstance() *LogSearchService {
	if __log_search_instance__ == nil {
		__log_search_instance__ = &LogSearchService{};
		__log_search_instance__.ds = GetDataServiceInstance();
	}
	return __log_search_instance__;
}

func (x LogSearchService) AddEvent(event map[string]string) error {
	if ! x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO LOGS (LOG_SOURCE, LOG_TEXT) VALUES ('%s', '%s')", event["source"], event["content"])) {
		return errors.New("Failed to ingest event");
	}
	return nil;
}

func (x LogSearchService) PurgeLogs() error {
	x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM LOGS"));
	return nil;
}

func (x LogSearchService) QueryLogs(filter map[string]string) (error, []*models.LogEvent) {
	var rst = []*models.LogEvent{};
	var _, query = (&parsers.FilterParser{}).Unparse(filter);
	var rows = x.ds.ServiceQuery(query).Rows();
	for _, row := range rows {
		var event = &models.LogEvent{};
		event.SetID(row.Column("LOG_ID").Value());
		event.SetTimestamp(row.Column("LAST_UPDATED_DATE").Value());
		event.SetSource(row.Column("LOG_SOURCE").Value());
		event.SetData(row.Column("LOG_TEXT").Value());
		for _, char := range []string{",", ";", " ",} {
			if len(strings.Split(event.Data(), char)) > 0 {
				for i, pair := range strings.Split(event.Data(), char) {
					event.AddCustomField(fmt.Sprintf("field%d", i), pair);
				}
			}
		}
		rst = append(rst, event);
	}
	return nil, rst;
}

