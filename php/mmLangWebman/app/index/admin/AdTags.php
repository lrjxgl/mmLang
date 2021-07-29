<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class AdTags
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","AdTags");
        $where="status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("tag_id","desc")
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
        

        $tag_id=intval($request->get("tag_id"));
        $data=[];
        if($tag_id){
            $fm=DBS::MM("index","AdTags");
            $data=$fm->find($tag_id);
            
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
       

        $tag_id=intval($request->post("tag_id"));
        $data=[];
        $fm=DBS::MM("index","AdTags");
        $indata=[];
        //处理发布内容
        
$indata["title"]=$request->post("title","");
$indata["orderindex"]=intval($request->post("orderindex","0"));
$indata["pid"]=intval($request->post("pid","0"));
$indata["status"]=intval($request->post("status","0"));
$indata["m"]=$request->post("m","");
$indata["a"]=$request->post("a","");
$indata["width"]=intval($request->post("width","0"));
$indata["height"]=intval($request->post("height","0"));
$indata["tagno"]=$request->post("tagno","");
        if($tag_id){
            $row=$fm->find($tag_id);
            
        }
        if($tag_id){
            
            $fm->where("tag_id",$tag_id)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            
			
            $tag_id=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$tag_id
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $tag_id=$request->get("tag_id");
        $fm=DBS::MM("index","AdTags");
        $row=$fm->where("tag_id",$tag_id)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($tag_id);
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
		

        $tag_id=$request->get("tag_id");
        $fm=DBS::MM("index","AdTags");
        $row=$fm->find($tag_id); 
        
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