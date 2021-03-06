package norway

import (
	"strings"
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

func TestParseEntryNoSubDirectoriesAdditionalNL(t *testing.T) {
	Convey("Given a single directory line that indicates no sub-directories", t, func() {
		entryLine := "D\n\n"
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
			Convey("Revision should match 10.24", func() {
				So(entry.Revision, ShouldEqual, "10.24")
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

func TestParseEntries(t *testing.T) {
	Convey("Given an Entries file", t, func() {
		entriesFile := "/bio_ssl.c/1.14/Mon Apr 21 16:34:43 2014//\n" +
			"/d1_both.c/1.12/Thu Apr 24 15:50:02 2014//\n" +
			"/d1_clnt.c/1.16/Wed Apr 23 22:26:26 2014//\n" +
			"/d1_enc.c/1.3/Mon Apr 14 14:16:33 2014//\n" +
			"/d1_lib.c/1.12/Sun Apr 20 14:14:52 2014//\n" +
			"/d1_meth.c/1.3/Sat Apr 19 08:52:32 2014//\n" +
			"/d1_pkt.c/1.16/Wed Apr 23 18:40:39 2014//\n" +
			"/d1_srtp.c/1.3/Sat Apr 19 08:52:32 2014//\n" +
			"/d1_srvr.c/1.18/Wed Apr 23 05:13:57 2014//\n" +
			"D\n\n"
		Convey("When the file is parsed", func() {
			entries, err := ParseEntries(strings.NewReader(entriesFile))
			Convey("That the file was parsed without errors", func() {
				So(err, ShouldBeNil)
			})
			Convey("bio_ssl.c is at revision 1.14", func() {
				So(entries["bio_ssl.c"].Revision, ShouldEqual, "1.14")
			})
			Convey("d1_enc.c is at revision 1.3", func() {
				So(entries["d1_enc.c"].Revision, ShouldEqual, "1.3")
			})
			Convey("d1_meth.c timestamp is Sat Apr 19 08:52:32 2014", func() {
				So(entries["d1_meth.c"].Timestamp, ShouldEqual, "Sat Apr 19 08:52:32 2014")
			})
		})
	})
}

func TestParseEntriesCRLF(t *testing.T) {
	Convey("Given an Entries file", t, func() {
		entriesFile := "/bio_ssl.c/1.14/Mon Apr 21 16:34:43 2014//\r\n" +
			"/d1_both.c/1.12/Thu Apr 24 15:50:02 2014//\r\n" +
			"/d1_clnt.c/1.16/Wed Apr 23 22:26:26 2014//\r\n" +
			"/d1_enc.c/1.3/Mon Apr 14 14:16:33 2014//\r\n" +
			"/d1_lib.c/1.12/Sun Apr 20 14:14:52 2014//\r\n" +
			"/d1_meth.c/1.3/Sat Apr 19 08:52:32 2014//\r\n" +
			"/d1_pkt.c/1.16/Wed Apr 23 18:40:39 2014//\r\n" +
			"/d1_srtp.c/1.3/Sat Apr 19 08:52:32 2014//\r\n" +
			"/d1_srvr.c/1.18/Wed Apr 23 05:13:57 2014//\r\n" +
			"D\r\n\r\n"
		Convey("When the file is parsed", func() {
			entries, err := ParseEntries(strings.NewReader(entriesFile))
			Convey("That the file was parsed without errors", func() {
				So(err, ShouldBeNil)
			})
			Convey("bio_ssl.c is at revision 1.14", func() {
				So(entries["bio_ssl.c"].Revision, ShouldEqual, "1.14")
			})
			Convey("d1_enc.c is at revision 1.3", func() {
				So(entries["d1_enc.c"].Revision, ShouldEqual, "1.3")
			})
			Convey("d1_meth.c timestamp is Sat Apr 19 08:52:32 2014", func() {
				So(entries["d1_meth.c"].Timestamp, ShouldEqual, "Sat Apr 19 08:52:32 2014")
			})
		})
	})
}

func TestEntriesSortedByTimestamp(t *testing.T) {
	Convey("Given an Entries file", t, func() {
		entriesFile := "/bio_ssl.c/1.14/Mon Apr 21 16:34:43 2014//\n" +
			"/d1_both.c/1.12/Thu Apr 24 15:50:02 2014//\n" +
			"/d1_clnt.c/1.16/Wed Apr 23 22:26:26 2014//\n" +
			"/d1_enc.c/1.3/Mon Apr 14 14:16:33 2014//\n" +
			"/d1_lib.c/1.12/Sun Apr 20 14:14:52 2014//\n" +
			"/d1_meth.c/1.3/Sat Apr 19 08:52:32 2014//\n" +
			"/d1_pkt.c/1.16/Wed Apr 23 18:40:39 2014//\n" +
			"/d1_srtp.c/1.3/Sat Apr 19 08:52:32 2014//\n" +
			"/d1_srvr.c/1.18/Wed Apr 23 05:13:57 2014//\n" +
			"D\n\n"
		Convey("When the file is parsed", func() {
			entries, err := ParseEntries(strings.NewReader(entriesFile))
			Convey("That the file was parsed without errors", func() {
				So(err, ShouldBeNil)
			})
			Convey("That only 10 entries were parsed", func() {
				So(len(entries), ShouldEqual, 10)
			})
			Convey("Then after sorting it by timestamp", func() {
				sorted := entries.SortedByTimestamp()
				Convey("That the sorted result still has 10 entries", func() {
					So(len(sorted), ShouldEqual, 10)
				})
				Convey("The first file is no longer bio_ssl.c", func() {
					So(sorted[0].FileName, ShouldNotEqual, "bio_ssl.c")
					So(sorted[0].FileName, ShouldEqual, "")
					So(sorted[0].IsDirectory, ShouldBeTrue)
				})
				Convey("d1_both.c is the first file after the directory", func() {
					So(sorted[1].FileName, ShouldEqual, "d1_both.c")
					So(sorted[1].Timestamp, ShouldEqual, "Thu Apr 24 15:50:02 2014")
				})
				Convey("d1_enc.c is the last file", func() {
					So(sorted[len(sorted)-1].FileName, ShouldEqual, "d1_enc.c")
				})
			})
		})
	})
}
