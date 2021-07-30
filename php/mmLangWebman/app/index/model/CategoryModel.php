<?php
namespace app\index\model;
use support\Model;
use ext\Help;
use ext\DBS; 
class CategoryModel extends Model{
	const UPDATED_AT= null;
	protected $table="category";
	protected $primaryKey = "catid";
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
	public function children($pid=0,$tablename="article",$status=0){
		$where=" tablename='".$tablename."' ";
		if($status==0){
			$where.=" AND status in(0,1,2) ";
		}else{
			$where.=" AND status=".$status;
		}

		$one=$this->whereRaw($where." AND pid=".$pid)->get();
		$two=[];
		$three=[];
		if(!empty($one)){
			foreach($one as $o){
				$oids[]=$o->catid;
			}
			$child=[];
			if(!empty($oids)){
				$two=$this->whereRaw($where." AND pid in(".Help::_implode($oids).") ")->get();
			if(!empty($two)){
				foreach($two as $t){
					$tids[]=$t->catid;
				}
				if(!empty($tids)){
					$three=$this->whereRaw($where." AND pid in(".Help::_implode($tids).") ")->get();
				} 
				
			}
			}
			
		}
		if(!empty($two)){
			foreach($two as $ti=>$t){
				$child=[];
				if(!empty($three)){
					foreach($three as $ei=>$e){
						if($t->catid==$e->pid){
							$child[]=$e;
							unset($three[$ei]);
						}	
					}
				}
				$two[$ti]["child"]=$child;
				
			}
		}

		if(!empty($one)){
			foreach($one as $ti=>$t){
				$child=[];
				if(!empty($two)){
					foreach($two as $ei=>$e){
						if($t->catid==$e->pid){
							$child[]=$e;
							unset($two[$ei]);
						}	
					}
				}
				$one[$ti]["child"]=$child;
				
			}
		}
		return $one;
		
	}

	public function id_family($id=0){
		$id=intval($id);
		$ids[]=$id;
		$ids1=$this->where("pid",$pid)->value("catid");
		if($ids1){
			$ids=array_merge($ids,$ids1);
			$ids2=$this->whereRaw(" pid in("._implode($ids1).") " )->value("catid");
			if($ids2){
				$ids=array_merge($ids,$ids2);
				$ids3=$this->whereRaw(" pid in("._implode($ids2).") " )->value("catid");
				if($ids3){
					$ids=array_merge($ids,$ids3);
				}
			}
		}
		return $ids;
		
	}
	
}