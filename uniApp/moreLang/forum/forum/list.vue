<template>
	<view v-if="pageLoad">
		<view class="tabs-border">
			<view  @click="setCat(0)" class="tabs-border-item" :class="0==catid?'tabs-border-active':''">推荐</view>
			<view @click="setCat(item.catid)" :class="item.catid==catid?'tabs-border-active':''" v-for="(item,key) in catList" :key="key" class="tabs-border-item">{{item.title}}</view>
			
		</view>
		<view class="sglist">
			<view class="emptyData" v-if="Object.keys(list).length==0">暂无帖子</view>
			<list-item :dlist="list"></list-item>
		</view>
		<go-top></go-top>
		<navigator :url="'../forum/add?gid='+gid+'&catid='+catid" class="fixedAdd">发布</navigator>
	</view> 
</template>

<script> 
	import listItem from "../list-item.vue";
	export default{
		components:{
			listItem
		},
		data:function(){
			return {
				pageLoad:false, 
				pageHide:false,
				gid:0,
				catid:0,
				list:[],
				catList:[],
				isFirst:true,
				per_page:0,
				group:{}
			}
			
		},
		onLoad:function(ops){
			if(ops.catid!=undefined){
				this.catid=ops.catid;
			}
			if(ops.gid!=undefined){
				this.gid=ops.gid;
			}
			this.getPage();
		},
		 
		onReachBottom:function(){
			this.getList();
		},
		onPullDownRefresh:function(){
			this.isFirst=true;
			this.per_page=0;
			this.getList();
			setTimeout(function(){
				uni.stopPullDownRefresh();
			},1000)
		},
		onShareAppMessage:function(){
			
		},
		onShareTimeline:function(){
			
		},
		methods:{
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/forum/list",
					data:{
						gid:this.gid,
						catid:this.catid
					},
					success:function(res){
						//登录
						if(res.error==1000){
							uni.navigateTo({
								url:"/pages/login/index",
							})
						}else{
							that.catList=res.catList;
							that.group=res.group;
							that.isFirst=false;
							that.pageLoad=true;
							that.list=res.list;
							that.per_page=res.per_page;
							 
							uni.setNavigationBarTitle({
								title: res.group.title
							});
						}
						 
					}
				})
			},
			setCat:function(cid){
				this.catid=cid;
				this.isFirst=true;
				this.per_page=0;
				this.getList();
			 },
			getList:function(){
				var that=this;
				if(!that.isFirst && that.per_page==0) return false;
				that.app.get({
					url:that.app.apiHost+"/forum/list",
					data:{
						per_page:that.per_page,
						catid:that.catid,
						gid:that.gid,
					},
					success:function(res){
						if(res.error){
							return false;
						}
						that.per_page=res.per_page; 
						if(that.isFirst){
							that.list=res.list;
							that.isFirst=false;
						}else{
							for(var i in res.list){
								that.list.push(res.list[i])
							}
							
						}
						
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
			}
			 
		},
	}
</script>

<style>
 @import url("../forum.css");
</style>
