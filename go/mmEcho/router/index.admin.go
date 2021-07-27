
	package router

	import (
		"app/index/indexAdmin"
		"github.com/labstack/echo/v4"
	)

	func IndexAdminRouter(e *echo.Echo) {
		/*Ad*/
		e.GET("/admin/ad/index", indexAdmin.AdIndex);
		e.GET("/admin/ad/add", indexAdmin.AdAdd);
		e.GET("/admin/ad/save", indexAdmin.AdSave);
		e.GET("/admin/ad/status", indexAdmin.AdStatus);
		e.GET("/admin/ad/delete", indexAdmin.AdDelete);
		/*AdTags*/
		e.GET("/admin/ad_tags/index", indexAdmin.AdTagsIndex);
		e.GET("/admin/ad_tags/add", indexAdmin.AdTagsAdd);
		e.GET("/admin/ad_tags/save", indexAdmin.AdTagsSave);
		e.GET("/admin/ad_tags/status", indexAdmin.AdTagsStatus);
		e.GET("/admin/ad_tags/delete", indexAdmin.AdTagsDelete);
		/*Admin*/
		e.GET("/admin/admin/index", indexAdmin.AdminIndex);
		e.GET("/admin/admin/add", indexAdmin.AdminAdd);
		e.GET("/admin/admin/save", indexAdmin.AdminSave);
		e.GET("/admin/admin/delete", indexAdmin.AdminDelete);
		/*AdminGroup*/
		e.GET("/admin/admin_group/index", indexAdmin.AdminGroupIndex);
		e.GET("/admin/admin_group/add", indexAdmin.AdminGroupAdd);
		e.GET("/admin/admin_group/save", indexAdmin.AdminGroupSave);
		e.GET("/admin/admin_group/delete", indexAdmin.AdminGroupDelete);
		/*Apppush*/
		e.GET("/admin/apppush/index", indexAdmin.ApppushIndex);
		e.GET("/admin/apppush/add", indexAdmin.ApppushAdd);
		e.GET("/admin/apppush/save", indexAdmin.ApppushSave);
		e.GET("/admin/apppush/delete", indexAdmin.ApppushDelete);
		/*ApppushPlan*/
		e.GET("/admin/apppush_plan/index", indexAdmin.ApppushPlanIndex);
		e.GET("/admin/apppush_plan/add", indexAdmin.ApppushPlanAdd);
		e.GET("/admin/apppush_plan/save", indexAdmin.ApppushPlanSave);
		e.GET("/admin/apppush_plan/status", indexAdmin.ApppushPlanStatus);
		e.GET("/admin/apppush_plan/delete", indexAdmin.ApppushPlanDelete);
		/*Article*/
		e.GET("/admin/article/index", indexAdmin.ArticleIndex);
		e.GET("/admin/article/add", indexAdmin.ArticleAdd);
		e.GET("/admin/article/save", indexAdmin.ArticleSave);
		e.GET("/admin/article/status", indexAdmin.ArticleStatus);
		e.GET("/admin/article/recommend", indexAdmin.ArticleRecommend);
		e.GET("/admin/article/delete", indexAdmin.ArticleDelete);
		/*ArticleComment*/
		e.GET("/admin/article_comment/index", indexAdmin.ArticleCommentIndex);
		e.GET("/admin/article_comment/add", indexAdmin.ArticleCommentAdd);
		e.GET("/admin/article_comment/save", indexAdmin.ArticleCommentSave);
		e.GET("/admin/article_comment/status", indexAdmin.ArticleCommentStatus);
		e.GET("/admin/article_comment/delete", indexAdmin.ArticleCommentDelete);
		/*ArticleData*/
		e.GET("/admin/article_data/index", indexAdmin.ArticleDataIndex);
		e.GET("/admin/article_data/add", indexAdmin.ArticleDataAdd);
		e.GET("/admin/article_data/save", indexAdmin.ArticleDataSave);
		e.GET("/admin/article_data/delete", indexAdmin.ArticleDataDelete);
		/*Attach*/
		e.GET("/admin/attach/index", indexAdmin.AttachIndex);
		e.GET("/admin/attach/add", indexAdmin.AttachAdd);
		e.GET("/admin/attach/save", indexAdmin.AttachSave);
		e.GET("/admin/attach/status", indexAdmin.AttachStatus);
		e.GET("/admin/attach/delete", indexAdmin.AttachDelete);
		/*Badip*/
		e.GET("/admin/badip/index", indexAdmin.BadipIndex);
		e.GET("/admin/badip/add", indexAdmin.BadipAdd);
		e.GET("/admin/badip/save", indexAdmin.BadipSave);
		e.GET("/admin/badip/delete", indexAdmin.BadipDelete);
		/*Badphone*/
		e.GET("/admin/badphone/index", indexAdmin.BadphoneIndex);
		e.GET("/admin/badphone/add", indexAdmin.BadphoneAdd);
		e.GET("/admin/badphone/save", indexAdmin.BadphoneSave);
		e.GET("/admin/badphone/delete", indexAdmin.BadphoneDelete);
		/*Bankcard*/
		e.GET("/admin/bankcard/index", indexAdmin.BankcardIndex);
		e.GET("/admin/bankcard/add", indexAdmin.BankcardAdd);
		e.GET("/admin/bankcard/save", indexAdmin.BankcardSave);
		e.GET("/admin/bankcard/status", indexAdmin.BankcardStatus);
		e.GET("/admin/bankcard/delete", indexAdmin.BankcardDelete);
		/*Blacklist*/
		e.GET("/admin/blacklist/index", indexAdmin.BlacklistIndex);
		e.GET("/admin/blacklist/add", indexAdmin.BlacklistAdd);
		e.GET("/admin/blacklist/save", indexAdmin.BlacklistSave);
		e.GET("/admin/blacklist/delete", indexAdmin.BlacklistDelete);
		/*BlacklistPost*/
		e.GET("/admin/blacklist_post/index", indexAdmin.BlacklistPostIndex);
		e.GET("/admin/blacklist_post/add", indexAdmin.BlacklistPostAdd);
		e.GET("/admin/blacklist_post/save", indexAdmin.BlacklistPostSave);
		e.GET("/admin/blacklist_post/delete", indexAdmin.BlacklistPostDelete);
		/*Category*/
		e.GET("/admin/category/index", indexAdmin.CategoryIndex);
		e.GET("/admin/category/add", indexAdmin.CategoryAdd);
		e.GET("/admin/category/save", indexAdmin.CategorySave);
		e.GET("/admin/category/status", indexAdmin.CategoryStatus);
		e.GET("/admin/category/delete", indexAdmin.CategoryDelete);
		/*Checkin*/
		e.GET("/admin/checkin/index", indexAdmin.CheckinIndex);
		e.GET("/admin/checkin/add", indexAdmin.CheckinAdd);
		e.GET("/admin/checkin/save", indexAdmin.CheckinSave);
		e.GET("/admin/checkin/delete", indexAdmin.CheckinDelete);
		/*CheckinIndex*/
		e.GET("/admin/checkin/index", indexAdmin.CheckinIndexIndex);
		e.GET("/admin/checkin_index/add", indexAdmin.CheckinIndexAdd);
		e.GET("/admin/checkin_index/save", indexAdmin.CheckinIndexSave);
		e.GET("/admin/checkin_index/delete", indexAdmin.CheckinIndexDelete);
		/*City*/
		e.GET("/admin/city/index", indexAdmin.CityIndex);
		e.GET("/admin/city/add", indexAdmin.CityAdd);
		e.GET("/admin/city/save", indexAdmin.CitySave);
		e.GET("/admin/city/status", indexAdmin.CityStatus);
		e.GET("/admin/city/recommend", indexAdmin.CityRecommend);
		e.GET("/admin/city/delete", indexAdmin.CityDelete);
		/*Config*/
		e.GET("/admin/config/index", indexAdmin.ConfigIndex);
		e.GET("/admin/config/add", indexAdmin.ConfigAdd);
		e.GET("/admin/config/save", indexAdmin.ConfigSave);
		e.GET("/admin/config/delete", indexAdmin.ConfigDelete);
		/*Coupon*/
		e.GET("/admin/coupon/index", indexAdmin.CouponIndex);
		e.GET("/admin/coupon/add", indexAdmin.CouponAdd);
		e.GET("/admin/coupon/save", indexAdmin.CouponSave);
		e.GET("/admin/coupon/status", indexAdmin.CouponStatus);
		e.GET("/admin/coupon/delete", indexAdmin.CouponDelete);
		/*CouponUser*/
		e.GET("/admin/coupon_user/index", indexAdmin.CouponUserIndex);
		e.GET("/admin/coupon_user/add", indexAdmin.CouponUserAdd);
		e.GET("/admin/coupon_user/save", indexAdmin.CouponUserSave);
		e.GET("/admin/coupon_user/status", indexAdmin.CouponUserStatus);
		e.GET("/admin/coupon_user/delete", indexAdmin.CouponUserDelete);
		/*Crontab*/
		e.GET("/admin/crontab/index", indexAdmin.CrontabIndex);
		e.GET("/admin/crontab/add", indexAdmin.CrontabAdd);
		e.GET("/admin/crontab/save", indexAdmin.CrontabSave);
		e.GET("/admin/crontab/status", indexAdmin.CrontabStatus);
		e.GET("/admin/crontab/delete", indexAdmin.CrontabDelete);
		/*Dataapi*/
		e.GET("/admin/dataapi/index", indexAdmin.DataapiIndex);
		e.GET("/admin/dataapi/add", indexAdmin.DataapiAdd);
		e.GET("/admin/dataapi/save", indexAdmin.DataapiSave);
		e.GET("/admin/dataapi/status", indexAdmin.DataapiStatus);
		e.GET("/admin/dataapi/delete", indexAdmin.DataapiDelete);
		/*Daysn*/
		e.GET("/admin/daysn/index", indexAdmin.DaysnIndex);
		e.GET("/admin/daysn/add", indexAdmin.DaysnAdd);
		e.GET("/admin/daysn/save", indexAdmin.DaysnSave);
		e.GET("/admin/daysn/delete", indexAdmin.DaysnDelete);
		/*Dbcache*/
		e.GET("/admin/dbcache/index", indexAdmin.DbcacheIndex);
		e.GET("/admin/dbcache/add", indexAdmin.DbcacheAdd);
		e.GET("/admin/dbcache/save", indexAdmin.DbcacheSave);
		e.GET("/admin/dbcache/delete", indexAdmin.DbcacheDelete);
		/*Dbsession*/
		e.GET("/admin/dbsession/index", indexAdmin.DbsessionIndex);
		e.GET("/admin/dbsession/add", indexAdmin.DbsessionAdd);
		e.GET("/admin/dbsession/save", indexAdmin.DbsessionSave);
		e.GET("/admin/dbsession/delete", indexAdmin.DbsessionDelete);
		/*District*/
		e.GET("/admin/district/index", indexAdmin.DistrictIndex);
		e.GET("/admin/district/add", indexAdmin.DistrictAdd);
		e.GET("/admin/district/save", indexAdmin.DistrictSave);
		e.GET("/admin/district/delete", indexAdmin.DistrictDelete);
		/*ExpressFee*/
		e.GET("/admin/express_fee/index", indexAdmin.ExpressFeeIndex);
		e.GET("/admin/express_fee/add", indexAdmin.ExpressFeeAdd);
		e.GET("/admin/express_fee/save", indexAdmin.ExpressFeeSave);
		e.GET("/admin/express_fee/delete", indexAdmin.ExpressFeeDelete);
		/*ExpressFeeCity*/
		e.GET("/admin/express_fee_city/index", indexAdmin.ExpressFeeCityIndex);
		e.GET("/admin/express_fee_city/add", indexAdmin.ExpressFeeCityAdd);
		e.GET("/admin/express_fee_city/save", indexAdmin.ExpressFeeCitySave);
		e.GET("/admin/express_fee_city/delete", indexAdmin.ExpressFeeCityDelete);
		/*Fav*/
		e.GET("/admin/fav/index", indexAdmin.FavIndex);
		e.GET("/admin/fav/add", indexAdmin.FavAdd);
		e.GET("/admin/fav/save", indexAdmin.FavSave);
		e.GET("/admin/fav/delete", indexAdmin.FavDelete);
		/*Follow*/
		e.GET("/admin/follow/index", indexAdmin.FollowIndex);
		e.GET("/admin/follow/add", indexAdmin.FollowAdd);
		e.GET("/admin/follow/save", indexAdmin.FollowSave);
		e.GET("/admin/follow/status", indexAdmin.FollowStatus);
		e.GET("/admin/follow/delete", indexAdmin.FollowDelete);
		/*Followed*/
		e.GET("/admin/followed/index", indexAdmin.FollowedIndex);
		e.GET("/admin/followed/add", indexAdmin.FollowedAdd);
		e.GET("/admin/followed/save", indexAdmin.FollowedSave);
		e.GET("/admin/followed/status", indexAdmin.FollowedStatus);
		e.GET("/admin/followed/delete", indexAdmin.FollowedDelete);
		/*Friend*/
		e.GET("/admin/friend/index", indexAdmin.FriendIndex);
		e.GET("/admin/friend/add", indexAdmin.FriendAdd);
		e.GET("/admin/friend/save", indexAdmin.FriendSave);
		e.GET("/admin/friend/delete", indexAdmin.FriendDelete);
		/*FriendApply*/
		e.GET("/admin/friend_apply/index", indexAdmin.FriendApplyIndex);
		e.GET("/admin/friend_apply/add", indexAdmin.FriendApplyAdd);
		e.GET("/admin/friend_apply/save", indexAdmin.FriendApplySave);
		e.GET("/admin/friend_apply/delete", indexAdmin.FriendApplyDelete);
		/*GoldLog*/
		e.GET("/admin/gold_log/index", indexAdmin.GoldLogIndex);
		e.GET("/admin/gold_log/add", indexAdmin.GoldLogAdd);
		e.GET("/admin/gold_log/save", indexAdmin.GoldLogSave);
		e.GET("/admin/gold_log/delete", indexAdmin.GoldLogDelete);
		/*GradeLog*/
		e.GET("/admin/grade_log/index", indexAdmin.GradeLogIndex);
		e.GET("/admin/grade_log/add", indexAdmin.GradeLogAdd);
		e.GET("/admin/grade_log/save", indexAdmin.GradeLogSave);
		e.GET("/admin/grade_log/delete", indexAdmin.GradeLogDelete);
		/*Guest*/
		e.GET("/admin/guest/index", indexAdmin.GuestIndex);
		e.GET("/admin/guest/add", indexAdmin.GuestAdd);
		e.GET("/admin/guest/save", indexAdmin.GuestSave);
		e.GET("/admin/guest/status", indexAdmin.GuestStatus);
		e.GET("/admin/guest/delete", indexAdmin.GuestDelete);
		/*Imgs*/
		e.GET("/admin/imgs/index", indexAdmin.ImgsIndex);
		e.GET("/admin/imgs/add", indexAdmin.ImgsAdd);
		e.GET("/admin/imgs/save", indexAdmin.ImgsSave);
		e.GET("/admin/imgs/delete", indexAdmin.ImgsDelete);
		/*Invite*/
		e.GET("/admin/invite/index", indexAdmin.InviteIndex);
		e.GET("/admin/invite/add", indexAdmin.InviteAdd);
		e.GET("/admin/invite/save", indexAdmin.InviteSave);
		e.GET("/admin/invite/delete", indexAdmin.InviteDelete);
		/*InviteAccount*/
		e.GET("/admin/invite_account/index", indexAdmin.InviteAccountIndex);
		e.GET("/admin/invite_account/add", indexAdmin.InviteAccountAdd);
		e.GET("/admin/invite_account/save", indexAdmin.InviteAccountSave);
		e.GET("/admin/invite_account/delete", indexAdmin.InviteAccountDelete);
		/*InviteAccountLog*/
		e.GET("/admin/invite_account_log/index", indexAdmin.InviteAccountLogIndex);
		e.GET("/admin/invite_account_log/add", indexAdmin.InviteAccountLogAdd);
		e.GET("/admin/invite_account_log/save", indexAdmin.InviteAccountLogSave);
		e.GET("/admin/invite_account_log/delete", indexAdmin.InviteAccountLogDelete);
		/*KefuMsg*/
		e.GET("/admin/kefu_msg/index", indexAdmin.KefuMsgIndex);
		e.GET("/admin/kefu_msg/add", indexAdmin.KefuMsgAdd);
		e.GET("/admin/kefu_msg/save", indexAdmin.KefuMsgSave);
		e.GET("/admin/kefu_msg/status", indexAdmin.KefuMsgStatus);
		e.GET("/admin/kefu_msg/delete", indexAdmin.KefuMsgDelete);
		/*KefuMsgindex*/
		e.GET("/admin/kefu_msgindex/index", indexAdmin.KefuMsgindexIndex);
		e.GET("/admin/kefu_msgindex/add", indexAdmin.KefuMsgindexAdd);
		e.GET("/admin/kefu_msgindex/save", indexAdmin.KefuMsgindexSave);
		e.GET("/admin/kefu_msgindex/status", indexAdmin.KefuMsgindexStatus);
		e.GET("/admin/kefu_msgindex/delete", indexAdmin.KefuMsgindexDelete);
		/*Love*/
		e.GET("/admin/love/index", indexAdmin.LoveIndex);
		e.GET("/admin/love/add", indexAdmin.LoveAdd);
		e.GET("/admin/love/save", indexAdmin.LoveSave);
		e.GET("/admin/love/delete", indexAdmin.LoveDelete);
		/*Maxid*/
		e.GET("/admin/maxid/index", indexAdmin.MaxidIndex);
		e.GET("/admin/maxid/add", indexAdmin.MaxidAdd);
		e.GET("/admin/maxid/save", indexAdmin.MaxidSave);
		e.GET("/admin/maxid/delete", indexAdmin.MaxidDelete);
		/*Model*/
		e.GET("/admin/model/index", indexAdmin.ModelIndex);
		e.GET("/admin/model/add", indexAdmin.ModelAdd);
		e.GET("/admin/model/save", indexAdmin.ModelSave);
		e.GET("/admin/model/status", indexAdmin.ModelStatus);
		e.GET("/admin/model/delete", indexAdmin.ModelDelete);
		/*ModelIndex*/
		e.GET("/admin/model/index", indexAdmin.ModelIndexIndex);
		e.GET("/admin/model_index/add", indexAdmin.ModelIndexAdd);
		e.GET("/admin/model_index/save", indexAdmin.ModelIndexSave);
		e.GET("/admin/model_index/delete", indexAdmin.ModelIndexDelete);
		/*Module*/
		e.GET("/admin/module/index", indexAdmin.ModuleIndex);
		e.GET("/admin/module/add", indexAdmin.ModuleAdd);
		e.GET("/admin/module/save", indexAdmin.ModuleSave);
		e.GET("/admin/module/delete", indexAdmin.ModuleDelete);
		/*Navbar*/
		e.GET("/admin/navbar/index", indexAdmin.NavbarIndex);
		e.GET("/admin/navbar/show", indexAdmin.NavbarShow);
		e.GET("/admin/navbar/add", indexAdmin.NavbarAdd);
		e.GET("/admin/navbar/save", indexAdmin.NavbarSave);
		e.GET("/admin/navbar/status", indexAdmin.NavbarStatus);
		e.GET("/admin/navbar/delete", indexAdmin.NavbarDelete);
		e.GET("/admin/navbar/index", indexAdmin.NavbarIndex);
		/*Notice*/
		e.GET("/admin/notice/index", indexAdmin.NoticeIndex);
		e.GET("/admin/notice/add", indexAdmin.NoticeAdd);
		e.GET("/admin/notice/save", indexAdmin.NoticeSave);
		e.GET("/admin/notice/status", indexAdmin.NoticeStatus);
		e.GET("/admin/notice/delete", indexAdmin.NoticeDelete);
		/*OpenAlioss*/
		e.GET("/admin/open_alioss/index", indexAdmin.OpenAliossIndex);
		e.GET("/admin/open_alioss/add", indexAdmin.OpenAliossAdd);
		e.GET("/admin/open_alioss/save", indexAdmin.OpenAliossSave);
		e.GET("/admin/open_alioss/status", indexAdmin.OpenAliossStatus);
		e.GET("/admin/open_alioss/delete", indexAdmin.OpenAliossDelete);
		/*OpenAlipay*/
		e.GET("/admin/open_alipay/index", indexAdmin.OpenAlipayIndex);
		e.GET("/admin/open_alipay/add", indexAdmin.OpenAlipayAdd);
		e.GET("/admin/open_alipay/save", indexAdmin.OpenAlipaySave);
		e.GET("/admin/open_alipay/status", indexAdmin.OpenAlipayStatus);
		e.GET("/admin/open_alipay/delete", indexAdmin.OpenAlipayDelete);
		/*OpenAlipayApp*/
		e.GET("/admin/open_alipay_app/index", indexAdmin.OpenAlipayAppIndex);
		e.GET("/admin/open_alipay_app/add", indexAdmin.OpenAlipayAppAdd);
		e.GET("/admin/open_alipay_app/save", indexAdmin.OpenAlipayAppSave);
		e.GET("/admin/open_alipay_app/status", indexAdmin.OpenAlipayAppStatus);
		e.GET("/admin/open_alipay_app/delete", indexAdmin.OpenAlipayAppDelete);
		/*OpenAlipayMini*/
		e.GET("/admin/open_alipay_mini/index", indexAdmin.OpenAlipayMiniIndex);
		e.GET("/admin/open_alipay_mini/add", indexAdmin.OpenAlipayMiniAdd);
		e.GET("/admin/open_alipay_mini/save", indexAdmin.OpenAlipayMiniSave);
		e.GET("/admin/open_alipay_mini/status", indexAdmin.OpenAlipayMiniStatus);
		e.GET("/admin/open_alipay_mini/delete", indexAdmin.OpenAlipayMiniDelete);
		/*OpenBaidu*/
		e.GET("/admin/open_baidu/index", indexAdmin.OpenBaiduIndex);
		e.GET("/admin/open_baidu/add", indexAdmin.OpenBaiduAdd);
		e.GET("/admin/open_baidu/save", indexAdmin.OpenBaiduSave);
		e.GET("/admin/open_baidu/status", indexAdmin.OpenBaiduStatus);
		e.GET("/admin/open_baidu/delete", indexAdmin.OpenBaiduDelete);
		/*OpenQq*/
		e.GET("/admin/open_qq/index", indexAdmin.OpenQqIndex);
		e.GET("/admin/open_qq/add", indexAdmin.OpenQqAdd);
		e.GET("/admin/open_qq/save", indexAdmin.OpenQqSave);
		e.GET("/admin/open_qq/status", indexAdmin.OpenQqStatus);
		e.GET("/admin/open_qq/delete", indexAdmin.OpenQqDelete);
		/*OpenToutiao*/
		e.GET("/admin/open_toutiao/index", indexAdmin.OpenToutiaoIndex);
		e.GET("/admin/open_toutiao/add", indexAdmin.OpenToutiaoAdd);
		e.GET("/admin/open_toutiao/save", indexAdmin.OpenToutiaoSave);
		e.GET("/admin/open_toutiao/status", indexAdmin.OpenToutiaoStatus);
		e.GET("/admin/open_toutiao/delete", indexAdmin.OpenToutiaoDelete);
		/*OpenWxapp*/
		e.GET("/admin/open_wxapp/index", indexAdmin.OpenWxappIndex);
		e.GET("/admin/open_wxapp/add", indexAdmin.OpenWxappAdd);
		e.GET("/admin/open_wxapp/save", indexAdmin.OpenWxappSave);
		e.GET("/admin/open_wxapp/status", indexAdmin.OpenWxappStatus);
		e.GET("/admin/open_wxapp/delete", indexAdmin.OpenWxappDelete);
		/*OpenWxnative*/
		e.GET("/admin/open_wxnative/index", indexAdmin.OpenWxnativeIndex);
		e.GET("/admin/open_wxnative/add", indexAdmin.OpenWxnativeAdd);
		e.GET("/admin/open_wxnative/save", indexAdmin.OpenWxnativeSave);
		e.GET("/admin/open_wxnative/status", indexAdmin.OpenWxnativeStatus);
		e.GET("/admin/open_wxnative/delete", indexAdmin.OpenWxnativeDelete);
		/*Openlogin*/
		e.GET("/admin/openlogin/index", indexAdmin.OpenloginIndex);
		e.GET("/admin/openlogin/add", indexAdmin.OpenloginAdd);
		e.GET("/admin/openlogin/save", indexAdmin.OpenloginSave);
		e.GET("/admin/openlogin/delete", indexAdmin.OpenloginDelete);
		/*Pagetpl*/
		e.GET("/admin/pagetpl/index", indexAdmin.PagetplIndex);
		e.GET("/admin/pagetpl/add", indexAdmin.PagetplAdd);
		e.GET("/admin/pagetpl/save", indexAdmin.PagetplSave);
		e.GET("/admin/pagetpl/delete", indexAdmin.PagetplDelete);
		/*PayLog*/
		e.GET("/admin/pay_log/index", indexAdmin.PayLogIndex);
		e.GET("/admin/pay_log/add", indexAdmin.PayLogAdd);
		e.GET("/admin/pay_log/save", indexAdmin.PayLogSave);
		e.GET("/admin/pay_log/delete", indexAdmin.PayLogDelete);
		/*Permission*/
		e.GET("/admin/permission/index", indexAdmin.PermissionIndex);
		e.GET("/admin/permission/add", indexAdmin.PermissionAdd);
		e.GET("/admin/permission/save", indexAdmin.PermissionSave);
		e.GET("/admin/permission/delete", indexAdmin.PermissionDelete);
		/*Pm*/
		e.GET("/admin/pm/index", indexAdmin.PmIndex);
		e.GET("/admin/pm/add", indexAdmin.PmAdd);
		e.GET("/admin/pm/save", indexAdmin.PmSave);
		e.GET("/admin/pm/status", indexAdmin.PmStatus);
		e.GET("/admin/pm/delete", indexAdmin.PmDelete);
		/*PmIndex*/
		e.GET("/admin/pm/index", indexAdmin.PmIndexIndex);
		e.GET("/admin/pm_index/add", indexAdmin.PmIndexAdd);
		e.GET("/admin/pm_index/save", indexAdmin.PmIndexSave);
		e.GET("/admin/pm_index/status", indexAdmin.PmIndexStatus);
		e.GET("/admin/pm_index/delete", indexAdmin.PmIndexDelete);
		/*Pv*/
		e.GET("/admin/pv/index", indexAdmin.PvIndex);
		e.GET("/admin/pv/add", indexAdmin.PvAdd);
		e.GET("/admin/pv/save", indexAdmin.PvSave);
		e.GET("/admin/pv/delete", indexAdmin.PvDelete);
		/*PvDay*/
		e.GET("/admin/pv_day/index", indexAdmin.PvDayIndex);
		e.GET("/admin/pv_day/add", indexAdmin.PvDayAdd);
		e.GET("/admin/pv_day/save", indexAdmin.PvDaySave);
		e.GET("/admin/pv_day/delete", indexAdmin.PvDayDelete);
		/*PvIp*/
		e.GET("/admin/pv_ip/index", indexAdmin.PvIpIndex);
		e.GET("/admin/pv_ip/add", indexAdmin.PvIpAdd);
		e.GET("/admin/pv_ip/save", indexAdmin.PvIpSave);
		e.GET("/admin/pv_ip/delete", indexAdmin.PvIpDelete);
		/*PvUv*/
		e.GET("/admin/pv_uv/index", indexAdmin.PvUvIndex);
		e.GET("/admin/pv_uv/add", indexAdmin.PvUvAdd);
		e.GET("/admin/pv_uv/save", indexAdmin.PvUvSave);
		e.GET("/admin/pv_uv/delete", indexAdmin.PvUvDelete);
		/*Queue*/
		e.GET("/admin/queue/index", indexAdmin.QueueIndex);
		e.GET("/admin/queue/add", indexAdmin.QueueAdd);
		e.GET("/admin/queue/save", indexAdmin.QueueSave);
		e.GET("/admin/queue/delete", indexAdmin.QueueDelete);
		/*Recharge*/
		e.GET("/admin/recharge/index", indexAdmin.RechargeIndex);
		e.GET("/admin/recharge/add", indexAdmin.RechargeAdd);
		e.GET("/admin/recharge/save", indexAdmin.RechargeSave);
		e.GET("/admin/recharge/status", indexAdmin.RechargeStatus);
		e.GET("/admin/recharge/delete", indexAdmin.RechargeDelete);
		/*Refund*/
		e.GET("/admin/refund/index", indexAdmin.RefundIndex);
		e.GET("/admin/refund/add", indexAdmin.RefundAdd);
		e.GET("/admin/refund/save", indexAdmin.RefundSave);
		e.GET("/admin/refund/status", indexAdmin.RefundStatus);
		e.GET("/admin/refund/delete", indexAdmin.RefundDelete);
		/*RefundApply*/
		e.GET("/admin/refund_apply/index", indexAdmin.RefundApplyIndex);
		e.GET("/admin/refund_apply/add", indexAdmin.RefundApplyAdd);
		e.GET("/admin/refund_apply/save", indexAdmin.RefundApplySave);
		e.GET("/admin/refund_apply/status", indexAdmin.RefundApplyStatus);
		e.GET("/admin/refund_apply/delete", indexAdmin.RefundApplyDelete);
		/*Seo*/
		e.GET("/admin/seo/index", indexAdmin.SeoIndex);
		e.GET("/admin/seo/add", indexAdmin.SeoAdd);
		e.GET("/admin/seo/save", indexAdmin.SeoSave);
		e.GET("/admin/seo/delete", indexAdmin.SeoDelete);
		/*Sgpay*/
		e.GET("/admin/sgpay/index", indexAdmin.SgpayIndex);
		e.GET("/admin/sgpay/add", indexAdmin.SgpayAdd);
		e.GET("/admin/sgpay/save", indexAdmin.SgpaySave);
		e.GET("/admin/sgpay/status", indexAdmin.SgpayStatus);
		e.GET("/admin/sgpay/delete", indexAdmin.SgpayDelete);
		/*Site*/
		e.GET("/admin/site/index", indexAdmin.SiteIndex);
		e.GET("/admin/site/add", indexAdmin.SiteAdd);
		e.GET("/admin/site/save", indexAdmin.SiteSave);
		e.GET("/admin/site/status", indexAdmin.SiteStatus);
		e.GET("/admin/site/delete", indexAdmin.SiteDelete);
		/*SiteCity*/
		e.GET("/admin/site_city/index", indexAdmin.SiteCityIndex);
		e.GET("/admin/site_city/add", indexAdmin.SiteCityAdd);
		e.GET("/admin/site_city/save", indexAdmin.SiteCitySave);
		e.GET("/admin/site_city/status", indexAdmin.SiteCityStatus);
		e.GET("/admin/site_city/delete", indexAdmin.SiteCityDelete);
		/*SmsLog*/
		e.GET("/admin/sms_log/index", indexAdmin.SmsLogIndex);
		e.GET("/admin/sms_log/add", indexAdmin.SmsLogAdd);
		e.GET("/admin/sms_log/save", indexAdmin.SmsLogSave);
		e.GET("/admin/sms_log/status", indexAdmin.SmsLogStatus);
		e.GET("/admin/sms_log/delete", indexAdmin.SmsLogDelete);
		/*Table*/
		e.GET("/admin/table/index", indexAdmin.TableIndex);
		e.GET("/admin/table/add", indexAdmin.TableAdd);
		e.GET("/admin/table/save", indexAdmin.TableSave);
		e.GET("/admin/table/status", indexAdmin.TableStatus);
		e.GET("/admin/table/delete", indexAdmin.TableDelete);
		/*TableData*/
		e.GET("/admin/table_data/index", indexAdmin.TableDataIndex);
		e.GET("/admin/table_data/add", indexAdmin.TableDataAdd);
		e.GET("/admin/table_data/save", indexAdmin.TableDataSave);
		e.GET("/admin/table_data/status", indexAdmin.TableDataStatus);
		e.GET("/admin/table_data/delete", indexAdmin.TableDataDelete);
		/*TableDataComment*/
		e.GET("/admin/table_data_comment/index", indexAdmin.TableDataCommentIndex);
		e.GET("/admin/table_data_comment/add", indexAdmin.TableDataCommentAdd);
		e.GET("/admin/table_data_comment/save", indexAdmin.TableDataCommentSave);
		e.GET("/admin/table_data_comment/status", indexAdmin.TableDataCommentStatus);
		e.GET("/admin/table_data_comment/delete", indexAdmin.TableDataCommentDelete);
		/*TableFields*/
		e.GET("/admin/table_fields/index", indexAdmin.TableFieldsIndex);
		e.GET("/admin/table_fields/add", indexAdmin.TableFieldsAdd);
		e.GET("/admin/table_fields/save", indexAdmin.TableFieldsSave);
		e.GET("/admin/table_fields/status", indexAdmin.TableFieldsStatus);
		e.GET("/admin/table_fields/delete", indexAdmin.TableFieldsDelete);
		/*Tixian*/
		e.GET("/admin/tixian/index", indexAdmin.TixianIndex);
		e.GET("/admin/tixian/add", indexAdmin.TixianAdd);
		e.GET("/admin/tixian/save", indexAdmin.TixianSave);
		e.GET("/admin/tixian/status", indexAdmin.TixianStatus);
		e.GET("/admin/tixian/delete", indexAdmin.TixianDelete);
		/*TixianLog*/
		e.GET("/admin/tixian_log/index", indexAdmin.TixianLogIndex);
		e.GET("/admin/tixian_log/add", indexAdmin.TixianLogAdd);
		e.GET("/admin/tixian_log/save", indexAdmin.TixianLogSave);
		e.GET("/admin/tixian_log/delete", indexAdmin.TixianLogDelete);
		/*User*/
		e.GET("/admin/user/index", indexAdmin.UserIndex);
		e.GET("/admin/user/add", indexAdmin.UserAdd);
		e.GET("/admin/user/save", indexAdmin.UserSave);
		e.GET("/admin/user/status", indexAdmin.UserStatus);
		e.GET("/admin/user/delete", indexAdmin.UserDelete);
		/*UserAddress*/
		e.GET("/admin/user_address/index", indexAdmin.UserAddressIndex);
		e.GET("/admin/userress/add", indexAdmin.UserAddressAdd);
		e.GET("/admin/user_address/save", indexAdmin.UserAddressSave);
		e.GET("/admin/user_address/status", indexAdmin.UserAddressStatus);
		e.GET("/admin/user_address/delete", indexAdmin.UserAddressDelete);
		/*UserAuth*/
		e.GET("/admin/user_auth/index", indexAdmin.UserAuthIndex);
		e.GET("/admin/user_auth/add", indexAdmin.UserAuthAdd);
		e.GET("/admin/user_auth/save", indexAdmin.UserAuthSave);
		e.GET("/admin/user_auth/status", indexAdmin.UserAuthStatus);
		e.GET("/admin/user_auth/delete", indexAdmin.UserAuthDelete);
		/*UserAuthNew*/
		e.GET("/admin/user_auth_new/index", indexAdmin.UserAuthNewIndex);
		e.GET("/admin/user_auth_new/add", indexAdmin.UserAuthNewAdd);
		e.GET("/admin/user_auth_new/save", indexAdmin.UserAuthNewSave);
		e.GET("/admin/user_auth_new/status", indexAdmin.UserAuthNewStatus);
		e.GET("/admin/user_auth_new/delete", indexAdmin.UserAuthNewDelete);
		/*UserBlack*/
		e.GET("/admin/user_black/index", indexAdmin.UserBlackIndex);
		e.GET("/admin/user_black/add", indexAdmin.UserBlackAdd);
		e.GET("/admin/user_black/save", indexAdmin.UserBlackSave);
		e.GET("/admin/user_black/delete", indexAdmin.UserBlackDelete);
		/*UserExpand*/
		e.GET("/admin/user_expand/index", indexAdmin.UserExpandIndex);
		e.GET("/admin/user_expand/add", indexAdmin.UserExpandAdd);
		e.GET("/admin/user_expand/save", indexAdmin.UserExpandSave);
		e.GET("/admin/user_expand/delete", indexAdmin.UserExpandDelete);
		/*UserGroup*/
		e.GET("/admin/user_group/index", indexAdmin.UserGroupIndex);
		e.GET("/admin/user_group/add", indexAdmin.UserGroupAdd);
		e.GET("/admin/user_group/save", indexAdmin.UserGroupSave);
		e.GET("/admin/user_group/delete", indexAdmin.UserGroupDelete);
		/*UserGroupPeople*/
		e.GET("/admin/user_group_people/index", indexAdmin.UserGroupPeopleIndex);
		e.GET("/admin/user_group_people/add", indexAdmin.UserGroupPeopleAdd);
		e.GET("/admin/user_group_people/save", indexAdmin.UserGroupPeopleSave);
		e.GET("/admin/user_group_people/delete", indexAdmin.UserGroupPeopleDelete);
		/*UserInvitecode*/
		e.GET("/admin/user_invitecode/index", indexAdmin.UserInvitecodeIndex);
		e.GET("/admin/user_invitecode/add", indexAdmin.UserInvitecodeAdd);
		e.GET("/admin/user_invitecode/save", indexAdmin.UserInvitecodeSave);
		e.GET("/admin/user_invitecode/delete", indexAdmin.UserInvitecodeDelete);
		/*UserLastaddr*/
		e.GET("/admin/user_lastaddr/index", indexAdmin.UserLastaddrIndex);
		e.GET("/admin/user_lastaddr/add", indexAdmin.UserLastaddrAdd);
		e.GET("/admin/user_lastaddr/save", indexAdmin.UserLastaddrSave);
		e.GET("/admin/user_lastaddr/delete", indexAdmin.UserLastaddrDelete);
		/*UserPassword*/
		e.GET("/admin/user_password/index", indexAdmin.UserPasswordIndex);
		e.GET("/admin/user_password/add", indexAdmin.UserPasswordAdd);
		e.GET("/admin/user_password/save", indexAdmin.UserPasswordSave);
		e.GET("/admin/user_password/delete", indexAdmin.UserPasswordDelete);
		/*UserRank*/
		e.GET("/admin/user_rank/index", indexAdmin.UserRankIndex);
		e.GET("/admin/user_rank/add", indexAdmin.UserRankAdd);
		e.GET("/admin/user_rank/save", indexAdmin.UserRankSave);
		e.GET("/admin/user_rank/delete", indexAdmin.UserRankDelete);
		/*UserVip*/
		e.GET("/admin/user_vip/index", indexAdmin.UserVipIndex);
		e.GET("/admin/user_vip/add", indexAdmin.UserVipAdd);
		e.GET("/admin/user_vip/save", indexAdmin.UserVipSave);
		e.GET("/admin/user_vip/status", indexAdmin.UserVipStatus);
		e.GET("/admin/user_vip/delete", indexAdmin.UserVipDelete);
		/*Weixin*/
		e.GET("/admin/weixin/index", indexAdmin.WeixinIndex);
		e.GET("/admin/weixin/add", indexAdmin.WeixinAdd);
		e.GET("/admin/weixin/save", indexAdmin.WeixinSave);
		e.GET("/admin/weixin/status", indexAdmin.WeixinStatus);
		e.GET("/admin/weixin/recommend", indexAdmin.WeixinRecommend);
		e.GET("/admin/weixin/delete", indexAdmin.WeixinDelete);
		/*WeixinCommand*/
		e.GET("/admin/weixin_command/index", indexAdmin.WeixinCommandIndex);
		e.GET("/admin/weixin_command/add", indexAdmin.WeixinCommandAdd);
		e.GET("/admin/weixin_command/save", indexAdmin.WeixinCommandSave);
		e.GET("/admin/weixin_command/delete", indexAdmin.WeixinCommandDelete);
		/*WeixinMenu*/
		e.GET("/admin/weixin_menu/index", indexAdmin.WeixinMenuIndex);
		e.GET("/admin/weixin_menu/add", indexAdmin.WeixinMenuAdd);
		e.GET("/admin/weixin_menu/save", indexAdmin.WeixinMenuSave);
		e.GET("/admin/weixin_menu/delete", indexAdmin.WeixinMenuDelete);
		/*WeixinReply*/
		e.GET("/admin/weixin_reply/index", indexAdmin.WeixinReplyIndex);
		e.GET("/admin/weixin_reply/add", indexAdmin.WeixinReplyAdd);
		e.GET("/admin/weixin_reply/save", indexAdmin.WeixinReplySave);
		e.GET("/admin/weixin_reply/status", indexAdmin.WeixinReplyStatus);
		e.GET("/admin/weixin_reply/delete", indexAdmin.WeixinReplyDelete);
		/*WeixinSucai*/
		e.GET("/admin/weixin_sucai/index", indexAdmin.WeixinSucaiIndex);
		e.GET("/admin/weixin_sucai/add", indexAdmin.WeixinSucaiAdd);
		e.GET("/admin/weixin_sucai/save", indexAdmin.WeixinSucaiSave);
		e.GET("/admin/weixin_sucai/status", indexAdmin.WeixinSucaiStatus);
		e.GET("/admin/weixin_sucai/delete", indexAdmin.WeixinSucaiDelete);
		/*WeixinUser*/
		e.GET("/admin/weixin_user/index", indexAdmin.WeixinUserIndex);
		e.GET("/admin/weixin_user/add", indexAdmin.WeixinUserAdd);
		e.GET("/admin/weixin_user/save", indexAdmin.WeixinUserSave);
		e.GET("/admin/weixin_user/status", indexAdmin.WeixinUserStatus);
		e.GET("/admin/weixin_user/delete", indexAdmin.WeixinUserDelete);

	}

	