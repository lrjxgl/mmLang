<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class UserGroup
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=4;
        $fm=DBS::MM("index","UserGroup");
        $where=" 1 ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
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
        

        $groupid=intval($request->get("groupid"));
        $row=[];
        if($groupid){
            $fm=DBS::MM("index","UserGroup");
            $row=$fm->find($groupid);
            
        }
        $redata=[
            "error" => 0, 
            "message" => "success",
            "data"=>$row 
        ];
		return json($redata);       
    } 
    
	
    /*@@save@@*/
	public function save(Request $request){
       

        $groupid=intval($request->get("groupid"));
        $data=[];
        $fm=DBS::MM("index","UserGroup");
        $indata=[];
        //处理发布内容
        
$indata["groupname"]=$request->post("groupname","");
$indata["access"]=$request->post("access","");
        if($groupid){
            $row=$fm->find($groupid);
            
        }
        if($groupid){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("groupid",$groupid)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
            $groupid=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$groupid
        ];
		return json($redata); 
    }

    /*@@delete@@*/
    public function delete(Request $request){
		

        $groupid=$request->get("groupid");
        $fm=DBS::MM("index","UserGroup");
        $row=$fm->find($groupid); 
        
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