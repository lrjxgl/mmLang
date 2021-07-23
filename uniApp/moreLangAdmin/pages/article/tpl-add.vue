<template>
	<view >
		<form @submit="submit" >
			<input type="text" name="id" class="none" v-model="data.id">
			  
				
				<div class="input-flex flex-ai-center">
					<div class="mgr-10">总价</div>
					<input class="input-flex-text" type="text" name="total_money" v-model="data.total_money" />
					<div class="mgl-5">万元</div>
				</div>	
				<div class="textarea-flex">
				 
					<textarea placeholder="请输入房源介绍"  class="textarea-flex-text h60" name="description" v-model="data.description"></textarea>
				</div>
			 
				<div v-if="pageLoad" class="bg-fff mgb-5">
					<input type="text" name="imgsdata" class="none" v-model="data.imgsdata" />
					<upimg-box :defaultImgsList="imgsList" :defaultImgsData="data.imgsdata" @call-parent="setImgs"></upimg-box>
				</div>
			 
			<button form-type="submit" class="btn-row-submit">保存</button>
		</form>
	</view>
</template>

<script>
	var id=0;
	import upimgBox from "../../components/upimgbox.vue";
	export default{
		components:{
			upimgBox
		},
		data:function(){
			return {
				pageLoad:false,
				data:{},
				imgsList:[]
			}
		},
		onLoad:function(ops){
			if(ops.id!=undefined){
				id=ops.id;
			}
			console.log(this.$data)
			this.getPage();
		},
		methods:{
			setImgs:function(e){
				this.data.imgsdata=e;
			},
			getPage:function(){
				var that=this;
				that.app.get({
					url:that.app.apiHost+"/module.php?m=house_resource&a=add",
					data:{
						id:id
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
					url:that.app.apiHost+"/module.php?m=house_resource&a=save",
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
