package cmd
import (
	"net"
	"log"
	"fmt"
	"bytes"
	"os/exec"
	"strconv"
	str "strings"
	"golang.org/x/crypto/ssh"
)

func runCmd() {

	var stdOut, stdErr bytes.Buffer

	cmd := exec.Command("ssh", "username@192.168.1.4", "if [ -d liujx/project ];then echo 0;else echo 1;fi")
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd exec failed: %s : %s", fmt.Sprint(err), stdErr.String())
	}

	fmt.Print(stdOut.String())
	ret, err := strconv.Atoi(str.Replace(stdOut.String(), "\n", "", -1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d, %s\n", ret, stdErr.String())
}
