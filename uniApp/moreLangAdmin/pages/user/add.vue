<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="userid" class="none" v-model="data.userid" >
 <table class="table-add">  <tr>
						<td>用户名：</td>		
						<td><input class="input-text" type="text" name="username"   v-model="data.username" ></td>
						</tr>
  <tr>
						<td>手机：</td>		
						<td><input class="input-text" type="text" name="telephone"   v-model="data.telephone" ></td>
						</tr>
  <tr>
						<td>昵称：</td>		
						<td><input class="input-text" type="text" name="nickname"   v-model="data.nickname" ></td>
						</tr>
  <tr>
						<td>金额：</td>		
						<td><input class="input-text" type="text" name="money"   v-model="data.money" ></td>
						</tr>
  <tr>
						<td>金币：</td>		
						<td><input class="input-text" type="text" name="gold"   v-model="data.gold" ></td>
						</tr>
  <tr>
						<td>积分：</td>		
						<td><input class="input-text" type="text" name="grade"   v-model="data.grade" ></td>
						</tr>
  <tr>
						<td>创建时间：</td>		
						<td><input class="input-text" type="text" name="createtime"   v-model="data.createtime" ></td>
						</tr>
  <tr>
						<td>更新时间：</td>		
						<td><input class="input-text" type="text" name="lasttime"   v-model="data.lasttime" ></td>
						</tr>
  <tr>
						<td>用户类型：</td>		
						<td><input class="input-text" type="text" name="user_type"   v-model="data.user_type" ></td>
						</tr>
  <tr>
						<td>状态：</td>		
						<td>
						<input type="text" class="none" name="status" v-model="data.status" />
						<radio-group>
							<radio type="text" class="mgr-5" checked="data.status==1"  value="1" > 上线</radio>
							<radio type="text" checked="data.status!=1"    value="2" > 下线</radio>
						</radio-group>	
						</td>
						</tr>
  <tr>
						<td>实名认证：</td>		
						<td><input class="input-text" type="text" name="is_auth"   v-model="data.is_auth" ></td>
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
						<td>关注数：</td>		
						<td><input class="input-text" type="text" name="follow_num"   v-model="data.follow_num" ></td>
						</tr>
  <tr>
						<td>粉丝数：</td>		
						<td><input class="input-text" type="text" name="followed_num"   v-model="data.followed_num" ></td>
						</tr>
  <tr>
						<td>性别：</td>		
						<td><input class="input-text" type="text" name="gender"   v-model="data.gender" ></td>
						</tr>
  <tr>
						<td>邀请人：</td>		
						<td><input class="input-text" type="text" name="invite_userid"   v-model="data.invite_userid" ></td>
						</tr>
  <tr>
						<td>出生日期：</td>		
						<td><input class="input-text" type="text" name="birthday"   v-model="data.birthday" ></td>
						</tr>
  <tr>
						<td>description：</td>		
						<td><input class="input-text" type="text" name="description"   v-model="data.description" ></td>
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
					url: that.app.apiHost + "/admin/user/add",
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
					url: that.app.apiHost + "/admin/user/save",
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