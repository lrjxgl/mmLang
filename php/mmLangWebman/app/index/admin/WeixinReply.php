<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class WeixinReply
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","WeixinReply");
        $where="status in(0,1,2) ";
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
            $fm=DBS::MM("index","WeixinReply");
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
        $fm=DBS::MM("index","WeixinReply");
        $indata=[];
        //处理发布内容
        
$indata["status"]=intval($request->post("status","0"));
$indata["openid"]=$request->post("openid","");
$indata["msgtype"]=$request->post("msgtype","");
$indata["content"]=$request->post("content","");
$indata["msgid"]=$request->post("msgid","");
$indata["picurl"]=$request->post("picurl","");
$indata["mediaid"]=$request->post("mediaid","");
$indata["thumbmediaid"]=$request->post("thumbmediaid","");
$indata["format"]=$request->post("format","");
$indata["location_x"]=$request->post("location_x","");
$indata["location_y"]=$request->post("location_y","");
$indata["scale"]=intval($request->post("scale","0"));
$indata["label"]=$request->post("label","");
$indata["title"]=$request->post("title","");
$indata["description"]=$request->post("description","");
$indata["url"]=$request->post("url","");
$indata["wid"]=intval($request->post("wid","0"));
$indata["shopid"]=intval($request->post("shopid","0"));
$indata["fromusername"]=$request->post("fromusername","");
$indata["tousername"]=$request->post("tousername","");
        if($id){
            $row=$fm->find($id);
            
        }
        if($id){
            
            $fm->where("id",$id)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            
			
            $id=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
            "insert_id"=>$id
        ];
		return json($redata); 
    }

    /*@@status@@*/
    public function Status(Request $request){
		

        $id=$request->get("id");
        $fm=DBS::MM("index","WeixinReply");
        $row=$fm->where("id",$id)->first();
		
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($id);
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
		

        $id=$request->get("id");
        $fm=DBS::MM("index","WeixinReply");
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