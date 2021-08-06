package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	FILE_URL_PREFIX = "https://raw.githubusercontent.com/yanzhoupan/poems/main/tang/"
	FILES_CNT_TANG  = 900
)

type Poem struct {
	Volume   string
	Sequence string
	Title    string
	Author   string
	Content  []string
}

// generate file name to download
func genFileName() string {
	randIdx := rand.Int31n(FILES_CNT_TANG-1) + 1
	digits := 1
	tmp := randIdx
	for tmp/10 > 0 {
		digits += 1
		tmp /= 10
	}

	suffix := strings.Repeat("0", 4-digits)
	return "zhs_" + suffix + fmt.Sprintf("%d", randIdx) + ".json"
}

// randomly select a poem and print it out
func printRandomPoem(fileName string) {
	poemFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file failed: %v\n", err)
	}
	defer poemFile.Close()

	byteValue, _ := ioutil.ReadAll(poemFile)
	var poems []Poem
	json.Unmarshal([]byte(byteValue), &poems)

	poemsLen := len(poems)
	randIdx := rand.Int31n(int32(poemsLen-1)) + 1
	randPoem := poems[randIdx]
	fmt.Println("题目：", randPoem.Title)
	fmt.Println("作者：", randPoem.Author)
	for _, line := range randPoem.Content {
		fmt.Println(line)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	fileName := genFileName()
	downloadURL := FILE_URL_PREFIX + fileName
	curl := exec.Command("curl", "-LO", downloadURL)
	curl.Run() // wait until curl finish

	printRandomPoem(fileName)

	// remove the downloaded json file
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("remove file filed: %v\n", err)
	}
}
