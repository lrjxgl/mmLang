<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>userid：</td>		
						<td><input class="input-text" type="text" name="userid"   v-model="data.userid" ></td>
						</tr>
  <tr>
						<td>收货地址：</td>		
						<td><input class="input-text" type="text" name="address"   v-model="data.address" ></td>
						</tr>
  <tr>
						<td>电话：</td>		
						<td><input class="input-text" type="text" name="telephone"   v-model="data.telephone" ></td>
						</tr>
  <tr>
						<td>收货人：</td>		
						<td><input class="input-text" type="text" name="truename"   v-model="data.truename" ></td>
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
						<td>邮编：</td>		
						<td><input class="input-text" type="text" name="zip_code"   v-model="data.zip_code" ></td>
						</tr>
  <tr>
						<td>lastid：</td>		
						<td><input class="input-text" type="text" name="lastid"   v-model="data.lastid" ></td>
						</tr>
  <tr>
						<td>省：</td>		
						<td><input class="input-text" type="text" name="province_id"   v-model="data.province_id" ></td>
						</tr>
  <tr>
						<td>市：</td>		
						<td><input class="input-text" type="text" name="city_id"   v-model="data.city_id" ></td>
						</tr>
  <tr>
						<td>县：</td>		
						<td><input class="input-text" type="text" name="town_id"   v-model="data.town_id" ></td>
						</tr>
  <tr>
						<td>默认：</td>		
						<td><input class="input-text" type="text" name="isdefault"   v-model="data.isdefault" ></td>
						</tr>
  <tr>
						<td>创建：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
						</tr>
  <tr>
						<td>省市地址：</td>		
						<td><input class="input-text" type="text" name="pct_address"   v-model="data.pct_address" ></td>
						</tr>
  <tr>
						<td>纬度：</td>		
						<td><input class="input-text" type="text" name="lat"   v-model="data.lat" ></td>
						</tr>
  <tr>
						<td>经度：</td>		
						<td><input class="input-text" type="text" name="lng"   v-model="data.lng" ></td>
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
					url: that.app.apiHost + "/admin/user_address/add",
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
					url: that.app.apiHost + "/admin/user_address/save",
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