<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class CategoryModel extends Model{
	 
	protected $table="category";
	protected $primaryKey = "catid";
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
		$list=$this->whereIn("catid",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->catid]=$v;
			}
		}
		return $reList;
	}
	
}