<?php
namespace app\forum\model;
use support\Model;
use ext\Help;
class ForumCategoryModel extends Model{
	const UPDATED_AT= null;
	protected $table="mod_forum_category";
	protected $primaryKey = "catid";
	protected $created_at="createtime"; 
	 
	public  function Dselect($list){
		if(empty($list)) return $list;
		foreach($list as &$v){
			if(isset($v->imgurl)){
				$v->imgurl=Help::images_site($v->imgurl);
			}
		}
		return $list; 
	} 
}