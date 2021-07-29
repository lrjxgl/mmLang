<?php
namespace app\index\model;
use support\Model;
use ext\Help;
class UserModel extends Model{
	const UPDATED_AT= null;
	protected $table="user";
	protected $primaryKey = "userid";
	protected $created_at="createtime"; 
	
	public function get($userid,$field=""){
		if($field==""){
			$field="userid,nickname,user_head,follow_num,followed_num"; 
		}
		$row=$this->where("userid",$userid)->selectRaw($field)->first();
		if(empty($row)){
			return [];
		}
		$row->user_head=Help::images_site($row->user_head);
		return $row;
	}
	public  function Dselect($list){
		if(empty($list)) return $list;
		foreach($list as &$v){
			if(isset($v->user_head)){
				$v->user_head=Help::images_site($v->user_head);
			}
		}
		return $list;   
	}  

	public function getListByIds($ids,$fields="userid,nickname,user_head"){
		if(empty($fields)){
			$fields="userid,nickname,user_head";
		}
		if(empty($ids)) return [];
		$list=$this->whereIn("userid",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->userid]=$v;
			}
		}
		return $reList;
	}

}