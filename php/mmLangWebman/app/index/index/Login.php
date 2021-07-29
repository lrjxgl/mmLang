<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\Help;
use ext\UserAccess;
use ext\Cache;
use ext\Sms;    
class Login{
     /*@@index@@*/
    public function index(Request $request){

    }
    /*@@save@@*/
    public function save(Request $request){
        $telephone=$request->post("telephone");
        $password=$request->post("password");
        $user=DBS::MM("index","user")->where("telephone",$telephone)->first();
        if(empty($user)){        
            return Help::success(1,"用户不存在".$telephone);		 
        }
        $up=DBS::MM("index","UserPassword")->where("userid",$user["userid"])->first();
        if($up["password"]!=Help::umd5($password.$up["salt"])){
            return Help::success(1,"登录密码错误");
        }
        $token=UserAccess::setToken($user["userid"],$up["password"]);
        $token["error"]=0;
        $token["message"]="登录成功";
        return json($token);
        

    }
    /*@@logout@@*/
    public function logout(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $iToken=$request->get("iToken","");
        $iRefreshToken=$request->get("iRefreshToken","");
        UserAccess::del($iToken);
        UserAccess::del($iRefreshToken);

    }

    /*@@findpwd@@*/
    public function findpwd(Request $request){
        return Help::success(0,"success");
    }
     /*@@findpwdsave@@*/
    public function findpwdSave(Request $request){
        $telephone=$request->post("telephone","");
        $yzm=$request->post("yzm","");
        $key="smsLogin".$telephone;
        $yzm2=Cache::get($key);
        if($yzm!=$yzm2){
            return Help::success(1,"短信验证码出错");
        }
        $password=$request->post("password");
        $user=DBS::MM("index","user")->where("telephone",$telephone)->first();
        if(empty($user)){        
            return Help::success(1,"用户不存在".$telephone);		 
        }
        $up=DBS::MM("index","UserPassword")->where("userid",$user["userid"])->first();
        $salt=rand(1000,9999);
        $pwd=Help::umd5($password.$salt);
        if($up){
            DBS::MM("index","UserPassword")->where("userid",$user["userid"])->update([
                "password"=>$pwd,
                "salt"=>$salt
            ]);
        }else{
            DBS::MM("index","UserPassword")->insert([
                "userid"=>$userid,
                "password"=>$pwd,
                "salt"=>$salt
            ]); 
        }
        return Help::success(0,"success");
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
        
        $key="smsLogin".$telephone;
        if(Sms::$TEST==1){
            Cache::set($key,$yzm,180);
            return  Help::success(0,"验证码：".$yzm);
        }
        $keyExpire="smsLoginExpire".$telephone;
        if(Cache::get($keyExpire)){
           // return  Help::success(1,"请稍后再发");
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