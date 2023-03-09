
// console.log(t);
var mydate = new Date();
udsfesr = localStorage.getItem('udsfesr');
var t=parseInt(mydate.getTime()/1000);

if(!udsfesr){
	window.location.href = './login.html';
	console.log(1);
}
// !@#@#$
arrlocalStorage = udsfesr.split('!@#@#$');
if( t- arrlocalStorage[1]>1469){
	 localStorage.removeItem('udsfesr');
	window.location.href = './login.html';
}
// 更新无操作时间
localStorage.setItem('udsfesr',arrlocalStorage[0]+"!@#@#$"+t);



var headHtml ="";
headHtml +='<header class="am-topbar am-topbar-inverse admin-header">';
headHtml +='<style type="text/css" media="screen">';
headHtml +='.am-topbar-inverse{background: #5eb95e; border: 0;}';
headHtml +='</style>';
headHtml +='<div class="am-topbar-brand">';
headHtml +='<strong style="font-size: 2.4rem; ">嘻嘻</strong> &nbsp &nbsp&nbsp <small>blog管理</small>';
headHtml +='</div>';
headHtml +='<button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only" data-am-collapse="{target: '+"#topbar-collapse"+'}"><span class="am-sr-only">导航切换</span> <span class="am-icon-bars"></span></button>';
headHtml +='<div class="am-collapse am-topbar-collapse" id="topbar-collapse">';
headHtml +='<ul class="am-nav am-nav-pills am-topbar-nav am-topbar-right admin-header-list">';
headHtml +='<li class="am-dropdown" data-am-dropdown>';
headHtml +='<a class="am-dropdown-toggle" data-am-dropdown-toggle href="javascript:;">';
headHtml +='<span class="am-icon-users"></span> 管理员 <span class="am-icon-caret-down"></span>';
headHtml +='</a>';
headHtml +='<ul class="am-dropdown-content">';
headHtml +='<li><a href="#"><span class="am-icon-user"></span> 资料</a></li>';
headHtml +='<li><a href="#"><span class="am-icon-cog"></span> 设置</a></li>';
headHtml +='<li><a href="#" class="outLogin"><span class="am-icon-power-off " ></span> 退出</a></li>';
headHtml +='</ul>';
headHtml +='</li>';
headHtml +='</ul>';
headHtml +='</div>';
headHtml +='</header>';

$(".header").html(headHtml);

$(".outLogin").bind("click");

$(".outLogin").click(function(){
	 localStorage.removeItem('udsfesr');
	window.location.href = './login.html';
});

var navList = "";
navList += '<div class="admin-sidebar am-offcanvas" id="admin-offcanvas">';
navList += '<div class="am-offcanvas-bar admin-offcanvas-bar">';
navList += '<ul class="am-list admin-sidebar-list">';
navList += '<li><a href="index.html"><span class="am-icon-home"></span> 首页</a></li>';
navList += '<li class="admin-parent">';
navList += '<a class="am-cf" data-am-collapse="{target: '+"#collapse-nav"+'}"><span class="am-icon-file"></span> 日记 & 笔记 <span class="am-icon-angle-right am-fr am-margin-right"></span></a>';
navList += '<ul class="am-list am-collapse admin-sidebar-sub am-in" id="collapse-nav">';
navList += '<li><a href="./article.html" class="am-cf"><span class="am-icon-check"></span> 文章 <span class="am-icon-star am-fr am-margin-right admin-icon-yellow"></span></a></li>';
navList += '<li><a href="./classify.html"><span class="am-icon-puzzle-piece"></span> 类别</a></li>';
navList += '</ul>';
navList += '</li>';
navList += '</ul>';

navList += '</div>';
navList += '';
navList += '<div class="am-panel am-panel-default admin-sidebar-panel">';
navList += '<div class="am-panel-bd">';
navList += '<p><span class="am-icon-tag"></span> wiki</p>';
navList += '<p>Welcome to golang int gin</p>';
navList += '</div>';
navList += '</div>';
navList += '</div>';
navList += '</div>';
$(".navList").html(navList);
var pageSize = 20;
// 点击分页

function pageFunc(total,clickPage, page=0){
	var html = "";
	if(total == 0 ){
		return ;
	}
	var html = "";
	pageSum = Math.ceil(total/pageSize);

	if(pageSum > 1){
		if(page != 0){
			html += "<li class='am-disabled'><a href='javascript:;' onclick='clickPage(0)' >«</a></li>";
		}

		if(pageSum<7){
			for (var i = 1; i <= pageSum; i++) {
					console.log(i)
				html += "<li><a href='javascript:;' onclick='clickPage("+(i-1)+")' >"+i+"</a></li>";
			}
		}

		if(pageSum >= 7 && page>=3){
			for (var D = page-2; D <= page; D++) {
				html += "<li><a href='javascript:;' onclick='clickPage("+(D-1)+")' >"+D+"</a></li>";
			}

			if(pageSum >= page+3 ){
				for (var F = page+1; F <= page+3; F++) {
					html += "<li><a href='javascript:;' onclick='clickPage("+(F-1)+")' >"+F+"</a></li>";
				}
			}

		}
		if(pageSum >= 7 && page<3){
			for (var S = 0; S < page; S++) {
				html += "<li><a href='javascript:;' onclick='clickPage("+(S-1)+")' >"+S+"</a></li>";
			}

			if(pageSum >= page+3 ){
				for (var F = page+1; F <= page+3; F++) {
					html += "<li><a href='javascript:;' onclick='clickPage("+(F-1)+")' >"+F+"</a></li>";
				}
			}

		}

		if(pageSum>2){
			html += "<li><a href='javascript:;' onclick='clickPage("+(pageSum-1)+")' >»</a></li>";
		}

	}
	console.log(pageSum)
	return html;
}
