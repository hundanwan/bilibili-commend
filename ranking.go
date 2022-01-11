package main

import "github.com/tidwall/gjson"

type Ranking struct{
	Index int `json:"index"`
	Title string `json:"title"`
	Pic   string `json:"pic"`
	Description  string `json:"description"`
	Bvid string `json:"bvid"`
	Type string `json:"type"`
}
type Rankings struct {
	Ranking []Ranking
	Type string
}
func (r *Ranking)setTitle(title string){
	r.Title = title
}
func (r *Ranking)setPic(pic string){
	r.Pic = pic
}
func (r *Ranking)setDesc(desc string){
	r.Description = desc
}
func (r *Ranking)setBvid(bvid string){
	r.Bvid = BiliLink + bvid
}

func (r *Rankings)setType(rid int)  {
	switch rid {
	case 1:
		r.Type="动 画"
	case 3:
		r.Type="音 乐"
	case 13:
		r.Type="番 剧"
	case 4:
		r.Type="游 戏"
		

	}

}
func (r *Ranking)setAll(json gjson.Result)  {
	r.setTitle(json.Get("title").String())
	r.setPic(json.Get("pic").String())
	r.setDesc(json.Get("description").String())
	r.setBvid(json.Get("bvid").String())
}

