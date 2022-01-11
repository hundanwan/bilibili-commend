package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

const CommendMax = 40
const CommendNeed = 3
const RankingMax = 11
const BiliLink = "https://www.bilibili.com/video/"
type Info struct {
	Commends [3]Commend
	Rankings []Rankings
}
func main(){
	//获取api的json数据

	//获取所有推荐视频数据

	resp, _ := http.Get("http://api.bilibili.com/x/web-interface/archive/related?aid=7")
	bytes, _ := ioutil.ReadAll(resp.Body)
	datas := gjson.Get(string(bytes),"data")

	//生成推荐视频类数组并赋值
	commends := [CommendNeed]Commend{}
	commendsDatas := [CommendMax]Commend{}
	for i := 0;i < CommendMax;i++ {
		p := "..0." + strconv.Itoa(i)
		data := datas.Get(p)
		commendsDatas[i].SetAll(data)
		commendsDatas[i].Index = i
	}

	//循环生成三个随机推荐视频
	for i := 0;i < CommendNeed;i++ {
		commends = RandVideo(commendsDatas)
	}


	//获取榜单视频json
	rids := []int{13,1,3,4}
	rankings := []Rankings{}
	for i,v := range rids {
		resp, _ = http.Get("http://api.bilibili.com/x/web-interface/ranking/region?rid=" + strconv.Itoa(v) + "&day=3")
		bytes, _ = ioutil.ReadAll(resp.Body)

		var ranking = Rankings{}

		for i := 0;i < RankingMax;i++ {
			p := "data."+strconv.Itoa(i)
			var rank = Ranking{}
			rank.setAll(gjson.Get(string(bytes),p))
			rank.Index = i

			ranking.Ranking = append(ranking.Ranking,rank)
		}
		rankings = append(rankings,ranking)
		rankings[i].setType(v)
	}

	all := Info{
		Commends: commends,
		Rankings: rankings,
	}
	//启动服务
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	router.StaticFS("/static",http.Dir("./static"))
	router.GET("/", func(context *gin.Context) {
		context.HTML(200,"index.html",all)
	})
	router.Run()



}

//func main()  {
//	body := StartCrawler("https://www.bilibili.com/")
//	fmt.Println(body)
//}