<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
use ext\Cache;
use ext\Sms;
class User
{ 
	/*@@index@@*/    
    public function index(Request $request)
    {
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $user=DBS::MM("index","User")->get($ssuserid,"userid,user_head,nickname,gold,money,grade");
        
        $redata=[
            "error" => 0, 
            "message" => "success",  
            "user"=>$user
        ];
		return json($redata); 
         
		   
    }
    /*@@set@@*/
    public function set(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $user=DBS::MM("index","User")->get($ssuserid,"userid,user_head,nickname,gold,money,grade");
        
        $redata=[
            "error" => 0, 
            "message" => "success",  
            "user"=>$user
        ];
		return json($redata);       
    }
    /*@@info@@*/
    public function info(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $user=DBS::MM("index","User")->get($ssuserid,"userid,nickname,description,telephone");
        
        $redata=[
            "error" => 0, 
            "message" => "success",  
            "user"=>$user
        ];
		return json($redata);       
    }  
    /*@@save@@*/
    public function save(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $fm=DBS::MM("index","User");
        $indata["nickname"]=$request->post("nickname",""); 
        $indata["description"]=$request->post("description","");
        $fm->where("userid",$ssuserid)->update($indata); 
        $redata=[
            "error" => 0, 
            "message" => "保存成功"
        ];
		return json($redata); 
    }
    /*@@head@@*/
    public function head(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $user=DBS::MM("index","User")->where("userid",$ssuserid)->first();
        $user->true_user_head=Help::images_site($user->user_head);
        $redata=[
            "error" => 0, 
            "message" => "success",  
            "user"=>$user
        ];
		return json($redata);       
    } 
    /*@@headsave@@*/
    public function headsave(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $fm=DBS::MM("index","User");
        $indata["user_head"]=Help::safeFileName($request->post("user_head","")); 
        
        $fm->where("userid",$ssuserid)->update($indata); 
        $redata=[
            "error" => 0, 
            "message" => "保存成功"
        ];
		return json($redata); 
    } 
    /*@@passwordSave@@*/
    public function passwordSave(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $oldpassword=$request->post("oldpassword");
        $password=$request->post("password");
        $password2=$request->post("password2");
        $salt=rand(1000,9999);
        if($password!=$password2){
            return Help::success(1,"两次密码不一致");
        }
        $row=DBS::MM("index","userPassword")->where("userid",$ssuserid)->first();
        if($row->password!=Help::umd5($oldpassword.$row->salt)){
            return Help::success(1,"旧密码出错");
        }
        $password= Help::umd5($password.$salt);
        DBS::MM("index","userPassword")->where("userid",$ssuserid)->update([
            "password"=>$password,
            "salt"=>$salt
        ]);
        $redata=[
            "error" => 0, 
            "message" => "success"
        ];
		return json($redata);       
    }
    
    /*@@paypwdSave@@*/
    public function paypwdSave(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $user=DBS::MM("index","User")->get($ssuserid,"userid,telephone");
        $telephone=$user->telephone;
        $yzm=$request->post("yzm","");
        $telephone=$request->post("telephone","");
        $key="smsUser".$telephone;
        $yzm2=Cache::get($key);
        if($yzm!=$yzm2){
            return Help::success(1,"短信验证码出错");
        }
        $password=$request->post("password","h");
        if(empty($password)){
            return Help::success(1,"请输入密码");
        }
        $paypwd=Help::umd5($password);
        DBS::MM("index","userPassword")->where("userid",$userid)->update([
            "paypwd"=>$paypwd
        ]);
        $redata=[
            "error" => 0, 
            "message" => "success"
        ];
		return json($redata);       
    }
     
    /*@@sendSms@@*/
    public function sendSms(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $user=DBS::MM("index","User")->get($ssuserid,"userid,nickname,description,telephone");
        $telephone=$request->get("telephone","");
        if($user->telephone!=''){ 
            $telephone=$user->telephone;
        }
        if(empty($telephone)){
            return Help::success(1,"请先绑定手机号码");
        }
        $yzm=rand(1000,9999);
        $opt=array(
            "telephone"=>$telephone,
            "content"=>"您的验证码".$yzm,
            "yzm"=>$yzm
        );
        
        $key="smsUser".$telephone;
        if(Sms::$TEST==1){
            Cache::set($key,$yzm,180);
            return  Help::success(0,"验证码：".$yzm);
        } 
        $keyExpire="smsUserExpire".$telephone;
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
    /*@@telephoneSave@@*/
    public function telephonesave(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $yzm=$request->post("yzm","");
        $telephone=$request->post("telephone","");
        $key="smsUser".$telephone;
        $yzm2=Cache::get($key);
        if($yzm!=$yzm2){
            return Help::success(1,"短信验证码出错");
        }
        $fm=DBS::MM("index","User");
        $user=$fm->get($ssuserid,"userid,nickname,description,telephone");
        if($user->telephone!=''){
            return Help::success(1,"账户已绑定手机");
        }
        
        if(empty($telephone)){
            return Help::success(1,"请输入手机");
        }
        $fm->where("userid",$ssuserid)->update(["telephone"=>$telephone]);
        return Help::success(0,"绑定成功");
    }
}

?>