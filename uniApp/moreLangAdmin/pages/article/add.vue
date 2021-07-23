<template>
	<view>
		<form v-if="pageLoad">

			<input type="hidden" name="id" style="display:none;" value="{$data.id}">
			<div class="tabs-box">
				<div class="js-tabs tabs-border">
					<a class="item active" href="#base">基本信息</a>

					<a class="item" href="#other">扩展</a>
				</div>
				<div class="tabs-item active" id="base">
					<table class="table-add">
						<tr>
							<td>标题：</td>
							<td><input class="input-text"  type="text" name="title" id="title" v-model="data.title"></td>
						</tr>
						<tr>
							<td>分类：</td>
							<td>
								<select name="catid" id="catid" class="input-flex-select w150">
									<option value="0">请选择</option>
									{foreach item=c from=$catlist}
									<option value="{$c.catid}" {if $data.catid eq $c.catid} selected="selected" {/if}>
										{$c.cname}</option>
									{foreach item=c_2 from=$c.child}
									<option value="{$c_2.catid}" {if $data.catid eq $c_2.catid} selected="selected"
										{/if} class="o_c_2">|__{$c_2.cname}</option>
									{foreach item=c_3 from=$c_2.child}
									<option value="{$c_3.catid}" {if $data.catid eq $c_3.catid} selected="selected"
										{/if} class="o_c_3"> |____{$c_3.cname}</option>
									{/foreach}
									{/foreach}
									{/foreach}
								</select>
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
							<td>作者</td>
							<td>
								<input type="text" name="author" v-model="data.author" />
							</td>
						</tr>
					 
							<tr v-if="data.id>0">
								<td>喜欢数：</td>
								<td><input type="text" name="love_num" id="love_num" v-model="data.love_num"></td>
							</tr>
							<tr v-if="data.id>0">
								<td>收藏：</td>
								<td><input type="text" name="fav_num" id="fav_num" v-model="data.fav_num"></td>
						  
						<tr>
							<td>描述：</td>
							<td>
								<textarea name="description" id="description"
									class="w600 h60">{$data.description}</textarea>
							</td>
						</tr>

						<tr>
							<td>推荐：</td>
							<td>
								<radio-group>
									<radio name="is_recommend" value="1" :checked="data.is_recommend==1"/> 推荐
									&nbsp;&nbsp;
									<radio  name="is_recommend" value="0" :checked="data.is_recommend!=1" /> 不推荐
								</radio-group>
								

							</td>
						</tr>
						<tr v-if="data.id>0">
							<td>访问数：</td>
							<td>{{data.view_num}}</td>
						</tr>



						<tr>
							<td>详情模板：</td>
							<td><input class="input-text" type="text" name="tpl" id="tpl" v-model="data.tpl">(如果需要可以指定模板)
							</td>
						</tr>
						<tr v-if="data.id>0">
							<td>创建时间：</td>
							<td>{{data.createtime}}</td>
						</tr>
						<tr>
							<td>内容</td>
							<td>
								<script type="text/plain" id="content" name="content">{$data.content}</script>
							</td>
						</tr>
					</table>
				</div>

				<div class="tabs-item active" id="other">
					<table class="table-add">
						<tr>
							<td class="w90">产品价格：</td>
							<td><input class="input-text" type="text" name="price" id="price" v-model="data.price"></td>
						</tr>
						<tr>
							<td>市场价格：</td>
							<td><input class="input-text" type="text" name="market_price" id="market_price" v-model="data.market_price">
							</td>
						</tr>
						<tr>
							<td>库存数：</td>
							<td><input class="input-text" type="text" name="total_num" id="total_num" v-model="data.total_num"></td>
						</tr>
						<tr>
							<td>销售数：</td>
							<td><input class="input-text" type="text" name="sold_num" id="sold_num" v-model="data.sold_num"></td>
						</tr>
						<tr>
							<td>积分：</td>
							<td><input class="input-text" type="text" name="grade" id="sold_num" v-model="data.grade"></td>
						</tr>

						<tr>
							<td>图集</td>
							<td>
								<input type="text" name="imgsdata" class="none" v-model="data.imgsdata" />
								<upimg-box :defaultImgsList="imgsList" :defaultImgsData="data.imgsdata"
									@call-parent="setImgs"></upimg-box>
							</td>
						</tr>

					</table>
				</div>
			</div>
			<div class="btn-row-submit js-submit">保存</div>
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
				this.data.imgurl = e;
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/admin/article/add",
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
					url: that.app.apiHost + "/admin/article/save",
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
