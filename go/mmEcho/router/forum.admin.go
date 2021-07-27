
	package router

	import (
		"app/forum/forumAdmin"
		"github.com/labstack/echo/v4"
	)

	func ForumAdminRouter(e *echo.Echo) {
		/*Forum*/
		e.GET("/admin/forum/index", forumAdmin.ForumIndex);
		e.GET("/admin/forum/add", forumAdmin.ForumAdd);
		e.GET("/admin/forum/save", forumAdmin.ForumSave);
		e.GET("/admin/forum/status", forumAdmin.ForumStatus);
		e.GET("/admin/forum/recommend", forumAdmin.ForumRecommend);
		e.GET("/admin/forum/delete", forumAdmin.ForumDelete);
		/*ForumCategory*/
		e.GET("/admin/forum_category/index", forumAdmin.ForumCategoryIndex);
		e.GET("/admin/forum_category/add", forumAdmin.ForumCategoryAdd);
		e.GET("/admin/forum_category/save", forumAdmin.ForumCategorySave);
		e.GET("/admin/forum_category/status", forumAdmin.ForumCategoryStatus);
		e.GET("/admin/forum_category/delete", forumAdmin.ForumCategoryDelete);
		/*ForumComment*/
		e.GET("/admin/forum_comment/index", forumAdmin.ForumCommentIndex);
		e.GET("/admin/forum_comment/add", forumAdmin.ForumCommentAdd);
		e.GET("/admin/forum_comment/save", forumAdmin.ForumCommentSave);
		e.GET("/admin/forum_comment/status", forumAdmin.ForumCommentStatus);
		e.GET("/admin/forum_comment/delete", forumAdmin.ForumCommentDelete);
		/*ForumData*/
		e.GET("/admin/forum_data/index", forumAdmin.ForumDataIndex);
		e.GET("/admin/forum_data/add", forumAdmin.ForumDataAdd);
		e.GET("/admin/forum_data/save", forumAdmin.ForumDataSave);
		e.GET("/admin/forum_data/delete", forumAdmin.ForumDataDelete);
		/*ForumFeeds*/
		e.GET("/admin/forum_feeds/index", forumAdmin.ForumFeedsIndex);
		e.GET("/admin/forum_feeds/add", forumAdmin.ForumFeedsAdd);
		e.GET("/admin/forum_feeds/save", forumAdmin.ForumFeedsSave);
		e.GET("/admin/forum_feeds/delete", forumAdmin.ForumFeedsDelete);
		/*ForumGroup*/
		e.GET("/admin/forum_group/index", forumAdmin.ForumGroupIndex);
		e.GET("/admin/forum_group/add", forumAdmin.ForumGroupAdd);
		e.GET("/admin/forum_group/save", forumAdmin.ForumGroupSave);
		e.GET("/admin/forum_group/status", forumAdmin.ForumGroupStatus);
		e.GET("/admin/forum_group/recommend", forumAdmin.ForumGroupRecommend);
		e.GET("/admin/forum_group/delete", forumAdmin.ForumGroupDelete);
		/*ForumGroupAdmin*/
		e.GET("/admin/forum_group_admin/index", forumAdmin.ForumGroupAdminIndex);
		e.GET("/admin/forum_group_admin/add", forumAdmin.ForumGroupAdminAdd);
		e.GET("/admin/forum_group_admin/save", forumAdmin.ForumGroupAdminSave);
		e.GET("/admin/forum_group_admin/status", forumAdmin.ForumGroupAdminStatus);
		e.GET("/admin/forum_group_admin/delete", forumAdmin.ForumGroupAdminDelete);
		/*ForumImg*/
		e.GET("/admin/forum_img/index", forumAdmin.ForumImgIndex);
		e.GET("/admin/forum_img/add", forumAdmin.ForumImgAdd);
		e.GET("/admin/forum_img/save", forumAdmin.ForumImgSave);
		e.GET("/admin/forum_img/delete", forumAdmin.ForumImgDelete);
		/*ForumTags*/
		e.GET("/admin/forum_tags/index", forumAdmin.ForumTagsIndex);
		e.GET("/admin/forum_tags/add", forumAdmin.ForumTagsAdd);
		e.GET("/admin/forum_tags/save", forumAdmin.ForumTagsSave);
		e.GET("/admin/forum_tags/status", forumAdmin.ForumTagsStatus);
		e.GET("/admin/forum_tags/delete", forumAdmin.ForumTagsDelete);
		/*ForumTagsIndex*/
		e.GET("/admin/forum_tags/index", forumAdmin.ForumTagsIndexIndex);
		e.GET("/admin/forum_tags_index/add", forumAdmin.ForumTagsIndexAdd);
		e.GET("/admin/forum_tags_index/save", forumAdmin.ForumTagsIndexSave);
		e.GET("/admin/forum_tags_index/delete", forumAdmin.ForumTagsIndexDelete);

	}

	