// getComic project doc.go

/*
getComic document
*/
package main

/*
	https://www.manhuaniu.com/manhua/2652/
	从这个漫画网址抽出以下网址写起来有点费事儿。。。就是手动处理了

	输出的格式是为了迎合我的git项目 https://github.com/soekchl/ImageOutHtml.git
*/

var reqList = []string{
	"https://www.manhuaniu.com/manhua/2652/54365.html",
	"https://www.manhuaniu.com/manhua/2652/54366.html",
	"https://www.manhuaniu.com/manhua/2652/54367.html",
	"https://www.manhuaniu.com/manhua/2652/54368.html",
	"https://www.manhuaniu.com/manhua/2652/54369.html",
	"https://www.manhuaniu.com/manhua/2652/54370.html",
	"https://www.manhuaniu.com/manhua/2652/54371.html",
	"https://www.manhuaniu.com/manhua/2652/54372.html",
	"https://www.manhuaniu.com/manhua/2652/54373.html",
	"https://www.manhuaniu.com/manhua/2652/54374.html",
	"https://www.manhuaniu.com/manhua/2652/54375.html",
	"https://www.manhuaniu.com/manhua/2652/54376.html",
	"https://www.manhuaniu.com/manhua/2652/54377.html",
	"https://www.manhuaniu.com/manhua/2652/54378.html",
	"https://www.manhuaniu.com/manhua/2652/54379.html",
	"https://www.manhuaniu.com/manhua/2652/54380.html",
	"https://www.manhuaniu.com/manhua/2652/54381.html",
	"https://www.manhuaniu.com/manhua/2652/54382.html",
	"https://www.manhuaniu.com/manhua/2652/54383.html",
	"https://www.manhuaniu.com/manhua/2652/54384.html",
	"https://www.manhuaniu.com/manhua/2652/54385.html",
	"https://www.manhuaniu.com/manhua/2652/54386.html",
	"https://www.manhuaniu.com/manhua/2652/54387.html",
	"https://www.manhuaniu.com/manhua/2652/54388.html",
	"https://www.manhuaniu.com/manhua/2652/54389.html",
	"https://www.manhuaniu.com/manhua/2652/54390.html",
	"https://www.manhuaniu.com/manhua/2652/54391.html",
	"https://www.manhuaniu.com/manhua/2652/54392.html",
	"https://www.manhuaniu.com/manhua/2652/54393.html",
	"https://www.manhuaniu.com/manhua/2652/54394.html",
	"https://www.manhuaniu.com/manhua/2652/54395.html",
	"https://www.manhuaniu.com/manhua/2652/54396.html",
	"https://www.manhuaniu.com/manhua/2652/54397.html",
	"https://www.manhuaniu.com/manhua/2652/54398.html",
	"https://www.manhuaniu.com/manhua/2652/54399.html",
	"https://www.manhuaniu.com/manhua/2652/54400.html",
	"https://www.manhuaniu.com/manhua/2652/54401.html",
	"https://www.manhuaniu.com/manhua/2652/54402.html",
	"https://www.manhuaniu.com/manhua/2652/54403.html",
	"https://www.manhuaniu.com/manhua/2652/54404.html",
	"https://www.manhuaniu.com/manhua/2652/54405.html",
	"https://www.manhuaniu.com/manhua/2652/54406.html",
	"https://www.manhuaniu.com/manhua/2652/54407.html",
	"https://www.manhuaniu.com/manhua/2652/54408.html",
	"https://www.manhuaniu.com/manhua/2652/54409.html",
	"https://www.manhuaniu.com/manhua/2652/54410.html",
	"https://www.manhuaniu.com/manhua/2652/54411.html",
	"https://www.manhuaniu.com/manhua/2652/54412.html",
	"https://www.manhuaniu.com/manhua/2652/54413.html",
	"https://www.manhuaniu.com/manhua/2652/54414.html",
	"https://www.manhuaniu.com/manhua/2652/54415.html",
	"https://www.manhuaniu.com/manhua/2652/54416.html",
	"https://www.manhuaniu.com/manhua/2652/54417.html",
	"https://www.manhuaniu.com/manhua/2652/54418.html",
	"https://www.manhuaniu.com/manhua/2652/54419.html",
	"https://www.manhuaniu.com/manhua/2652/54420.html",
	"https://www.manhuaniu.com/manhua/2652/54421.html",
	"https://www.manhuaniu.com/manhua/2652/54422.html",
	"https://www.manhuaniu.com/manhua/2652/54423.html",
	"https://www.manhuaniu.com/manhua/2652/54424.html",
	"https://www.manhuaniu.com/manhua/2652/54425.html",
	"https://www.manhuaniu.com/manhua/2652/54426.html",
	"https://www.manhuaniu.com/manhua/2652/54427.html",
	"https://www.manhuaniu.com/manhua/2652/54428.html",
	"https://www.manhuaniu.com/manhua/2652/54429.html",
	"https://www.manhuaniu.com/manhua/2652/54430.html",
	"https://www.manhuaniu.com/manhua/2652/54431.html",
	"https://www.manhuaniu.com/manhua/2652/54432.html",
	"https://www.manhuaniu.com/manhua/2652/54433.html",
	"https://www.manhuaniu.com/manhua/2652/54434.html",
	"https://www.manhuaniu.com/manhua/2652/54435.html",
	"https://www.manhuaniu.com/manhua/2652/54436.html",
	"https://www.manhuaniu.com/manhua/2652/54437.html",
	"https://www.manhuaniu.com/manhua/2652/54438.html",
	"https://www.manhuaniu.com/manhua/2652/54439.html",
	"https://www.manhuaniu.com/manhua/2652/54440.html",
	"https://www.manhuaniu.com/manhua/2652/54441.html",
	"https://www.manhuaniu.com/manhua/2652/54442.html",
	"https://www.manhuaniu.com/manhua/2652/54443.html",
	"https://www.manhuaniu.com/manhua/2652/54444.html",
	"https://www.manhuaniu.com/manhua/2652/54445.html",
	"https://www.manhuaniu.com/manhua/2652/54446.html",
	"https://www.manhuaniu.com/manhua/2652/54447.html",
	"https://www.manhuaniu.com/manhua/2652/54448.html",
	"https://www.manhuaniu.com/manhua/2652/54449.html",
	"https://www.manhuaniu.com/manhua/2652/54450.html",
	"https://www.manhuaniu.com/manhua/2652/54451.html",
	"https://www.manhuaniu.com/manhua/2652/54452.html",
	"https://www.manhuaniu.com/manhua/2652/54453.html",
	"https://www.manhuaniu.com/manhua/2652/54454.html",
	"https://www.manhuaniu.com/manhua/2652/54455.html",
	"https://www.manhuaniu.com/manhua/2652/54456.html",
	"https://www.manhuaniu.com/manhua/2652/54457.html",
	"https://www.manhuaniu.com/manhua/2652/54458.html",
	"https://www.manhuaniu.com/manhua/2652/54459.html",
	"https://www.manhuaniu.com/manhua/2652/54460.html",
	"https://www.manhuaniu.com/manhua/2652/54461.html",
	"https://www.manhuaniu.com/manhua/2652/54462.html",
	"https://www.manhuaniu.com/manhua/2652/54463.html",
	"https://www.manhuaniu.com/manhua/2652/54464.html",
	"https://www.manhuaniu.com/manhua/2652/222493.html",
	"https://www.manhuaniu.com/manhua/2652/339375.html",
	"https://www.manhuaniu.com/manhua/2652/339376.html",
	"https://www.manhuaniu.com/manhua/2652/339377.html",
	"https://www.manhuaniu.com/manhua/2652/339378.html",
	"https://www.manhuaniu.com/manhua/2652/339379.html",
	"https://www.manhuaniu.com/manhua/2652/339380.html",
	"https://www.manhuaniu.com/manhua/2652/339381.html",
	"https://www.manhuaniu.com/manhua/2652/339382.html",
	"https://www.manhuaniu.com/manhua/2652/339383.html",
	"https://www.manhuaniu.com/manhua/2652/339384.html",
	"https://www.manhuaniu.com/manhua/2652/339385.html",
	"https://www.manhuaniu.com/manhua/2652/339386.html",
	"https://www.manhuaniu.com/manhua/2652/339387.html",
	"https://www.manhuaniu.com/manhua/2652/339388.html",
	"https://www.manhuaniu.com/manhua/2652/339389.html",
	"https://www.manhuaniu.com/manhua/2652/339390.html",
	"https://www.manhuaniu.com/manhua/2652/339391.html",
	"https://www.manhuaniu.com/manhua/2652/339392.html",
	"https://www.manhuaniu.com/manhua/2652/339393.html",
	"https://www.manhuaniu.com/manhua/2652/339394.html",
	"https://www.manhuaniu.com/manhua/2652/339395.html",
	"https://www.manhuaniu.com/manhua/2652/339396.html",
	"https://www.manhuaniu.com/manhua/2652/339397.html",
	"https://www.manhuaniu.com/manhua/2652/339398.html",
	"https://www.manhuaniu.com/manhua/2652/339399.html",
	"https://www.manhuaniu.com/manhua/2652/339400.html",
	"https://www.manhuaniu.com/manhua/2652/339401.html",
	"https://www.manhuaniu.com/manhua/2652/339402.html",
	"https://www.manhuaniu.com/manhua/2652/339403.html",
	"https://www.manhuaniu.com/manhua/2652/339404.html",
	"https://www.manhuaniu.com/manhua/2652/339405.html",
	"https://www.manhuaniu.com/manhua/2652/339406.html",
	"https://www.manhuaniu.com/manhua/2652/339407.html",
	"https://www.manhuaniu.com/manhua/2652/339408.html",
	"https://www.manhuaniu.com/manhua/2652/339409.html",
	"https://www.manhuaniu.com/manhua/2652/339410.html",
	"https://www.manhuaniu.com/manhua/2652/341814.html",
	"https://www.manhuaniu.com/manhua/2652/344360.html",
	"https://www.manhuaniu.com/manhua/2652/347370.html",
	"https://www.manhuaniu.com/manhua/2652/349580.html",
	"https://www.manhuaniu.com/manhua/2652/353146.html",
	"https://www.manhuaniu.com/manhua/2652/356889.html",
	"https://www.manhuaniu.com/manhua/2652/359221.html",
	"https://www.manhuaniu.com/manhua/2652/361688.html",
	"https://www.manhuaniu.com/manhua/2652/363855.html",
	"https://www.manhuaniu.com/manhua/2652/365858.html",
	"https://www.manhuaniu.com/manhua/2652/368232.html",
	"https://www.manhuaniu.com/manhua/2652/371589.html",
	"https://www.manhuaniu.com/manhua/2652/375341.html",
	"https://www.manhuaniu.com/manhua/2652/383379.html",
	"https://www.manhuaniu.com/manhua/2652/383380.html",
	"https://www.manhuaniu.com/manhua/2652/383381.html",
	"https://www.manhuaniu.com/manhua/2652/385879.html",
	"https://www.manhuaniu.com/manhua/2652/388096.html",
}
