package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/GoToUse/treeprint"
)

// 存储大小之间的单位
const unit = 1024

// command line arguments
var (
	// 目标地址
	folderPath string
	// 排除在外的地址数组
	excludeDirArray excludeDirs
	// 是否输出size
	humanRead bool
	// 是否部分匹配
	partialMatch bool
)

// global variable
var (
	tp      = treeprint.New()
	folders int32
	files   int32
)

type excludeDirs []string

func (e *excludeDirs) String() string {
	return fmt.Sprint(*e)
}

// multiExists check if there are more than `multi` bool values in the evalArray.
func multiExists(evalArray []bool, multi int) bool {
	var n int
	for _, _b := range evalArray {
		if _b {
			n++
		}
	}

	if n > multi {
		return true
	}
	return false
}

func (e *excludeDirs) Set(value string) error {
	commaRegex := regexp.MustCompile(`,`)
	commaMatched := commaRegex.MatchString(value)

	spaceRegex := regexp.MustCompile(`\s`)
	spaceMatched := spaceRegex.MatchString(value)

	semicolonRegex := regexp.MustCompile(`;`)
	semicolonMatched := semicolonRegex.MatchString(value)

	switch {
	case multiExists([]bool{commaMatched, spaceMatched, semicolonMatched}, 1):
		return errors.New("spaces and commas and semicolons cannot be included at the same time, " +
			"there is only one type: `spaces` or `commas` or `semicolons`")
	case commaMatched:
		commas := strings.Split(value, ",")
		*e = commas
		return nil
	case spaceMatched:
		spaces := strings.Split(value, " ")
		*e = spaces
		return nil
	case semicolonMatched:
		semicolons := strings.Split(value, ";")
		*e = semicolons
		return nil
	default:
		*e = append(*e, value)
		return nil
	}
}

// convertToAbsPath convert a passed in path to an absolute path.
func convertToAbsPath(root string) (path string, err error) {
	path, err = filepath.Abs(root)
	return path, err
}

func folderInExcludeArrays(subDir os.DirEntry) bool {
	if subDir.IsDir() {
		name := subDir.Name()
		for _, dir := range excludeDirArray {
			// This will exclude all `dir`s whose names are included in `name`.
			if partialMatch && strings.Contains(name, dir) {
				return true
			} else if name == dir {
				return true
			}
		}
	}
	return false
}

func catchError() {
	err := recover()

	if err != nil {
		e := fmt.Errorf("error: %v", err)
		fmt.Println(e.Error())
		return
	}
}

func calc(entry fs.DirEntry, wg *sync.WaitGroup, folder string, total *int64, tree treeprint.Tree) {
	defer wg.Done()

	if entry.IsDir() {
		size, err := Parallel(path.Join(folder, entry.Name()), tree)
		defer catchError()
		if err != nil {
			panic(err)
		}
		atomic.AddInt32(&folders, 1)
		atomic.AddInt64(total, size)
		return
	}

	info, err := entry.Info()
	defer catchError()
	if err != nil {
		panic(err)
	}

	size := info.Size()
	atomic.AddInt32(&files, 1)
	atomic.AddInt64(total, size)

	if humanRead {
		tree.AddNode(fmt.Sprintf("%s (%s)", entry.Name(), ByteCountIEC(size)))
	} else {
		tree.AddNode(entry.Name())
	}
}

// Parallel execution, fast enough
func Parallel(folder string, tree treeprint.Tree) (total int64, e error) {
	var wg sync.WaitGroup
	entryS, err := os.ReadDir(folder)
	// Do not record the size of directory.
	var branch treeprint.Tree
	if folder == folderPath {
		branch = tree
	} else {
		baseFolder := path.Base(folder)
		branch = tree.AddBranch(baseFolder)
	}

	if err != nil {
		return 0, err
	}

	entrySLen := len(entryS)

	if entrySLen == 0 {
		return 0, nil
	}

	// wg.Add(entrySLen)

	for i := 0; i < entrySLen; i++ {
		subFolder := entryS[i]
		if !folderInExcludeArrays(subFolder) {
			wg.Add(1)
			go calc(subFolder, &wg, folder, &total, branch)
		}
	}

	wg.Wait()
	return total, nil
}

// ByteCountIEC is based on 1024, converts the bytes to corresponding units such as KB.
func ByteCountIEC(b int64) string {
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func init() {
	flag.StringVar(&folderPath, "f", ".", "Folder path.")
	flag.Var(&excludeDirArray, "e", "Exclude directories.")
	flag.BoolVar(&humanRead, "h", false, "Print the size in a more human readable way.")
	flag.BoolVar(&partialMatch, "p", false, "Support partial match.")
}

func main() {
	flag.Parse()
	size, err := Parallel(folderPath, tp)
	defer catchError()
	if err != nil {
		panic(err)
	}

	rootPath, err := convertToAbsPath(folderPath)
	defer catchError()
	if err != nil {
		panic(err)
	}

	for _, d := range []string{rootPath, tp.String(), fmt.Sprintf("\033[1mSummary:\033[0m Total folders: \033[31m%d\033[0m Total files: \033[32m%d\033[0m Total size: \033[34m%s\033[0m", folders, files, ByteCountIEC(size))} {
		fmt.Println(d)
	}
}
