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
	@import "./common/iconfont.css";
		@import "./common/dt-ui-uni.css";
	
		/*修复uniapp flex*/
		.uni-page-head {
			display: flex;
			flex-direction: row;
		}
	
		.uni-tabbar,
		.uni-picker-view-wrapper {
			flex-direction: row !important;
		}
	
		.uni-swiper-dots {
			flex-direction: row;
		}
	
		.uni-scroll-view>div {
			flex-direction: row;
		}
	
		.uni-modal__ft {
			flex-direction: row;
		}
	
		/**/
		.imgWidth {
			width: 100%;
		}
	
		.fixToHome {
			position: fixed;
			top: 10px;
			left: 10px;
			background-color: #efefef;
			border-radius: 5px;
			cursor: pointer;
			padding: 5px 10px;
			color: #333;
			font-size: 14px;
			z-index: 999;
		}
	
		.badge {
			right: 6px;
			top: 3px;
			font-size: 12px;
			background-color: #dd524d;
			line-height: 1;
			border-radius: 50%;
			height: 14px;
			line-height: 1;
			padding-top: 1px;
			width: 14px;
			overflow: hidden;
			box-sizing: border-box;
			color: #fff;
		}
	
		.badge-abs {
			position: absolute;
		}
	
		.badge-mini {
			width: 6px;
			height: 6px;
			;
		}
	
		.input-flex-btn {
			z-index: 2;
		}
	
		.header,
		.header-row {
			display: none;
		}
	
		.fixedAdd {
			position: fixed;
			bottom: 200upx;
			right: 7upx;
			width: 92upx;
			height: 92upx;
			text-align: center;
			box-sizing: border-box;
			background-color: rgba(240, 85, 75, .82);
			color: #fff;
			font-family: iconfont;
			font-size: 32upx;
			padding-top: 11upx;
			border-radius: 22upx;
	
		}
	
		.fixedAdd:before {
			content: "\e7e8";
			display: block;
			font-size: 32upx;
		}
	
		.back-top {
			position: fixed;
			right: 10px;
			bottom: 240px;
			background-color: #009688;
			width: 40px;
			height: 40px;
			border-radius: 50%;
			display: flex;
			justify-content: center;
			align-items: center;
		}
	
		.back-top-icon {
			font-family: iconfont;
			line-height: 1;
		}
	
		.back-top-icon:after {
			content: "\e603";
			font-size: 12px;
			color: #fff;
		}
	
		.back-top-text {
			font-size: 12px;
			line-height: 1;
			font-weight: 300;
			color: #fff;
		}
	
		.buy-btn {
			display: inline-block;
			height: 30px;
			line-height: 30px;
			text-align: center;
			width: 30px;
			border: 1px solid #83c44e;
			color: #83c44e;
			border-radius: 50%;
			float: right;
			cursor: pointer;
		}
	
		.buy-btn-active {
			background-color: #83c44e;
			color: #fff;
		}
	
		.btn-mini-cycle {
			line-height: 12px;
			font-size: 12px;
			padding: 2px 5px;
			border-radius: 5px;
			border: 1px solid #f60;
			color: #f60;
		}
	
		.btn-type,
		.btn-type-wm {
			line-height: 12px;
			font-size: 12px;
			padding: 2px 5px;
			border-radius: 5px;
			border: 1px solid #007BFF;
			color: #0000FF;
		}
	
		.btn-type {
			width: 22px;
		}
	
		.unLoginBox {
			position: absolute;
			top: 0;
			left: 0;
			right: 0;
			bottom: 0;
			background-color: #fff;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
		}
		.adBox{
			display: flex;
			flex-direction: row;
		}
		.adBox-item{
			flex: 1;
		}
		.adBox-img{
			width: 100%;
		}
</style>
