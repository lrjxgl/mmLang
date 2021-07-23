<template>
	<view class="leftBox">
		<div class="menu-bar" id="menuBar"></div>
		<div v-if="menuPage=='main'" class="">
			<div v-for="(item,index) in list" :key="index" :class="item.on?'menu-active':''" class="menu">
				<div class="menu-title"  @click="item.on=item.on?false:true"><i class="iconfont icon-home"></i> {{item.title}}</div>
				<div class="menu-box">
					<div v-for="(cc,ii) in item.child" :key="ii" @click="go(cc)"   :class="actItem==cc.m+cc.a?'menu-item-active':''" class="menu-item">{{cc.title}}</div>
				</div>
			</div>
		</div>
		<forum-menu v-if="menuPage=='forum'"></forum-menu>
	</view>
</template>

<script>
	import forumMenu from "../forum/menu.vue";
	export default{
		components:{
			forumMenu
		},
		data:function(){
			return {
				list:[],
				actItem:"",
				menuPage:"main"
			}
		},
		created:function(){
			this.getPage();
			var that=this;
			uni.$on('leftWindowPage',function(data){
				that.menuPage=data.page;
			})
		},
		
		methods:{
			go:function(cc){
				this.actItem=cc.m+cc.a;
				var m=cc.m;
				var a=cc.a;
				if(a=='default'){
					a="index";
				}
				var path="/pages";
				var url=path+"/"+m+"/"+a;
				
				uni.redirectTo({
					url:url
				})
			},
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/admin/navbar/get?group_id=2",
					success:function(res){
						for(var i in res.list){
							res.list[i].on=true;
							for(var j in res.list[i].child){
								res.list[i].child[j].selected=false;
							}
						}
						that.list=res.list;
					}
				})
			}
		}
	}
</script>

<style>
	@import url("../common/admin-menu.css");
	.leftBox{
		 
		height: 100%;
		overflow: auto;
		background-color: #212529;
	}
	.sitename {
		background-color: #AA3130 !important;
		color: #fff;
		font-size: 16px;
		line-height: 50px;
		text-align: center;
	}
	 
	
</style>
