<template>
	<view>
	<div class="tabs-border">
			<div @click="gourl('index')" class="item">列表</div>
			<div @click="gourl('add')" class="item active">添加</div>
		</div>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>签到类型：</td>		
						<td><input class="input-text" type="text" name="type_id"   v-model="data.type_id" ></td>
						</tr>
  <tr>
						<td>用户id：</td>		
						<td><input class="input-text" type="text" name="userid"   v-model="data.userid" ></td>
						</tr>
  <tr>
						<td>时间：</td>		
						<td><input class="input-text" type="text" name="dateline"   v-model="data.dateline" ></td>
						</tr>
  <tr>
						<td>积分：</td>		
						<td><input class="input-text" type="text" name="grade"   v-model="data.grade" ></td>
						</tr>
  <tr>
						<td>last_day：</td>		
						<td><input class="input-text" type="text" name="last_day"   v-model="data.last_day" ></td>
						</tr>
  <tr>
						<td>签到ip：</td>		
						<td><input class="input-text" type="text" name="last_ip"   v-model="data.last_ip" ></td>
						</tr>
  <tr>
						<td>签到总次数：</td>		
						<td><input class="input-text" type="text" name="all_times"   v-model="data.all_times" ></td>
						</tr>
  <tr>
						<td>1有效：</td>		
						<td><input class="input-text" type="text" name="is_valid"   v-model="data.is_valid" ></td>
						</tr>
  <tr>
						<td>金币：</td>		
						<td><input class="input-text" type="text" name="gold"   v-model="data.gold" ></td>
						</tr>
  <tr>
						<td>连续签到次数：</td>		
						<td><input class="input-text" type="text" name="days"   v-model="data.days" ></td>
						</tr>
  <tr>
						<td>分站id：</td>		
						<td><input class="input-text" type="text" name="siteid"   v-model="data.siteid" ></td>
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
					url: that.app.apiHost + "/admin/checkin_index/add",
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
					url: that.app.apiHost + "/admin/checkin_index/save",
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