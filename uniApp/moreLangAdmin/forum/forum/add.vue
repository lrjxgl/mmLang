<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>title：</td>		
						<td><input class="input-text" type="text" name="title"   v-model="data.title" ></td>
						</tr>
  <tr>
						<td>userid：</td>		
						<td><input class="input-text" type="text" name="userid"   v-model="data.userid" ></td>
						</tr>
  <tr>
						<td>gid：</td>		
						<td><input class="input-text" type="text" name="gid"   v-model="data.gid" ></td>
						</tr>
  <tr>
						<td>catid：</td>		
						<td><input class="input-text" type="text" name="catid"   v-model="data.catid" ></td>
						</tr>
  <tr>
						<td>love_num：</td>		
						<td><input class="input-text" type="text" name="love_num"   v-model="data.love_num" ></td>
						</tr>
  <tr>
						<td>fav_num：</td>		
						<td><input class="input-text" type="text" name="fav_num"   v-model="data.fav_num" ></td>
						</tr>
  <tr>
						<td>forward_num：</td>		
						<td><input class="input-text" type="text" name="forward_num"   v-model="data.forward_num" ></td>
						</tr>
  <tr>
						<td>keywords：</td>		
						<td><input class="input-text" type="text" name="keywords"   v-model="data.keywords" ></td>
						</tr>
  <tr>
						<td>description：</td>		
						<td><input class="input-text" type="text" name="description"   v-model="data.description" ></td>
						</tr>
  <tr>
						<td>dateline：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
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
						<td>comment_num：</td>		
						<td><input class="input-text" type="text" name="comment_num"   v-model="data.comment_num" ></td>
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
						<td>last_time：</td>		
						<td><input class="input-text" type="text" name="last_time"   v-model="data.last_time" ></td>
						</tr>
  <tr>
						<td>grade：</td>		
						<td><input class="input-text" type="text" name="grade"   v-model="data.grade" ></td>
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
						<td>view_num：</td>		
						<td><input class="input-text" type="text" name="view_num"   v-model="data.view_num" ></td>
						</tr>
  <tr>
						<td>isnew：</td>		
						<td>
						<input type="text" class="none" name="isnew" v-model="data.isnew" />
						<radio-group>
							<radio type="text" class="mgr-5" checked="data.isnew==1"  value="1" > 是</radio>
							<radio type="text" class="mgr-5" checked="data.isnew!=1"    value="0" > 否</radio>
						</radio-group>	
						</td>
						</tr>
  <tr>
						<td>tags：</td>		
						<td><input class="input-text" type="text" name="tags"   v-model="data.tags" ></td>
						</tr>
  <tr>
						<td>videourl：</td>		
						<td><input class="input-text" type="text" name="videourl"   v-model="data.videourl" ></td>
						</tr>
  <tr>
						<td>money：</td>		
						<td><input class="input-text" type="text" name="money"   v-model="data.money" ></td>
						</tr>
  <tr>
						<td>奖励：</td>		
						<td><input class="input-text" type="text" name="gold"   v-model="data.gold" ></td>
						</tr>
  <tr>
						<td>imgsdata：</td>		
						<td><input class="input-text" type="text" name="imgsdata"   v-model="data.imgsdata" ></td>
						</tr>
  <tr>
						<td>createtime：</td>		
						<td><input class="input-text" type="text" name="createtime"   v-model="data.createtime" ></td>
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
					url: that.app.apiHost + "/admin/forum/add",
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
					url: that.app.apiHost + "/admin/forum/save",
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