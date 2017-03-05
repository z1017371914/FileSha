package main

import "fmt"
import (
	"crypto/sha1"
	"flag"
	"goim/libs/bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type FileAndSha struct {
	name string
	size int64
	sha  []byte
}

var fileChan = make(chan FileAndSha)

func CheckError(err error) {

	if err != nil {
		fmt.Println(err)
	}

}

var sema = make(chan struct{}, 4) //开启四个线程

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	dir_list, e := ioutil.ReadDir(dir)
	CheckError(e)
	return dir_list
}

func walkDir(dir string, wg *sync.WaitGroup, fileChan chan<- FileAndSha) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, wg, fileChan)
		} else {
			subDir := filepath.Join(dir, entry.Name())
			file, err := os.Open(subDir)
			CheckError(err)
			mySha := sha1.New()
			io.Copy(mySha, file)
			file.Close()
			fileAndSha := FileAndSha{subDir, entry.Size(), mySha.Sum(nil)}
			fileChan <- fileAndSha
		}
	}

}

func main() {

	root := flag.String("root", "./", "choose the directoty")
	_ = flag.String("filters", "", "the files or directotys that will be ignored")
	resultDir := flag.String("resultDir","./","the directory where will  store the result file")
	resultFileName := flag.String("resultFileName","result.txt","the result file name")
	flag.Args()

	flag.Parse()

	var wg sync.WaitGroup

	//1.递归遍历把结果存放到数组里
	wg.Add(1)
	go walkDir(*root, &wg, fileChan)

	go func() {
		wg.Wait()
		close(fileChan)
	}()
	//2.接收数据放到缓存里
	f, err := os.Create(filepath.Join(*resultDir, *resultFileName))
	CheckError(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	for {
		fileShaModel, ok := <-fileChan
		if !ok {
			break
		}
		lineStr := fmt.Sprintf("%s,%d,%q", fileShaModel.name, fileShaModel.size, fileShaModel.sha)
		fmt.Fprintln(w, lineStr)
	}
	//3.写入文件
	err = w.Flush()
	CheckError(err)

}
