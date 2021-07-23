<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>openid：</td>		
						<td><input class="input-text" type="text" name="openid"   v-model="data.openid" ></td>
						</tr>
  <tr>
						<td>dateline：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
						</tr>
  <tr>
						<td>add_time：</td>		
						<td><input class="input-text" type="text" name="add_time"   v-model="data.add_time" ></td>
						</tr>
  <tr>
						<td>last_time：</td>		
						<td><input class="input-text" type="text" name="last_time"   v-model="data.last_time" ></td>
						</tr>
  <tr>
						<td>del_time：</td>		
						<td><input class="input-text" type="text" name="del_time"   v-model="data.del_time" ></td>
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
						<td>nickname：</td>		
						<td><input class="input-text" type="text" name="nickname"   v-model="data.nickname" ></td>
						</tr>
  <tr>
						<td>sex：</td>		
						<td><input class="input-text" type="text" name="sex"   v-model="data.sex" ></td>
						</tr>
  <tr>
						<td>city：</td>		
						<td><input class="input-text" type="text" name="city"   v-model="data.city" ></td>
						</tr>
  <tr>
						<td>country：</td>		
						<td><input class="input-text" type="text" name="country"   v-model="data.country" ></td>
						</tr>
  <tr>
						<td>province：</td>		
						<td><input class="input-text" type="text" name="province"   v-model="data.province" ></td>
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
						<td>update_time：</td>		
						<td><input class="input-text" type="text" name="update_time"   v-model="data.update_time" ></td>
						</tr>
  <tr>
						<td>reply_num：</td>		
						<td><input class="input-text" type="text" name="reply_num"   v-model="data.reply_num" ></td>
						</tr>
  <tr>
						<td>wid：</td>		
						<td><input class="input-text" type="text" name="wid"   v-model="data.wid" ></td>
						</tr>
  <tr>
						<td>shopid：</td>		
						<td><input class="input-text" type="text" name="shopid"   v-model="data.shopid" ></td>
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
					url: that.app.apiHost + "/admin/weixin_user/add",
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
					url: that.app.apiHost + "/admin/weixin_user/save",
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