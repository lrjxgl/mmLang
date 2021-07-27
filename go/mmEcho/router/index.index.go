
	package router

	import (
		"app/index/indexIndex"
		"github.com/labstack/echo/v4"
	)

	func IndexIndexRouter(e *echo.Echo) {
		/*Ad*/
		e.GET("/ad/index", indexIndex.AdIndex);
		e.GET("/ad/list", indexIndex.AdList);
		e.GET("/ad/show", indexIndex.AdShow);
		/*AdTags*/
		e.GET("/ad_tags/index", indexIndex.AdTagsIndex);
		e.GET("/ad_tags/list", indexIndex.AdTagsList);
		e.GET("/ad_tags/show", indexIndex.AdTagsShow);
		/*Admin*/
		e.GET("/admin/index", indexIndex.AdminIndex);
		e.GET("/admin/list", indexIndex.AdminList);
		e.GET("/admin/show", indexIndex.AdminShow);
		/*AdminGroup*/
		e.GET("/admin_group/index", indexIndex.AdminGroupIndex);
		e.GET("/admin_group/list", indexIndex.AdminGroupList);
		e.GET("/admin_group/show", indexIndex.AdminGroupShow);
		/*Apppush*/
		e.GET("/apppush/index", indexIndex.ApppushIndex);
		e.GET("/apppush/list", indexIndex.ApppushList);
		e.GET("/apppush/show", indexIndex.ApppushShow);
		e.GET("/apppush/my", indexIndex.ApppushMy);
		e.GET("/apppush/add", indexIndex.ApppushAdd);
		e.GET("/apppush/save", indexIndex.ApppushSave);
		e.GET("/apppush/delete", indexIndex.ApppushDelete);
		/*ApppushPlan*/
		e.GET("/apppush_plan/index", indexIndex.ApppushPlanIndex);
		e.GET("/apppush_plan/list", indexIndex.ApppushPlanList);
		e.GET("/apppush_plan/show", indexIndex.ApppushPlanShow);
		/*Article*/
		e.GET("/article/index", indexIndex.ArticleIndex);
		e.GET("/article/list", indexIndex.ArticleList);
		e.GET("/article/show", indexIndex.ArticleShow);
		/*ArticleComment*/
		e.GET("/article_comment/index", indexIndex.ArticleCommentIndex);
		e.GET("/article_comment/list", indexIndex.ArticleCommentList);
		e.GET("/article_comment/show", indexIndex.ArticleCommentShow);
		e.GET("/article_comment/my", indexIndex.ArticleCommentMy);
		e.GET("/article_comment/add", indexIndex.ArticleCommentAdd);
		e.GET("/article_comment/save", indexIndex.ArticleCommentSave);
		e.GET("/article_comment/status", indexIndex.ArticleCommentStatus);
		e.GET("/article_comment/delete", indexIndex.ArticleCommentDelete);
		/*ArticleData*/
		e.GET("/article_data/index", indexIndex.ArticleDataIndex);
		e.GET("/article_data/list", indexIndex.ArticleDataList);
		e.GET("/article_data/show", indexIndex.ArticleDataShow);
		/*Attach*/
		e.GET("/attach/index", indexIndex.AttachIndex);
		e.GET("/attach/list", indexIndex.AttachList);
		e.GET("/attach/show", indexIndex.AttachShow);
		e.GET("/attach/my", indexIndex.AttachMy);
		e.GET("/attach/add", indexIndex.AttachAdd);
		e.GET("/attach/save", indexIndex.AttachSave);
		e.GET("/attach/status", indexIndex.AttachStatus);
		e.GET("/attach/delete", indexIndex.AttachDelete);
		/*Badip*/
		e.GET("/badip/index", indexIndex.BadipIndex);
		e.GET("/badip/list", indexIndex.BadipList);
		e.GET("/badip/show", indexIndex.BadipShow);
		/*Badphone*/
		e.GET("/badphone/index", indexIndex.BadphoneIndex);
		e.GET("/badphone/list", indexIndex.BadphoneList);
		e.GET("/badphone/show", indexIndex.BadphoneShow);
		/*Bankcard*/
		e.GET("/bankcard/index", indexIndex.BankcardIndex);
		e.GET("/bankcard/list", indexIndex.BankcardList);
		e.GET("/bankcard/show", indexIndex.BankcardShow);
		e.GET("/bankcard/my", indexIndex.BankcardMy);
		e.GET("/bankcard/add", indexIndex.BankcardAdd);
		e.GET("/bankcard/save", indexIndex.BankcardSave);
		e.GET("/bankcard/status", indexIndex.BankcardStatus);
		e.GET("/bankcard/delete", indexIndex.BankcardDelete);
		/*Blacklist*/
		e.GET("/blacklist/index", indexIndex.BlacklistIndex);
		e.GET("/blacklist/list", indexIndex.BlacklistList);
		e.GET("/blacklist/show", indexIndex.BlacklistShow);
		e.GET("/blacklist/my", indexIndex.BlacklistMy);
		e.GET("/blacklist/add", indexIndex.BlacklistAdd);
		e.GET("/blacklist/save", indexIndex.BlacklistSave);
		e.GET("/blacklist/delete", indexIndex.BlacklistDelete);
		/*BlacklistPost*/
		e.GET("/blacklist_post/index", indexIndex.BlacklistPostIndex);
		e.GET("/blacklist_post/list", indexIndex.BlacklistPostList);
		e.GET("/blacklist_post/show", indexIndex.BlacklistPostShow);
		e.GET("/blacklist_post/my", indexIndex.BlacklistPostMy);
		e.GET("/blacklist_post/add", indexIndex.BlacklistPostAdd);
		e.GET("/blacklist_post/save", indexIndex.BlacklistPostSave);
		e.GET("/blacklist_post/delete", indexIndex.BlacklistPostDelete);
		/*Category*/
		e.GET("/category/index", indexIndex.CategoryIndex);
		e.GET("/category/list", indexIndex.CategoryList);
		e.GET("/category/show", indexIndex.CategoryShow);
		/*Checkin*/
		e.GET("/checkin/index", indexIndex.CheckinIndex);
		e.GET("/checkin/list", indexIndex.CheckinList);
		e.GET("/checkin/show", indexIndex.CheckinShow);
		e.GET("/checkin/my", indexIndex.CheckinMy);
		e.GET("/checkin/add", indexIndex.CheckinAdd);
		e.GET("/checkin/save", indexIndex.CheckinSave);
		e.GET("/checkin/delete", indexIndex.CheckinDelete);
		/*CheckinIndex*/
		e.GET("/checkin/index", indexIndex.CheckinIndexIndex);
		e.GET("/checkin_index/list", indexIndex.CheckinIndexList);
		e.GET("/checkin_index/show", indexIndex.CheckinIndexShow);
		e.GET("/checkin_index/my", indexIndex.CheckinIndexMy);
		e.GET("/checkin_index/add", indexIndex.CheckinIndexAdd);
		e.GET("/checkin_index/save", indexIndex.CheckinIndexSave);
		e.GET("/checkin_index/delete", indexIndex.CheckinIndexDelete);
		/*City*/
		e.GET("/city/index", indexIndex.CityIndex);
		e.GET("/city/list", indexIndex.CityList);
		e.GET("/city/show", indexIndex.CityShow);
		/*Config*/
		e.GET("/config/index", indexIndex.ConfigIndex);
		e.GET("/config/list", indexIndex.ConfigList);
		e.GET("/config/show", indexIndex.ConfigShow);
		/*Coupon*/
		e.GET("/coupon/index", indexIndex.CouponIndex);
		e.GET("/coupon/list", indexIndex.CouponList);
		e.GET("/coupon/show", indexIndex.CouponShow);
		/*CouponUser*/
		e.GET("/coupon_user/index", indexIndex.CouponUserIndex);
		e.GET("/coupon_user/list", indexIndex.CouponUserList);
		e.GET("/coupon_user/show", indexIndex.CouponUserShow);
		e.GET("/coupon_user/my", indexIndex.CouponUserMy);
		e.GET("/coupon_user/add", indexIndex.CouponUserAdd);
		e.GET("/coupon_user/save", indexIndex.CouponUserSave);
		e.GET("/coupon_user/status", indexIndex.CouponUserStatus);
		e.GET("/coupon_user/delete", indexIndex.CouponUserDelete);
		/*Crontab*/
		e.GET("/crontab/index", indexIndex.CrontabIndex);
		e.GET("/crontab/list", indexIndex.CrontabList);
		e.GET("/crontab/show", indexIndex.CrontabShow);
		/*Dataapi*/
		e.GET("/dataapi/index", indexIndex.DataapiIndex);
		e.GET("/dataapi/list", indexIndex.DataapiList);
		e.GET("/dataapi/show", indexIndex.DataapiShow);
		/*Daysn*/
		e.GET("/daysn/index", indexIndex.DaysnIndex);
		e.GET("/daysn/list", indexIndex.DaysnList);
		e.GET("/daysn/show", indexIndex.DaysnShow);
		/*Dbcache*/
		e.GET("/dbcache/index", indexIndex.DbcacheIndex);
		e.GET("/dbcache/list", indexIndex.DbcacheList);
		e.GET("/dbcache/show", indexIndex.DbcacheShow);
		/*Dbsession*/
		e.GET("/dbsession/index", indexIndex.DbsessionIndex);
		e.GET("/dbsession/list", indexIndex.DbsessionList);
		e.GET("/dbsession/show", indexIndex.DbsessionShow);
		/*District*/
		e.GET("/district/index", indexIndex.DistrictIndex);
		e.GET("/district/list", indexIndex.DistrictList);
		e.GET("/district/show", indexIndex.DistrictShow);
		/*ExpressFee*/
		e.GET("/express_fee/index", indexIndex.ExpressFeeIndex);
		e.GET("/express_fee/list", indexIndex.ExpressFeeList);
		e.GET("/express_fee/show", indexIndex.ExpressFeeShow);
		/*ExpressFeeCity*/
		e.GET("/express_fee_city/index", indexIndex.ExpressFeeCityIndex);
		e.GET("/express_fee_city/list", indexIndex.ExpressFeeCityList);
		e.GET("/express_fee_city/show", indexIndex.ExpressFeeCityShow);
		/*Fav*/
		e.GET("/fav/index", indexIndex.FavIndex);
		e.GET("/fav/list", indexIndex.FavList);
		e.GET("/fav/show", indexIndex.FavShow);
		e.GET("/fav/my", indexIndex.FavMy);
		e.GET("/fav/add", indexIndex.FavAdd);
		e.GET("/fav/save", indexIndex.FavSave);
		e.GET("/fav/delete", indexIndex.FavDelete);
		/*Follow*/
		e.GET("/follow/index", indexIndex.FollowIndex);
		e.GET("/follow/list", indexIndex.FollowList);
		e.GET("/follow/show", indexIndex.FollowShow);
		e.GET("/follow/my", indexIndex.FollowMy);
		e.GET("/follow/add", indexIndex.FollowAdd);
		e.GET("/follow/save", indexIndex.FollowSave);
		e.GET("/follow/status", indexIndex.FollowStatus);
		e.GET("/follow/delete", indexIndex.FollowDelete);
		/*Followed*/
		e.GET("/followed/index", indexIndex.FollowedIndex);
		e.GET("/followed/list", indexIndex.FollowedList);
		e.GET("/followed/show", indexIndex.FollowedShow);
		e.GET("/followed/my", indexIndex.FollowedMy);
		e.GET("/followed/add", indexIndex.FollowedAdd);
		e.GET("/followed/save", indexIndex.FollowedSave);
		e.GET("/followed/status", indexIndex.FollowedStatus);
		e.GET("/followed/delete", indexIndex.FollowedDelete);
		/*Friend*/
		e.GET("/friend/index", indexIndex.FriendIndex);
		e.GET("/friend/list", indexIndex.FriendList);
		e.GET("/friend/show", indexIndex.FriendShow);
		e.GET("/friend/my", indexIndex.FriendMy);
		e.GET("/friend/add", indexIndex.FriendAdd);
		e.GET("/friend/save", indexIndex.FriendSave);
		e.GET("/friend/delete", indexIndex.FriendDelete);
		/*FriendApply*/
		e.GET("/friend_apply/index", indexIndex.FriendApplyIndex);
		e.GET("/friend_apply/list", indexIndex.FriendApplyList);
		e.GET("/friend_apply/show", indexIndex.FriendApplyShow);
		e.GET("/friend_apply/my", indexIndex.FriendApplyMy);
		e.GET("/friend_apply/add", indexIndex.FriendApplyAdd);
		e.GET("/friend_apply/save", indexIndex.FriendApplySave);
		e.GET("/friend_apply/delete", indexIndex.FriendApplyDelete);
		/*GoldLog*/
		e.GET("/gold_log/index", indexIndex.GoldLogIndex);
		e.GET("/gold_log/list", indexIndex.GoldLogList);
		e.GET("/gold_log/show", indexIndex.GoldLogShow);
		e.GET("/gold_log/my", indexIndex.GoldLogMy);
		e.GET("/gold_log/add", indexIndex.GoldLogAdd);
		e.GET("/gold_log/save", indexIndex.GoldLogSave);
		e.GET("/gold_log/delete", indexIndex.GoldLogDelete);
		/*GradeLog*/
		e.GET("/grade_log/index", indexIndex.GradeLogIndex);
		e.GET("/grade_log/list", indexIndex.GradeLogList);
		e.GET("/grade_log/show", indexIndex.GradeLogShow);
		e.GET("/grade_log/my", indexIndex.GradeLogMy);
		e.GET("/grade_log/add", indexIndex.GradeLogAdd);
		e.GET("/grade_log/save", indexIndex.GradeLogSave);
		e.GET("/grade_log/delete", indexIndex.GradeLogDelete);
		/*Guest*/
		e.GET("/guest/index", indexIndex.GuestIndex);
		e.GET("/guest/list", indexIndex.GuestList);
		e.GET("/guest/show", indexIndex.GuestShow);
		e.GET("/guest/my", indexIndex.GuestMy);
		e.GET("/guest/add", indexIndex.GuestAdd);
		e.GET("/guest/save", indexIndex.GuestSave);
		e.GET("/guest/status", indexIndex.GuestStatus);
		e.GET("/guest/delete", indexIndex.GuestDelete);
		/*Imgs*/
		e.GET("/imgs/index", indexIndex.ImgsIndex);
		e.GET("/imgs/list", indexIndex.ImgsList);
		e.GET("/imgs/show", indexIndex.ImgsShow);
		e.GET("/imgs/my", indexIndex.ImgsMy);
		e.GET("/imgs/add", indexIndex.ImgsAdd);
		e.GET("/imgs/save", indexIndex.ImgsSave);
		e.GET("/imgs/delete", indexIndex.ImgsDelete);
		/*Invite*/
		e.GET("/invite/index", indexIndex.InviteIndex);
		e.GET("/invite/list", indexIndex.InviteList);
		e.GET("/invite/show", indexIndex.InviteShow);
		e.GET("/invite/my", indexIndex.InviteMy);
		e.GET("/invite/add", indexIndex.InviteAdd);
		e.GET("/invite/save", indexIndex.InviteSave);
		e.GET("/invite/delete", indexIndex.InviteDelete);
		/*InviteAccount*/
		e.GET("/invite_account/index", indexIndex.InviteAccountIndex);
		e.GET("/invite_account/list", indexIndex.InviteAccountList);
		e.GET("/invite_account/show", indexIndex.InviteAccountShow);
		e.GET("/invite_account/my", indexIndex.InviteAccountMy);
		e.GET("/invite_account/add", indexIndex.InviteAccountAdd);
		e.GET("/invite_account/save", indexIndex.InviteAccountSave);
		e.GET("/invite_account/delete", indexIndex.InviteAccountDelete);
		/*InviteAccountLog*/
		e.GET("/invite_account_log/index", indexIndex.InviteAccountLogIndex);
		e.GET("/invite_account_log/list", indexIndex.InviteAccountLogList);
		e.GET("/invite_account_log/show", indexIndex.InviteAccountLogShow);
		e.GET("/invite_account_log/my", indexIndex.InviteAccountLogMy);
		e.GET("/invite_account_log/add", indexIndex.InviteAccountLogAdd);
		e.GET("/invite_account_log/save", indexIndex.InviteAccountLogSave);
		e.GET("/invite_account_log/delete", indexIndex.InviteAccountLogDelete);
		/*KefuMsg*/
		e.GET("/kefu_msg/index", indexIndex.KefuMsgIndex);
		e.GET("/kefu_msg/list", indexIndex.KefuMsgList);
		e.GET("/kefu_msg/show", indexIndex.KefuMsgShow);
		e.GET("/kefu_msg/my", indexIndex.KefuMsgMy);
		e.GET("/kefu_msg/add", indexIndex.KefuMsgAdd);
		e.GET("/kefu_msg/save", indexIndex.KefuMsgSave);
		e.GET("/kefu_msg/status", indexIndex.KefuMsgStatus);
		e.GET("/kefu_msg/delete", indexIndex.KefuMsgDelete);
		/*KefuMsgindex*/
		e.GET("/kefu_msgindex/index", indexIndex.KefuMsgindexIndex);
		e.GET("/kefu_msgindex/list", indexIndex.KefuMsgindexList);
		e.GET("/kefu_msgindex/show", indexIndex.KefuMsgindexShow);
		e.GET("/kefu_msgindex/my", indexIndex.KefuMsgindexMy);
		e.GET("/kefu_msgindex/add", indexIndex.KefuMsgindexAdd);
		e.GET("/kefu_msgindex/save", indexIndex.KefuMsgindexSave);
		e.GET("/kefu_msgindex/status", indexIndex.KefuMsgindexStatus);
		e.GET("/kefu_msgindex/delete", indexIndex.KefuMsgindexDelete);
		/*Love*/
		e.GET("/love/index", indexIndex.LoveIndex);
		e.GET("/love/list", indexIndex.LoveList);
		e.GET("/love/show", indexIndex.LoveShow);
		e.GET("/love/my", indexIndex.LoveMy);
		e.GET("/love/add", indexIndex.LoveAdd);
		e.GET("/love/save", indexIndex.LoveSave);
		e.GET("/love/delete", indexIndex.LoveDelete);
		/*Maxid*/
		e.GET("/maxid/index", indexIndex.MaxidIndex);
		e.GET("/maxid/list", indexIndex.MaxidList);
		e.GET("/maxid/show", indexIndex.MaxidShow);
		/*Model*/
		e.GET("/model/index", indexIndex.ModelIndex);
		e.GET("/model/list", indexIndex.ModelList);
		e.GET("/model/show", indexIndex.ModelShow);
		/*ModelIndex*/
		e.GET("/model/index", indexIndex.ModelIndexIndex);
		e.GET("/model_index/list", indexIndex.ModelIndexList);
		e.GET("/model_index/show", indexIndex.ModelIndexShow);
		/*Module*/
		e.GET("/module/index", indexIndex.ModuleIndex);
		e.GET("/module/list", indexIndex.ModuleList);
		e.GET("/module/show", indexIndex.ModuleShow);
		/*Notice*/
		e.GET("/notice/index", indexIndex.NoticeIndex);
		e.GET("/notice/list", indexIndex.NoticeList);
		e.GET("/notice/show", indexIndex.NoticeShow);
		e.GET("/notice/my", indexIndex.NoticeMy);
		e.GET("/notice/add", indexIndex.NoticeAdd);
		e.GET("/notice/save", indexIndex.NoticeSave);
		e.GET("/notice/status", indexIndex.NoticeStatus);
		e.GET("/notice/delete", indexIndex.NoticeDelete);
		/*OpenAlioss*/
		e.GET("/open_alioss/index", indexIndex.OpenAliossIndex);
		e.GET("/open_alioss/list", indexIndex.OpenAliossList);
		e.GET("/open_alioss/show", indexIndex.OpenAliossShow);
		/*OpenAlipay*/
		e.GET("/open_alipay/index", indexIndex.OpenAlipayIndex);
		e.GET("/open_alipay/list", indexIndex.OpenAlipayList);
		e.GET("/open_alipay/show", indexIndex.OpenAlipayShow);
		/*OpenAlipayApp*/
		e.GET("/open_alipay_app/index", indexIndex.OpenAlipayAppIndex);
		e.GET("/open_alipay_app/list", indexIndex.OpenAlipayAppList);
		e.GET("/open_alipay_app/show", indexIndex.OpenAlipayAppShow);
		/*OpenAlipayMini*/
		e.GET("/open_alipay_mini/index", indexIndex.OpenAlipayMiniIndex);
		e.GET("/open_alipay_mini/list", indexIndex.OpenAlipayMiniList);
		e.GET("/open_alipay_mini/show", indexIndex.OpenAlipayMiniShow);
		/*OpenBaidu*/
		e.GET("/open_baidu/index", indexIndex.OpenBaiduIndex);
		e.GET("/open_baidu/list", indexIndex.OpenBaiduList);
		e.GET("/open_baidu/show", indexIndex.OpenBaiduShow);
		/*OpenQq*/
		e.GET("/open_qq/index", indexIndex.OpenQqIndex);
		e.GET("/open_qq/list", indexIndex.OpenQqList);
		e.GET("/open_qq/show", indexIndex.OpenQqShow);
		/*OpenToutiao*/
		e.GET("/open_toutiao/index", indexIndex.OpenToutiaoIndex);
		e.GET("/open_toutiao/list", indexIndex.OpenToutiaoList);
		e.GET("/open_toutiao/show", indexIndex.OpenToutiaoShow);
		/*OpenWxapp*/
		e.GET("/open_wxapp/index", indexIndex.OpenWxappIndex);
		e.GET("/open_wxapp/list", indexIndex.OpenWxappList);
		e.GET("/open_wxapp/show", indexIndex.OpenWxappShow);
		/*OpenWxnative*/
		e.GET("/open_wxnative/index", indexIndex.OpenWxnativeIndex);
		e.GET("/open_wxnative/list", indexIndex.OpenWxnativeList);
		e.GET("/open_wxnative/show", indexIndex.OpenWxnativeShow);
		/*Openlogin*/
		e.GET("/openlogin/index", indexIndex.OpenloginIndex);
		e.GET("/openlogin/list", indexIndex.OpenloginList);
		e.GET("/openlogin/show", indexIndex.OpenloginShow);
		e.GET("/openlogin/my", indexIndex.OpenloginMy);
		e.GET("/openlogin/add", indexIndex.OpenloginAdd);
		e.GET("/openlogin/save", indexIndex.OpenloginSave);
		e.GET("/openlogin/delete", indexIndex.OpenloginDelete);
		/*Pagetpl*/
		e.GET("/pagetpl/index", indexIndex.PagetplIndex);
		e.GET("/pagetpl/list", indexIndex.PagetplList);
		e.GET("/pagetpl/show", indexIndex.PagetplShow);
		/*PayLog*/
		e.GET("/pay_log/index", indexIndex.PayLogIndex);
		e.GET("/pay_log/list", indexIndex.PayLogList);
		e.GET("/pay_log/show", indexIndex.PayLogShow);
		e.GET("/pay_log/my", indexIndex.PayLogMy);
		e.GET("/pay_log/add", indexIndex.PayLogAdd);
		e.GET("/pay_log/save", indexIndex.PayLogSave);
		e.GET("/pay_log/delete", indexIndex.PayLogDelete);
		/*Permission*/
		e.GET("/permission/index", indexIndex.PermissionIndex);
		e.GET("/permission/list", indexIndex.PermissionList);
		e.GET("/permission/show", indexIndex.PermissionShow);
		/*Pm*/
		e.GET("/pm/index", indexIndex.PmIndex);
		e.GET("/pm/list", indexIndex.PmList);
		e.GET("/pm/show", indexIndex.PmShow);
		e.GET("/pm/my", indexIndex.PmMy);
		e.GET("/pm/add", indexIndex.PmAdd);
		e.GET("/pm/save", indexIndex.PmSave);
		e.GET("/pm/status", indexIndex.PmStatus);
		e.GET("/pm/delete", indexIndex.PmDelete);
		/*PmIndex*/
		e.GET("/pm/index", indexIndex.PmIndexIndex);
		e.GET("/pm_index/list", indexIndex.PmIndexList);
		e.GET("/pm_index/show", indexIndex.PmIndexShow);
		e.GET("/pm_index/my", indexIndex.PmIndexMy);
		e.GET("/pm_index/add", indexIndex.PmIndexAdd);
		e.GET("/pm_index/save", indexIndex.PmIndexSave);
		e.GET("/pm_index/status", indexIndex.PmIndexStatus);
		e.GET("/pm_index/delete", indexIndex.PmIndexDelete);
		/*Pv*/
		e.GET("/pv/index", indexIndex.PvIndex);
		e.GET("/pv/list", indexIndex.PvList);
		e.GET("/pv/show", indexIndex.PvShow);
		/*PvDay*/
		e.GET("/pv_day/index", indexIndex.PvDayIndex);
		e.GET("/pv_day/list", indexIndex.PvDayList);
		e.GET("/pv_day/show", indexIndex.PvDayShow);
		/*PvIp*/
		e.GET("/pv_ip/index", indexIndex.PvIpIndex);
		e.GET("/pv_ip/list", indexIndex.PvIpList);
		e.GET("/pv_ip/show", indexIndex.PvIpShow);
		/*PvUv*/
		e.GET("/pv_uv/index", indexIndex.PvUvIndex);
		e.GET("/pv_uv/list", indexIndex.PvUvList);
		e.GET("/pv_uv/show", indexIndex.PvUvShow);
		/*Queue*/
		e.GET("/queue/index", indexIndex.QueueIndex);
		e.GET("/queue/list", indexIndex.QueueList);
		e.GET("/queue/show", indexIndex.QueueShow);
		/*Recharge*/
		e.GET("/recharge/index", indexIndex.RechargeIndex);
		e.GET("/recharge/list", indexIndex.RechargeList);
		e.GET("/recharge/show", indexIndex.RechargeShow);
		e.GET("/recharge/my", indexIndex.RechargeMy);
		e.GET("/recharge/add", indexIndex.RechargeAdd);
		e.GET("/recharge/save", indexIndex.RechargeSave);
		e.GET("/recharge/status", indexIndex.RechargeStatus);
		e.GET("/recharge/delete", indexIndex.RechargeDelete);
		/*Refund*/
		e.GET("/refund/index", indexIndex.RefundIndex);
		e.GET("/refund/list", indexIndex.RefundList);
		e.GET("/refund/show", indexIndex.RefundShow);
		e.GET("/refund/my", indexIndex.RefundMy);
		e.GET("/refund/add", indexIndex.RefundAdd);
		e.GET("/refund/save", indexIndex.RefundSave);
		e.GET("/refund/status", indexIndex.RefundStatus);
		e.GET("/refund/delete", indexIndex.RefundDelete);
		/*RefundApply*/
		e.GET("/refund_apply/index", indexIndex.RefundApplyIndex);
		e.GET("/refund_apply/list", indexIndex.RefundApplyList);
		e.GET("/refund_apply/show", indexIndex.RefundApplyShow);
		e.GET("/refund_apply/my", indexIndex.RefundApplyMy);
		e.GET("/refund_apply/add", indexIndex.RefundApplyAdd);
		e.GET("/refund_apply/save", indexIndex.RefundApplySave);
		e.GET("/refund_apply/status", indexIndex.RefundApplyStatus);
		e.GET("/refund_apply/delete", indexIndex.RefundApplyDelete);
		/*Seo*/
		e.GET("/seo/index", indexIndex.SeoIndex);
		e.GET("/seo/list", indexIndex.SeoList);
		e.GET("/seo/show", indexIndex.SeoShow);
		/*Sgpay*/
		e.GET("/sgpay/index", indexIndex.SgpayIndex);
		e.GET("/sgpay/list", indexIndex.SgpayList);
		e.GET("/sgpay/show", indexIndex.SgpayShow);
		e.GET("/sgpay/my", indexIndex.SgpayMy);
		e.GET("/sgpay/add", indexIndex.SgpayAdd);
		e.GET("/sgpay/save", indexIndex.SgpaySave);
		e.GET("/sgpay/status", indexIndex.SgpayStatus);
		e.GET("/sgpay/delete", indexIndex.SgpayDelete);
		/*Site*/
		e.GET("/site/index", indexIndex.SiteIndex);
		e.GET("/site/list", indexIndex.SiteList);
		e.GET("/site/show", indexIndex.SiteShow);
		/*SiteCity*/
		e.GET("/site_city/index", indexIndex.SiteCityIndex);
		e.GET("/site_city/list", indexIndex.SiteCityList);
		e.GET("/site_city/show", indexIndex.SiteCityShow);
		/*SmsLog*/
		e.GET("/sms_log/index", indexIndex.SmsLogIndex);
		e.GET("/sms_log/list", indexIndex.SmsLogList);
		e.GET("/sms_log/show", indexIndex.SmsLogShow);
		e.GET("/sms_log/my", indexIndex.SmsLogMy);
		e.GET("/sms_log/add", indexIndex.SmsLogAdd);
		e.GET("/sms_log/save", indexIndex.SmsLogSave);
		e.GET("/sms_log/status", indexIndex.SmsLogStatus);
		e.GET("/sms_log/delete", indexIndex.SmsLogDelete);
		/*Table*/
		e.GET("/table/index", indexIndex.TableIndex);
		e.GET("/table/list", indexIndex.TableList);
		e.GET("/table/show", indexIndex.TableShow);
		/*TableData*/
		e.GET("/table_data/index", indexIndex.TableDataIndex);
		e.GET("/table_data/list", indexIndex.TableDataList);
		e.GET("/table_data/show", indexIndex.TableDataShow);
		e.GET("/table_data/my", indexIndex.TableDataMy);
		e.GET("/table_data/add", indexIndex.TableDataAdd);
		e.GET("/table_data/save", indexIndex.TableDataSave);
		e.GET("/table_data/status", indexIndex.TableDataStatus);
		e.GET("/table_data/delete", indexIndex.TableDataDelete);
		/*TableDataComment*/
		e.GET("/table_data_comment/index", indexIndex.TableDataCommentIndex);
		e.GET("/table_data_comment/list", indexIndex.TableDataCommentList);
		e.GET("/table_data_comment/show", indexIndex.TableDataCommentShow);
		e.GET("/table_data_comment/my", indexIndex.TableDataCommentMy);
		e.GET("/table_data_comment/add", indexIndex.TableDataCommentAdd);
		e.GET("/table_data_comment/save", indexIndex.TableDataCommentSave);
		e.GET("/table_data_comment/status", indexIndex.TableDataCommentStatus);
		e.GET("/table_data_comment/delete", indexIndex.TableDataCommentDelete);
		/*TableFields*/
		e.GET("/table_fields/index", indexIndex.TableFieldsIndex);
		e.GET("/table_fields/list", indexIndex.TableFieldsList);
		e.GET("/table_fields/show", indexIndex.TableFieldsShow);
		/*Tixian*/
		e.GET("/tixian/index", indexIndex.TixianIndex);
		e.GET("/tixian/list", indexIndex.TixianList);
		e.GET("/tixian/show", indexIndex.TixianShow);
		/*TixianLog*/
		e.GET("/tixian_log/index", indexIndex.TixianLogIndex);
		e.GET("/tixian_log/list", indexIndex.TixianLogList);
		e.GET("/tixian_log/show", indexIndex.TixianLogShow);
		/*User*/
		e.GET("/user/index", indexIndex.UserIndex);
		e.GET("/user/list", indexIndex.UserList);
		e.GET("/user/show", indexIndex.UserShow);
		e.GET("/user/my", indexIndex.UserMy);
		e.GET("/user/add", indexIndex.UserAdd);
		e.GET("/user/save", indexIndex.UserSave);
		e.GET("/user/status", indexIndex.UserStatus);
		e.GET("/user/delete", indexIndex.UserDelete);
		/*UserAddress*/
		e.GET("/user_address/index", indexIndex.UserAddressIndex);
		e.GET("/user_address/list", indexIndex.UserAddressList);
		e.GET("/user_address/show", indexIndex.UserAddressShow);
		e.GET("/user_address/my", indexIndex.UserAddressMy);
		e.GET("/userress/add", indexIndex.UserAddressAdd);
		e.GET("/user_address/save", indexIndex.UserAddressSave);
		e.GET("/user_address/status", indexIndex.UserAddressStatus);
		e.GET("/user_address/delete", indexIndex.UserAddressDelete);
		/*UserAuth*/
		e.GET("/user_auth/index", indexIndex.UserAuthIndex);
		e.GET("/user_auth/list", indexIndex.UserAuthList);
		e.GET("/user_auth/show", indexIndex.UserAuthShow);
		e.GET("/user_auth/my", indexIndex.UserAuthMy);
		e.GET("/user_auth/add", indexIndex.UserAuthAdd);
		e.GET("/user_auth/save", indexIndex.UserAuthSave);
		e.GET("/user_auth/status", indexIndex.UserAuthStatus);
		e.GET("/user_auth/delete", indexIndex.UserAuthDelete);
		/*UserAuthNew*/
		e.GET("/user_auth_new/index", indexIndex.UserAuthNewIndex);
		e.GET("/user_auth_new/list", indexIndex.UserAuthNewList);
		e.GET("/user_auth_new/show", indexIndex.UserAuthNewShow);
		e.GET("/user_auth_new/my", indexIndex.UserAuthNewMy);
		e.GET("/user_auth_new/add", indexIndex.UserAuthNewAdd);
		e.GET("/user_auth_new/save", indexIndex.UserAuthNewSave);
		e.GET("/user_auth_new/status", indexIndex.UserAuthNewStatus);
		e.GET("/user_auth_new/delete", indexIndex.UserAuthNewDelete);
		/*UserBlack*/
		e.GET("/user_black/index", indexIndex.UserBlackIndex);
		e.GET("/user_black/list", indexIndex.UserBlackList);
		e.GET("/user_black/show", indexIndex.UserBlackShow);
		e.GET("/user_black/my", indexIndex.UserBlackMy);
		e.GET("/user_black/add", indexIndex.UserBlackAdd);
		e.GET("/user_black/save", indexIndex.UserBlackSave);
		e.GET("/user_black/delete", indexIndex.UserBlackDelete);
		/*UserExpand*/
		e.GET("/user_expand/index", indexIndex.UserExpandIndex);
		e.GET("/user_expand/list", indexIndex.UserExpandList);
		e.GET("/user_expand/show", indexIndex.UserExpandShow);
		/*UserGroup*/
		e.GET("/user_group/index", indexIndex.UserGroupIndex);
		e.GET("/user_group/list", indexIndex.UserGroupList);
		e.GET("/user_group/show", indexIndex.UserGroupShow);
		/*UserGroupPeople*/
		e.GET("/user_group_people/index", indexIndex.UserGroupPeopleIndex);
		e.GET("/user_group_people/list", indexIndex.UserGroupPeopleList);
		e.GET("/user_group_people/show", indexIndex.UserGroupPeopleShow);
		e.GET("/user_group_people/my", indexIndex.UserGroupPeopleMy);
		e.GET("/user_group_people/add", indexIndex.UserGroupPeopleAdd);
		e.GET("/user_group_people/save", indexIndex.UserGroupPeopleSave);
		e.GET("/user_group_people/delete", indexIndex.UserGroupPeopleDelete);
		/*UserInvitecode*/
		e.GET("/user_invitecode/index", indexIndex.UserInvitecodeIndex);
		e.GET("/user_invitecode/list", indexIndex.UserInvitecodeList);
		e.GET("/user_invitecode/show", indexIndex.UserInvitecodeShow);
		e.GET("/user_invitecode/my", indexIndex.UserInvitecodeMy);
		e.GET("/user_invitecode/add", indexIndex.UserInvitecodeAdd);
		e.GET("/user_invitecode/save", indexIndex.UserInvitecodeSave);
		e.GET("/user_invitecode/delete", indexIndex.UserInvitecodeDelete);
		/*UserLastaddr*/
		e.GET("/user_lastaddr/index", indexIndex.UserLastaddrIndex);
		e.GET("/user_lastaddr/list", indexIndex.UserLastaddrList);
		e.GET("/user_lastaddr/show", indexIndex.UserLastaddrShow);
		e.GET("/user_lastaddr/my", indexIndex.UserLastaddrMy);
		e.GET("/user_lastaddr/add", indexIndex.UserLastaddrAdd);
		e.GET("/user_lastaddr/save", indexIndex.UserLastaddrSave);
		e.GET("/user_lastaddr/delete", indexIndex.UserLastaddrDelete);
		/*UserPassword*/
		e.GET("/user_password/index", indexIndex.UserPasswordIndex);
		e.GET("/user_password/list", indexIndex.UserPasswordList);
		e.GET("/user_password/show", indexIndex.UserPasswordShow);
		e.GET("/user_password/my", indexIndex.UserPasswordMy);
		e.GET("/user_password/add", indexIndex.UserPasswordAdd);
		e.GET("/user_password/save", indexIndex.UserPasswordSave);
		e.GET("/user_password/delete", indexIndex.UserPasswordDelete);
		/*UserRank*/
		e.GET("/user_rank/index", indexIndex.UserRankIndex);
		e.GET("/user_rank/list", indexIndex.UserRankList);
		e.GET("/user_rank/show", indexIndex.UserRankShow);
		/*UserVip*/
		e.GET("/user_vip/index", indexIndex.UserVipIndex);
		e.GET("/user_vip/list", indexIndex.UserVipList);
		e.GET("/user_vip/show", indexIndex.UserVipShow);
		e.GET("/user_vip/my", indexIndex.UserVipMy);
		e.GET("/user_vip/add", indexIndex.UserVipAdd);
		e.GET("/user_vip/save", indexIndex.UserVipSave);
		e.GET("/user_vip/status", indexIndex.UserVipStatus);
		e.GET("/user_vip/delete", indexIndex.UserVipDelete);
		/*Weixin*/
		e.GET("/weixin/index", indexIndex.WeixinIndex);
		e.GET("/weixin/list", indexIndex.WeixinList);
		e.GET("/weixin/show", indexIndex.WeixinShow);
		e.GET("/weixin/my", indexIndex.WeixinMy);
		e.GET("/weixin/add", indexIndex.WeixinAdd);
		e.GET("/weixin/save", indexIndex.WeixinSave);
		e.GET("/weixin/status", indexIndex.WeixinStatus);
		e.GET("/weixin/delete", indexIndex.WeixinDelete);
		/*WeixinCommand*/
		e.GET("/weixin_command/index", indexIndex.WeixinCommandIndex);
		e.GET("/weixin_command/list", indexIndex.WeixinCommandList);
		e.GET("/weixin_command/show", indexIndex.WeixinCommandShow);
		e.GET("/weixin_command/my", indexIndex.WeixinCommandMy);
		e.GET("/weixin_command/add", indexIndex.WeixinCommandAdd);
		e.GET("/weixin_command/save", indexIndex.WeixinCommandSave);
		e.GET("/weixin_command/delete", indexIndex.WeixinCommandDelete);
		/*WeixinMenu*/
		e.GET("/weixin_menu/index", indexIndex.WeixinMenuIndex);
		e.GET("/weixin_menu/list", indexIndex.WeixinMenuList);
		e.GET("/weixin_menu/show", indexIndex.WeixinMenuShow);
		e.GET("/weixin_menu/my", indexIndex.WeixinMenuMy);
		e.GET("/weixin_menu/add", indexIndex.WeixinMenuAdd);
		e.GET("/weixin_menu/save", indexIndex.WeixinMenuSave);
		e.GET("/weixin_menu/delete", indexIndex.WeixinMenuDelete);
		/*WeixinReply*/
		e.GET("/weixin_reply/index", indexIndex.WeixinReplyIndex);
		e.GET("/weixin_reply/list", indexIndex.WeixinReplyList);
		e.GET("/weixin_reply/show", indexIndex.WeixinReplyShow);
		/*WeixinSucai*/
		e.GET("/weixin_sucai/index", indexIndex.WeixinSucaiIndex);
		e.GET("/weixin_sucai/list", indexIndex.WeixinSucaiList);
		e.GET("/weixin_sucai/show", indexIndex.WeixinSucaiShow);
		e.GET("/weixin_sucai/my", indexIndex.WeixinSucaiMy);
		e.GET("/weixin_sucai/add", indexIndex.WeixinSucaiAdd);
		e.GET("/weixin_sucai/save", indexIndex.WeixinSucaiSave);
		e.GET("/weixin_sucai/status", indexIndex.WeixinSucaiStatus);
		e.GET("/weixin_sucai/delete", indexIndex.WeixinSucaiDelete);
		/*WeixinUser*/
		e.GET("/weixin_user/index", indexIndex.WeixinUserIndex);
		e.GET("/weixin_user/list", indexIndex.WeixinUserList);
		e.GET("/weixin_user/show", indexIndex.WeixinUserShow);

	}

	