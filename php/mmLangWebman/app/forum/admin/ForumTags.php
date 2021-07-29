<?php
namespace app\forum\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class ForumTags
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("forum","ForumTags");
        $where="status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("tagid","desc")
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
        

        $tagid=intval($request->get("tagid"));
        $row=[];
        if($tagid){
            $fm=DBS::MM("forum","ForumTags");
            $row=$fm->find($tagid);
            
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
       

        $tagid=intval($request->post("tagid"));
        $data=[];
        $fm=DBS::MM("forum","ForumTags");
        $indata=[];
        //处理发布内容
        
$indata["title"]=$request->post("title","");
$indata["status"]=intval($request->post("status","0"));
$indata["total_num"]=intval($request->post("total_num","0"));
$indata["gkey"]=$request->post("gkey","");
$indata["gnum"]=intval($request->post("gnum","0"));
        if($tagid){
            $row=$fm->find($tagid);
            
        }
        if($tagid){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("tagid",$tagid)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
            $tagid=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$tagid
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $tagid=$request->get("tagid");
        $fm=DBS::MM("forum","ForumTags");
        $row=$fm->where("tagid",$tagid)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($tagid);
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
		

        $tagid=$request->get("tagid");
        $fm=DBS::MM("forum","ForumTags");
        $row=$fm->find($tagid); 
        
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