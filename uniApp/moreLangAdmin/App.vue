<script>
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	
	@import url("//at.alicdn.com/t/font_902011_f1rcztainbm.css");
	@import url("common/dt-ui-h5.css");
	@import url("common/admin.css");
	uni-left-window{
		position: fixed;
		left: 0;
		top: 60px;
		bottom: 0px;
		width: 200px;
	}
	.uni-top-window{
		z-index: 999;
		background-color: #efefef;
	}
	uni-main{
		margin-left: 210px;
	}
</style>
