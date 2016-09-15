package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dereking/wuxing/models"
)

const URL_FORMAT = "http://hanyu.baidu.com/hanyu/ajax/search_list?wd=%s&ptype=&pn=%d"

func main() {
	keys := []string{
		"属金的字",
		"属木的字",
		"属水的字",
		"属火的字",
		"属土的字",
	}

	for _, k := range keys {

		url := fmt.Sprintf(URL_FORMAT, k, 0)
		cnt, pagecnt, _ := fetchPage(k, 0, url)

		log.Println(k, "cnt=", cnt)

		var all = make([]models.JsonDataChar, 0)
		for i := 1; i <= pagecnt; i++ {
			url := fmt.Sprintf(URL_FORMAT, k, i)
			_, _, j := fetchPage(k, i, url)

			if j != nil {
				all = append(all, j.Ret_array...)
			}
		}

		saveData(k, all)

		for _, c := range all {
			fmt.Print(c.Name[0])
		}
		fmt.Println("")
	}

}

func saveData(key string, dat []models.JsonDataChar) {
	d, err := json.Marshal(dat)
	if err != nil {
		log.Println(err)
	} else {
		ioutil.WriteFile(key+".json", d, 0777)
	}
}

func fetchPage(key string, page int, url string) (allcnt, pagecnt int, ret *models.JsonData) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return 0, 0, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, 0, nil
	}

	if page > 0 {
		ioutil.WriteFile(fmt.Sprintf("%s_%d.json", key, page), body, 0777)
	}

	var j models.JsonData
	json.Unmarshal(body, &j)

	return j.Ret_num, j.Page_total, &j
}
