<?php
namespace app\index\model;
use support\Model;
use ext\Help;
class NavbarModel extends Model{
	const UPDATED_AT= null;
	protected $table="navbar";
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
	public function groupList(){
		return [
			1=>["gid"=>1,"title"=>"后台顶部"],
			2=>["gid"=>2,"title"=>"后台左边"]
		];

	}
	
}