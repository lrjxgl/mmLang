<?php
namespace app\forum\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class ForumModel extends Model{
	const UPDATED_AT= null;
	protected $table="mod_forum";
	protected $primaryKey = "id";
	protected $created_at="createtime"; 
	 
	public  function Dselect($list){
		if(empty($list)) return [];
		$uids=[];
		foreach($list as $v){
			$uids[]=$v->userid; 
		}
		
		$us=DBS::MM("index","user")->getListByIds($uids);
		foreach($list as &$v){
			if(isset($v->imgurl)){
				$v->imgurl=Help::images_site($v->imgurl);
			}
			if(isset($us[$v->userid])){
				$v["user"]=$us[$v->userid];
			}else{
				$v["user"]=[];
			}
			$v->timeago=Help::timeago($v->createtime); 
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