<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class AdTagsModel extends Model{
	const UPDATED_AT= null;
	protected $table="ad_tags";
	protected $primaryKey = "tag_id";
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
		$list=$this->whereIn("tag_id",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->tag_id]=$v;
			}
		}
		return $reList;
	}
	
}