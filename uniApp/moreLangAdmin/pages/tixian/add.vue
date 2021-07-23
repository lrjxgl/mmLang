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
						<td>类别：</td>		
						<td><input class="input-text" type="text" name="k"   v-model="data.k" ></td>
						</tr>
  <tr>
						<td>金额：</td>		
						<td><input class="input-text" type="text" name="money"   v-model="data.money" ></td>
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
						<td>创建时间：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
						</tr>
  <tr>
						<td>备注：</td>		
						<td><input class="input-text" type="text" name="info"   v-model="data.info" ></td>
						</tr>
  <tr>
						<td>回复：</td>		
						<td><input class="input-text" type="text" name="reply"   v-model="data.reply" ></td>
						</tr>
  <tr>
						<td>回复时间：</td>		
						<td><input class="input-text" type="text" name="redateline"   v-model="data.redateline" ></td>
						</tr>
  <tr>
						<td>siteid：</td>		
						<td><input class="input-text" type="text" name="siteid"   v-model="data.siteid" ></td>
						</tr>
  <tr>
						<td>银行名称：</td>		
						<td><input class="input-text" type="text" name="yhk_name"   v-model="data.yhk_name" ></td>
						</tr>
  <tr>
						<td>银行卡号码：</td>		
						<td><input class="input-text" type="text" name="yhk_haoma"   v-model="data.yhk_haoma" ></td>
						</tr>
  <tr>
						<td>银行卡户名：</td>		
						<td><input class="input-text" type="text" name="yhk_huming"   v-model="data.yhk_huming" ></td>
						</tr>
  <tr>
						<td>电话：</td>		
						<td><input class="input-text" type="text" name="telephone"   v-model="data.telephone" ></td>
						</tr>
  <tr>
						<td>开户地址：</td>		
						<td><input class="input-text" type="text" name="yhk_address"   v-model="data.yhk_address" ></td>
						</tr>
  <tr>
						<td>管理员：</td>		
						<td><input class="input-text" type="text" name="adminid"   v-model="data.adminid" ></td>
						</tr>
  <tr>
						<td>支付方式：</td>		
						<td><input class="input-text" type="text" name="paytype"   v-model="data.paytype" ></td>
						</tr>
  <tr>
						<td>银行卡：</td>		
						<td><input class="input-text" type="text" name="bankid"   v-model="data.bankid" ></td>
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
					url: that.app.apiHost + "/admin/tixian/add",
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
					url: that.app.apiHost + "/admin/tixian/save",
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