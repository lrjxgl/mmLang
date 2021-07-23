<template>
	<view>
		
	</view>
</template>

<script>
	
	export default{
		data:function(){
			return {
				pageLoad:false,
				list:[],
				per_page:0,
				isFirst:true
			}
		},
		onLoad:function(){
			this.getPage();
		},
		onReachBottom:function(){
			this.getList();
		},
		onPullDownRefresh:function(){
			this.getPage();
			uni.stopPullDownRefresh();
		},
		onShareAppMessage:function(){
			
		},
		onShareTimeline:function(){
			
		},
		methods: {
			gourl:function(url){
				uni.navigateTo({
					url:url
				})
			},
			getPage:function() {
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/",
					success:function(res){
						that.pageLoad=true;
						that.list=res.list;
						that.per_page=res.per_page;
					}
				})
			},
			getList:function() {
				var that=this;
				if(that.per_page==0 && !that.isFirst){
					return false;
				}
				that.app.get({
					url:that.app.apiHost+"/",
					data:{
						per_page:that.per_page
					},
					success:function(res){						 
						that.per_page=res.per_page;
						if(that.isFirst){
							that.list=res.list;
							that.isFirst=false;
						}else{
							for(var i in res.list){
								that.list.push(res.list[i]);
							}							
						}
						
					}
				})
			}
		},
	}
</script>

<style>
</style>
