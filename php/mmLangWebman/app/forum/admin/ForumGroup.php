<?php
namespace app\forum\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class ForumGroup
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=4;
        $fm=DBS::MM("forum","ForumGroup");
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
        

        $gid=intval($request->get("gid"));
        $row=[];
        if($gid){
            $fm=DBS::MM("forum","ForumGroup");
            $row=$fm->find($gid);
            
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
       

        $gid=intval($request->get("gid"));
        $data=[];
        $fm=DBS::MM("forum","ForumGroup");
        $indata=[];
        //处理发布内容
        
$indata["title"]=$request->post("title","");
$indata["imgurl"]=$request->post("imgurl","");
$indata["description"]=$request->post("description","");
$indata["orderindex"]=intval($request->post("orderindex","0"));
$indata["status"]=intval($request->post("status","0"));
$indata["topic_num"]=intval($request->post("topic_num","0"));
        if($gid){
            $row=$fm->find($gid);
            
        }
        if($gid){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("gid",$gid)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
            $gid=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$gid
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $gid=$request->get("gid");
        $fm=DBS::MM("forum","ForumGroup");
        $row=$fm->where("gid",$gid)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($gid);
        $up->status=$status;
        $up->save();
        $redata=[
            "error" => 0, 
            "message" => "success",
            "status"=>$status
        ];
		return json($redata); 
    }

    /*@@recommend@@*/
    public function recommend(Request $request){
		

        $gid=$request->get("gid");
       $fm=DBS::MM("forum","ForumGroup");
        
        $row=$fm->where("gid",$gid)->first();
		
        if($row->isrecommend==1){
            $isrecommend=0;
        }else{
            $isrecommend=1;
        }
         
        $row->isrecommend=$isrecommend;
        $row->save();
        $redata=[
            "error" => 0, 
            "message" => "success",
            "isrecommend"=>$isrecommend
        ];
		return json($redata); 
    }

    /*@@delete@@*/
    public function delete(Request $request){
		

        $gid=$request->get("gid");
        $fm=DBS::MM("forum","ForumGroup");
        $row=$fm->find($gid); 
        
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