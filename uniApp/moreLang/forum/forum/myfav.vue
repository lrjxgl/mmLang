<template>
	<view>
		<div v-if="pageLoad">
			<list-item   :dlist="list"></list-item>
			<div @click="getList" v-if="per_page>0" class="loadMore">点我加载更多...</div>
		</div>
	</view>
</template>

<script>
	var tablename = "mod_forum";
	import listItem from "../list-item.vue";
	export default{
		components:{
			listItem
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
		},
		onShareAppMessage: function() {

		},
		methods: {
			goDetail: function (id) {
				uni.navigateTo({
					url: "../forum/show?id=" + id
				})
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/forum_fav/index",
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
					url: that.app.apiHost + "/forum_fav/index",
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
