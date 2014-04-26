package norway

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strings"
	"time"
)


type EntriesSorted []Entry

func (e EntriesSorted) Len() int {
	return len(e)
}

func (e EntriesSorted) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e EntriesSorted) Less(i, j int) bool {
	if e[i].Timestamp == "" {
		return true
	}
	if e[j].Timestamp == "" {
		return false
	}

	it, err := time.Parse(time.ANSIC, e[i].Timestamp)
	if err != nil {
		log.Fatal(err)
	}
	jt, err := time.Parse(time.ANSIC, e[j].Timestamp)
	if err != nil {
		log.Fatal(err)
	}
	return jt.Before(it)
}

type Entries map[string]Entry

func (e Entries) SortedByTimestamp() EntriesSorted {
	s := make(EntriesSorted, 0, len(e))
	for _, entry := range e {
		s = append(s, entry)
	}
	sort.Sort(s)
	return s
}

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
	fields := strings.Split(strings.TrimSpace(entryLine), "/")
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

func ParseEntries(r io.Reader) (Entries, error) {
	entries := Entries{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			// Skipping empty entry
			continue
		}
		entry := ParseEntry(scanner.Text())
		entries[entry.FileName] = entry
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}
