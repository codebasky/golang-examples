package pipeline

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

const (
	NUMOFDIGESTER = 10
)

type Result struct {
	Path   string
	Digest [md5.Size]byte
	Err    error
}

func FilePathGenerator(done <-chan struct{}, search string) (<-chan string, <-chan error) {
	paths := make(chan string)
	err := make(chan error, 1)
	walk := func(root string) {
		defer close(paths)
		err <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			if info.IsDir() {
				fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
				return nil
			}

			select {
			case paths <- path:
				//fmt.Printf("file path found %s\n", path)
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}
	go walk(search)
	return paths, err
}

func Digester(done <-chan struct{}, paths <-chan string, results chan<- Result) {
	for path := range paths {
		b, err := os.ReadFile(path)
		//fmt.Printf("Read the file at : %s path\n", path)
		select {
		case results <- Result{
			Path:   path,
			Digest: md5.Sum(b),
			Err:    err,
		}:
		case <-done:
			return
		}
	}
}

func MD5ALL(done <-chan struct{}, root string) (map[string][md5.Size]byte, error) {
	paths, err := FilePathGenerator(done, root)
	results := make(chan Result)
	var wg sync.WaitGroup

	wg.Add(NUMOFDIGESTER)
	for i := 0; i < NUMOFDIGESTER; i++ {
		go func() {
			defer wg.Done()
			Digester(done, paths, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
		fmt.Println("closed the result channel")
	}()

	digest := make(map[string][md5.Size]byte)
	for result := range results {
		if result.Err != nil {
			fmt.Printf("Error %s on processing file %s\n", result.Err, result.Path)
			return nil, result.Err
		}
		digest[result.Path] = result.Digest
	}
	return digest, <-err
}
