package main


import (
	"github.com/tidwall/gjson"

	"math/rand"
	"time"
)
type Commend struct{
	Index int `json:"index"`
	Title string `json:"title"`
	Pic   string `json:"pic"`
	Desc  string `json:"desc"`
	Bvid  string `json:"bvid"`
	//face string `json:"face"`
}

func (c *Commend)SetTitle(title string){
	c.Title = title
}
func (c *Commend)SetPic(pic string){
	c.Pic = pic
}
func (c *Commend)SetDesc(desc string){
	c.Desc = desc
}
func (c *Commend)SetBvid(bvid string){
	c.Bvid = BiliLink + bvid
}
func (c *Commend) SetAll(json gjson.Result)  {
	c.SetTitle(json.Get("title").String())
	c.SetPic(json.Get("pic").String())
	c.SetDesc(json.Get("desc").String())
	c.SetBvid(json.Get("bvid").String())
}
func RandVideo(commendsData [CommendMax]Commend) [CommendNeed]Commend {
	c := [CommendNeed]Commend{}
	rand.Seed(time.Now().Unix())
	var random [CommendNeed]int
	flag := true
	for flag{
		for i := 0;i < CommendNeed;i++{
			random[i] = rand.Intn(CommendMax)
		}
		if random[0] != random[1]{
			if random[0] != random[2] && random[1] != random[2]{
				flag = false
			}
		}
	}

	for i := 0;i < CommendNeed;i++{
		c[i] = commendsData[random[i]]
	}
	return c
}
