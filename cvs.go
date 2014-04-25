package norway

import (
	"strings"
)

type Entry struct {
	IsDirectory bool
	IsRemoved   bool
	IsAdded     bool
	FileName    string
	Revision    string
	Timestamp   string
	Options     string
	Tagdate     string
}

func ParseEntry(entryLine string) Entry {
	fields := strings.Split(entryLine, "/")
	newEntry := Entry{}

	if fields[0] == "D" {
		newEntry.IsDirectory = true
	}

	if len(fields) == 1 {
		return newEntry
	}

	newEntry.FileName = fields[1]
	newEntry.Revision = fields[2]
	newEntry.Timestamp = fields[3]
	newEntry.Options = fields[4]
	newEntry.Tagdate = fields[5]

	return newEntry
}
