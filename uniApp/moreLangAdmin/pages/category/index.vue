<template>
	<view>
		<div class="tabs-border">
			<div @click="gourl('index')" class="item active">列表</div>
			<div @click="gourl('add')" class="item">添加</div>
		</div>
		 <table class="tbs">
<thead>  <tr>
   <td>catid</td>
   <td>表名</td>
   <td>上级分类</td>
   <td>名称</td>
   <td>排序</td>
   <td>类型</td>
   <td>栏目模板</td>
   <td>列表模板</td>
   <td>详情模板</td>
<td>操作</td></tr>
  </tr>
</thead> <tr v-for="(item,index) in list" :key="index">
   <td>{{item.catid}}</td>
   <td>{{item.tablename}}</td>
   <td>{{item.pid}}</td>
   <td>{{item.cname}}</td>
   <td>{{item.orderindex}}</td>
   <td>{{item.type_id}}</td>
   <td>{{item.cat_tpl}}</td>
   <td>{{item.list_tpl}}</td>
   <td>{{item.show_tpl}}</td>
<td>
	<div class="btn-small mgr-5" @click="goAdd(item.catid)">编辑</div>

					<div class="btn-small mgr-5" @click="goShow(item.catid)">查看</div>
					<div class="btn-small btn-danger" @click="del(item)">删除</div>
</td>
  </tr>
 </table>

	</view>
</template>
<script>
	export default {
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
					url: that.app.apiHost + "/admin/category/index",
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
					url: that.app.apiHost + "/admin/category/index",
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
			},
			toggleStatus:function(item){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/admin/category/status",
					data:{
						id:item.id
					},
					success:function(res){
						item.status=res.status;
					}
				})
			},
			toggleRecommend:function(item){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/admin/category/recommend",
					data:{
						id:item.id
					},
					success:function(res){
						item.is_recommend=res.is_recommend;
					}
				})
			},
			del:function(item){
				var that=this;
				uni.showModal({
					content:"确认删除吗",
					success:function(res){
						if(!res.confirm){
							return false;
						}
						that.app.get({
							url:that.app.apiHost+"/admin/category/delete",
							data:{
								id:item.id
							},
							success:function(res){
								if(res.error){
									return false;
								}
								var list=[];
								for(var i in that.list){
									if(that.list[i].id!=item.id){
										list.push(that.list[i]);
									}
								}
								that.list=list;
							}
						})
					}
				})
			},
			goAdd:function(id){
				uni.navigateTo({
					url:"add?id="+id
				})
			},
			goShow:function(id){
				uni.navigateTo({
					url:"show?id="+id
				})
			}
		},
	}
</script>

<style>
</style>