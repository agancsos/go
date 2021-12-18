package models

// ServiceState
type ServiceState int;
const (
	SS_NONE ServiceState = iota
    SS_UNKNOWN
    SS_INITIALIZED
);
/*****************************************************************************/

// NodeMemberState
type NodeMemberState int;
const (
	NMS_NONE NodeMemberState = iota
    NMS_INITIALIZING
    NMS_INITIALIZED
    NMS_STOPPING
	NMS_STOPPED
);
/*****************************************************************************/

// UserState
type UserState int;
const (
	US_NONE UserState = iota
    US_ACTIVE
    US_BLOCKED
);
/*****************************************************************************/

// PostUrlType
type PostUrlType int;
const (
	PUT_NONE PostUrlType = iota
    PUT_SITE
    PUT_IMAGE
    PUT_VIDEO
);
/*****************************************************************************/

// RuleType
type RuleType int;
const (
	RT_NONE RuleType = iota
	RT_KEYWORD
)
/*****************************************************************************/


// EnumHelper
type EnumHelper struct {}
func (x ServiceState) Name() string {
	return "";
}
/*****************************************************************************/

