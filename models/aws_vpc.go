package models

type VpcNormalizedData struct {
	Cloud           string //custom
	Type            string //custom
	Version         int
	AccountID       string
	InterfaceID     string
	SourceIP        string
	DestinationIP   string
	DestinationPort int
	SourcePort      int
	Protocol        int
	Packets         int
	Bytes           int
	StartTime       int
	EndTime         int
	Action          string
	LogStatus       string
}
