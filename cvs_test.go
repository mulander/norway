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
