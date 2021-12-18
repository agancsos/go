package services
import (
	"../../incryptEncode"
)

type SecurityService struct {
	encoder    incryptEncode.IncryptEncode
}

var __security_service__ *SecurityService;

func GetSecurityServiceInstance() *SecurityService {
	if __security_service__ == nil {
		__security_service__ = &SecurityService{};
		__security_service__.encoder = &incryptEncode.IncryptEncodeFullBinary{};
	}
	return __security_service__;
}

func (x *SecurityService) GetEncoded(a string) string {
	return x.encoder.GetEncoded(a);
}

func (x *SecurityService) GetDecoded(a string) string {
	return x.encoder.GetDecoded(a);
}

