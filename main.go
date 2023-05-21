package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/GoToUse/treeprint"
)

// 存储大小之间的单位
const unit = 1024

// command line arguments
var (
	flag_folder_path string
	exclude_dirs     excludeDirs
	humanRead        bool
)

// global variable
var (
	te      = treeprint.New()
	folders int32
	files   int32
)

type excludeDirs []string

func (e *excludeDirs) String() string {
	return fmt.Sprint(*e)
}

func (e *excludeDirs) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func convertToAbsPath(root string) (path string, err error) {
	path, err = filepath.Abs(root)
	return path, err
}

func folderInExcludeArrays(name string) bool {
	for _, dir := range exclude_dirs {
		if name == dir {
			return true
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
	entrys, err := os.ReadDir(folder)
	// 不记录子目录的大小
	var branch treeprint.Tree
	if folder == flag_folder_path {
		branch = tree
	} else {
		baseFolder := path.Base(folder)
		branch = tree.AddBranch(baseFolder)
	}

	if err != nil {
		return 0, err
	}

	entrysLen := len(entrys)

	if entrysLen == 0 {
		return 0, nil
	}

	// wg.Add(entrysLen)

	for i := 0; i < entrysLen; i++ {
		subFolder := entrys[i]
		if !folderInExcludeArrays(subFolder.Name()) {
			wg.Add(1)
			go calc(subFolder, &wg, folder, &total, branch)
		}
	}

	wg.Wait()
	return total, nil
}

// ByteCountIEC, 以1024作为基数, 将字节转为对应的KB等单位
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
	flag.StringVar(&flag_folder_path, "f", ".", "Folder path.")
	flag.Var(&exclude_dirs, "e", "Exclude directories.")
	flag.BoolVar(&humanRead, "h", false, "Print the size in a more human readable way.")
}

func main() {
	flag.Parse()
	size, err := Parallel(flag_folder_path, te)
	defer catchError()
	if err != nil {
		panic(err)
	}

	rootPath, err := convertToAbsPath(flag_folder_path)
	defer catchError()
	if err != nil {
		panic(err)
	}

	for _, d := range []string{rootPath, te.String(), fmt.Sprintf("\033[1mSummary:\033[0m Total folders: \033[31m%d\033[0m Total files: \033[32m%d\033[0m Total size: \033[34m%s\033[0m", folders, files, ByteCountIEC(size))} {
		fmt.Println(d)
	}
}
