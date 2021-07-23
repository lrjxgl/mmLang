<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="userid" class="none" v-model="data.userid" >
 <table class="table-add">  <tr>
						<td>真实姓名：</td>		
						<td><input class="input-text" type="text" name="truename"   v-model="data.truename" ></td>
						</tr>
  <tr>
						<td>身份证号码：</td>		
						<td><input class="input-text" type="text" name="user_card"   v-model="data.user_card" ></td>
						</tr>
  <tr>
						<td>全身照：</td>		
						<td><input class="input-text" type="text" name="true_user_head"   v-model="data.true_user_head" ></td>
						</tr>
  <tr>
						<td>收入：</td>		
						<td><input class="input-text" type="text" name="income"   v-model="data.income" ></td>
						</tr>
  <tr>
						<td>创建时间：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
						</tr>
  <tr>
						<td>更新时间：</td>		
						<td><input class="input-text" type="text" name="lasttime"   v-model="data.lasttime" ></td>
						</tr>
  <tr>
						<td>手机号码：</td>		
						<td><input class="input-text" type="text" name="telephone"   v-model="data.telephone" ></td>
						</tr>
  <tr>
						<td>is_auth：</td>		
						<td><input class="input-text" type="text" name="is_auth"   v-model="data.is_auth" ></td>
						</tr>
  <tr>
						<td>status：</td>		
						<td>
						<input type="text" class="none" name="status" v-model="data.status" />
						<radio-group>
							<radio type="text" class="mgr-5" checked="data.status==1"  value="1" > 上线</radio>
							<radio type="text" checked="data.status!=1"    value="2" > 下线</radio>
						</radio-group>	
						</td>
						</tr>
  <tr>
						<td>详细信息：</td>		
						<td><input class="input-text" type="text" name="content"   v-model="data.content" ></td>
						</tr>
  <tr>
						<td>简介：</td>		
						<td><input class="input-text" type="text" name="info"   v-model="data.info" ></td>
						</tr>
  <tr>
						<td>简介：</td>		
						<td><input class="input-text" type="text" name="sky_info"   v-model="data.sky_info" ></td>
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
					url: that.app.apiHost + "/admin/user_auth/add",
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
					url: that.app.apiHost + "/admin/user_auth/save",
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