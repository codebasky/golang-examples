package pipeline

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Result struct {
	Path   string
	Digest [md5.Size]byte
	Err    error
}

func MD5ALL(root string) (<-chan Result, <-chan error) {
	result := make(chan Result)
	err := make(chan error, 1)
	walk := func(root string) {
		defer close(result)
		err <- filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			if info.IsDir() {
				fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
				return nil
			}
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			result <- Result{
				Path:   path,
				Digest: md5.Sum(b),
			}
			fmt.Printf("visited file or dir: %q\n", path)
			return nil
		})
	}
	go walk(root)
	return result, err
}
