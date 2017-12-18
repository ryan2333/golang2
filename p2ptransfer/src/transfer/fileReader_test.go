package transfer

import (
	"log"
	"os"
)

func main() {
	var fh TranseData
	f, err := os.Open("D:\\workspace\testText/a.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fw, err1 := os.Create("D:\\workspace\testText/a1.txt")
	if err1 != nil {
		log.Println(err1)
	}
	defer fw.Close()

	ferr := fh.Read(f)
	if ferr != nil {
		log.Fatal(ferr)
	}
	fwerr := fh.Write(fw)
	if fwerr != nil {
		log.Fatal(fwerr)
	}
}
