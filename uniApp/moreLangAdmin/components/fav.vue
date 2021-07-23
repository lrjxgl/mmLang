<template>
	<view>
		<view class="btn-fav mgr-10" @click="favToggle()"  :class="isFav?'btn-fav-active':''">收藏</view>
		
	</view>
</template>

<script>
	export default{
		props:{
			dtable:"",
			dobjectid:0
		},
		data:function(){
			return {
				tablename:"",
				objectid:"",
				isFav:0
			}
		},
		created:function(){
			this.tablename=this.dtable;
			this.objectid=this.dobjectid;
			this.getPage()
		},
		watch:{
			dobjectid:function(n,o){
				this.objectid=n;
				this.getPage()
			}
		},
		methods:{
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/fav/get",
					data:{
						objectid:this.objectid,
						tablename:this.tablename
					},
					success:function(res){
						if(res.action=='delete'){
							that.isFav=false;
						}else{
							that.isFav=true;
						}
						
					}
				})
			},
			favToggle:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/fav/toggle",
					data:{
						objectid:this.objectid,
						tablename:this.tablename
					},
					success:function(res){
						if(res.action=='delete'){
							that.isFav=false;
						}else{
							that.isFav=true;
						}
						
					}
				})
			}
		}
	}
</script>

<style>
</style>
