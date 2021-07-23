<template>
	<view class="row-box">
		<div class="sky-tools"></div> 
		<editor id="editor" @input="change" class="sky-editor" :placeholder="placeholder" @ready="onEditorReady"></editor> 
	</view>
</template>

<script>
	var editorCtx;
	export default{
		props:{
			html:""
		},
		data:function(){
			return {
				content:"",
				placeholder:"请输入内容"
			}
		},
		created:function(){
			this.content=this.html;
			 
		},
		 
		watch:{
			html:function(n,o){
				this.content=n;
			}
		},
		methods:{
			onEditorReady() {
				var that=this;
				// #ifdef MP-BAIDU
				editorCtx = requireDynamicLib('editorLib').createEditorContext('editorId');
				// #endif

				// #ifdef APP-PLUS || H5 ||MP-WEIXIN
				uni.createSelectorQuery().select('#editor').context((res) => {
					 
					editorCtx = res.context
					editorCtx.setContents({
					  html:that.content
				  });
				}).exec()
				// #endif
			},
			change(e){
				 
				this.$emit("call-parent",e.detail.html);
			}
		}
	}
</script>

<style>
</style>
