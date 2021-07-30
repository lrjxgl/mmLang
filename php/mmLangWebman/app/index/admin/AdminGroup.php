<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class AdminGroup
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","AdminGroup");
        $where=" 1 ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("id","desc")
                ->get();
        $list=$fm->Dselect($list);
        $rscount=$fm->whereRaw($where)->count();
        $per_page=$start+$limit;
        $per_page=$per_page>$rscount?0:$per_page;
        $redata=[
            "error" => 0, 
            "message" => "success",
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount

        ];
		return json($redata); 
         
		   
    }

    /*@@add@@*/
	public function add(Request $request){
        

        $id=intval($request->get("id"));
        $data=[];
        $permissionList=DBS::MM("index","permission")->get();
        $uAccess=[];
        if($id){
            $fm=DBS::MM("index","AdminGroup");
            $data=$fm->find($id);
            $uAccess=explode(",",$data->pids);
            
        }
        if(!empty($permissionList)){
            foreach($permissionList as $k=>$v){
                if(!empty($uAccess)){
                    if(in_array($v->id,$uAccess)){
                        $v->checked=true;
                    }else{
                        $v->checked=false;
                    }
                }else{
                    $v->checked=false;
                }
                $permissionList[$k]=$v;
            }
        }
        
        $redata=[
            "error" => 0, 
            "message" => "success",
            "data"=>$data,
            "permissionList"=>$permissionList 
        ];
		return json($redata);       
    } 
    
	
    /*@@save@@*/
	public function save(Request $request){
       

        $id=intval($request->post("id"));
        $data=[];
        $fm=DBS::MM("index","AdminGroup");
        $indata=[];
        //处理发布内容   
        $indata["title"]=$request->post("title","");
        $indata["orderindex"]=intval($request->post("orderindex","0"));
        $pids=$request->post("pids","");
        $arr=explode(",",$pids);
        foreach($arr as $k=>$v){
            
            $arr[$k]=intval($v);
            if($arr[$k]==0){
                unset($arr[$k]);
            }
        }
        //处理权限
        $pers=DBS::MM("index","permission")->getListByIds($arr);
        $access=[];
        if(!empty($pers)){
            foreach($pers as $per){
                if(isset($access[$per->m])){
                    $access[$per->m]+=explode(",",$per->access);
                }else{
                    $access[$per->m]=explode(",",$per->access);
                }
                
            }
        }
        $indata["content"]=json_encode($access);
        $pids=implode(",",$arr);
        $indata["pids"]=$pids;
        if($id){
            $row=$fm->find($id);
            
        }
        if($id){
            
            $fm->where("id",$id)->update($indata);
        }else{       
            
            
            
			
            $id=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$id
        ];
		return json($redata); 
    }

    /*@@delete@@*/
    public function delete(Request $request){
		

        $id=$request->get("id");
        $fm=DBS::MM("index","AdminGroup");
        $row=$fm->find($id); 
        
        $row->status=11;
        $row->save();
        $redata=[
            "error" => 0, 
            "message" => "success"
        ];
		return json($redata); 
    }
      
}

?>