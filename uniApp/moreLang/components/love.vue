<template>
	<view>
		<view class="btn-love  mgr-10" @click="loveToggle()"  :class="isLove?'btn-love-active':''">
			喜欢
		</view>
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
				isLove:0
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
					url:that.app.apiHost+"/love/get",
					data:{
						objectid:this.objectid,
						tablename:this.tablename
					},
					success:function(res){
						if(res.action=='delete'){
							that.isLove=false;
						}else{
							that.isLove=true;
						}
						
					}
				})
			},
			loveToggle:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/love/toggle",
					data:{
						objectid:this.objectid,
						tablename:this.tablename
					},
					success:function(res){
						if(res.action=='delete'){
							that.isLove=false;
						}else{
							that.isLove=true;
						}
						
					}
				})
			}
		}
	}
</script>

<style>
</style>
