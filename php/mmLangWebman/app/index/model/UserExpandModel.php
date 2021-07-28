<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class UserExpandModel extends Model{
	const UPDATED_AT= null;
	protected $table="user_expand";
	protected $primaryKey = "uid";
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
		$list=$this->whereIn("uid",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->uid]=$v;
			}
		}
		return $reList;
	}
	
}