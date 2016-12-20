/*下载图片, 这个功能很久以前已经想做了, 但要真正做出来还是比较难的.

*/
package main

import (
"bytes"
"fmt"
"io"
"io/ioutil"
"net/http"
"os"
"regexp"
"strings"
)

var folder string//图片存放文件夹

func main() {
	//在当前程序下建立一个存放图片的文件夹
	folder = "tmp"
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	}

	//下载html
	response, err := http.Get("https://www.wstx.com/p-22052-9")
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	source := string(body)
	pattern1 := `\<img(.+?)\/\>`
	part := regexp.MustCompile(pattern1)
	match := part.FindAllStringSubmatch(source, -1)//找出所有的img标签
	for _, v := range match {
		findaddr(v[1])
		if hasimage < -1 {
			break
		}
	}
}

var hasimage int = 0 //测试时候用, 一般就下载两幅图片就停止,检查程序是否运行正确.

func findaddr(imghtml string) {//imghtml是<img ?>标签, 还需要进一步提取
	pattern2 := `http://(.+?).jpg`
	part := regexp.MustCompile(pattern2)
	match := part.FindStringSubmatch(imghtml)//找到图片地址, 如果是gif,png啊需要另外处理好像是 [.jpg|.gif|.png]
	if len(match) > 0 {
		fmt.Println(match[0])//真实路径 //fmt.Println("#", imghtml)
		getImg(match[0])
		hasimage++
	}

}
func getImg(url string) (n int64, err error) {
	//找到文件名, 如果没必要就自己随机起一个名字
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	//fmt.Println(name)

	//如果不下载就能够知道图片大小就最好, 现在是先把图片下载到本地,然后只保存100K以上的
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if len(pix) > 1000*100 { //大于100k才保存
		out, err := os.Create(folder + "/" + name)
		if err != nil {
			fmt.Printf("%s", err)
		}
		defer out.Close()
		n, err = io.Copy(out, bytes.NewReader(pix))
	}
	return
}