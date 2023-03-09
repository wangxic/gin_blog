config = new Array();
config['baseUrl'] = 'http://127.0.0.1:861'
/**
 * 
 * @param {"权限类型" 1 = 需要权限；0=无需权限验证接口} as_toke  
 * @returns http 头
 */
function getHeaders(as_toke=1) {
	udsfesr = localStorage.getItem('udsfesr');
	arrlocalStorage = udsfesr.split('!@#@#$');
	var headers = {
			"Access-Control-Request-Header":"DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type, Accept-Language, Origin, Accept-Encoding,Token",
			"Access-Control-Allow-Method":"POST",
			"Access-Control-Allow-Origin":config['baseUrl'],
			'Token':arrlocalStorage[0],
			"Content-Type":"application/x-www-form-urlencoded;charset:utf-8;"
		};
	
	return headers;
}

// 封装的请求
function Request(url,data ='',type = "POST",as_toke=1){
		udsfesr = localStorage.getItem('udsfesr');
	arrlocalStorage = udsfesr.split('!@#@#$');
	var return_data = '';
	Header = getHeaders(as_toke);
	if(data!=''){
		$.ajax({
			async:false,
			type: type,
			url: config['baseUrl']+url,
			headers:Header, 
			data: data,
			success: function(data){
				return_data =data;
			},
			error:function(error){
				if(error['status']==0){
					return_data ={code:404};
				}else{
					return_data ={code:error['status']};
				}
			}
		});
	}else{
		$.ajax({
			async:false,
			type: type,
			url: config['baseUrl']+url,
			// headers:Header,       
			success: function(data){
				return_data =data;
			},
			error:function(error){
				if(error['status']==0){
					return_data ={code:404};
				}else{
					return_data ={code:error['status']};
				}
			}
		});
	}

 	return return_data;
}
let user_centent =  localStorage.getItem("user_centent");
var user_info='';
if( user_centent != undefined  && user_centent != "" && user_centent != 'undefined' ){
	user_info = JSON.parse(user_centent);
}
// 错误处理
function askRequest(url,data,type="POST"){
	let as_toke = 0;
	// if(route['url']!= undefined && route['url']!= "undefined"  && route['url']['type'] == 1){
	// 	as_toke = 1;
	// }
	// if(as_toke==1){
	// 	let user_centent =  localStorage.getItem("user_centent");
	// 	let user_info = JSON.parse(user_centent);
	// 	if(user_info==null || user_info=='' || user_info.token==2){
	// 		layerhint(config['error']['900001'],2);
	// 		setInterval(function(){
	// 			window.parent.parent.location.href="../login.html";
	// 		},2000);
	// 		return false;
	// 	}
	// }
	_data = Request(url,data,type,as_toke);
	// console.log(_data);
	if(_data['result'] == 90001){
		// localStorage.removeItem("user_centent");
		udsfesr = localStorage.removeItem('udsfesr');
		// layerhint("请重新登录",2);
		alert("登录失效,请重新登录");
		setInterval(function(){
			window.parent.parent.location.href="./login.html";
		},50);
		return false;
	}
	
	// if( _data['code'] != 0 && config['error'][_data['code']] != 'undefined'){
	// 	layerhint(config['error'][_data['code']],2);
	// }
	return _data;
}
// 弹出提示
function layerhint(centent,minute){
	layer.alert(centent, {
	  time: minute*1000
	  ,success: function(layero, index){
	    var timeNum = this.time/1000, setText = function(start){
	      layer.title((start ? timeNum : --timeNum) + ' 秒后关闭', index);
	    };
	    setText(!0);
	    this.timer = setInterval(setText, 1000);
	    if(timeNum <= 0) clearInterval(this.timer);
	  }
	  ,end: function(){
	    clearInterval(this.timer);
	  }
	});
}

function _import(url){
    var script = document.createElement('script');
    script.setAttribute('type','text/javascript');
    script.setAttribute('src',url);
    document.getElementsByTagName('head')[0].appendChild(script);
}


// 获取get传值
function getQueryString(name) {
	var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
	var reg_rewrite = new RegExp("(^|/)" + name + "/([^/]*)(/|$)", "i");
	var r = window.location.search.substr(1).match(reg);
	var q = window.location.pathname.substr(1).match(reg_rewrite);
	if(r != null){
		return unescape(r[2]);
	}else if(q != null){
		return unescape(q[2]);
	}else{
		return null;
	}
}
