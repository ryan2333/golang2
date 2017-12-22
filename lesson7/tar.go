//tar打包工具，go run tar.go dirname > tarName.tar
package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

func mktar(dir string, w io.Writer) error {
	base := filepath.Base(dir)
	tr := tar.NewWriter(w)
	defer tr.Close()
	return filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		h, err := tar.FileInfoHeader(info, "") //获取文件头信息
		if err != nil {
			return err
		}
		p, err := filepath.Rel(base, name)
		h.Name = filepath.Join(base, p)          //拼接文件全路径
		if err = tr.WriteHeader(h); err != nil { //写入文件头信息
			return err
		}
		if info.Mode().IsRegular() {
			io.Copy(tr, f)
		}
		return nil
	})
}

func main() {
	err := mktar(os.Args[1], os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
