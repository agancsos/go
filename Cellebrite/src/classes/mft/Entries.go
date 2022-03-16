package mft

type MFTItem interface {
	TypeName()	  string
	Content()	  string
}

type MFTDevice struct {
	DevicePath	          string
	Platform	          string
	Version		          string
	Descriptor	          string
	TrackSectors          int
	TotalSectorCount      int
	ClusterSectors        int
	FileRecordSegmentSize int
	content		          string
}
type MFTEntry struct {
	FileName	   string
	FilePath	   string
	content		   string
}
func (x MFTDevice) TypeName() string { return "Device"; }
func (x MFTDevice) Content() string { return x.content; }
func (x MFTEntry) TypeName() string { return "Entry"; }
func (x MFTEntry) Content() string { return x.content; }

