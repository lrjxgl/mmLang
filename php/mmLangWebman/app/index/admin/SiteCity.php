<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class SiteCity
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","SiteCity");
        $where="status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("sc_id","desc")
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
        

        $sc_id=intval($request->get("sc_id"));
        $data=[];
        if($sc_id){
            $fm=DBS::MM("index","SiteCity");
            $data=$fm->find($sc_id);
            
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
       

        $sc_id=intval($request->post("sc_id"));
        $data=[];
        $fm=DBS::MM("index","SiteCity");
        $indata=[];
        //处理发布内容
        
$indata["title"]=$request->post("title","");
$indata["cityid"]=intval($request->post("cityid","0"));
$indata["lat"]=floatval($request->post("lat","0"));
$indata["lng"]=floatval($request->post("lng","0"));
$indata["orderindex"]=intval($request->post("orderindex","0"));
$indata["status"]=intval($request->post("status","0"));
$indata["pid"]=intval($request->post("pid","0"));
$indata["siteid"]=intval($request->post("siteid","0"));
        if($sc_id){
            $row=$fm->find($sc_id);
            
        }
        if($sc_id){
            
            $fm->where("sc_id",$sc_id)->update($indata);
        }else{       
            
            
            
			
            $sc_id=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$sc_id
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $sc_id=$request->get("sc_id");
        $fm=DBS::MM("index","SiteCity");
        $row=$fm->where("sc_id",$sc_id)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($sc_id);
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
		

        $sc_id=$request->get("sc_id");
        $fm=DBS::MM("index","SiteCity");
        $row=$fm->find($sc_id); 
        
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