<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class AttachModel extends Model{
	const UPDATED_AT= null;
	protected $table="attach";
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

	public function Add($ops){
		if(!isset($ops["url"]) && empty($ops["url"]) && file_exists($ops["url"])){
			return false;
		}
		$data=array();
		$data["file_url"]=$ops["url"];
		$data["file_name"]=basename($ops["url"]);
		if(!isset($ops["file_size"])){
			$data["file_size"]=filesize($data["file_url"])/1024;
		}else{
			$data["file_size"]=$ops["file_size"]/1024;
		}
		
		$data["file_type"]=strtolower(trim(substr(strrchr($ops["url"], '.'), 1)));
		$data["file_group"]=$this->group($data["file_type"]);
		$data["userid"]=$ops["userid"];
		$data["createtime"]=date("Y-m-d H:i:s");
		return $this->insert($data);
	}
	public function group($type){
		switch($type){
			case "png":
			case "bmp":
			case "jpg":
			case "jpeg":
			case "gif":
			case "webp":
				return "img";
				break;
			case "mp3":
			case "wma":
			case "aac":
			case "wav":
			case "midi":
				return "audio";
				break;
			case "mp4":
			case "wmv":
			case "flv":
			case "rm":
			case "mpeg":
			case "avi":
			case "3gp":
				return "video";
				break;
			case "txt":
			case "doc":
			case "docx":
			case "csv":
			case "xsl":
				return "text";
				break;
			case "zip":
			case "rar":
				return "zip";
				break;
			default:
				return "other";
				break;
		}
	}
	
}