<?php
namespace app\forum\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class ForumTagsModel extends Model{
	const UPDATED_AT= null;
	protected $table="mod_forum_tags";
	protected $primaryKey = "tagid";
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
	public function getListByIds($ids,$fields="*"){
		if(empty($ids)) return [];
		$list=$this->whereIn("tagid",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->tagid]=$v;
			}
		}
		return $reList;
	}

	public function GetForumByKey($gkey,$limit=12){
		$where=[
			["gkey",$gkey],
			["status",1]
		];
		$tag=$this->where($where)->first();
		if(empty($tag)){
			return [];
		}
		$list=DBS::MM("forum","ForumTagsIndex")->where("tagid",$tag->tagid)->get();
		if(empty($list)){
			return [];
		}
		$ids=[];
		foreach($list as $v){
			$ids[]=$v->objectid;
		}
		$res=DBS::MM("forum","forum")->getListByIds($ids);
		$reList=[];
		foreach($list as $v){
			if(isset($res[$v->objectid])){
				$reList[]=$res[$v->objectid];
			} 
			
		}
		return $reList;

	}
	
}