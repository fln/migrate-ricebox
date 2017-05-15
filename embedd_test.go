package ricebox

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "1_foobar.down.sql",
		FileModTime: time.Unix(1494531677, 0),
		Content:     string("1 down\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "1_foobar.up.sql",
		FileModTime: time.Unix(1494531666, 0),
		Content:     string("1 up\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "3_foobar.up.sql",
		FileModTime: time.Unix(1494531685, 0),
		Content:     string("3 up\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "4_foobar.down.sql",
		FileModTime: time.Unix(1494531709, 0),
		Content:     string("4 down\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "4_foobar.up.sql",
		FileModTime: time.Unix(1494531699, 0),
		Content:     string("4 up\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "5_foobar.down.sql",
		FileModTime: time.Unix(1494531718, 0),
		Content:     string("5 down\n"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "7_foobar.down.sql",
		FileModTime: time.Unix(1494531736, 0),
		Content:     string("7 down\n"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    "7_foobar.up.sql",
		FileModTime: time.Unix(1494531726, 0),
		Content:     string("7 up\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1494531736, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "1_foobar.down.sql"
			file3, // "1_foobar.up.sql"
			file4, // "3_foobar.up.sql"
			file5, // "4_foobar.down.sql"
			file6, // "4_foobar.up.sql"
			file7, // "5_foobar.down.sql"
			file8, // "7_foobar.down.sql"
			file9, // "7_foobar.up.sql"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`testbox-embedd`, &embedded.EmbeddedBox{
		Name: `testbox-embedd`,
		Time: time.Unix(1494531736, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"1_foobar.down.sql": file2,
			"1_foobar.up.sql":   file3,
			"3_foobar.up.sql":   file4,
			"4_foobar.down.sql": file5,
			"4_foobar.up.sql":   file6,
			"5_foobar.down.sql": file7,
			"7_foobar.down.sql": file8,
			"7_foobar.up.sql":   file9,
		},
	})
}
