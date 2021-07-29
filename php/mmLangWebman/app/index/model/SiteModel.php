<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class SiteModel extends Model{
	const UPDATED_AT= null;
	protected $table="site";
	protected $primaryKey = "siteid";
	protected $created_at="createtime"; 
	
	public function get(){
		$row=$this->first();
		if(empty($row)){
			$this->insert(["sitename"=>"deituicms"]);
			$row=$this->first();
		}
		return $row;
	}
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
		$list=$this->whereIn("siteid",$ids)->selectRaw($fields)->get();
		$list=$this->Dselect($list);
		$reList=[];
		if($list){
			foreach($list as $v){
				$reList[$v->siteid]=$v;
			}
		}
		return $reList;
	}
	
}