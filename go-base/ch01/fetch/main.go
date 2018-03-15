// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// output:

// ChinaDreams:fetch kangcunhua$ ./fetch http://www.git.com.cn
// <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
// <html xmlns="http://www.w3.org/1999/xhtml">
// <head>
// <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
// <title>高伟达</title>
// <link href="templates/css/style.css" rel="stylesheet" type="text/css" />
// <link href="templates/css/font.css" rel="stylesheet" type="text/css" />
// <script type="text/javascript" src="templates/js/jquery.js"></script>
// <script type="text/javascript" src="templates/js/xl.js"></script>
// <script type="text/javascript" src="templates/js/changeimg.js"></script>
// <script src="templates/js/jq_scroll.js" type="text/javascript"></script>
// <script type="text/javascript">
// $(document).ready(function(){
//         $("#scrollDiv").Scroll({line:1,speed:500,timer:3000,up:"but_up",down:"but_down"});
// });
// </script>
// <style type="text/css">
// #scrollDiv{width:590px;height:20px; overflow:hidden;}/*这里的高度和超出隐藏是必须的*/
// #scrollDiv li{height:20px; vertical-align:bottom; zoom:1;}

// </style>
// </head>

// <body>
// <div class="topleft">
//   <div class="topleft-01" style="background-image:url(templates/images/topbj.png);">
//         <div class="zo">
//           <div class="top">
//                 <div class="logo"><a href="index.php"><img src="templates/images/logo.jpg" width="158" height="66" border="0" alt="" /></a></div>
//                 <div class="dh">
//                     <div class="dh-01">
//                         <div class="dh-02"><a href="index.php" class="bjbj2">首页</a></div>
//                         <div class="dh-02 content"><a href="index.php?m=content&amp;id=400
//                               " class="bjbj3">关于高伟达</a>
//                             <ul class="subMenuul noneul">
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=content&amp;id=400
//                               ">公司简介</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=history&amp;id=440&amp;aid=537
// ">发展历程</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=honer&amp;id=409&amp;aid=410
//                               ">资质荣誉</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=honer&amp;id=491&amp;aid=492
//                               ">合作伙伴</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=download&amp;id=405&amp;aid=406
// ">资料下载</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=content&amp;id=413
//                               ">联系我们</a></li>
//                                                         </ul>
//                         </div>
//                         <div class="dh-02 content"><a h1ref="index.php?m=business&amp;id=414" class="bjbj3">我们的业务</a>
//                             <ul class="subMenuul noneul">

//                                                                                     <li class="subMenuli linkxl">
//                                                           <a href="index.php?m=products&oneid=415&id=416">产品及解决方案</a>

//                                                           </li>
//                                                                    <li class="subMenuli linkxl">
//                                                           <a href="index.php?m=products&oneid=415&id=417">云计算数据中心</a>

//                                                           </li>
//                                                                    <li class="subMenuli linkxl">
//                                                           <a href="index.php?m=products&oneid=415&id=442">服务</a>

//                                                           </li>
//                                                                    <li class="subMenuli linkxl">
//                                                           <a href="index.php?m=products&oneid=415&id=446">咨询</a>

//                                                           </li>

//                                     <li class="subMenuli linkxl">
//                                                           <a href="index.php?m=operation&oneid=418&id=419">合作客户</a>

//                                                           </li>

//                                             <li class="subMenuli linkxl">
//                                                           <a href="index.php?m=cgan&oneid=418&id=490">成功案例</a>

//                                                           </li>

//                                                     </ul>
//                         </div>
//                         <div class="dh-02 content"><a href="index.php?m=news&amp;id=423" class="bjbj3">新闻中心</a>
//                             <ul class="subMenuul noneul">
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=news&amp;id=423">公司新闻</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=news&amp;id=424">媒体之声</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=news&amp;id=476">高伟达视点</a></li>
//                                                         </ul>
//                         </div>
//                         <div class="dh-02 content"><a href="index.php?m=viewpoint&lmid=478&id=425" class="bjbj3">投资者关系</a>

//                                                  <ul class="subMenuul noneul">
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=viewpoint&amp;lmid=478&id=425">股价信息</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=viewpoint&amp;lmid=479&id=425">信息披露</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=viewpoint&amp;lmid=480&id=425">公司治理</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=viewpoint&amp;lmid=481&id=425">股本信息</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=viewpoint&amp;lmid=482&id=425">公司章程</a></li>
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=viewpoint&amp;lmid=483&id=425">投资者交流</a></li>
//                                                         </ul>

//                         </div>
//                         <div class="dh-02 content"><a href="index.php?m=joins
//                               " class="bjbj3">人才招聘</a>
//                             <ul class="subMenuul noneul">
//                                                           <li class="subMenuli linkxl"><a href="index.php?m=joins
//                               ">加入高伟达</a></li>
//                                                           <li class="subMenuli linkxl"><a href="     index.php?m=joincontent&amp;oneid=430&amp;id=432
//                               ">成长通道</a></li>
//                                                           <li class="subMenuli linkxl"><a href="     index.php?m=joincontent&amp;oneid=433&amp;id=434
//                               ">人在高伟达</a></li>
//                                                         </ul>
//                         </div>
//                     </div>
//                 </div>
//             <div class="zy">
//                     <div class="zy-01">
//                     <table width="108" height="20" border="0" cellspacing="0" cellpadding="0">
//                       <tr align="right">
//                         <td width="63" class="linkzw"><a href="/"><strong>中文</strong></a></td>
//                         <td width="12" class="en-x"><strong>|</strong></td>
//                         <td width="33" class="linken"><a href="/"><strong>En</strong></a></td>
//                       </tr>
//                     </table>
//               </div>
//             </div>
//           </div>
//                 </div>
//   </div>
// </div>

// <div id="bigbanner" style="margin-top:81px;">
//   <div id="banners">
//         <div class="currents" id="banner-1" style="width:100%;height:474px;background-image:url(uploadfiles/1418807488.jpg); background-repeat:no-repeat; margin:0 auto; background-position:top center;">
//         <div class="zt">
//                 <div class="zt1">
//                         <img src="templates/images/ztbj.png" width="980" height="474" border="0" alt="" />
//                         </div>
//         </div>
//       </div>
//         <div class="banner" id="banner-2" style="width:100%;height:474px;background-image:url(uploadfiles/1418263402.jpg); background-repeat:no-repeat; margin:0 auto; background-position:top center;">
//         <div class="zt">
//                 <div class="zt1">
//                         <img src="templates/images/ztbj.png" width="980" height="474" border="0" alt="" />
//                         </div>
//         </div>
//       </div>
//         <div class="banner" id="banner-3" style="width:100%;height:474px;background-image:url(uploadfiles/1418263390.jpg); background-repeat:no-repeat; margin:0 auto; background-position:top center;">
//         <div class="zt">
//                 <div class="zt1">
//                         <img src="templates/images/ztbj.png" width="980" height="474" border="0" alt="" />
//                         </div>
//         </div>
//       </div>

//       <ul id="controls">
//                   <li class="active"><a href="#" rel="banner-1"></a></li>
//                   <li ><a href="#" rel="banner-2"></a></li>
//                   <li ><a href="#" rel="banner-3"></a></li>

//         </ul>
//     </div>
// </div>

// <div class="index">
//     <div class="index-01">
//         <div class="index-02">
//             <div class="index-03">
//                   <div class="index-04 bai-z"><strong>最新动态</strong></div>
//               <div class="index-05">
//               <div id="scrollDiv">
//                                 <ul class="noneul">
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=717" target="_blank">2016金融云探索与创新金秋研讨会圆满落幕</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=638" target="_blank">坚持高品质，高伟达实力荣膺2项“云”大奖</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=637" target="_blank">基于押品产品的落地实施——华润银行押品管理项目报道</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=636" target="_blank">我们的创新业务落地实施——农商行资产证券化项目报道</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=635" target="_blank">高伟达：以持续创新，从金融互联网迈向互联网金融</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=558" target="_blank">高伟达上市一周年 成为中国金融业信息化市场的中坚力量</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=557" target="_blank">电子商业汇票业务在金融中的运用培训讲座于北京二十一世纪酒店顺利举行</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=556" target="_blank">“把爱传递下去”高伟达收到来自多所乡村学校发来的感谢信</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=555" target="_blank">公司喜讯连连 陆续收到多封来自客户的表扬信</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=553" target="_blank">高伟达公司荣获“中国方案商百强、2015中国十大金融行业方案商、行业云应用开发商”称号</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=552" target="_blank">合作成就共赢 专业铸就梦想</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=526" target="_blank">高伟达软件股份有限公司上市成功</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=429" target="_blank">高伟达公司成为“中国电子信息化行业联合会”会员单位</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=394" target="_blank">高伟达成为中国软件行业协会系统与软件过程改进分会副主席单位</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=393" target="_blank">葫芦岛银行新核心业务系统全新上线</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=392" target="_blank">高伟达公司顺利通过IS09000质量管理体系监督审核</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=391" target="_blank">建设银行IT服务管理云平台项目当选2013年中国金融信息化十件大事</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=390" target="_blank">高伟达顺利通过CMMI3级复评估</a></li>
//                                                         <li class="linkgg"><a href="index.php?m=newscon&amp;id=423&amp;aid=389" target="_blank">高伟达中标北京银行零售押品业务咨询及系统实施项目</a></li>

//                     </ul>
//               </div>
//               </div>
//               <div class="index-06">
//               <form name="searchform" id="searchform" action="index.php?m=search" method="post">
//                 <div class="index-07"><input name="kwords" id="kwords" type="text" class="wby" value="输入您要搜索的内容" onFocus="this.value=''" onblur='if (value ==""){value="输入您要搜索的内容"}' /></div>
//                 <div class="index-08"><input name="" type="image" src="templates/images/fdj.jpg" /></div>
//               </form>
//               </div>
//             </div>
//         </div>
//         <div class="index-09">
//             <div class="index-10">
//                 <div class="index-11">
//                     <div class="index-12 xw-z linkindex"><a href="index.php?m=news&amp;id=423" target="_blank"><strong>新闻中心</strong></a>&nbsp;<strong class="en-z">/ Media Room</strong></div>
//                     <div class="index-13"><a href="index.php?m=news&amp;id=423" target="_blank"><img src="templates/images/index.jpg" width="35" height="17" border="0" alt="" /></a></div>
//                 </div>
//               <div class="index-14">
//                 <div class="index-15">
//                     <div class="index-16"><a href="index.php?m=newscon&amp;id=423&amp;aid=747" target="_blank"><img src="uploadfiles/1513238814.jpg" width="121" height="77" border="0" alt="" /></a></div>
//                     <div class="index-17">
//                       <div class="index-18 linkred"><a href="index.php?m=newscon&amp;id=423&amp;aid=747" target="_blank">高伟达中标安徽省农信产品创新服务平台项目！</a></div>
//                         <div class="index-19 date-z">2017-9-30</div>
//                     </div>
//                 </div>
//                 <div class="index-20"></div>
//                                 <div class="index-21">
//                   <div class="index-22 linkn"><img src="templates/images/index-05.jpg" width="5" height="5" border="0" alt="" style="margin-bottom:3px;" />&nbsp;<a href="index.php?m=newscon&amp;id=423&amp;aid=746" target="_blank">祝贺宁夏银行押品管理项目...</a></div>
//                   <div class="index-23 date-z2">2017-9-29</div>
//                 </div>
//                                                 <div class="index-21">
//                   <div class="index-22 linkn"><img src="templates/images/index-05.jpg" width="5" height="5" border="0" alt="" style="margin-bottom:3px;" />&nbsp;<a href="index.php?m=newscon&amp;id=423&amp;aid=745" target="_blank">高伟达携ECIF系统出席...</a></div>
//                   <div class="index-23 date-z2">2017-9-22</div>
//                 </div>
//                                                 <div class="index-21">
//                   <div class="index-22 linkn"><img src="templates/images/index-05.jpg" width="5" height="5" border="0" alt="" style="margin-bottom:3px;" />&nbsp;<a href="index.php?m=newscon&amp;id=423&amp;aid=744" target="_blank">开拓创新、锐意进取--中...</a></div>
//                   <div class="index-23 date-z2">2017-9-14</div>
//                 </div>
//                                                 <div class="index-21">
//                   <div class="index-22 linkn"><img src="templates/images/index-05.jpg" width="5" height="5" border="0" alt="" style="margin-bottom:3px;" />&nbsp;<a href="index.php?m=newscon&amp;id=423&amp;aid=743" target="_blank">云领未来，高伟达亮相华为...</a></div>
//                   <div class="index-23 date-z2">2017-9-8</div>
//                 </div>
//                                               </div>
//             </div>
//             <div class="index-10">
//                 <div class="index-11">
//                     <div class="index-12 xw-z linkindex"><a href="index.php?m=products&oneid=415&id=416" target="_blank"><strong>业务中心</strong></a>&nbsp;<strong class="en-z">/ Our  Business</strong></div>
//                 </div>
//                 <div class="index-24">
//                         <div class="index-25"><a href="index.php?m=products&oneid=415&id=416" target="_blank"><img src="uploadfiles/1418268744.jpg" width="157" height="99" border="0" alt="" /></a></div>
//                     <div class="index-26">
//                           <div class="index-27 linkn"><a href="index.php?m=products&oneid=415&id=416" target="_blank">高伟达经过多年的开发与积累，面向金融行业提供的完整、...</a><br><a href="index.php?m=business&amp;id=414" target="_blank"><img src="templates/images/more.jpg" width="49" height="12" border="0" alt="" style="margin-top:6px;" /></a></div>
//                     </div>
//                 </div>
//               <div class="index-28">
//                         <div class="index-29">
//                                                 <div class="index-30"><a href="index.php?m=products&oneid=415&id=417" target="_blank"><img src="uploadfiles/1418280017.jpg" width="143" height="99" border="0" alt="" /></a>
//                                 <div class="index-31 linkxl"><a href="http://www.git.com.cn/index.php?m=procon&oneid=415&id=417&twoid=330&aid=136" target="_blank"><strong>云管理平台</strong></a></div>
//                         </div>
//                                             </div>
//                 <div class="index-29" style="margin-right:0;">

//                                                         <div class="index-30"><a href="index.php?m=products&oneid=415&id=417" target="_blank"><img src="uploadfiles/1418280034.jpg" width="143"height="99" border="0" alt="" /></a>
//                                 <div class="index-31 linkxl"><a href="index.php?m=products&oneid=415&id=416" target="_blank"><strong>新一代核心系统</strong></a></div>
//                         </div>
//                                             </div>
//               </div>
//             </div>
//           <div class="index-10" style="margin-right:0;">
//                 <div class="index-11">
//                     <div class="index-12 xw-z linkindex"><a href="index.php?m=history&amp;id=440&amp;aid=537" target="_blank"><strong>历史足迹</strong></a>&nbsp;<strong class="en-z">/ HistoryTrail</strong></div>
//                 </div>
//                 <div class="index-32"><a href="index.php?m=history&amp;id=440&amp;aid=537" target="_blank"><img src="uploadfiles/1418869832.jpg" width="299" height="135" border="0" alt="" /></a></div>
//                         <div class="index-33 linkcg"><a href="index.php?m=history&amp;id=440&amp;aid=537" target="_blank">高伟达软件股份有限公司————发展历程</a></div>
//                 <div class="index-34"><a href="index.php?m=history&amp;id=440&amp;aid=537" target="_blank"><img src="templates/images/index-08.jpg" width="299" height="16" border="0" alt="" /></a></div>
//           </div>
//       </div>
//     </div>
// </div>

// <div class="bottom">
//         <div class="bottom-01">
//         <div class="bottom-02">
//           <div class="bottom-03 linkbot bot-z"><a href="index.php?m=content&amp;id=413" target="_blank">联系我们</a>&nbsp;&nbsp;|&nbsp;&nbsp;<a href="index.php?m=maps" target="_blank">网站地图</a>&nbsp;&nbsp;|&nbsp;&nbsp;<a href="index.php?m=content&amp;id=436" target="_blank">法律声明</a></div>
//           <div class="bottom-04 bot-z">版权所有 © 高伟达  京ICP备 15004154 号</div>
//       </div>
//     </div>
// </div>

// </body>
