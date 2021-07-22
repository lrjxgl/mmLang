<template>
	<view v-if="pageLoad">
		 
		<view class="main-body pd-5 bg-fff">
			<d-userbox :user="author"></d-userbox>
			<view class="d-title">{{data.title}}</view>
			<view class="d-tools">
				<view class="mgr-10 cl2">作者：{{author.nickname}}</view>
				<view class="cl2">{{data.timeago}}</view>
			</view>
			<view class="row-box">
				<div v-if="data.videourl!=''" class="flex flex-center mgb-5">
					<video :src="data.videourl" ></video>
				</div>
				<div class="flex-col flex-center mgb-5">
					<image style="width: 100%;" @longpress="downImg(item)" class="wmax mgb-5" mode="widthFix" v-for="(item,key) in imgsList" :key="key" :src="item"></image> 
				</div>
				
			</view>
			<rich-text class="d-content" :nodes="data.content">
				 
			</rich-text>

			<view class="flex flex-center mgb-10">
				 
				<love dtable="mod_forum" :dobjectid="data.id"></love> 
				<fav dtable="mod_forum" :dobjectid="data.id"></fav> 
				 
			</view>
		</view>
		<view class="comment-hd">跟帖列表</view>
		<!--评论-->
		
		<cmform tablename="mod_forum" :objectid="data.id"></cmform>
		<navigator :url="'../forum/add?gid='+data.gid+'&catid='+data.catid" class="fixedAdd">发布</navigator>
	</view>
</template>

<script>
	import dUserbox from "../d-userbox.vue";
	import cmform from "../../components/cmform.vue";
	import love from "../../components/love.vue";
	import fav from "../../components/fav.vue"; 
	var id;
	export default {
		components: {
			dUserbox,
			cmform,
			love,
			fav
		},
		data:function(){
			return {
				pageLoad:false, 
				pageHide:false,
				data:{},
				author:{},
				imgsList:[],
				isLove:0,
				isFav:0
			}
			
		},
		onLoad: function (ops) {
			 
			 
			if(ops.id!=undefined){
				id=ops.id;
			}
			if(ops.scene!=undefined){
				id=ops.scene.id;
			}
			this.getPage();
			this.addClick();
			
		},
		onShareAppMessage:function(){
			
		},
		onShareTimeline:function(){
			
		},
		onPullDownRefresh:function(){
			this.refresh();
		} ,
		methods: {
			refresh:function(){
				this.getPage();
				setTimeout(function(){
					uni.stopPullDownRefresh();
				},1000)
			},
			addClick:function(){
				this.app.get({
					url: this.app.apiHost + "/forum/addclick?id=" + id,
					success: function (res) {
					}
				})
			},
			downImg:function(url){
				uni.downloadFile({
				    url: url,  
				    success: (res) => {
				        if (res.statusCode === 200) {
				            uni.saveImageToPhotosAlbum({
				            	filePath:res.tempFilePath,
								success:function(){
									uni.showToast({
										title:"图片保存成功"
									})
								}
				            })
				        }
				    }
				});
			},
			getPage: function () {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/forum/show?id=" + id,
					 
					success: function (res) {
						
						that.pageLoad = true;
						that.data = res.data;
						that.author=res.author;
						that.imgsList=res.data.imgsList;
						uni.setNavigationBarTitle({
							title:res.data.title
						})
					}
				})
			}
			 
			
		},
	}
</script>
<style>
	.wmax{
		width: 100%;
	}
</style>

 
