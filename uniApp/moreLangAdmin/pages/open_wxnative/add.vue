<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>名称：</td>		
						<td><input class="input-text" type="text" name="title"   v-model="data.title" ></td>
						</tr>
  <tr>
						<td>创建时间：</td>		
						<td><input class="input-text" type="text" name="createtime"   v-model="data.createtime" ></td>
						</tr>
  <tr>
						<td>appid：</td>		
						<td><input class="input-text" type="text" name="appid"   v-model="data.appid" ></td>
						</tr>
  <tr>
						<td>appkey：</td>		
						<td><input class="input-text" type="text" name="appkey"   v-model="data.appkey" ></td>
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
						<td>支付cer证书：</td>		
						<td><input class="input-text" type="text" name="sslcert_path"   v-model="data.sslcert_path" ></td>
						</tr>
  <tr>
						<td>支付key证书：</td>		
						<td><input class="input-text" type="text" name="sslkey_path"   v-model="data.sslkey_path" ></td>
						</tr>
  <tr>
						<td>商户号：</td>		
						<td><input class="input-text" type="text" name="mchid"   v-model="data.mchid" ></td>
						</tr>
  <tr>
						<td>商户支付密钥：</td>		
						<td><input class="input-text" type="text" name="mchkey"   v-model="data.mchkey" ></td>
						</tr>
  <tr>
						<td>openlogin：</td>		
						<td><input class="input-text" type="text" name="openlogin"   v-model="data.openlogin" ></td>
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
					url: that.app.apiHost + "/admin/open_wxnative/add",
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
					url: that.app.apiHost + "/admin/open_wxnative/save",
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