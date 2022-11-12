package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"my-spider/pkg"
)

var reHtml = regexp.MustCompile(`<.*?>`)
var output *os.File

func init() {
	fo, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}

	output = fo

	// 寫入表頭
	var data []string
	data = append(data, []string{
		"學校名稱",
		"重要招生資訊",
		"「在校學業成績」全校排名百分比標準",
		"轉系規定",
		"重要事項說明",
		"校系名稱及代碼",
		"招生名額",
		"外加名額",
		"學群類別",
		"招生名額各學群可選填志願數",
		"外加名額各學群可選填志願數",
		"校系分則詳細資料",

		"校系代碼",
		"學群類別",
		"招生名額",
		"可填志願數",
		"外加名額",
		"可填志願數",
		"學測、英聽檢定",
		"科目1",
		"科目2",
		"科目3",
		"科目4",
		"科目5",
		"科目6",
		"科目7",
		"標準1",
		"標準2",
		"標準3",
		"標準4",
		"標準5",
		"標準6",
		"標準7",
		"分發比序項目",
		"備註",
	}...)
	output.WriteString(pkg.Utf8Big5("\"" + strings.Join(data, "\",\"") + "\"\n"))
	data = nil
}

func main() {
	defer output.Close()

	// 取得首頁
	client := &http.Client{}

	link := "https://www.cac.edu.tw/star112/system/8ColQry_xfor112Star_Z84eH3ep/TotalGsdShow.htm"
	param := url.Values{}
	req, err := http.NewRequest("GET", link, strings.NewReader(param.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("DNT", "1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// 取得 html
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(ai int, a *goquery.Selection) {
		txt := a.Text()
		txt = strings.ReplaceAll(txt, " ", "")
		txt = strings.ReplaceAll(txt, "\n", "")
		fmt.Println(ai, txt)

		href, _ := a.Attr("href")
		get(client, txt, href)
	})
}

func get(client *http.Client, title, link string) {
	link = "https://www.cac.edu.tw/star112/system/8ColQry_xfor112Star_Z84eH3ep/" + link
	param := url.Values{}
	req, err := http.NewRequest("GET", link, strings.NewReader(param.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	req.Header.Set("Host", "www.cac.edu.tw")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("DNT", "1")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Set("Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://www.cac.edu.tw/star112/system/8ColQry_xfor112Star_Z84eH3ep/TotalGsdShow.htm")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-TW,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-CN;q=0.6,ru;q=0.5")
	req.Header.Set("Cookie", "fwchk=SVx18uV/8u1RedeKui9qaf/pyPc0002")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// 取得 html
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	data0 := ""
	html := ""
	data0 += "\"'" + pkg.Utf8Big5(title) + "\","

	// 第一張表
	table := doc.Find("table:nth-of-type(1)")

	// 重要招生資訊
	data0 += "\"\","

	// 「在校學業成績」全校排名百分比標準
	html, _ = table.Find("tr:nth-of-type(2)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data0 += "\"'" + pkg.Utf8Big5(html) + "\","

	// 轉系規定
	html, _ = table.Find("tr:nth-of-type(3)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data0 += "\"'" + pkg.Utf8Big5(html) + "\","

	// 第二張表
	table = doc.Find("table:nth-of-type(2)")

	// 重要事項說明
	html, _ = table.Find("tr:nth-of-type(2)").Find("td:nth-of-type(1)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data0 += "\"'" + pkg.Utf8Big5(html) + "\","

	// 第三張表
	table = doc.Find("table:nth-of-type(3)")
	table.Find("tr").Each(func(tri int, tr *goquery.Selection) {
		if tri == 0 {
			return
		}

		data := ""
		tr.Find("td").Each(func(tdi int, td *goquery.Selection) {
			txt := td.Text()
			txt = strings.ReplaceAll(txt, " ", "")
			txt = strings.ReplaceAll(txt, "\n", "")
			data += "\"" + pkg.Utf8Big5(txt) + "\","

			if tdi == 0 {
				fmt.Println(txt)
			}
		})

		tr.Find("a").Each(func(ai int, a *goquery.Selection) {
			href, _ := a.Attr("href")
			data += getDetail(client, href)
		})

		output.WriteString(data0)
		output.WriteString(data)
		output.WriteString("\n")
	})
}

func getDetail(client *http.Client, link string) string {
	link = "https://www.cac.edu.tw/star112/system/8ColQry_xfor112Star_Z84eH3ep/" + link
	param := url.Values{}
	req, err := http.NewRequest("GET", link, strings.NewReader(param.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	req.Header.Set("Host", "www.cac.edu.tw")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("DNT", "1")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Set("Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer",
		"https://www.cac.edu.tw/star112/system/8ColQry_xfor112Star_Z84eH3ep/ShowSchGsd.php?colno=001")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-TW,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-CN;q=0.6,ru;q=0.5")
	req.Header.Set("Cookie", "fwchk=8+V3+KZJX6N63ijiyUpwFWepE7o0002")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// 取得 html
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	data := ""
	html := ""

	// 第一張表
	table := doc.Find("table:nth-of-type(1)")

	// 校系代碼
	html, _ = table.Find("tr:nth-of-type(3)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"'" + pkg.Utf8Big5(html) + "\","

	// 學群類別
	html, _ = table.Find("tr:nth-of-type(4)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	// 招生名額
	html, _ = table.Find("tr:nth-of-type(5)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	// 可填志願數
	html, _ = table.Find("tr:nth-of-type(6)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	// 外加名額
	html, _ = table.Find("tr:nth-of-type(7)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	// 可填志願數
	html, _ = table.Find("tr:nth-of-type(8)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	// 學測、英聽檢定
	data += "\"\","

	// 科目
	html, _ = table.Find("tr:nth-of-type(3)").Find("td:nth-of-type(3)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	arr := strings.Split(html, "\n")
	for _, v := range arr {
		data += "\"" + pkg.Utf8Big5(v) + "\","
	}

	for i := len(arr); i < 7; i++ {
		data += "\"\","
	}

	// 標準
	html, _ = table.Find("tr:nth-of-type(3)").Find("td:nth-of-type(4)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	arr = strings.Split(html, "\n")
	for _, v := range arr {
		data += "\"" + pkg.Utf8Big5(v) + "\","
	}

	for i := len(arr); i < 7; i++ {
		data += "\"\","
	}

	// 分發比序項目
	html, _ = table.Find("tr:nth-of-type(2)").Find("td:nth-of-type(3)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	// 備註
	html, _ = table.Find("tr:nth-of-type(9)").Find("td:nth-of-type(2)").Html()
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "<br/>", "\n")
	html = reHtml.ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "\"", "")
	html = strings.Trim(html, " ")
	data += "\"" + pkg.Utf8Big5(html) + "\","

	return data
}
