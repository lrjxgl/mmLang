<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>用户：</td>		
						<td><input class="input-text" type="text" name="userid"   v-model="data.userid" ></td>
						</tr>
  <tr>
						<td>平台类型：</td>		
						<td><input class="input-text" type="text" name="xfrom"   v-model="data.xfrom" ></td>
						</tr>
  <tr>
						<td>三方用户：</td>		
						<td><input class="input-text" type="text" name="openid"   v-model="data.openid" ></td>
						</tr>
  <tr>
						<td>三方授权：</td>		
						<td><input class="input-text" type="text" name="accesstoken"   v-model="data.accesstoken" ></td>
						</tr>
  <tr>
						<td>createtime：</td>		
						<td><input class="input-text" type="text" name="createtime"   v-model="data.createtime" ></td>
						</tr>
  <tr>
						<td>昵称：</td>		
						<td><input class="input-text" type="text" name="nickname"   v-model="data.nickname" ></td>
						</tr>

					<tr>
							<td>图片：</td>
							<td>
								<input class="none" type="text" name="user_head" v-model="data.user_head" />
								<sky-upimg @call-parent="setImgurl" :imgurl="data.user_head" :trueimgurl="data.trueuser_head">
								</sky-upimg>
							</td>
						</tr>
				
  <tr>
						<td>性别：</td>		
						<td><input class="input-text" type="text" name="gender"   v-model="data.gender" ></td>
						</tr>
  <tr>
						<td>unionid：</td>		
						<td><input class="input-text" type="text" name="unionid"   v-model="data.unionid" ></td>
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
					url: that.app.apiHost + "/admin/openlogin/add",
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
					url: that.app.apiHost + "/admin/openlogin/save",
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