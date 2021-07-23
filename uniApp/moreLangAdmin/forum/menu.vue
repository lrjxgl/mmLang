<template>
	<view>
		<div v-for="(item,index) in list" :key="index" :class="item.on?'menu-active':''" class="menu">
			<div class="menu-title"  @click="item.on=item.on?false:true"><i class="iconfont icon-home"></i> {{item.title}}</div>
			<div class="menu-box">
				<div v-for="(cc,ii) in item.child" :key="ii" @click="go(cc.url)"   :class="actItem==cc.url?'menu-item-active':''" class="menu-item">{{cc.title}}</div>
			</div>
		</div>
		
	</view>
</template>

<script>
	export default{
		data:function(){
			return {
				actItem:"",
				list:[{
					 
					title:"论坛",
					on:1,
					child:[{
						url:"/forum/forum/index",
						title:"论坛"
					},{
						url:"/forum/forum_comment/index",
						title:"评论"
					},{
						url:"/forum/forum_group/index",
						title:"板块"
					}]
				},{
					title:"聚合标签",
					on:0,
					child:[{
						url:"/forum/forum_tags/index",
						title:"标签"
					}]
				}]
			}
		},
		methods:{
			go:function(url){
				this.actItem=url;
				uni.redirectTo({
					url:url
				})
			},
		}
	}
</script>

<style>
	@import url("../common/admin-menu.css");
</style>
