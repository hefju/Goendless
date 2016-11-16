package explorer
import (
	"time"
	"os/exec"
	"fmt"
)

func OpenExplorer(url string) {
	time.Sleep(100)
	datapath := "start vivaldi.exe "+url// 127.0.0.1:12345/hello" //chrome.exe iexplore.exe firefox.exe vivaldi.exe
	cmd := exec.Command("cmd.exe", "/c", datapath)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}