package leetcode

import (
	"fmt"
	"regexp"
	"strings"
)

type Q0388 struct{}

type Q0388PathTokenType uint8

const (
	tRoot Q0388PathTokenType = iota
	tDir
	tFile
)

type Q0388PathToken struct {
	Depth int
	Type  Q0388PathTokenType
	Name  string
	Text  string
}

func (t Q0388PathToken) String() string {
	return t.Text
}

func (t Q0388PathToken) Info() string {
	return fmt.Sprintf("`%s (%d)`", t.Name, t.Depth)
}

func (t Q0388PathToken) Is(pt Q0388PathTokenType) bool {
	return t.Type == pt
}

type FileNode struct {
	Type         Q0388PathTokenType
	Name         string
	ParentFolder *FileNode
	ChildFolder  []*FileNode
	File         []string
}

type Q0388FilePathTextParser struct{ fileRegExp *regexp.Regexp }

func (q Q0388) NewFilePathTextParser() Q0388FilePathTextParser {
	return Q0388FilePathTextParser{fileRegExp: regexp.MustCompile(`[ a-zA-Z0-9]+\.[ a-zA-Z0-9]+`)}
}

func (ftp Q0388FilePathTextParser) Parse(filepath string) []*Q0388PathToken {
	lines := strings.Split(filepath, "\n")

	rootType := tRoot
	tokens := make([]*Q0388PathToken, 0, len(lines)+1)
	tokens = append(tokens, &Q0388PathToken{
		Depth: -1,
		Type:  rootType,
	})

	for _, l := range lines {
		pttkn := &Q0388PathToken{
			Depth: strings.Count(l, "\t"),
			Name:  strings.TrimLeft(l, "\t"),
			Text:  l,
		}
		if isfile := ftp.fileRegExp.MatchString(l); isfile {
			pttkn.Type = tFile
		} else {
			pttkn.Type = tDir
		}
		fmt.Println(pttkn.Info())
		tokens = append(tokens, pttkn)
	}
	return tokens
}

func (ftp Q0388FilePathTextParser) Lext(tokens []*Q0388PathToken) *FileNode {
	var root, curNode *FileNode
	var curDepth int

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		switch t.Type {
		case tRoot:
			root = &FileNode{Type: tRoot, Name: t.Name}
			curNode = root
			curDepth = t.Depth
		case tDir:
			// find the parenet of the folder
			for curDepth >= t.Depth {
				curNode = curNode.ParentFolder
				curDepth--
			}
			// fmt.Printf("Append folder %s (%d) to %s (%d)\n", t.Name, t.Depth, curNode.Name, curDepth)
			newNode := &FileNode{
				Type:         tDir,
				ParentFolder: curNode,
				Name:         t.Name,
			}
			curNode.ChildFolder = append(curNode.ChildFolder, newNode)
			curDepth = t.Depth
			curNode = newNode
			// fmt.Printf("-> Current node %s (%d)\n", curNode.Name, curDepth)
		case tFile:
			for curDepth >= t.Depth {
				curNode = curNode.ParentFolder
				curDepth--
			}
			// fmt.Printf("Append file %s to %s\n", t.Name, curNode.Name)
			curNode.File = append(curNode.File, t.Name)
		default:
			panic("Unknown token")
		}
	}
	return root
}

func (ftp Q0388FilePathTextParser) FileNodeTreeView(fn *FileNode) {
	root := fn
	for _, fn := range root.File {
		fmt.Println("└─", fn)
	}

	for _, cf := range root.ChildFolder {
		ftp.fileNodeTreeView(cf, 0)
	}
}

func (ftp Q0388FilePathTextParser) fileNodeTreeView(fn *FileNode, d int) {
	if fn == nil {
		return
	}

	if d == 0 {
		fmt.Println("└─", fn.Name)
	} else {
		fmt.Printf("   %s%s %s\n", strings.Repeat("     ", d-1), "└───", fn.Name)
	}

	for _, f := range fn.File {
		fmt.Printf("   %s%s %s\n", strings.Repeat("     ", d), "└───", f)
	}
	for _, cf := range fn.ChildFolder {
		ftp.fileNodeTreeView(cf, d+1)
	}
}

func (ftp Q0388FilePathTextParser) Files(fn *FileNode) []string {
	fs := []string{}
	for _, f := range fn.File {
		fs = append(fs, f)
	}
	for _, cf := range fn.ChildFolder {
		ftp.files(cf, cf.Name+"/", &fs)
	}
	return fs
}

func (ftp Q0388FilePathTextParser) files(fn *FileNode, preffix string, files *[]string) {
	if fn == nil {
		return
	}

	for _, f := range fn.File {
		(*files) = append((*files), preffix+f)
	}

	for _, cf := range fn.ChildFolder {
		ftp.files(cf, preffix+cf.Name+"/", files)
	}
}

func (ftp Q0388FilePathTextParser) LengthLongestFilePath(fn *FileNode) int {
	root := fn
	path := 0
	for _, f := range root.File {
		if len(f) > path {
			path = len(f)
		}
	}
	for _, cf := range root.ChildFolder {
		ftp.lengthLongestFilePath(cf, len(cf.Name)+1, &path)
	}
	return path
}

func (ftp Q0388FilePathTextParser) lengthLongestFilePath(fn *FileNode, preffixlen int, longestpath *int) {
	if fn == nil {
		return
	}

	if len(fn.File) > 0 {
		fileNameLength := 0
		for _, f := range fn.File {
			if len(f) > fileNameLength {
				fileNameLength = len(f)
			}
		}
		if fileNameLength+preffixlen > *longestpath {
			*longestpath = fileNameLength + preffixlen
		}
	}
	for _, cf := range fn.ChildFolder {
		ftp.lengthLongestFilePath(cf, preffixlen+len(cf.Name)+1, longestpath)
	}
}
