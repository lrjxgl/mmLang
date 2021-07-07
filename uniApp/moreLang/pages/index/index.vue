<template>
	<view>

		<swiper :style="{height:swiperHeight+'px'}" :indicator-dots="true" :autoplay="true" :interval="3000"
			:duration="1000">
			<swiper-item v-for="(item,key) in flashList" :key="key">
				<view class="swiper-item">
					<image @click="gourl(item.link1)" :src="item.imgurl" style="width:100%" mode="widthFix"></image>
				</view>
			</swiper-item>

		</swiper>
		<view v-if="Object.keys(navList).length>0" class="m-navPic mgb-5">

			<navigator v-for="(item,index) in navList" :key="index" :url="item.link_url" class="m-navPic-item">
				<image mode="widthFix" class="m-navPic-img" :src="item.logo"></image>
				<view class="m-navPic-title">{{item.title}}</view>
			</navigator>
		</view>
		<view v-if="Object.keys(adList).length>0" class="adBox bg-white mgb-5">
			<navigator v-for="(item,index) in adList" :key="index" :url="item.link1" class="adBox-item">
				<image :src="item.imgurl" mode="widthFix" class="adBox-img"></image>
			</navigator>
		</view>
		<scroll-view scroll-x="true">
			<div :style="tabClass" class="tabNav">
				<div @click="setTab('article')" :class="tab=='article'?'tabNav-item-active':''" class="tabNav-item ">头条
				</div>
				<div @click="setTab('freeshop')" :class="tab=='freeshop'?'tabNav-item-active':''" class="tabNav-item ">
					闲时优惠</div>
				<div @click="setTab('forum')" :class="tab=='forum'?'tabNav-item-active':''" class="tabNav-item ">论坛
				</div>
				<div @click="setTab('sblog')" :class="tab=='sblog'?'tabNav-item-active':''" class="tabNav-item ">好友圈
				</div>
				<div @click="setTab('fenlei')" :class="tab=='fenlei'?'tabNav-item-active':''" class="tabNav-item ">同城信息
				</div>
				<div @click="setTab('household')" :class="tab=='household'?'tabNav-item-active':''"
					class="tabNav-item ">家电维护</div>
			</div>

		</scroll-view>
		<div v-if="fixTab" style="height: 40px;"></div>
		<div v-if="tab=='article'">
			<view class="flexlist">
				<view v-if="Object.keys(articleList).length==0" class="emptyData">暂无文章</view>
				<view v-else @click="gourl('../../pages/article/show?id='+item.id)" class="flexlist-item pdb-10"
					v-for="(item,index) in articleList" :key="index">
					<image v-if="item.imgurl!=''" class="flexlist-img" :src="item.imgurl+'.100x100.jpg'"></image>
					<view class="flex-1">
						<view class="flexlist-title f16">{{item.title}}</view>
						<view class="flexlist-desc cl2 f14">{{item.description}}</view>
					</view>
				</view>
				<div @click="gourl('../../pages/article/list?catid=696')" class="loadMore">更多头条资讯</div>
			</view>
		</div>
		<div v-if="tab=='forum'" class="sglist">
			<view v-if="Object.keys(forumList).length==0" class="emptyData">暂无帖子</view>
			<view v-else @click="gourl('../../pageforum/forum/show?id='+item.id)" class="sglist-item"
				v-for="(item,fkey) in forumList" :key="fkey">
				<view class="flex mgb-5">
					<image :src="item.user_head+'.100x100.jpg'" class="wh-40 mgr-5 bd-radius-50"></image>
					<view class="flex-1">
						<view class="f14 mgb-5">{{item.nickname}}</view>
						<view class="f12 cl3">{{item.timeago}}</view>
					</view>
				</view>
				<div class="flex mgb-5">
					<div v-if="item.videourl" class="iconfont cl-red mgr-5 icon-video"></div>
					<div class="flex-1">{{item.title}}</div>
				</div>
				<view class="sglist-imglist" v-if="item.imgslist">
					<image v-for="(img,imgIndex) in item.imgslist" :key="imgIndex" :src="img+'.100x100.jpg'"
						class="sglist-imglist-img" mode="widthFix"></image>
				</view>

				<view class="flex sglist-ft">
					<view class="sglist-ft-love">
						{{item.love_num}}
					</view>
					<view class="sglist-ft-cm">
						{{item.comment_num}}
					</view>
					<view class="sglist-ft-view">
						{{item.view_num}}
					</view>
				</view>
			</view>
			<div @click="gourl('../../pageforum/forum/index')" class="loadMore">更多论坛信息</div>
		</div>
		<div v-if="tab=='sblog'">
			<view v-if="Object.keys(blogList).length==0" class="emptyData">暂无帖子</view>
			<view v-else v-for="(item,index) in blogList" :key="index" @click="goBlog(item.id)" class="sglist-item">
				<view class="flex mgb-5">
					<image :src="item.user.user_head+'.100x100.jpg'" class="wh-40 mgr-5 bd-radius-50"></image>
					<view class="flex-1">
						<view class="f14 mgb-5">{{item.user.nickname}}</view>
						<view class="f12 cl3">{{item.timeago}}</view>
					</view>
				</view>
				<view class="sglist-title block" v-html="item.content"></view>
				<view class="sglist-imglist">

					<img v-for="(img,imgIndex) in item.imgslist" :key="imgIndex" :src="img+'.100x100.jpg'"
						class="sglist-imglist-img" />

				</view>
				<view class="sglist-ft">
					<view class="sglist-ft-love">{{item.love_num}}</view>
					<view class="sglist-ft-cm">{{item.comment_num}}</view>
					<view class="sglist-ft-view">{{item.view_num}}</view>
				</view>
			</view>
			<div @click="gourl('../../pagesblog/sblog/index')" class="loadMore">来交些好友</div>
		</div>
		<div v-if="tab=='fenlei'">
			<view v-if="Object.keys(fenleiList).length==0" class="emptyData">暂无帖子</view>
			<div v-else v-for="(item,index) in fenleiList" :key="index"
				@click="gourl('../../pagefenlei/fenlei/show?id='+item.id)" class="flexlist-item">
				<img v-if="item.imgurl!=''" class="flexlist-img" :src="item.imgurl+'.100x100.jpg'" />

				<div class="flex-1">
					<div class="flexlist-title">{{item.title}}</div>

					<div v-if="item.money>0" class="flexlist-row">
						<div class="cl-money">￥{{item.money}}</div>
					</div>

					<div class="flexlist-desc">
						{{item.description}}
					</div>
				</div>
			</div>
			<div @click="gourl('../../pagefenlei/fenlei/index')" class="loadMore">更多同城信息</div>
		</div>
		<div v-if="tab=='household'">
			<view class="flexlist">
				<view v-if="Object.keys(householdList).length==0" class="emptyData">暂无帖子</view>
				<div v-else @click="gourl('../../household/household_product/show?id='+item.id)"
					v-for="(item,index) in householdList" :key="index" class="flexlist-item">
					<image class="flexlist-img" mode="widthFix" :src="item.imgurl+'.100x100.jpg'"></image>
					<div class="flex-1">
						<div class="flexlist-title">{{item.title}}</div>
						<div class="flex flex-ai-center mgb-5">
							<div class="f12 cl2">{{item.addr}}</div>
							<div class="flex-1"></div>
							<div class="cl-money">￥{{item.price}}</div>
						</div>
						<div class="flexlist-desc">{{item.description}}</div>
					</div>
				</div>
			</view>
			<div @click="gourl('../../household/household/index')" class="loadMore">更多家政服务</div>
		</div>
		<div v-if="tab=='freeshop'">
			<view v-if="Object.keys(freeshopList).length==0" class="emptyData">暂无帖子</view>
			<div v-else v-for="(item,index) in  freeshopList" :key="index"
				@click="gourl('../../freeshop/freeshop_product/show?productid='+item.productid)"
				class="sglist-item pointer">
				<div class="flex flex-ai-start  mgb-5">
					<div v-if="item.catid" class="btn-type mgr-5">{{item.catid_name}}</div>
					<div v-if="item.invite_money>0" class="btn-type cl-money mgr-5">赏</div>
					<div class="flex-1" v-html="item.content"></div>
				</div>

				<div class="sglist-imglist">

					<img v-for="(img,imgIndex) in item.imgslist" :key="imgIndex" :src="img+'.100x100.jpg'"
						class="sglist-imglist-img" />

				</div>
				<div class="flex mgb-10 flex-ai-center">
					<div class="mgr-5 cl2">优惠价</div>
					<div class="cl-money mgr-10">￥{{item.price}}</div>
					<div class="mgr-5 cl2"></div>
					<div class="cl2 mgr-10">{{item.discount}}折</div>
					<div class="flex-1"></div>
					<div class="cl3 f12">原价</div>
					<div class="market-price">￥{{item.market_price}}</div>
				</div>
				<div class="flex flex-ai-center">
					<div class="mgr-5 cl2">销量</div>
					<div class="cl-num mgr-10">{{item.buynum}}</div>
					<div class="mgr-5 cl2">库存</div>
					<div class="cl-num mgr-10">{{item.maxnum}}</div>
					<div class="flex-1"></div>
					<div class="cl-status">{{item.status_name}}</div>
				</div>

			</div>
			<div @click="gourl('../../freeshop/freeshop/index')" class="loadMore">更多优惠活动</div>
		</div>
		<mt-footer tab="home"></mt-footer>
		<go-top></go-top>
	</view>
</template>

<script>
	import mtFooter from "../../components/footer.vue";
	var winWidth = 750;
	export default ({
		components: {
			mtFooter
		},
		data: function() {
			return {
				flashList: [],
				navList: [],
				adList: [],
				pageLoad: false,
				articleList: [],
				forumList: [],
				blogList: [],
				fenleiList: [],
				householdList: [],
				freeshopList: [],
				tab: "article",
				swiperHeight: 200,
				tabClass: "",
				fixTab: false,
				pageCacheKey:"page-index"
			}
		},
		onLoad: function() {
			var sys = uni.getSystemInfoSync();
			this.swiperHeight = sys.windowWidth / 2;
			console.log(this.swiperHeight);
			if(!this.getCache()){
				this.getPage();
			}
		},
		onPageScroll: function(e) {
			if (e.scrollTop > this.swiperHeight + 70) {
				this.fixTab = true;
				// #ifdef H5
				this.tabClass = "position: fixed;left:0;right:0;top:44px;";
				// #endif
				//#ifndef H5
				this.tabClass = "position: fixed;left:0;right:0;top:0px;";
				// #endif

			} else {
				this.tabClass = "";
				this.fixTab = false;
			}

		},
		onPullDownRefresh: function() {
			this.getPage();
			uni.stopPullDownRefresh();
		},
		onShareAppMessage: function() {

		},
		onShareTimeline: function() {

		},
		 
		methods: {
			setCache: function() {
				var val=this.$data;
				val.expire= Date.parse(new Date()) / 1000 + 300;
				uni.setStorageSync(this.pageCacheKey, JSON.stringify(val));
			},
			getCache: function() {
				var that=this;
				var val = uni.getStorageSync(this.pageCacheKey);
				if (!val) return false;
				var time = Date.parse(new Date()) / 1000;
				if (val.expire < time) {
					return false;
				}
				
				if(val.shopid!=this.shopid){
					return false;
				}
				var obj1 = JSON.parse(val);				
				var obj2=this.$data;
				Object.keys(obj1).forEach((key) => {
					if(key!='swiperHeight'){
						 that[key]=obj1[key] ;
					}
				   
				 });	 
				return true;
			},
			goTop: function() {
				uni.pageScrollTo({
					scrollTop: 0
				})
			},
			getIndex: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost,
					success: function(res) {
						if (!res.error) {
							that.navList = res.data.navList;
							that.flashList = res.data.flashList;
							that.articleList = res.data.articleList;
							if(res.data.adList){
								that.adList = res.data.adList;
							}
							that.setCache()
							
						}

					}
				})
			},
			gourl: function(url) {
				console.log(url)
				uni.navigateTo({
					url: url
				})
			},
			setTab: function(tab) {
				console.log(tab)
				this.tab = tab;
				this.setCache()
			},
			goBlog: function(id) {
				uni.navigateTo({
					url: "../../pagesblog/sblog_blog/show?id=" + id
				})
			},
			goTopic: function(title) {
				uni.navigateTo({
					url: "../../pagesblog/sblog_blog/topic?title=" + title
				})
			},
			getPage: function() {
				
				this.getIndex();
				this.getBlogList();
				this.getForumList();
				this.getFenleiList();
				this.getHouseholdList();
				this.getFreeshopList();
			},
			getBlogList: function() {
				var that = this;
				this.app.get({
					url: that.app.apiHost + "/module.php?m=sblog_blog&a=list&type=recommend&ajax=1",
					dataType: "json",
					success: function(res) {
						if (!res.error) {
							that.blogList = res.data.list;
							that.setCache()
						}
					}
				})
			},
			getForumList: function() {
				var that = this;
				this.app.get({
					url: that.app.apiHost + "/index.php?a=forum&ajax=1",
					dataType: "json",
					success: function(res) {
						if (!res.error) {
							that.forumList = res.data.recList;
							that.setCache()
						}
					}
				})
			},
			getFenleiList: function() {
				var that = this;
				this.app.get({
					url: that.app.apiHost + "/index.php?a=fenlei&ajax=1",
					dataType: "json",
					success: function(res) {
						if (!res.error) {
							that.fenleiList = res.data.recList;
							that.setCache()
						}
					}
				})
			},
			getHouseholdList: function() {
				var that = this;
				this.app.get({
					url: that.app.apiHost + "/index.php?a=household&ajax=1",
					dataType: "json",
					success: function(res) {
						if (!res.error) {
							that.householdList = res.data.recList;
							that.setCache()
						}
					}
				})
			},
			getFreeshopList: function() {
				var that = this;
				this.app.get({
					url: that.app.apiHost + "/module.php?m=freeshop_product&a=list&type=recommend&ajax=1",
					dataType: "json",
					success: function(res) {
						if (!res.error) {
							that.freeshopList = res.data.list;
							that.setCache()
						}
					}
				})
			},
		}
	})
</script>

<style>
	@import "../../common/sblog.css";

	.tabNav {
		padding: 10px;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: center;
		background-color: #fff;
		border-bottom: 1px solid #eee;
		background-color: #ff;
		z-index: 999;
		position: relative;
		width: 100%;
	}

	.tabNav-item {
		cursor: pointer;
		margin-right: 10px;
		font-weight: 600;
	}

	.tabNav-item-active {
		color: #f60;
		padding-bottom: 3px;

	}

	.swiper-item {
		height: 100%;
	}

	.mtt10 {
		margin-top: -10px;
	}

	.adBox {
		display: flex;
		flex-direction: row;

	}

	.adBox-item {
		flex: 1;
		margin: 5px;
	}

	.adBox-img {
		width: 100%;
		border-radius: 10px;
	}

	.m-navPic-img {
		border-radius: 50%;
	}
</style>
