package main


import "fmt"
import "io/ioutil"
import "os/exec"


func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))


	grepCmd := exec.Command("grep","-o", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))


	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(">ls -alh")
	fmt.Println(string(lsOut))

	psCmd := exec.Command("bash","-c","psf", "auefdfdsfsd")
	psOut, err := psCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(psOut))
}
