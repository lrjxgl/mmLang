<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class UserRank
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=4;
        $fm=DBS::MM("index","UserRank");
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
        

        $id=intval($request->get("id"));
        $row=[];
        if($id){
            $fm=DBS::MM("index","UserRank");
            $row=$fm->find($id);
            
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
       

        $id=intval($request->get("id"));
        $data=[];
        $fm=DBS::MM("index","UserRank");
        $indata=[];
        //处理发布内容
        
$indata["rank_name"]=$request->post("rank_name","");
$indata["min_grade"]=intval($request->post("min_grade","0"));
$indata["max_grade"]=intval($request->post("max_grade","0"));
$indata["discount"]=intval($request->post("discount","0"));
$indata["logo"]=$request->post("logo","");
        if($id){
            $row=$fm->find($id);
            
        }
        if($id){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("id",$id)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
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
        $fm=DBS::MM("index","UserRank");
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