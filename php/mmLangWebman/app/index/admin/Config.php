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
	     
        $data=DBS::MM("index","config")->get();
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
        
        $fm=DBS::MM("index","Config");
        $data=$fm->get();
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
        $fm->whereRaw(" 1=1 ")->update($indata);
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功"
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