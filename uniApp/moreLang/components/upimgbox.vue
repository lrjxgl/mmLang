<template>
	<view>
		<view class="upimg-box bg-fff">
			 
			<view class="upimg-item" v-for="(img,imgIndex) in imgsList" :key="imgIndex">
				<image class="upimg-img" :src="img.trueimgurl+'.100x100.jpg'"></image>
				<view class="upimg-del" @click="delImg(imgIndex)"></view>
			</view>
			 
			<view @click="upImg()" class="upimg-btn">
				<i class="upimg-btn-icon"></i>
			</view>
		</view>
	</view>
</template>

<script>
	 
	export default {
		name:"upimg-box",
		props:{
			defaultImgsList:{}, 
			defaultImgsData:{},
		},
		data:function(){
			return {
				imgsData:"",
				imgsList:[],
			}
		},
		created:function(){
			if(this.defaultImgsData!=undefined){
				this.imgsData=this.defaultImgsData;
			}
			if(this.defaultImgsList!=undefined){
				this.imgsList=this.defaultImgsList;
			}
			
			console.log(this.defaultImgsList)
		},
		
		 
		methods:{
			upImg:function(){
				var that=this;
				
				uni.chooseImage({
					success: function (chooseImageRes) {
						const tempFilePaths = chooseImageRes.tempFilePaths;
						const len=tempFilePaths.length;
						for(var i=0;i<len;i++){
							uni.uploadFile({
								url: that.app.apiHost+'/upload/img?token='+that.app.getToken(),  
								filePath: tempFilePaths[i],
								name: 'upimg',
								dataType:"json",
								success: function (rs) {
									var res=JSON.parse(rs.data);
									if(!res.error){
										
										var json={
											imgurl:res.imgurl,
											trueimgurl:res.trueimgurl
										};
										console.log(json)
										that.imgsList.push(json);
										if(that.imgsData=="" || that.imgsData==undefined ){
											that.imgsData=res.imgurl;											
										}else{
											that.imgsData=that.imgsData+","+res.imgurl;
										}
										that.$emit("call-parent",that.imgsData);
									}
								}
							});
						}
					}
				});
			},
			delImg:function(index){
				var len=this.imgsList.length;
				var imgslist=new Array();
				var imgsData="";
				var mgs=this.imgsList;
				for(var i in mgs){
					if(i!=index){
						imgslist.push(mgs[i]);						
					}
				}
				for(var i=0;i<imgslist.length;i++){
					if(i>0){
						imgsData+=",";
					}
					imgsData+=imgslist[i].imgurl;
				}
			 
				this.imgsData=imgsData;
				this.imgsList=imgslist;
				this.$emit("call-parent",this.imgsData);
				
			}
		}
	}
</script>

<style>

</style>
