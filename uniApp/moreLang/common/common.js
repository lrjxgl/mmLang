module.exports = {
	//apiHost:"http://fd175.skymvc.com/",
	apiHost: "http://127.0.0.1:8080",
	appRoot: "https://www.fd175.com/uniapp/h5/",
	wsHost: "wss://wss.deituicms.com:8282",
	pmTimer: 0,
	inPost: false,
	 
	json_add: function(a, b) {

		if (a == undefined || a.length == 0) return b;
		if (b == undefined || b.length == 0) return a;

		var s_a = JSON.stringify(a);
		var s_b = JSON.stringify(b);
		var c = s_a.substring(0, s_a.length - 1) + "," + s_b.substring(1);
		return JSON.parse(c);
	},
	goBack: function() {
		console.log(Router);
		this.$router.go(-1);
	},
	getCookie: function(name) {
		// #ifdef H5
		var strcookie = document.cookie;
		var arrcookie = strcookie.split("; ");
		for (var i = 0; i < arrcookie.length; i++) {
			var arr = arrcookie[i].split("=");
			if (arr[0] == name) {
				return arr[1];
			}
		}
		return "";
		// #endif

	},
	getToken: function() {
		var authcode = uni.getStorageSync("token"); 
		return authcode;
	},
	setToken: function($authcode) {
		uni.setStorageSync("authcode", $authcode);
	},
	 
	setOpenid: function(openid) {
		uni.setStorageSync("openid", openid)
	},
	getOpenid: function(openid) {
		uni.getStorageSync("openid")
	},
	getLoginCode: function() {
		return uni.getStorageSync("loginCode");
	},
	setLoginCode: function(code) {
		uni.setStorageSync("loginCode", code)
	},
	getS2cScid: function() {
		return uni.getStorageSync("s2c_scid");
	},
	setS2cScid: function(v) {
		uni.setStorageSync("s2c_scid", v);
	},
	getGps: function() {
		var gps = uni.getStorageSync("appGps");
		if (gps == "") {
			return {
				lat: 39.850836,
				lng: 116.471564
			};
		}
		return JSON.parse(gps);

	},
	setGps: function(lat, lng) {
		var s = {
			lat: lat,
			lng: lng
		};
		var str = JSON.stringify(s);
		uni.setStorageSync("appGps", str);
	},
	setCity: function(city, cityid) {
		uni.setStorageSync("city", city);
		uni.setStorageSync("cityid", cityid);
	},
	getCity: function() {
		var city = uni.getStorageSync("city");
		if (!city) {
			city = "厦门";
		}
		return city;
	},

	getCityId: function() {
		return uni.setStorageSync("cityid", cityid);
	},
	getLocalAddress: function() {
		var city = uni.getStorageSync("appLocalAddress");
		if (!city) {
			city = "瑞景商业广场";
		}
		return city;
	},
	setLocalAddress: function(val) {
		uni.setStorageSync("appLocalAddress", val);
	},
	fromapp: function() {
		//var $paltform= uni.platform();
		return "uniapp";
	},
	isWeixin: function() {
		// #ifdef H5
		var ua = window.navigator.userAgent.toLowerCase();
		if (ua.match(/MicroMessenger/i) == 'micromessenger') {
			return true;
		}
		// #endif
		return false;
	},
	get: function(ops) {
		var ops = ops;
		if (ops.data == undefined) {
			ops.data = {
				token: this.getToken(),
				fromapp: this.fromapp(),
				ajax: 1
			};
		} else {
			ops.data.token = this.getToken();
			ops.data.fromapp = this.fromapp();
			ops.data.ajax = 1;
		}

		uni.request({
			url: ops.url,
			method: "GET",
			data: ops.data,
			success: (rs) => {

				if (rs.data.error == 1000 && ops.unLogin == undefined) {
					uni.navigateTo({
						url: "../../pages/login/index"
					})
				} else {
					ops.success(rs.data);
				}
			}
		})
	},
	post: function(ops) {
		var callback = callback;
		var ops = ops;
		if (ops.url.indexOf("?") >= 0) {
			ops.url += "&ajax=1&token=" + this.getToken() + "&fromapp=" + this.fromapp();
		} else {
			ops.url += "?ajax=1&token=" + this.getToken() + "&fromapp=" + this.fromapp();
		}

		uni.request({
			url: ops.url,
			data: ops.data,
			method: "POST",
			header: {
				"content-type": "application/x-www-form-urlencoded"
			},
			success: (rs) => {
				if (rs.data.error == 1000 && ops.unLogin == undefined) {
					uni.navigateTo({
						url: "../../pages/login/index"
					})
				} else {
					ops.success(rs.data);
				}
			}
		})
	},

	canPost: function() {
		var that = this;
		if (!this.inPost) {
			this.inPost = true;

			setTimeout(function() {
				that.inPost = false;
			}, 1000)
			return true;
		} else {
			setTimeout(function() {
				that.inPost = false;
			}, 1000)
			return false;
		}

	},
	goHome: function() {
		uni.reLaunch({
			url: "../../pages/index/index"
		})
		/*
		uni.switchTab({
			url:"../index/index",
		})
		*/
	},

	goUser: function() {
		uni.switchTab({
			url: "../user/index",
		})
	},

	goLogin: function() {
		uni.navigateTo({
			url: "../../pages/login/index"
		})
	},
	goH5WeiXin: function(backurl) {
		// #ifdef H5
		var url = document.location.href.split('#')[0] + "#" + backurl;
		window.location = this.apiHost + '/index.php?m=open_weixin&a=Geturl&backurl=' + encodeURIComponent(url);
		// #endif

	},
	html: function(html) {
		if (html == '' || html == null) return '';
		html = html.replace(' style="', '  xstyle="');
		html = html.replace(/<img /g, '<img style="max-width:100%;height:auto;"');
		html = html.replace(/&hellip;/g, '');
		html = html.replace(' class="brush:html;toolbar:false"', " ");
		html = html.replace(/<pre /g,
			'<pre style="border:1px solid #eee;padding:10px;max-height:300px;overflow:auto;" ');

		html = html.replace(/<pre /g, "<view ");
		html = html.replace("</pre>", "</view>");
		html = html.replace(/\n/g, "<br />");
		html = html.replace(/<iframe /, '<view style="display:none" ');
		html = html.replace(/<\/iframe>/, '</view>');
		html = html.replace("<section ", "<view ");
		html = html.replace("</section>", "</view>");

		//console.log(html);
		return html;
	}


}