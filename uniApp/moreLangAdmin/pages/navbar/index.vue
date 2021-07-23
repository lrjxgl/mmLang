<template>
	<view>
		<nav-tab tab="index"></nav-tab>
		<div class="flex pd-10">
			<div @click="setGid(2)" :class="group_id==2?'cl-primary':''" class="mgr-10 pointer">后台左边</div>
			<div  @click="setGid(1)" :class="group_id==1?'cl-primary':''" class="pointer" >后台顶部</div>
		</div>
		<table class="tbs">
			<thead>
				<tr>
					<td width="5%" align="center">ID</td>
					<td width="10%" align="center">名称</td>
					<td width="5%">状态</td>
					<td width="4%" align="center">m</td>
					<td width="4%" align="center">a</td>

					<td width="7%" align="center">目标</td>

					<td width="7%" height="30" align="center">排序</td>
					<td width="16%" height="30" align="center">操作</td>
				</tr>
			</thead>
			<tbody v-for="(item,index) in list" :key="item.id">
				<tr>
					<td align="center">{{item.id}}</td>
					<td align="left"><a :href="item.link_url">{{item.title}}</a></td>
					<td>
						<div @click="toggleStatus(item)" :class="item.status==1?'yes':'no'"></div>
					</td>
					<td align="center">{{item.m}}</td>
					<td align="center">{{item.a}}</td>

					<td align="center">{{item.target}}</td>

					<td height="25" align="center"><input type="text" v-model="item.orderindex" class="input-small" />
					</td>
					<td height="25" align="center">
						<div class="btn-mini mgr-5" @click="gourl('add?pid='+item.id)">添加子类</div>
						<div class="btn-mini mgr-5" @click="gourl('add?id='+item.id)">编辑</div>
						<div @click="del(item)" class="btn-mini btn-danger">删除</div>
					</td>
				</tr>

				<tr v-for="(cc,ii) in item.child" :key="cc.id">
					<td align="center">{{cc.id}}</td>
					<td align="left">|__<a :href="cc.link_url" target="_blank">{{cc.title}}</a> </td>
					<td> 
					<div  @click="toggleStatus(cc)"  :class="cc.status==1?'yes':'no'"></div>
					</td>
					<td align="center">{{cc.m}}</td>
					<td align="center">{{cc.a}}</td>

					<td align="center">{{cc.target}}</td>

					<td height="25" align="center"><input name="orderindex[]" type="text" v-model="cc.orderindex"
							class="input-small" /></td>
					<td height="25" align="center">
						<div class="btn-mini mgr-5" @click="gourl('add?id='+cc.id)">编辑</div>
						<div @click="del(cc)" class="btn-mini btn-danger">删除</div>
					</td>
				</tr>
			</tbody>



		</table>
	</view>
</template>

<script>
	import navTab from "./nav-tab.vue"
	export default {
		components: {
			navTab
		},
		data: function() {
			return {
				pageLoad: false,
				list: [],
				per_page: 0,
				isFirst: true,
				group_id:2
			}
		},
		onLoad: function() {
			this.getPage();
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
			setGid:function(gid){
				this.group_id=gid;
				this.getPage();
			},
			gourl: function(url) {
				uni.navigateTo({
					url: url
				})
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/admin/navbar/index",
					data:{
						group_id:this.group_id
					},
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
					url: that.app.apiHost + "/admin/navbar/index",
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
			del: function(item) {
				var that=this;
				uni.showModal({
					content:"确认删除吗",
					success:function(res){
						if(!res.confirm){
							return false;
						}
						that.app.get({
							url:that.app.apiHost+"/admin/navbar/delete",
							data:{
								id:item.id
							},
							success:function(res){
								uni.showToast({
									title:res.message,
									icon:"none"
								})
								that.getPage();
							}
						})
					}
				})	
			},
			toggleStatus:function(item){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/admin/navbar/status",
					data:{
						id:item.id
					},
					success:function(res){
						item.status=res.status;
					}
				})
			},
		},
	}
</script>

<style>
</style>
