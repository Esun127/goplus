package main


import "fmt"


import "os"


func main() {
	
	f := createFile("/tmp/file.log")
	defer closeFile(f)
	writeFile(f)
}



func createFile(p string) *os.File {
	fmt.Println("createing ...")	
	f, err := os.Create(p)
	if err != nil {
		panic(err)	
	}
	return f
}


func writeFile(f *os.File) {
	fmt.Println("writeing ...")
	fmt.Fprintln(f, "data")
}


func closeFile(f *os.File) {
	fmt.Println("closing ...")
	f.Close()
}

