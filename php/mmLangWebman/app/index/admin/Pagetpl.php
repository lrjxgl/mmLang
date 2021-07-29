<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Pagetpl
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","Pagetpl");
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
        if($id){
            $fm=DBS::MM("index","Pagetpl");
            $data=$fm->find($id);
            
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
       

        $id=intval($request->post("id"));
        $data=[];
        $fm=DBS::MM("index","Pagetpl");
        $indata=[];
        //处理发布内容
        
$indata["m"]=$request->post("m","");
$indata["a"]=$request->post("a","");
$indata["tpl"]=$request->post("tpl","");
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
        $fm=DBS::MM("index","Pagetpl");
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