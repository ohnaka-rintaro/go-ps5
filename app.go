package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// // 対象のURL
	// url := ""
	// res, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer res.Body.Close()
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// buf := bytes.NewBuffer(body)
	// html := buf.String()
	// fmt.Println(html)

	res, err := http.Get("https://books.rakuten.co.jp/rb/16462860/?bkts=1")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}
	fmt.Println(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	// ここのFind()等でデータの指定をする
	doc.Find("#itemReview").Each(func(i int, s *goquery.Selection) {
        // url, _ := s.Attr("itemprop")
		name, _ := s.Html()
		fmt.Println(s.Text())
        fmt.Printf("name = %s\n", name)
	})
}
