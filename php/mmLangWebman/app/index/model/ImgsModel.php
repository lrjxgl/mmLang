<?php
namespace app\index\model;
use support\Model;
use ext\Help;
class ImgsModel extends Model{
	const UPDATED_AT= null;
	protected $table="imgs";
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
}