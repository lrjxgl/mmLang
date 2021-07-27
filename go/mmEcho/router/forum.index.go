
	package router

	import (
		"app/forum/forumIndex"
		"github.com/labstack/echo/v4"
	)

	func ForumIndexRouter(e *echo.Echo) {
		/*Forum*/
		e.GET("/forum/index", forumIndex.ForumIndex);
		e.GET("/forum/list", forumIndex.ForumList);
		e.GET("/forum/show", forumIndex.ForumShow);
		e.GET("/forum/my", forumIndex.ForumMy);
		e.GET("/forum/add", forumIndex.ForumAdd);
		e.GET("/forum/save", forumIndex.ForumSave);
		e.GET("/forum/status", forumIndex.ForumStatus);
		e.GET("/forum/delete", forumIndex.ForumDelete);
		/*ForumCategory*/
		e.GET("/forum_category/index", forumIndex.ForumCategoryIndex);
		e.GET("/forum_category/list", forumIndex.ForumCategoryList);
		e.GET("/forum_category/show", forumIndex.ForumCategoryShow);
		/*ForumComment*/
		e.GET("/forum_comment/index", forumIndex.ForumCommentIndex);
		e.GET("/forum_comment/list", forumIndex.ForumCommentList);
		e.GET("/forum_comment/show", forumIndex.ForumCommentShow);
		e.GET("/forum_comment/my", forumIndex.ForumCommentMy);
		e.GET("/forum_comment/add", forumIndex.ForumCommentAdd);
		e.GET("/forum_comment/save", forumIndex.ForumCommentSave);
		e.GET("/forum_comment/status", forumIndex.ForumCommentStatus);
		e.GET("/forum_comment/delete", forumIndex.ForumCommentDelete);
		/*ForumData*/
		e.GET("/forum_data/index", forumIndex.ForumDataIndex);
		e.GET("/forum_data/list", forumIndex.ForumDataList);
		e.GET("/forum_data/show", forumIndex.ForumDataShow);
		/*ForumFeeds*/
		e.GET("/forum_feeds/index", forumIndex.ForumFeedsIndex);
		e.GET("/forum_feeds/list", forumIndex.ForumFeedsList);
		e.GET("/forum_feeds/show", forumIndex.ForumFeedsShow);
		e.GET("/forum_feeds/my", forumIndex.ForumFeedsMy);
		e.GET("/forum_feeds/add", forumIndex.ForumFeedsAdd);
		e.GET("/forum_feeds/save", forumIndex.ForumFeedsSave);
		e.GET("/forum_feeds/delete", forumIndex.ForumFeedsDelete);
		/*ForumGroup*/
		e.GET("/forum_group/index", forumIndex.ForumGroupIndex);
		e.GET("/forum_group/list", forumIndex.ForumGroupList);
		e.GET("/forum_group/show", forumIndex.ForumGroupShow);
		/*ForumGroupAdmin*/
		e.GET("/forum_group_admin/index", forumIndex.ForumGroupAdminIndex);
		e.GET("/forum_group_admin/list", forumIndex.ForumGroupAdminList);
		e.GET("/forum_group_admin/show", forumIndex.ForumGroupAdminShow);
		e.GET("/forum_group_admin/my", forumIndex.ForumGroupAdminMy);
		e.GET("/forum_group_admin/add", forumIndex.ForumGroupAdminAdd);
		e.GET("/forum_group_admin/save", forumIndex.ForumGroupAdminSave);
		e.GET("/forum_group_admin/status", forumIndex.ForumGroupAdminStatus);
		e.GET("/forum_group_admin/delete", forumIndex.ForumGroupAdminDelete);
		/*ForumImg*/
		e.GET("/forum_img/index", forumIndex.ForumImgIndex);
		e.GET("/forum_img/list", forumIndex.ForumImgList);
		e.GET("/forum_img/show", forumIndex.ForumImgShow);
		/*ForumTags*/
		e.GET("/forum_tags/index", forumIndex.ForumTagsIndex);
		e.GET("/forum_tags/list", forumIndex.ForumTagsList);
		e.GET("/forum_tags/show", forumIndex.ForumTagsShow);
		/*ForumTagsIndex*/
		e.GET("/forum_tags/index", forumIndex.ForumTagsIndexIndex);
		e.GET("/forum_tags_index/list", forumIndex.ForumTagsIndexList);
		e.GET("/forum_tags_index/show", forumIndex.ForumTagsIndexShow);

	}

	