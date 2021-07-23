<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>objectid：</td>		
						<td><input class="input-text" type="text" name="objectid"   v-model="data.objectid" ></td>
						</tr>
  <tr>
						<td>shopid：</td>		
						<td><input class="input-text" type="text" name="shopid"   v-model="data.shopid" ></td>
						</tr>
  <tr>
						<td>orderindex：</td>		
						<td><input class="input-text" type="text" name="orderindex"   v-model="data.orderindex" ></td>
						</tr>
  <tr>
						<td>createtime：</td>		
						<td><input class="input-text" type="text" name="createtime"   v-model="data.createtime" ></td>
						</tr>

					<tr>
							<td>图片：</td>
							<td>
								<input class="none" type="text" name="imgurl" v-model="data.imgurl" />
								<sky-upimg @call-parent="setImgurl" :imgurl="data.imgurl" :trueimgurl="data.trueimgurl">
								</sky-upimg>
							</td>
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
					url: that.app.apiHost + "/admin/forum_img/add",
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
					url: that.app.apiHost + "/admin/forum_img/save",
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