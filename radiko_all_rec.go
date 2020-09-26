package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	data := httpGet("http://radiko.jp/v3/program/station/weekly/TBS.xml")
	recorded_path := "/home/chinachu/chinachu/recorded/radio/"

	result := Radiko{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, progs := range result.Stations.Station.Progs {
		for _, prog := range progs.Prog {

			t, err := time.Parse("20060102150405", prog.Ft)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("%s", t.Format("04 15 02 01 "))
			fmt.Printf("* /root/docker_rec_radiko/radiko_docker_run.sh TBS ")
			fmt.Printf("%v %v ", (prog.Dur/60)+1, recorded_path) //radikoの場合1分ぐらい遅く終了しないと、番組の最後まで入らない
			fmt.Printf("\"%v\" \n", prog.Title)

		}
	}

}

type Radiko struct {
	XMLName  xml.Name `xml:"radiko"`
	Text     string   `xml:",chardata"`
	Ttl      string   `xml:"ttl"`
	Srvtime  string   `xml:"srvtime"`
	Stations struct {
		Text    string `xml:",chardata"`
		Station struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Name  string `xml:"name"`
			Progs []struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Prog []struct {
					Text         string `xml:",chardata"`
					ID           string `xml:"id,attr"`
					MasterID     string `xml:"master_id,attr"`
					Ft           string `xml:"ft,attr"`
					To           string `xml:"to,attr"`
					Ftl          string `xml:"ftl,attr"`
					Tol          string `xml:"tol,attr"`
					Dur          int    `xml:"dur,attr"`
					Title        string `xml:"title"`
					URL          string `xml:"url"`
					FailedRecord string `xml:"failed_record"`
					TsInNg       string `xml:"ts_in_ng"`
					TsOutNg      string `xml:"ts_out_ng"`
					Desc         string `xml:"desc"`
					Info         string `xml:"info"`
					Pfm          string `xml:"pfm"`
					Img          string `xml:"img"`
					Metas        struct {
						Text string `xml:",chardata"`
						Meta struct {
							Text  string `xml:",chardata"`
							Name  string `xml:"name,attr"`
							Value string `xml:"value,attr"`
						} `xml:"meta"`
					} `xml:"metas"`
				} `xml:"prog"`
			} `xml:"progs"`
		} `xml:"station"`
	} `xml:"stations"`
}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}
