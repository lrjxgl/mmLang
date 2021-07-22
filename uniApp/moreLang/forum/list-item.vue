<template>
	<view class="sglist">
		<view  class="sglist-item" v-for="(item,fkey) in  list" :key="fkey">
			<view @click="goUser(item.userid)"  class="flex mgb-5">
				<image :src="item.user.user_head+'.100x100.jpg'" class="wh-40 mgr-5 bd-radius-50"></image>
				<view class="flex-1">
					<view class="f14 mgb-5">{{item.user.nickname}}</view>
					<view class="f12 cl3">{{item.timeago}}</view>
				</view>
			</view>
			<div @click="goForum(item.id)" class="flex mgb-5">
				<div v-if="item.videourl" class="iconfont cl-red mgr-5 icon-video"></div>
				<div class="flex-1">{{item.title}}</div>
			</div>		
			<view @click="goForum(item.id)" class="sglist-imglist" v-if="item.imgList">                   
				<image v-for="(img,imgIndex) in item.imgList" :key="imgIndex" :src="img+'.100x100.jpg'" class="sglist-imglist-img"  mode="widthFix" ></image>
			</view>
			
			<view class="flex sglist-ft">
				<view class="sglist-ft-love">
					{{item.love_num}} </view>
				<view class="sglist-ft-cm">
					{{item.comment_num}} </view>
				<view class="sglist-ft-view">
					{{item.view_num}} </view>
			</view>
		</view>
	</view>
</template>

<script>
	export default{
		props:{
			dlist:{}
		},
		data:function(){
			return {
				list:[]
			}
		},
		created:function(){
			this.list=this.dlist;
		},
		watch:{
			dlist:function(n,o){
				this.list=n;
			}
		},
		methods:{
			goForum:function(id){
				uni.navigateTo({
					url:"../forum/show?id="+id
				})
			},
			goUser:function(userid){
				uni.navigateTo({
					url:"../forum_home/index?userid="+userid
				})
			}
		}
	}
</script>

<style>
</style>
