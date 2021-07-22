<template>
	<view v-if="pageLoad">
		<div v-if="!list || list.length==0" class="emptyData">暂无发帖</div>
		<list-item v-else :dlist="list"></list-item>
	</view> 
</template>

<script> 
	 
	 
	var catid=0;
	var gid=0;
	import listItem from "../list-item.vue";
	export default{
		components:{
			listItem
		},
	
		data:function(){
			return {
				pageLoad:false, 
				pageHide:false,
				list:[],
				isFirst:true,
				per_page:0
			}
			
		},
		onLoad:function(option){
			 
			uni.setNavigationBarTitle({
				title: '新帖子'
			});
			this.getPage();
		},
		 
		onShow:function(){
			if(this.pageHide){
				this.pageHide=false;
				this.getPage();
			}			
		},
		onHide:function(){
			this.pageHide=true;
		}, 
		onReachBottom:function(){
			this.getList();
		},
		onPullDownRefresh:function(){
			this.refresh();
		},
		methods:{
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/forum/new",
					success:function(res){
						if(res.error){
							return false;
						}
						that.isFirst=false;
						that.pageLoad=true;
						that.list=res.list;
						that.per_page=res.per_page;
						 
					}
				})
			},
			 
			getList:function(){
				var that=this;
				if(!that.isFIrst && that.per_page==0) return false;
				that.app.get({
					url:that.app.apiHost+"/forum/new",
					data:{
						per_page:that.per_page
					},
					success:function(res){
						if(res.error){
							return false;
						}
						if(that.isFirst){
							that.isFirst=false;
							that.list=res.list;
						}else{
							for(var i in res.list){
								that.list.push(res.list[i]);
							}
						}
						 
						that.per_page=res.per_page;
						 
					}
				})
			},
			goForum: function (id) {
				uni.navigateTo({
					url: "../forum/show?id=" + id
				})
			},
			goUser: function(userid) {
				uni.navigateTo({
					url: "../forum_home/index?userid=" + userid
				})
			},
			refresh:function(){
				this.getPage();
				setTimeout(function(){
					uni.stopPullDownRefresh();
				},1000)
			},
			loadMore:function(){
				this.getList();
			}
			 
		},
	}
</script>

<style>
 
</style>
