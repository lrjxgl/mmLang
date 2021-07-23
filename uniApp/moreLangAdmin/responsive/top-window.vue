<template>
	<view>
		<div class="">
			<div class="topa">
				 
				<div @click="goBack()" class="top-item js-back" title="后退">
					<i class="iconfont icon-back"></i>
				</div>

				<div @click="goHome()" class="top-item-btn" title="后台主页">
					<i class="iconfont icon-home"></i> 主后台
				</div>
				<div v-for="(item,index) in list" @click="goMenu(item.m)"  class="top-item-btn">
					<i class="iconfont "></i> {{item.title}}
				</div>
				<a   class="top-item-btn">
					<i class="iconfont icon-repair"></i> 插件扩展
				</a>
				<div class="flex-1"></div>

				<div @click="goRefresh()" class="top-item" title="刷新">
					<i class="iconfont icon-refresh"></i>
				</div>
				<a href="/" target="_blank" class="top-item" title="前台首页">
					<i class="iconfont icon-global"></i>
				</a>
				<div class="top-item" id="nickname">
					老雷

				</div>
				<div onclick="logout()" class="top-item" title="注销">
					<i class="iconfont icon-logout"></i>
				</div>
				<div class="top-item" title="更多">
					<i class="iconfont icon-moreandroid"></i>
				</div>

			</div>
		</div>
	</view>
</template>

<script>
	export default{
		data:function(){
			return {
				list:[]
				 
			}
		},
		created:function(){
			this.getPage();
			 
		},
		methods:{
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/admin/navbar/get?group_id=1",
					success:function(res){
						 
						that.list=res.list;
					}
				})
			},
			goBack:function(){
				console.log("goBack")
				uni.navigateBack();
			},
			goRefresh:function(){
				var url=this.$route.fullPath;
				uni.redirectTo({
					url:url
				})
			},
			goHome:function(){
				uni.$emit("leftWindowPage",{page:"main"})
				uni.reLaunch({
					url:"/pages/index/index"
				})
			},
			goMenu:function(page){
				console.log(page);
				uni.$emit("leftWindowPage",{page:page})
			}
		}
	}
</script>

<style>
	.topa {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		 
		border-bottom: 1px solid #eee;
		overflow: auto;
		 
		align-items: center;
	}

	.top-item {
		height: 50px;
		line-height: 50px;
		min-width: 50px;
		cursor: pointer;
		box-sizing: border-box;
		text-align: center;
		display: inline-block;
		text-decoration: none;
		position: relative;
	}

	.top-item:hover {
		border-bottom: 3px solid #ddd;
	}

	.top-item-btn {
		position: relative;

		height: 30px;
		line-height: 30px;
		padding: 0px 10px;
		border-radius: 10px;
		cursor: pointer;
		box-sizing: border-box;
		text-align: center;
		display: inline-block;
		text-decoration: none;
		color: #646464;
		border: 1px solid #ddd;
		margin-right: 10px;
		font-size: 14px;

	}

	.top-item-btn.active {
		color: #007aff;
	}

	.top-item .iconfont {
		caption-side: #323232;
		font-size: 16px;
	}

	.abs-menu {
		position: absolute;
		top: 45px;
		background-color: #fff;
		display: block;
		border: 1px solid #ddd;
		width: 100px;
		left: -30px;
		padding: 10px 0px;
	}

	.abs-menu-item {
		line-height: 36px;
		padding: 0px 10px;
		cursor: pointer;
	}
</style>
