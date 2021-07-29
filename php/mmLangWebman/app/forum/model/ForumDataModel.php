<?php
namespace app\forum\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class ForumDataModel extends Model{
	 
	protected $table="mod_forum_data";
	protected $primaryKey = "did";
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
		$list=$this->whereIn("did",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->did]=$v;
			}
		}
		return $reList;
	}
	
}