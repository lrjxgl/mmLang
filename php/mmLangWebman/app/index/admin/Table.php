<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Table
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","Table");
        $where="status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("tableid","desc")
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
        

        $tableid=intval($request->get("tableid"));
        $data=[];
        if($tableid){
            $fm=DBS::MM("index","Table");
            $data=$fm->find($tableid);
            
        }
        $redata=[
            "error" => 0, 
            "message" => "success",
            "data"=>$data 
        ];
		return json($redata);       
    } 
    
	
    /*@@save@@*/
	public function save(Request $request){
       

        $tableid=intval($request->post("tableid"));
        $data=[];
        $fm=DBS::MM("index","Table");
        $indata=[];
        //处理发布内容
        
$indata["title"]=$request->post("title","");
$indata["tablename"]=$request->post("tablename","");
$indata["description"]=$request->post("description","");
$indata["status"]=intval($request->post("status","0"));
$indata["showTpl"]=$request->post("showTpl","");
$indata["listTpl"]=$request->post("listTpl","");
$indata["addTpl"]=$request->post("addTpl","");
        if($tableid){
            $row=$fm->find($tableid);
            
        }
        if($tableid){
            
            $fm->where("tableid",$tableid)->update($indata);
        }else{       
            
            
            
			
            $tableid=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$tableid
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $tableid=$request->get("tableid");
        $fm=DBS::MM("index","Table");
        $row=$fm->where("tableid",$tableid)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($tableid);
        $up->status=$status;
        $up->save();
        $redata=[
            "error" => 0, 
            "message" => "success",
            "status"=>$status
        ];
		return json($redata); 
    }

    /*@@delete@@*/
    public function delete(Request $request){
		

        $tableid=$request->get("tableid");
        $fm=DBS::MM("index","Table");
        $row=$fm->find($tableid); 
        
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