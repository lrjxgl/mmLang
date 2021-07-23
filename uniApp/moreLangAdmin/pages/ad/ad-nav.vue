<template>
	<view>
		<div class="tabs-border">
			<div @click="gourl('../ad/index')" :class="tab=='adIndex'?'tabs-border-active':''" class="tabs-border-item" >广告管理</div>
			<div @click="gourl('../ad/add')" :class="tab=='adAdd'?'tabs-border-active':''" class="tabs-border-item">广告添加</div>
			<div  @click="gourl('../ad_tags/index')" :class="tab=='adTagsIndex'?'tabs-border-active':''"  class="tabs-border-item">广告分类管理</div>
			<div  @click="gourl('../ad_tags/add')" :class="tab=='adTagsAdd'?'tabs-border-active':''"  class="tabs-border-item">广告分类添加</div>
		</div>
	</view>
</template>

<script>
	export default{
		props:{
			tab:""
		},
		data:function(){
			return {
				a:1
			}
		},
		created:function(){
			
		},
		methods:{
			gourl:function(url){
				uni.navigateTo({
					url:url
				})
			}
		}
	}
</script>

<style>
</style>
