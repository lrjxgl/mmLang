<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class AdModel extends Model{
	 
	const UPDATED_AT= null;
	protected $table="ad";
	protected $primaryKey = "id";
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

	function listByNo($tagno,$limit=4){
		$tags=DBS::MM("index","adTags")->where("tagno",$tagno)->first();
		if(empty($tags)){
			return []; 
		}
		$list=$this->where("tag_id_2nd",$tags->tag_id)->get();
		$list=$this->Dselect($list);
		return $list;
	}

	
}