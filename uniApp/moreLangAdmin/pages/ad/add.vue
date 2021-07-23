<template>
	<view >
		<ad-nav tab="adAdd"></ad-nav>
		<form >
			<div class="none">
				<input type="text" name="id"  v-model="id" >
				<input type="text" id="img_width" value="width" />
				<input type="text" id="img_height" value="height" />
			</div>
		
		  <table class="table-add">
		    
		    <tr>
		      <td width="100">分类id：</td>
		      <td>
		      <select class="w100" name="tag_id" id="ajax_tag_id">
		      <option value="0">请选择</option>
		      {foreach item=c key=k from=$tag_list}
		      <option value="{$k}" {if $k eq $data.tag_id or $k eq get_post("tag_id")} selected="selected"{/if}>{$c.title}({$c.width}*{$c.height})</option>
		      {/foreach}
		      </select>
		    	
		        <select class="w150" name="tag_id_2nd" id="ajax_tag_id_2nd" {if !$tag_list_2nd}style="display:none"{/if}>
		      <option value="0">请选择</option>
		      {foreach item=c key=k from=$tag_list_2nd}
		      <option value="{$k}" {if $k eq $data.tag_id_2nd or $k eq get_post("tag_2nd_id")} selected="selected"{/if}>{$c.title}({$c.width}*{$c.height})</option>
		      {/foreach}
		      </select>
		    
		    </td> 
		    </tr>
		    <tr>
		      <td>标题：</td>
		      <td><input type="text" name="title" id="title" class="w600" value="{$data.title}" ></td>
		    </tr>
		    <tr>
		      <td>描述：</td>
		      <td><textarea name="info" id="info"  class="w600" >{$data.info}</textarea></td>
		    </tr>
		    <tr>
		      <td>链接1：</td>
		      <td><input type="text" name="link1" id="link1"  class="w600"  value="{$data.link1}" ></td>
		    </tr>
		    <tr>
		      <td>链接2：</td>
		      <td><input type="text" name="link2" id="link2"  class="w600"  value="{$data.link2}" ></td>
		    </tr>
		    <tr>
		      <td>开始时间：</td>
		      <td><input readonly="" type="text" name="starttime" id="starttime" value="{if $data}{$data.starttime|date:Y-m-d H:m:s}{else}2019-07-17 09:03:01{/if}" ></td>
		    </tr>
		    <tr>
		      <td>结束时间：</td>
		      <td><input readonly="" type="text" name="endtime" id="endtime" value="{if $data}{$data.endtime|date:Y-m-d H:m:s}{else}2029-07-17 09:03:01{/if}" ></td>
		    </tr>
		    <tr>
		      <td>图片地址：</td>
		      <td>
							<div class="js-upload-item">
								<input type="file" id="upa" class="js-upload-file" style="display: none;" />
								<div class="upimg-btn js-upload-btn">+</div>
								<input type="hidden" name="imgurl" class="js-imgurl" value="{$data.imgurl}" />
								<div class="js-imgbox">
									{if $data.imgurl}
									<img src="{$data.imgurl|images_site}.100x100.jpg">
									{/if}
								</div>
							</div>
							
						</td>
		    </tr>
		    
		    <tr>
		      <td>图片地址2：</td>
		      <td>
						
				</td>
		    </tr>
		    
		    <tr>
		      <td>排序：</td>
		      <td><input type="text" name="orderindex" id="orderindex" value="{$data.orderindex}" ></td>
		    </tr>
		    <tr>
		      <td>状态：</td>
		      <td><input type="radio" name="status" value="1" {if $data.status eq 1} checked="checked"{/if} />隐藏 &nbsp; 
		      <input type="radio" name="status" value="2" {if !$data || $data.status eq 2} checked="checked"{/if} />显示</td>
		    </tr>
		    {if $data}
		    <tr>
		      <td>添加时间：</td>
		      <td>{$data.dateline|date:Y-m-d}</td>
		    </tr>
		    {/if}
		    <tr>
		      <td>价格：</td>
		      <td><input type="text" name="price" id="price" value="{$data.price}" ></td>
		    </tr>
		    <tr>
		      <td>对象ID：</td>
		      <td><input type="text" name="object_id" id="object_id" value="{$data.object_id}" ></td>
		    </tr>
		    
		
		  </table>
		  <div class="btn-row-submit js-submit">保存</div>
		</form>
	</view>
</template>

<script>
	 
	import upImg from "../../components/skyupimg.vue";
	import adNav from "../ad/ad-nav.vue"  
	export default{
		components:{
			adNav,
			upImg
		},
		data:function(){
			return {
				id:0,
				pageLoad:false,
				data:{},
				imgsList:[]
			}
		},
		onLoad:function(ops){
			if(ops.id!=undefined){
				this.id=ops.id;
			}
			console.log(this.$data)
			//this.getPage();
		},
		methods:{
			setImgs:function(e){
				this.data.imgsdata=e;
			},
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/admin/ad/add",
					data:{
						id:this.id
					},
					success:function(res){
						if(res.data.data){
							that.data=res.data.data;
							that.imgsList=res.data.imgsdata;
						}else{
							that.data={
								id:0,
								title:"",
								typeid:0,
								total_money:0,
								description:"",
								imgsdata:""
							}
						}
						
						that.pageLoad=true;
					}
				})
			},
			submit:function(e){
				var that=this;
				that.app.post({
					url:that.app.apiHost+"/admin/ad/save",
					data:e.detail.value,
					success:function(res){
						uni.showToast({
							title:res.message
						})
					}
				})
				
			}
		}
	}
</script>

<style>
</style>
