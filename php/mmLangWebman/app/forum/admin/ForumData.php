<?php
namespace app\forum\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class ForumData
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("forum","ForumData");
        $where=" 1 ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
				->orderBy("did","desc")
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
        

        $did=intval($request->get("did"));
        $data=[];
        if($did){
            $fm=DBS::MM("forum","ForumData");
            $data=$fm->find($did);
            
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
       

        $did=intval($request->post("did"));
        $data=[];
        $fm=DBS::MM("forum","ForumData");
        $indata=[];
        //处理发布内容
        
$indata["id"]=intval($request->post("id","0"));
$indata["content"]=$request->post("content","");
        if($did){
            $row=$fm->find($did);
            
        }
        if($did){
            
            $fm->where("did",$did)->update($indata);
        }else{       
            
            
            
			
            $did=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$did
        ];
		return json($redata); 
    }

    /*@@delete@@*/
    public function delete(Request $request){
		

        $did=$request->get("did");
        $fm=DBS::MM("forum","ForumData");
        $row=$fm->find($did); 
        
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