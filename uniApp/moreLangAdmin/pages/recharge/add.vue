<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>siteid：</td>		
						<td><input class="input-text" type="text" name="siteid"   v-model="data.siteid" ></td>
						</tr>
  <tr>
						<td>用户：</td>		
						<td><input class="input-text" type="text" name="userid"   v-model="data.userid" ></td>
						</tr>
  <tr>
						<td>金额：</td>		
						<td><input class="input-text" type="text" name="money"   v-model="data.money" ></td>
						</tr>
  <tr>
						<td>支付方式：</td>		
						<td><input class="input-text" type="text" name="pay_type"   v-model="data.pay_type" ></td>
						</tr>
  <tr>
						<td>用户标识：</td>		
						<td><input class="input-text" type="text" name="openid"   v-model="data.openid" ></td>
						</tr>
  <tr>
						<td>支付的订单的订单号：</td>		
						<td><input class="input-text" type="text" name="pay_orderno"   v-model="data.pay_orderno" ></td>
						</tr>
  <tr>
						<td>类型：</td>		
						<td><input class="input-text" type="text" name="type_id"   v-model="data.type_id" ></td>
						</tr>
  <tr>
						<td>创建时间：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
						</tr>
  <tr>
						<td>1支付成功：</td>		
						<td>
						<input type="text" class="none" name="status" v-model="data.status" />
						<radio-group>
							<radio type="text" class="mgr-5" checked="data.status==1"  value="1" > 上线</radio>
							<radio type="text" checked="data.status!=1"    value="2" > 下线</radio>
						</radio-group>	
						</td>
						</tr>
  <tr>
						<td>订单号：</td>		
						<td><input class="input-text" type="text" name="orderno"   v-model="data.orderno" ></td>
						</tr>
  <tr>
						<td>订单信息：</td>		
						<td><input class="input-text" type="text" name="orderinfo"   v-model="data.orderinfo" ></td>
						</tr>
  <tr>
						<td>表名称：</td>		
						<td><input class="input-text" type="text" name="tablename"   v-model="data.tablename" ></td>
						</tr>
  <tr>
						<td>订单处理内容：</td>		
						<td><input class="input-text" type="text" name="orderdata"   v-model="data.orderdata" ></td>
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
					url: that.app.apiHost + "/admin/recharge/add",
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
					url: that.app.apiHost + "/admin/recharge/save",
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