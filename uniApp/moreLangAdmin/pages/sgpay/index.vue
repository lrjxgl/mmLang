<template>
	<view>
		<div class="tabs-border">
			<div @click="gourl('index')" class="item active">列表</div>
			<div @click="gourl('add')" class="item">添加</div>
		</div>
		 <table class="tbs">
<thead>  <tr>
   <td>id</td>
   <td>userid</td>
   <td>银行卡名称</td>
   <td>卡号</td>
   <td>yhk_user</td>
   <td>金额</td>
   <td>时间</td>
   <td>status</td>
<td>操作</td></tr>
  </tr>
</thead> <tr v-for="(item,index) in list" :key="index">
   <td>{{item.id}}</td>
   <td>{{item.userid}}</td>
   <td>{{item.yhk_name}}</td>
   <td>{{item.yhk_num}}</td>
   <td>{{item.yhk_user}}</td>
   <td>{{item.money}}</td>
   <td>{{item.createtime}}</td>
   <td>{{item.status}}</td>
<td>
	<div class="btn-small mgr-5" @click="goAdd(item.id)">编辑</div>

					<div class="btn-small mgr-5" @click="goShow(item.id)">查看</div>
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
					url: that.app.apiHost + "/admin/sgpay/index",
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
					url: that.app.apiHost + "/admin/sgpay/index",
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
					url:that.app.apiHost+"/admin/sgpay/status",
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
					url:that.app.apiHost+"/admin/sgpay/recommend",
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
							url:that.app.apiHost+"/admin/sgpay/delete",
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