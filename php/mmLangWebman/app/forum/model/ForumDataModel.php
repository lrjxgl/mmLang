<?php
namespace app\forum\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class ForumDataModel extends Model{
	const UPDATED_AT= null;
	const CREATED_At = null;
	protected $table="mod_forum_data";
	protected $primaryKey = "id";
	 
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
		$list=$this->whereIn("id",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->id]=$v;
			}
		}
		return $reList;
	}
	
}