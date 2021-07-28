<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Site
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=4;
        $fm=DBS::MM("index","Site");
        $where="status in(0,1,2) ";
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
        

        $siteid=intval($request->get("siteid"));
        $row=[];
        if($siteid){
            $fm=DBS::MM("index","Site");
            $row=$fm->find($siteid);
            
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
       

        $siteid=intval($request->get("siteid"));
        $data=[];
        $fm=DBS::MM("index","Site");
        $indata=[];
        //处理发布内容
        
$indata["sitename"]=$request->post("sitename","");
$indata["domain"]=$request->post("domain","");
$indata["title"]=$request->post("title","");
$indata["keywords"]=$request->post("keywords","");
$indata["description"]=$request->post("description","");
$indata["close_why"]=$request->post("close_why","");
$indata["logo"]=$request->post("logo","");
$indata["icp"]=$request->post("icp","");
$indata["status"]=intval($request->post("status","0"));
$indata["template"]=$request->post("template","");
$indata["statjs"]=$request->post("statjs","");
$indata["index_template"]=$request->post("index_template","");
$indata["wap_template"]=$request->post("wap_template","");
$indata["wapbg"]=$request->post("wapbg","");
        if($siteid){
            $row=$fm->find($siteid);
            
        }
        if($siteid){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("siteid",$siteid)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
            $siteid=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$siteid
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $siteid=$request->get("siteid");
        $fm=DBS::MM("index","Site");
        $row=$fm->where("siteid",$siteid)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($siteid);
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
		

        $siteid=$request->get("siteid");
        $fm=DBS::MM("index","Site");
        $row=$fm->find($siteid); 
        
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