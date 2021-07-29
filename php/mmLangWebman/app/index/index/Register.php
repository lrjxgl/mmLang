<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
use ext\Cache;
use ext\Sms; 
class Register{
    /*@@index@@*/
    public function Index(Request $request){
       return  Help::success(0,"success");
    }
    /*@@save@@*/
    public function Save(Request $request){
        $yzm=$request->post("yzm","");
        $telephone=$request->post("telephone","");
        $key="smsReg".$telephone;
        $yzm2=Cache::get($key);
        if($yzm!=$yzm2){
            return Help::success(1,"短信验证码出错");
        }
        $fm=DBS::MM("index","user");
        $row=$fm->where("telephone",$telephone)->first();
        if(!empty($row)){
            return Help::success(1,"手机已存在");
        }
        $username=$nickname=$request->post("nickname","h");
        $where=" username='".$username."' or nickname='".$nickname."' ";
        $row=$fm->whereRaw($where)->first();
        if(!empty($row)){
            return Help::success(1,"昵称已存在");
        }
        $password=$request->post("password","h");
        $password2=$request->post("password2","h"); 
        if(empty($password) || $password!=$password2){
            return Help::success(1,"请确认密码");
        }
        $userid=$fm->insertGetId([
            "username"=>$username,
            "nickname"=>$nickname,
            "telephone"=>$telephone
        ]);
        $salt=rand(1000,9999);
        $password=Help::umd5($password);
        DBS::MM("index","userPassword")->insert([
            "userid"=>$userid,
            "password"=>$password,
            "salt"=>$salt
        ]);
        $token=UserAccess::setToken($userid,$password);
        $token["error"]=0;
        $token["message"]="注册成功";
        return json($token);
    }

     /*@@sendSms@@*/
     public function sendSms(Request $request){
        
        $telephone=$request->get("telephone","");
         
        if(empty($telephone)){
            return Help::success(1,"请确认手机号码");
        }
        $yzm=rand(1000,9999);
        $opt=array(
            "telephone"=>$telephone,
            "content"=>"您的验证码".$yzm,
            "yzm"=>$yzm
        );
        
        $key="smsReg".$telephone;
        if(Sms::$TEST==1){
            Cache::set($key,$yzm,180);
            return  Help::success(0,"验证码：".$yzm);
        }
        $keyExpire="smsRegExpire".$telephone;
        if(Cache::get($keyExpire)){
            return  Help::success(1,"请稍后再发");
        }
        $res=Sms::send($opt);
        if(!$res){
            return Help::success(1,"短信发送失败");
        }

        Cache::set($key,$yzm,180);
        Cache::set($keyExpire,1,60);

        $redata=[
            "error" => 0, 
            "message" => "success",     
        ];
		return json($redata);
    }
}