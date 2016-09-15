package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/dereking/wuxing/models"
)

type AllChar [5][]models.JsonDataChar

var all AllChar

func main() {

	for i := 0; i < 5; i++ {
		all[i] = make([]models.JsonDataChar, 0)
	}

	dat, _ := ioutil.ReadFile("属金的字.json")
	json.Unmarshal(dat, &all[0])

	dat, _ = ioutil.ReadFile("属木的字.json")
	json.Unmarshal(dat, &all[1])

	dat, _ = ioutil.ReadFile("属水的字.json")
	json.Unmarshal(dat, &all[2])

	dat, _ = ioutil.ReadFile("属火的字.json")
	json.Unmarshal(dat, &all[3])

	dat, _ = ioutil.ReadFile("属土的字.json")
	json.Unmarshal(dat, &all[4])

	prompt()
}

func prompt() {
	running := true

	r := bufio.NewReader(os.Stdin)
	for running {
		fmt.Print(">输入要查询的五行或汉字: 1：金；2：木；3：水；4：火；5：土；")
		d, _, _ := r.ReadLine()
		cmd := string(d)
		i, err := strconv.ParseInt(cmd, 10, 32)

		if (err == nil) || ((i >= 1) && (i <= 4)) {

			fmt.Println(i-1, err)
			for _, c := range all[i-1] {
				fmt.Print(c.Name[0])
			}
			fmt.Println("")
		} else {
			SearchChar(cmd)
		}
	}
}

func SearchChar(ch string) {
	for _, cs := range all {
		for _, c := range cs {
			if strings.Compare(c.Name[0], ch) == 0 {
				PrintStrings("汉字", c.Name)
				PrintStrings("拼音", c.Pinyin)
				PrintStrings("五行", c.Line_type)
				PrintStrings("笔画", c.Stroke_count)
				PrintStrings("部首", c.Radicals)
				PrintStrings("解释", c.Definition)
			}
		}
	}
}

func PrintStrings(title string, strs []string) {
	fmt.Print(title + " : ")
	if len(strs) == 1 {
		fmt.Println(strs[0])
	} else {
		fmt.Println()
		for i, line := range strs {
			fmt.Println("\t", i+1, ".", line)
		}

	}
}
