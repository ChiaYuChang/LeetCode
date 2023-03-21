package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFilePathParser(t *testing.T) {
	ftp := leetcode.Q0388{}.NewFilePathTextParser()

	type testCase struct {
		name     string
		filepath []string
		text     string
		ans      int
	}

	tcs := []testCase{
		{
			name: "Case 1",
			text: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			filepath: []string{
				"dir/subdir2/file.ext",
			},
			ans: 20,
		},
		{
			name: "Case 2",
			text: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			filepath: []string{
				"dir/subdir1/file1.ext",
				"dir/subdir2/subsubdir2/file2.ext",
			},
			ans: 32,
		},
		{
			name:     "Case 3 no file exist",
			text:     "dir",
			filepath: []string{},
			ans:      0,
		},
		{
			name: "Case 4 all files in root",
			text: "file1.txt\nfile2.txt\nlongfile.txt",
			filepath: []string{
				"file1.txt",
				"file2.txt",
				"longfile.txt",
			},
			ans: 12,
		},
		{
			name: "Case 5 space in filename",
			text: "dir\n        file.txt",
			filepath: []string{
				"        file.txt",
			},
			ans: 16,
		},
		{
			name: "Case 5 space in filename",
			text: "dir1\n\t   subdir\n\t\tfile1.txt\ndir2\n\tfile2.txt",
			filepath: []string{
				"dir1/   subdir/file1.txt",
				"dir2/file2.txt",
			},
			ans: 24,
		},
		{
			name: "Case 6 unsort dir and file at root",
			text: "file1.tar\nfile2.jpeg\ndir1\nfile3.txt\ndir2\n\tdir2sub\n\t\tfile4.csv",
			filepath: []string{
				"file1.tar",
				"file2.jpeg",
				"file3.txt",
				"dir2/dir2sub/file4.csv",
			},
			ans: 22,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				tkns := ftp.Parse(tc.text)
				require.NotNil(t, tkns)
				tree := ftp.Lext(tkns)
				require.NotNil(t, tree)

				ftp.FileNodeTreeView(tree)

				require.ElementsMatch(t, tc.filepath, ftp.Files(tree))
				require.Equal(t, tc.ans, ftp.LengthLongestFilePath(tree))
			},
		)
	}
}
