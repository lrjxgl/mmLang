<template>
	<view>
		<ad-nav tab="adIndex"></ad-nav>
		<div class="main-body">
			<table class="tbs">
				<thead>
					<tr>
						<td width="50">id</td>

						<td>名称</td>
						<td>图片地址</td>
						<td>tag_id</td>
						<td>链接1</td>




						<td>排序</td>
						<td>状态</td>


						<td>操作</td>
					</tr>
					</tr>
				</thead>
				<tr v-for="(item,index) in list" :key="index">
					<td>{{item.id}}</td>

					<td>{{item.title}}</td>
					<td>
						<image :src="item.imgurl+'.100x100.jpg'" mode="widthFix" class="wh-60"></image>
					</td>
					<td>{{item.tag_id}}</td>
					<td>{{item.link1}}</td>




					<td>{{item.orderindex}}</td>
					<td>{{item.status}}</td>


					<td>
						<div class="btn-small mgr-5" @click="goAdd(item.id)">编辑</div>
						<div class="btn-small mgr-5" @click="goShow(item.id)">查看</div>
						<div class="btn-small btn-danger" @click="del(item)">删除</div>
					</td>
				</tr>
				{/foreach}
			</table>
			 
		</div>
	</view>
</template>

<script>
	import adNav from "./ad-nav.vue"
	export default {
		components: {
			adNav
		},
		data: function() {
			return {
				pageLoad: false,
				list: [],
				per_page: 0,
				isFirst: true
			}
		},
		onLoad: function() {
			this.getPage();
		},
		onReachBottom: function() {
			this.getList();
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
			gourl: function(url) {
				uni.navigateTo({
					url: url
				})
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/admin/ad/index",
					success: function(res) {
						that.pageLoad = true;
						that.list = res.list;
						that.per_page = res.per_page;
					}
				})
			},
			getList: function() {
				var that = this;
				if (that.per_page == 0 && !that.isFirst) {
					return false;
				}
				that.app.get({
					url: that.app.apiHost + "/admin/ad/index",
					data: {
						per_page: that.per_page
					},
					success: function(res) {
						that.per_page = res.per_page;
						if (that.isFirst) {
							that.list = res.list;
							that.isFirst = false;
						} else {
							for (var i in res.list) {
								that.list.push(res.list[i]);
							}
						}

					}
				})
			}
		},
	}
</script>

<style>
</style>
