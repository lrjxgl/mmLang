<?php
namespace app\forum\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class ForumGroupModel extends Model{
	 
	protected $table="mod_forum_group";
	protected $primaryKey = "gid";
	const CREATED_AT= null;
	const UPDATED_AT= null;  
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
		$list=$this->whereIn("gid",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->gid]=$v;
			}
		}
		return $reList;
	}
	
}