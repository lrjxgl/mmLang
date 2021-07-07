<template>
	<view>
		<form @submit="submit">
			<view class="input-flex">
				<view class="input-flex-label">支付密码</view>
				<input password="password" class="input-flex-text" name="paypwd">
			</view>
			<view class="input-flex">
				<view class="input-flex-label">手机</view>
				<div class="input-flex-txt" >{{telephone}}</div>
			</view>

			<view class="input-flex">
				<view class="input-flex-label">验证码</view>
				<input type="text" class="input-flex-text" name="yzm" v-model="yzm">
				<view class="input-flex-btn" v-bind:class="yzmClass" @click="getYzm()">{{yzmStatus}}</view>
			</view>

			<button form-type="submit" class="btn-row-submit">保存</button>
		</form>
	</view>
</template>

<script>
	var yzmTimer=60,yzmTime=59,yzmEnable=true;
	export default {
		data: function() {
			return {
				pageLoad: false,
				yzm: "",
				telephone: "",
				yzmClass: "",
				yzmStatus: "获取验证码",
			}
		},
		onLoad: function(option) {
			uni.setNavigationBarTitle({
				title: "支付密码"
			})
			this.getPage();
		},

		methods: {
			downTimer: function() {
				var that = this;
				var it = setInterval(function() {
					yzmEnable = false;
					that.yzmStatus = "倒计时" + yzmTime + "秒";
					that.yzmClass = "yzmDisable";
					yzmTime--;
					if (yzmTime == 0) {
						yzmTime = 59;
						yzmEnable = true;
						that.yzmClass = "";
						that.yzmStatus = "获取验证码";
						clearInterval(it);
					}
				}, 1000);
			},
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/user/info",
					 
					success:function(res){
						that.telephone=res.user.telephone;
					}
				})
			},
			getYzm: function() {
				if (!yzmEnable) return false;
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/user/sendsms",
					success: function(res) {
						uni.showToast({
							title: res.message,
						})
						if (!res.error) {
							that.downTimer();
						}

					}
				})
			},
			submit: function(e) {
				var that = this;
				that.app.post({
					url: this.app.apiHost + "/user/paypwdsave",
					data: e.detail.value,
					success: function(res) {
						uni.showToast({
							"title": res.message
						})
						if (!res.error) {
							setTimeout(function() {
								uni.navigateBack()
							}, 600)

						}

					}
				})
			}
		},
	}
</script>

<style>
</style>
