package engine

import "fmt"

const (
	SnakeCase = NameConv(iota)
	CamelCase
)

var (
	ErrUnkNameConv = fmt.Errorf("unknown name convension")

	logTypes = map[string]*LogType{
		"kunai":    &TypeKunai,
		"winevt":   &TypeWinevt,
		"maltrace": &TypeMaltrace,
	}
)

// NameConv defines a custom type for identifying
// naming convention
type NameConv int

type LogType struct {
	FieldNameConv NameConv
	Data          *XPath
	Source        *XPath
	EventID       *XPath
	Hostname      *XPath
	GeneInfo      *XPath
	Timestamp     *XPath
}

// Windows Event Format
var (
	systemPath = Path("/Event/System")

	TypeWinevt = LogType{
		FieldNameConv: CamelCase,
		Data:          eventDataPath,
		Source:        systemPath.Append("Channel"),
		EventID:       systemPath.Append("EventID"),
		Hostname:      systemPath.Append("Computer"),
		GeneInfo:      Path("/Event/GeneInfo"),
		Timestamp:     systemPath.Append("TimeCreated").Append("SystemTime"),
	}
)

// Kunai's log Format
var (
	TypeKunai = LogType{
		FieldNameConv: SnakeCase,
		Data:          Path("/data"),
		Source:        Path("/info/event/source"),
		EventID:       Path("/info/event/id"),
		Hostname:      Path("/info/host/hostname"),
		GeneInfo:      Path("/gene_info"),
		Timestamp:     Path("/info/utc_time"),
	}
)

// Maltrace's Event Format
var (
	TypeMaltrace = LogType{
		FieldNameConv: SnakeCase,
		Data:          Path("/process"),
		Source:        Path("/event_type"),
		EventID:       Path("/event_id"),
		Hostname:      Path("/event_id"),
		Timestamp:     Path("/timestamp"),
	}
	// {
	// 	"event_id": 7019535073598018000,
	// 	"event_type": "syscall",
	// 	"process": {
	// 	  "filename": "/etc/localtime",
	// 	  "flags": [
	// 		"O_RDONLY"
	// 	  ],
	// 	  "mode": 0,
	// 	  "parent_pid": 14329,
	// 	  "pid": 17367
	// 	},
	// 	"syscall_name": "openat",
	// 	"timestamp": "0001-01-01T00:00:00Z"
	//   },
	// {
	// 	"event_id": 7607674595460999000,
	// 	"event_type": "syscall",
	// 	"process": {
	// 	  "cmdline": "maltrace /usr/bin/ls",
	// 	  "filename": "/usr/bin/ls",
	// 	  "hash": "7987cf330ff5bb94015dfbb9eae5a99f",
	// 	  "parent_pid": 17373,
	// 	  "pid": 17377
	// 	},
	// 	"syscall_name": "execve",
	// 	"timestamp": "2025-05-17T12:12:36.67788862-04:00"
	//   },

)
