<template>
	<view>
		<view class="footer-row"></view>
		<view class="footer">
			<view class="footer-item icon-news" v-bind:class="{'footer-active':tab=='home'}" @click="goHome()">首页</view>
			 
			<view class="footer-item icon-message_light" v-bind:class="{'footer-active':tab=='pm'}" @click="goPm()">私信<text v-if="pm_num>0" class="badge  badge-abs">{{pm_num}}</text></view> 
			<view class="footer-item icon-my_light" v-bind:class="{'footer-active':tab=='user'}"  @click="goUser()">我的</view>
		</view>
	</view>
</template>

<script>
	 
	export default{
		props:{
			tab:""
		},
		data:function(){
			return {
				pm_num:0
			}
		},
		created:function(){
			var that=this;
			//this.getPm();
			if(this.app.timer){
				//clearInterval(this.app.timer)
			}
			this.app.timer=setInterval(function(){
				//that.getPm();
			},10000)
			
		},
		
		methods:{
			getPm:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/index.php?m=pm&a=getnewindex&ajax=1",
					unLogin:true,
					success:function(res){
						that.pm_num=res.num;
					}
				})	
			},
			goHome:function(){
				uni.reLaunch({
					url:"../../pages/index/index"
				})
			},
			goPeople:function(){
				uni.reLaunch({
					url:"../../pages/sblog_people/index"
				})
			},
			goChat:function(){
				uni.navigateTo({
					url:"../../pagesblog/sblog_chat/index"
				})
			},
			goUser:function(){
				uni.navigateTo({
					url:"../../pages/user/index"
				})
			},
			goPm:function(){
				uni.reLaunch({
					url:"../../pages/pm/index"
				})
			},
			goMap:function(){
				uni.reLaunch({
					url:"../../pageshopmap/shopmap/index"
				})
			}
		}
	}
</script>

 
