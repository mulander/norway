package norway

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
	return Entry{}
}
