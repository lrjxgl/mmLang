 <template>
	<view>
<form @submit="submit">
<input type="text" name="id" class="none" v-model="data.id" >
 <table class="table-add">  <tr>
						<td>siteid：</td>		
						<td><input class="input-text" type="text" name="siteid"   v-model="data.siteid" ></td>
						</tr>
  <tr>
						<td>phone_on：</td>		
						<td><input class="input-text" type="text" name="phone_on"   v-model="data.phone_on" ></td>
						</tr>
  <tr>
						<td>phone_reg：</td>		
						<td><input class="input-text" type="text" name="phone_reg"   v-model="data.phone_reg" ></td>
						</tr>
  <tr>
						<td>phone_user：</td>		
						<td><input class="input-text" type="text" name="phone_user"   v-model="data.phone_user" ></td>
						</tr>
  <tr>
						<td>phone_pwd：</td>		
						<td><input class="input-text" type="text" name="phone_pwd"   v-model="data.phone_pwd" ></td>
						</tr>
  <tr>
						<td>phone_num：</td>		
						<td><input class="input-text" type="text" name="phone_num"   v-model="data.phone_num" ></td>
						</tr>
  <tr>
						<td>smtphost：</td>		
						<td><input class="input-text" type="text" name="smtphost"   v-model="data.smtphost" ></td>
						</tr>
  <tr>
						<td>smtpport：</td>		
						<td><input class="input-text" type="text" name="smtpport"   v-model="data.smtpport" ></td>
						</tr>
  <tr>
						<td>smtpemail：</td>		
						<td><input class="input-text" type="text" name="smtpemail"   v-model="data.smtpemail" ></td>
						</tr>
  <tr>
						<td>smtpuser：</td>		
						<td><input class="input-text" type="text" name="smtpuser"   v-model="data.smtpuser" ></td>
						</tr>
  <tr>
						<td>smtppwd：</td>		
						<td><input class="input-text" type="text" name="smtppwd"   v-model="data.smtppwd" ></td>
						</tr>
  <tr>
						<td>water_on：</td>		
						<td><input class="input-text" type="text" name="water_on"   v-model="data.water_on" ></td>
						</tr>
  <tr>
						<td>water_type：</td>		
						<td><input class="input-text" type="text" name="water_type"   v-model="data.water_type" ></td>
						</tr>
  <tr>
						<td>water_pos：</td>		
						<td><input class="input-text" type="text" name="water_pos"   v-model="data.water_pos" ></td>
						</tr>
  <tr>
						<td>water_str：</td>		
						<td><input class="input-text" type="text" name="water_str"   v-model="data.water_str" ></td>
						</tr>
  <tr>
						<td>water_img：</td>		
						<td><input class="input-text" type="text" name="water_img"   v-model="data.water_img" ></td>
						</tr>
  <tr>
						<td>water_size：</td>		
						<td><input class="input-text" type="text" name="water_size"   v-model="data.water_size" ></td>
						</tr>
  <tr>
						<td>rewrite_on：</td>		
						<td><input class="input-text" type="text" name="rewrite_on"   v-model="data.rewrite_on" ></td>
						</tr>
  <tr>
						<td>spread_on：</td>		
						<td><input class="input-text" type="text" name="spread_on"   v-model="data.spread_on" ></td>
						</tr>
  <tr>
						<td>spread_discount：</td>		
						<td><input class="input-text" type="text" name="spread_discount"   v-model="data.spread_discount" ></td>
						</tr>
  <tr>
						<td>spread_money：</td>		
						<td><input class="input-text" type="text" name="spread_money"   v-model="data.spread_money" ></td>
						</tr>
  <tr>
						<td>spread_type：</td>		
						<td><input class="input-text" type="text" name="spread_type"   v-model="data.spread_type" ></td>
						</tr>
  <tr>
						<td>grade_on：</td>		
						<td><input class="input-text" type="text" name="grade_on"   v-model="data.grade_on" ></td>
						</tr>
  <tr>
						<td>hotsearch：</td>		
						<td><input class="input-text" type="text" name="hotsearch"   v-model="data.hotsearch" ></td>
						</tr>
  <tr>
						<td>kf_qq：</td>		
						<td><input class="input-text" type="text" name="kf_qq"   v-model="data.kf_qq" ></td>
						</tr>
  <tr>
						<td>kf_ww：</td>		
						<td><input class="input-text" type="text" name="kf_ww"   v-model="data.kf_ww" ></td>
						</tr>
  <tr>
						<td>短信签名：</td>		
						<td><input class="input-text" type="text" name="sms_qianming"   v-model="data.sms_qianming" ></td>
						</tr>
  <tr>
						<td>短信类型：</td>		
						<td><input class="input-text" type="text" name="sms_type"   v-model="data.sms_type" ></td>
						</tr>
  <tr>
						<td>s_bank_name：</td>		
						<td><input class="input-text" type="text" name="s_bank_name"   v-model="data.s_bank_name" ></td>
						</tr>
  <tr>
						<td>s_bank_user：</td>		
						<td><input class="input-text" type="text" name="s_bank_user"   v-model="data.s_bank_user" ></td>
						</tr>
  <tr>
						<td>s_bank_num：</td>		
						<td><input class="input-text" type="text" name="s_bank_num"   v-model="data.s_bank_num" ></td>
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
				this.data.imgurl = e;
			},
			getPage: function() {
				var that = this;
				that.app.get({
					url: that.app.apiHost + "/admin/config/add",
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
					url: that.app.apiHost + "/admin/config/save",
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
