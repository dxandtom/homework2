package main

import "fmt"

//lv1,打开bilibili，随便点一个视频，尽可能优雅的创建一个”视频详情页“结构体。
//lv2,为LV1的”视频详情页“结构体的”视频部分“结构体写以下几个方法
//  ○ 点赞
//  ○ 收藏
//  ○ 投币
//  ○ 一键三连
//这几个方法分别能对接收者的对应值进行更改
//并且写一个发布视频函数，入参为作者名，视频名，返回一个视频结构体

//type Author struct { //作者部分
//	Name2     string //视频名
//	Name      string //名字
//	VIP       bool   //是否是高贵的带会员
//	Signature string //签名
//	Focus     int    //关注人数
//}
//type recommendations struct { //推荐部分
//	Name     string
//	Title    string
//	Number   int //播放数量
//	barrages int //弹幕数量
//}
//
//func res(p, m, n recommendations) {
//	fmt.Println(p, m, n)
//}
//func Publish(a, b string) Author { //发布视频
//	return Author{
//		Name:  a,
//		Name2: b,
//	}
//}
//func Praise(m int) { //点赞
//
//	for i := 0; i < m; i++ {
//		if m <= i {
//
//		} else {
//			fmt.Println("the number of Praise", i)
//
//		}
//	}
//
//}
//func Collect(m int) { //收藏
//	n := 1
//
//	for i := 0; i < 99999; i++ {
//		if n <= i {
//
//		} else {
//			fmt.Println("the number of Collect", i)
//			break
//		}
//	}
//	n++
//}
//func coin(m int) { //投币
//	n := 1
//	for i := 0; i < 99999; i++ {
//		if n <= i {
//
//		} else {
//			fmt.Println("the number of Collect", i)
//			break
//		}
//	}
//	n++
//}
//func main() {
//	//作者部分
//	var Animenz = Author{
//		Name:      "Animenzz",
//		VIP:       true,
//		Signature: "Hallo 我是A叔！",
//		Focus:     1796000,
//	}
//	fmt.Println("作者部分：")
//	fmt.Println(Animenz)
//	//推荐部分
//	a := recommendations{
//		Name:     "Animenzz",
//		Title:    "unravel: piano version",
//		Number:   18840000, //播放数量
//		barrages: 104000,   //弹幕数量
//	}
//	var b = recommendations{
//		Name:     "Animenzz",
//		Title:    "secret base: piano version",
//		Number:   1881000, //播放数量
//		barrages: 5363,    //弹幕数量
//	}
//	var c = recommendations{
//		Name:     "Animenzz",
//		Title:    "kenan: piano version",
//		Number:   2364000, //播放数量
//		barrages: 4907,    //弹幕数量
//	}
//	fmt.Println("视频推荐部分：")
//	res(a, b, c)
//	fmt.Println("点赞")
//	Praise(4)
//}

//lv3
//写一个接收者函数:func Receiver(v interface{})  {
//    switch v.(Type)
//    case ...
//}该接收者能够判断传入参数的类型，并作出不同的反应
//Receiver("你好吗")
//"这个是string"
//Receiver(32)
//"这个是int"
//Receiver(true)
//"这个是bool"

func Receiver(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("这个是int类型")
	case string:
		fmt.Println("这个是string类型")
	case bool:
		fmt.Println("这个是bool类型")
	}
}
func main() {
	Receiver("dx")
	Receiver(32)
	Receiver(true)
}
