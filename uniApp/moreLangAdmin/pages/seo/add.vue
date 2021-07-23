<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>m：</td>		
						<td><input class="input-text" type="text" name="m"   v-model="data.m" ></td>
						</tr>
  <tr>
						<td>a：</td>		
						<td><input class="input-text" type="text" name="a"   v-model="data.a" ></td>
						</tr>
  <tr>
						<td>cname：</td>		
						<td><input class="input-text" type="text" name="cname"   v-model="data.cname" ></td>
						</tr>
  <tr>
						<td>title：</td>		
						<td><input class="input-text" type="text" name="title"   v-model="data.title" ></td>
						</tr>
  <tr>
						<td>description：</td>		
						<td><input class="input-text" type="text" name="description"   v-model="data.description" ></td>
						</tr>
  <tr>
						<td>keywords：</td>		
						<td><input class="input-text" type="text" name="keywords"   v-model="data.keywords" ></td>
						</tr>
  <tr>
						<td>object_id：</td>		
						<td><input class="input-text" type="text" name="object_id"   v-model="data.object_id" ></td>
						</tr>
</table> <button form-type="submit" class="btn-row-submit">保存</button> 
</form>
	</view>
</template>
<script>
	import upimgBox from "../../components/upimgbox.vue";
	import skyUpimg from "../../components/skyupimg.vue";
	export default {
		components: {
			upimgBox,
			skyUpimg
		},
		data: function() {
			return {
				id: 0,
				pageLoad: false,
				data: {},
				imgsList: []
			}
		},
		onLoad: function(ops) {
			if (ops.id != undefined) {
				this.id = ops.id;
			}
			console.log("onLoad")
			this.getPage();
		},
		methods: {
			gourl: function(url) {
				uni.navigateTo({
					url: url
				})
			},
			setImgs: function(e) {
				this.data.imgsdata = e;
			},
			setImgurl: function(e) {
				this.data.imgurl = e;
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/admin/seo/add",
					data: {
						id: this.id
					},
					success: function(res) {
						if (Object.keys(res.data).length > 0) {
							that.data = res.data;
							that.imgsList = res.imgsdata;
						} else {
							that.data = {
								id: 0,
								title: "",
								typeid: 0,
								total_money: 0,
								description: "",
								imgsdata: ""
							}
						}

						that.pageLoad = true;
					}
				})
			},
			submit: function(e) {
				var that = this;
				that.app.post({
					url: that.app.apiHost + "/admin/seo/save",
					data: e.detail.value,
					success: function(res) {
						uni.showToast({
							title: res.message
						})
					}
				})

			}
		}
	}
</script>

<style>
</style>