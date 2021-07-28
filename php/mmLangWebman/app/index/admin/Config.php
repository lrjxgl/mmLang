<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Config
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=4;
        $fm=DBS::MM("index","Config");
        $where=" 1 ";
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
        

        $id=intval($request->get("id"));
        $row=[];
        if($id){
            $fm=DBS::MM("index","Config");
            $row=$fm->find($id);
            
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
       

        $id=intval($request->get("id"));
        $data=[];
        $fm=DBS::MM("index","Config");
        $indata=[];
        //处理发布内容
        
$indata["siteid"]=intval($request->post("siteid","0"));
$indata["phone_on"]=intval($request->post("phone_on","0"));
$indata["phone_reg"]=intval($request->post("phone_reg","0"));
$indata["phone_user"]=$request->post("phone_user","");
$indata["phone_pwd"]=$request->post("phone_pwd","");
$indata["phone_num"]=$request->post("phone_num","");
$indata["smtphost"]=$request->post("smtphost","");
$indata["smtpport"]=intval($request->post("smtpport","0"));
$indata["smtpemail"]=$request->post("smtpemail","");
$indata["smtpuser"]=$request->post("smtpuser","");
$indata["smtppwd"]=$request->post("smtppwd","");
$indata["water_on"]=intval($request->post("water_on","0"));
$indata["water_type"]=intval($request->post("water_type","0"));
$indata["water_pos"]=intval($request->post("water_pos","0"));
$indata["water_str"]=$request->post("water_str","");
$indata["water_img"]=$request->post("water_img","");
$indata["water_size"]=intval($request->post("water_size","0"));
$indata["rewrite_on"]=intval($request->post("rewrite_on","0"));
$indata["spread_on"]=intval($request->post("spread_on","0"));
$indata["spread_discount"]=intval($request->post("spread_discount","0"));
$indata["spread_money"]=intval($request->post("spread_money","0"));
$indata["spread_type"]=intval($request->post("spread_type","0"));
$indata["grade_on"]=intval($request->post("grade_on","0"));
$indata["hotsearch"]=$request->post("hotsearch","");
$indata["kf_qq"]=$request->post("kf_qq","");
$indata["kf_ww"]=$request->post("kf_ww","");
$indata["sms_qianming"]=$request->post("sms_qianming","");
$indata["sms_type"]=$request->post("sms_type","");
$indata["s_bank_name"]=$request->post("s_bank_name","");
$indata["s_bank_user"]=$request->post("s_bank_user","");
$indata["s_bank_num"]=$request->post("s_bank_num","");
        if($id){
            $row=$fm->find($id);
            
        }
        if($id){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("id",$id)->update($indata);
        }else{       
            
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
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
        $fm=DBS::MM("index","Config");
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