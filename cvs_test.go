package norway

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseEntryNoSubDirectories(t *testing.T) {
	Convey("Given a single directory line that indicates no sub-directories", t, func() {
		entryLine := "D"
		Convey("When the entry is parsed", func() {
			entry := ParseEntry(entryLine)
			Convey("IsDirectory should be true", func() {
				So(entry.IsDirectory, ShouldBeTrue)
			})
		})
	})
}

func TestParseEntryFile(t *testing.T) {
	Convey("Given a regular file entry", t, func() {
		entryLine := "/bio_ssl.c/1.14/Mon Apr 21 16:34:43 2014//"
		Convey("When the entry is parsed", func() {
			entry := ParseEntry(entryLine)
			Convey("IsDirectory should be false", func() {
				So(entry.IsDirectory, ShouldBeFalse)
			})
			Convey("IsRemoved should be false", func() {
				So(entry.IsRemoved, ShouldBeFalse)
			})
			Convey("IsAdded should be false", func() {
				So(entry.IsAdded, ShouldBeFalse)
			})
			Convey("FileName should match bio_ssl.c", func() {
				So(entry.FileName, ShouldEqual, "bio_ssl.c")
			})
			Convey("Revision should match 1.14", func() {
				So(entry.Revision, ShouldEqual, "1.14")
			})
			Convey("Timestamp should match Mon Apr 21 16:34:43 2014", func() {
				So(entry.Timestamp, ShouldEqual, "Mon Apr 21 16:34:43 2014")
			})
			Convey("Options should be an empty string", func() {
				So(entry.Options, ShouldEqual, "")
			})
			Convey("Tagdate should be an empty string", func() {
				So(entry.Tagdate, ShouldEqual, "")
			})
		})
	})
}

func TestParseEntryFileWithOptions(t *testing.T) {
	Convey("Given a regular file entry with options", t, func() {
		entryLine := "/some_file.c/10.24/Fri Apr 11 08:17:05 2014/-kkv/"
		Convey("When the entry is parsed", func() {
			entry := ParseEntry(entryLine)
			Convey("IsDirectory should be false", func() {
				So(entry.IsDirectory, ShouldBeFalse)
			})
			Convey("IsRemoved should be false", func() {
				So(entry.IsRemoved, ShouldBeFalse)
			})
			Convey("IsAdded should be false", func() {
				So(entry.IsAdded, ShouldBeFalse)
			})
			Convey("FileName should match some_file.c", func() {
				So(entry.FileName, ShouldEqual, "some_file.c")
			})
			Convey("Revision should match 1.24", func() {
				So(entry.Revision, ShouldEqual, "1.24")
			})
			Convey("Timestamp should match Fri Apr 11 08:17:05 2014", func() {
				So(entry.Timestamp, ShouldEqual, "Fri Apr 11 08:17:05 2014")
			})
			Convey("Options should be equal to -kkv", func() {
				So(entry.Options, ShouldEqual, "-kkv")
			})
			Convey("Tagdate should be an empty string", func() {
				So(entry.Tagdate, ShouldEqual, "")
			})
		})
	})
}

func TestParseEntryFileWithTag(t *testing.T) {
	Convey("Given a regular file entry with options", t, func() {
		entryLine := "/some_file.bin/10.1.4.1.22.1/Tue Apr  8 14:00:44 2014/-kb/TB_01_05X_000Y_snapshots"
		Convey("When the entry is parsed", func() {
			entry := ParseEntry(entryLine)
			Convey("IsDirectory should be false", func() {
				So(entry.IsDirectory, ShouldBeFalse)
			})
			Convey("IsRemoved should be false", func() {
				So(entry.IsRemoved, ShouldBeFalse)
			})
			Convey("IsAdded should be false", func() {
				So(entry.IsAdded, ShouldBeFalse)
			})
			Convey("FileName should match some_file.bin", func() {
				So(entry.FileName, ShouldEqual, "some_file.bin")
			})
			Convey("Revision should match 10.1.4.1.22.1", func() {
				So(entry.Revision, ShouldEqual, "10.1.4.1.22.1")
			})
			Convey("Timestamp should match Tue Apr  8 14:00:44 2014", func() {
				So(entry.Timestamp, ShouldEqual, "Tue Apr  8 14:00:44 2014")
			})
			Convey("Options should be equal to -kb", func() {
				So(entry.Options, ShouldEqual, "-kb")
			})
			Convey("Tagdate should be equal to ", func() {
				So(entry.Tagdate, ShouldEqual, "TB_01_05X_000Y_snapshots")
			})
		})
	})
}

func TestParseEntryDirectory(t *testing.T) {
	Convey("Given a regular file entry with options", t, func() {
		entryLine := "D/some_directory////"
		Convey("When the entry is parsed", func() {
			entry := ParseEntry(entryLine)
			Convey("IsDirectory should be true", func() {
				So(entry.IsDirectory, ShouldBeTrue)
			})
			Convey("IsRemoved should be false", func() {
				So(entry.IsRemoved, ShouldBeFalse)
			})
			Convey("IsAdded should be false", func() {
				So(entry.IsAdded, ShouldBeFalse)
			})
			Convey("FileName should match some_directory", func() {
				So(entry.FileName, ShouldEqual, "some_directory")
			})
			Convey("Revision should be an empty string", func() {
				So(entry.Revision, ShouldEqual, "")
			})
			Convey("Timestamp should be an empty string", func() {
				So(entry.Timestamp, ShouldEqual, "")
			})
			Convey("Options should be an empty string", func() {
				So(entry.Options, ShouldEqual, "")
			})
			Convey("Tagdate should be an empty string", func() {
				So(entry.Tagdate, ShouldEqual, "")
			})
		})
	})
}
