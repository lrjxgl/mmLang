<template>
 	<view v-if="pageLoad">
 		<form @submit="submit">
 			<input type="text" name="siteid" class="none" v-model="data.siteid">
 			<table class="table-add">
 				<tr>
 					<td>名称：</td>
 					<td><input class="input-text" type="text" name="sitename" v-model="data.sitename"></td>
 				</tr>
 				
 				<tr>
 					<td>标题：</td>
 					<td><input class="input-text" type="text" name="title" v-model="data.title"></td>
 				</tr>
 				<tr>
 					<td>关键字：</td>
 					<td><input class="input-text" type="text" name="keywords" v-model="data.keywords"></td>
 				</tr>
 				<tr>
 					<td>描述：</td>
 					<td><input class="input-text" type="text" name="description" v-model="data.description"></td>
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
					<td>域名：</td>
					<td><input class="input-text" type="text" name="domain" v-model="data.domain"></td>
				</tr>
				<tr>
					<td>是否开启：</td>
					<td>
						<input type="text" class="none" name="is_open" v-model="data.is_open" />
						<radio-group>
							<radio type="text" class="mgr-5" checked="data.is_open==1" value="1"> 是</radio>
							<radio type="text" class="mgr-5" checked="data.is_open!=1" value="0"> 否</radio>
						</radio-group>
					</td>
				</tr>
				<tr>
					<td>关闭原因：</td>
					<td><input class="input-text" type="text" name="close_why" v-model="data.close_why"></td>
				</tr>
 				<tr>
 					<td>备案号：</td>
 					<td><input class="input-text" type="text" name="icp" v-model="data.icp"></td>
 				</tr>
 				<tr>
 					<td>状态：</td>
 					<td>
 						<input type="text" class="none" name="status" v-model="data.status" />
 						<radio-group>
 							<radio type="text" class="mgr-5" checked="data.status==1" value="1"> 上线</radio>
 							<radio type="text" checked="data.status!=1" value="2"> 下线</radio>
 						</radio-group>
 					</td>
 				</tr>
 				<tr>
 					<td>模板：</td>
 					<td><input class="input-text" type="text" name="template" v-model="data.template"></td>
 				</tr>
 				<tr>
 					<td>统计代码：</td>
 					<td><input class="input-text" type="text" name="statjs" v-model="data.statjs"></td>
 				</tr>
 				<tr>
 					<td>pc模板：</td>
 					<td><input class="input-text" type="text" name="index_template" v-model="data.index_template"></td>
 				</tr>
 				<tr>
 					<td>wap_template：</td>
 					<td><input class="input-text" type="text" name="wap_template" v-model="data.wap_template"></td>
 				</tr>
 				<tr>
 					<td>wapbg：</td>
 					<td><input class="input-text" type="text" name="wapbg" v-model="data.wapbg"></td>
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
 			setImgs: function(e) {
 				this.data.imgsdata = e;
 			},
 			setImgurl: function(e) {
 				this.data.logo = e;
 			},
 			getPage: function() {
 				var that = this;
 				that.app.get({
 					url: that.app.apiHost + "/admin/site/index",
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
 					url: that.app.apiHost + "/admin/site/save",
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
