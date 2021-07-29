<?php
namespace app\forum\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class ForumCategory
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("forum","ForumCategory");
        $where="status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("catid","desc")
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
        

        $catid=intval($request->get("catid"));
        $data=[];
        if($catid){
            $fm=DBS::MM("forum","ForumCategory");
            $data=$fm->find($catid);
            
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
       

        $catid=intval($request->post("catid"));
        $data=[];
        $fm=DBS::MM("forum","ForumCategory");
        $indata=[];
        //处理发布内容
        
$indata["gid"]=intval($request->post("gid","0"));
$indata["title"]=$request->post("title","");
$indata["description"]=$request->post("description","");
$indata["orderindex"]=intval($request->post("orderindex","0"));
$indata["status"]=intval($request->post("status","0"));
$indata["imgurl"]=$request->post("imgurl","");
        if($catid){
            $row=$fm->find($catid);
            
        }
        if($catid){
            
            $fm->where("catid",$catid)->update($indata);
        }else{       
            
            
            
			
            $catid=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$catid
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $catid=$request->get("catid");
        $fm=DBS::MM("forum","ForumCategory");
        $row=$fm->where("catid",$catid)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($catid);
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
		

        $catid=$request->get("catid");
        $fm=DBS::MM("forum","ForumCategory");
        $row=$fm->find($catid); 
        
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