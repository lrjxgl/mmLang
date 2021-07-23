<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>token：</td>		
						<td><input class="input-text" type="text" name="token"   v-model="data.token" ></td>
						</tr>
  <tr>
						<td>title：</td>		
						<td><input class="input-text" type="text" name="title"   v-model="data.title" ></td>
						</tr>
  <tr>
						<td>dateline：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
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
						<td>userid：</td>		
						<td><input class="input-text" type="text" name="userid"   v-model="data.userid" ></td>
						</tr>
  <tr>
						<td>domain：</td>		
						<td><input class="input-text" type="text" name="domain"   v-model="data.domain" ></td>
						</tr>
  <tr>
						<td>catid：</td>		
						<td><input class="input-text" type="text" name="catid"   v-model="data.catid" ></td>
						</tr>

					<tr>
							<td>图片：</td>
							<td>
								<input class="none" type="text" name="logo" v-model="data.logo" />
								<sky-upimg @call-parent="setImgurl" :imgurl="data.logo" :trueimgurl="data.truelogo">
								</sky-upimg>
							</td>
						</tr>
				

					<tr>
							<td>图片：</td>
							<td>
								<input class="none" type="text" name="imgurl" v-model="data.imgurl" />
								<sky-upimg @call-parent="setImgurl" :imgurl="data.imgurl" :trueimgurl="data.trueimgurl">
								</sky-upimg>
							</td>
						</tr>
				
  <tr>
						<td>isrecommend：</td>		
						<td>
						<input type="text" class="none" name="isrecommend" v-model="data.isrecommend" />
						<radio-group>
							<radio type="text" class="mgr-5" checked="data.isrecommend==1"  value="1" > 是</radio>
							<radio type="text" class="mgr-5" checked="data.isrecommend!=1"    value="0" > 否</radio>
						</radio-group>	
						</td>
						</tr>
  <tr>
						<td>imgsdata：</td>		
						<td><input class="input-text" type="text" name="imgsdata"   v-model="data.imgsdata" ></td>
						</tr>
  <tr>
						<td>appid：</td>		
						<td><input class="input-text" type="text" name="appid"   v-model="data.appid" ></td>
						</tr>
  <tr>
						<td>公众帐号secert：</td>		
						<td><input class="input-text" type="text" name="appkey"   v-model="data.appkey" ></td>
						</tr>
  <tr>
						<td>isshow：</td>		
						<td><input class="input-text" type="text" name="isshow"   v-model="data.isshow" ></td>
						</tr>
  <tr>
						<td>ysid：</td>		
						<td><input class="input-text" type="text" name="ysid"   v-model="data.ysid" ></td>
						</tr>
  <tr>
						<td>shopid：</td>		
						<td><input class="input-text" type="text" name="shopid"   v-model="data.shopid" ></td>
						</tr>
  <tr>
						<td>wx_username：</td>		
						<td><input class="input-text" type="text" name="wx_username"   v-model="data.wx_username" ></td>
						</tr>
  <tr>
						<td>wx_pwd：</td>		
						<td><input class="input-text" type="text" name="wx_pwd"   v-model="data.wx_pwd" ></td>
						</tr>
  <tr>
						<td>enkey：</td>		
						<td><input class="input-text" type="text" name="enkey"   v-model="data.enkey" ></td>
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
						<td>支付cer证书：</td>		
						<td><input class="input-text" type="text" name="sslcert_path"   v-model="data.sslcert_path" ></td>
						</tr>
  <tr>
						<td>支付key证书：</td>		
						<td><input class="input-text" type="text" name="sslkey_path"   v-model="data.sslkey_path" ></td>
						</tr>
  <tr>
						<td>0直接登录：</td>		
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
					url: that.app.apiHost + "/admin/weixin/add",
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
					url: that.app.apiHost + "/admin/weixin/save",
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