package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
	"sync/atomic"

	"github.com/GoToUse/treeprint"
)

// 存储大小之间的单位
const unit = 1024

// command line arguments
var flag_folder_path string

// global variable
var te = treeprint.New()

func calc(entry fs.DirEntry, wg *sync.WaitGroup, folder string, total *int64, tree treeprint.Tree) {
	defer wg.Done()

	if entry.IsDir() {
		size, err := Parallel(path.Join(folder, entry.Name()), tree)
		if err != nil {
			panic(err)
		}
		atomic.AddInt64(total, size)
		return
	}

	info, err := entry.Info()
	if err != nil {
		panic(err)
	}

	size := info.Size()
	atomic.AddInt64(total, size)
	tree.AddNode(fmt.Sprintf("%s (%s)", path.Base(entry.Name()), ByteCountIEC(size)))
}

// Parallel execution, fast enough
func Parallel(folder string, tree treeprint.Tree) (total int64, e error) {
	var wg sync.WaitGroup
	entrys, err := os.ReadDir(folder)
	// 不记录子目录的大小
	branch := tree.AddBranch(path.Base(folder))

	// catch panic
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	if err != nil {
		return 0, err
	}

	entrysLen := len(entrys)

	if entrysLen == 0 {
		return 0, nil
	}

	wg.Add(entrysLen)

	for i := 0; i < entrysLen; i++ {
		go calc(entrys[i], &wg, folder, &total, branch)
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
	flag.Parse()
}

func main() {
	size, err := Parallel(flag_folder_path, te)
	if err != nil {
		panic(err)
	}

	for _, d := range []string{fmt.Sprintf("Total size: %s", ByteCountIEC(size)), "Output in tree format:", te.String()} {
		fmt.Println(d)
	}
}
