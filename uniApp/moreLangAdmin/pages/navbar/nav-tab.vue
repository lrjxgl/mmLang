<template>
	<view>
		<div class="tabs-border">
			<div @click="go('index')" :class="tab=='index'?'active':''" class="item">导航管理</div>
			<div @click="go('add')" :class="tab=='add'?'active':''"   class="item">导航添加</div>
		</div>
	</view>
</template>

<script>
	export default{
		props:{
			tab:""
		},
		methods:{
			go:function(url){
				uni.navigateTo({
					url:url
				})
			}
		}
	}
</script>

<style>
</style>
