<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
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
						<td>openid：</td>		
						<td><input class="input-text" type="text" name="openid"   v-model="data.openid" ></td>
						</tr>
  <tr>
						<td>msgtype：</td>		
						<td><input class="input-text" type="text" name="msgtype"   v-model="data.msgtype" ></td>
						</tr>
  <tr>
						<td>createtime：</td>		
						<td><input class="input-text" type="text" name="createtime"   v-model="data.createtime" ></td>
						</tr>
  <tr>
						<td>content：</td>		
						<td><input class="input-text" type="text" name="content"   v-model="data.content" ></td>
						</tr>
  <tr>
						<td>msgid：</td>		
						<td><input class="input-text" type="text" name="msgid"   v-model="data.msgid" ></td>
						</tr>
  <tr>
						<td>picurl：</td>		
						<td><input class="input-text" type="text" name="picurl"   v-model="data.picurl" ></td>
						</tr>
  <tr>
						<td>mediaid：</td>		
						<td><input class="input-text" type="text" name="mediaid"   v-model="data.mediaid" ></td>
						</tr>
  <tr>
						<td>thumbmediaid：</td>		
						<td><input class="input-text" type="text" name="thumbmediaid"   v-model="data.thumbmediaid" ></td>
						</tr>
  <tr>
						<td>format：</td>		
						<td><input class="input-text" type="text" name="format"   v-model="data.format" ></td>
						</tr>
  <tr>
						<td>location_x：</td>		
						<td><input class="input-text" type="text" name="location_x"   v-model="data.location_x" ></td>
						</tr>
  <tr>
						<td>location_y：</td>		
						<td><input class="input-text" type="text" name="location_y"   v-model="data.location_y" ></td>
						</tr>
  <tr>
						<td>scale：</td>		
						<td><input class="input-text" type="text" name="scale"   v-model="data.scale" ></td>
						</tr>
  <tr>
						<td>label：</td>		
						<td><input class="input-text" type="text" name="label"   v-model="data.label" ></td>
						</tr>
  <tr>
						<td>title：</td>		
						<td><input class="input-text" type="text" name="title"   v-model="data.title" ></td>
						</tr>
  <tr>
						<td>description：</td>		
						<td><input class="input-text" type="text" name="description"   v-model="data.description" ></td>
						</tr>
  <tr>
						<td>url：</td>		
						<td><input class="input-text" type="text" name="url"   v-model="data.url" ></td>
						</tr>
  <tr>
						<td>wid：</td>		
						<td><input class="input-text" type="text" name="wid"   v-model="data.wid" ></td>
						</tr>
  <tr>
						<td>shopid：</td>		
						<td><input class="input-text" type="text" name="shopid"   v-model="data.shopid" ></td>
						</tr>
  <tr>
						<td>fromusername：</td>		
						<td><input class="input-text" type="text" name="fromusername"   v-model="data.fromusername" ></td>
						</tr>
  <tr>
						<td>tousername：</td>		
						<td><input class="input-text" type="text" name="tousername"   v-model="data.tousername" ></td>
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
					url: that.app.apiHost + "/admin/weixin_reply/add",
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
					url: that.app.apiHost + "/admin/weixin_reply/save",
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