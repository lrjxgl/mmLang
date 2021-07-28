<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\Help;
use ext\UserAccess;
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
        $token["message"]="success";
        return json($token);
        

    }

    /*@@findpwd@@*/
    public function findpwd(Request $request){
        return Help::success(0,"success");
    }
     /*@@findpwdsave@@*/
    public function findpwdSave(Request $request){
        $telephone=$request->post("telephone");
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

}