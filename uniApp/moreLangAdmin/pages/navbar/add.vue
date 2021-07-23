<template>
	<view>
		<nav-tab tab="add"></nav-tab>
		<form @submit="submit">
			<input type="text" name="id" class="none" v-model="data.id">
			<table class="table-add">
				<tr>
					<td>位置：</td>
					<td><select name="group_id" v-model="group_id">
						<option value="0">请选择</option> 
						<option v-for="(v,k) in groupList" :key="k" :value="v.gid">{{v.title}}</option>
						</select>
					</td>
				</tr>
				<tr>
					<td>上级：</td>
					<td><select name="pid" v-model="pid">
						<option value="0">作为父级</option>
						<option v-for="(v,k) in parentList" :key="k" :value="v.id">{{v.title}}</option>
						</select>
					</td>
				</tr>
				<tr>
					<td>名称：</td>
					<td><input class="input-text" type="text" name="title" v-model="data.title"></td>
				</tr>
				
				<tr>
					<td>链接地址：</td>
					<td><input class="input-text" type="text" name="link_url" v-model="data.link_url"></td>
				</tr>
				<tr>
					<td>跳转目标：</td>
					<td>
						<select name="target" v-model="data.target" >
								<option value="_self">self</option>
								<option value="_blank">_blank </option>
						</select>
					</td>
				</tr>
				
				 
				<tr>
					<td>m：</td>
					<td><input class="input-text" type="text" name="m" v-model="data.m"></td>
				</tr>
				<tr>
					<td>a：</td>
					<td><input class="input-text" type="text" name="a" v-model="data.a"></td>
				</tr>
				<tr>
					<td>状态：</td>
					<td>
						<input type="text" class="none" name="status" v-model="data.status" />
						<radio-group @change="changeStatus">
							<radio type="text" class="mgr-5" checked="data.status==1" value="1"> 上线</radio>
							<radio type="text" checked="data.status!=1" value="2"> 下线</radio>
						</radio-group>
					</td>
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
					<td>图标：</td>
					<td><input class="input-text" type="text" name="icon" v-model="data.icon"></td>
				</tr>
				<tr>
					<td>排序：</td>
					<td><input class="input-text" type="text" name="orderindex" v-model="data.orderindex"></td>
				</tr>
			</table> <button form-type="submit" class="btn-row-submit">保存</button>
		</form>
	</view>
</template>

<script>
	import skyUpimg from "../../components/skyupimg.vue";
	import navTab from "./nav-tab.vue"
	export default {
		components: {
			navTab,
			skyUpimg
		},
		data: function() {
			return {
				id: 0,
				pid:0,
				pageLoad: false,
				data: {},
				imgsList: [],
				parentList:[],
				groupList:[],
				group_id:0
			}
		},
		onLoad: function(ops) {
			if (ops.id != undefined) {
				this.id = ops.id;
			}
			if(ops.pid!=undefined){
				this.pid=ops.pid;
			}
			console.log(this.$data)
			this.getPage();
		},
		methods: {
			changeStatus:function(e){
				this.data.status=e.detail.value;
			},
			setImgs: function(e) {
				this.data.imgsdata = e;
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/admin/navbar/add",
					data: {
						id: this.id,
						pid:this.pid
					},
					success: function(res) {
						if (res.data) {
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
						that.parentList=res.parentList;
						that.groupList=res.groupList;
						that.group_id=res.group_id;
						that.pid=res.pid;
						that.pageLoad = true;
					}
				})
			},
			submit: function(e) {
				var that = this;
				var formValue=e.detail.value;
				formValue.pid=this.pid;
				formValue.group_id=this.group_id;
				that.app.post({
					url: that.app.apiHost + "/admin/navbar/save",
					data: e.detail.value,
					success: function(res) {
						uni.showToast({
							title: res.message
						})
						if(!res.error){
							uni.navigateBack()
						}
					}
				})

			}
			
		}
	}
</script>

<style>
</style>
