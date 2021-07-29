<?php
namespace ext;
use ext\Cache;
use ext\Help;
class UserAccess{

    public static function checkAccess($req){
        $token=$req->get("token");
        if(empty($token)){
            $token=$req->post("token");
        } 
        if(empty($token)){
            return 0;
        } 
        $id=Cache::get($token);
       
        if($id==""){
            return 0;
        }
        return intval($id);
    }

    public static function setToken($userid,$password){
        $salt=rand(1000,9999);
        $password=\md5($password.$salt);
        $token=$userid.".user".$password;
        $token_expire=3600*24*3;
        $refresh_token=$userid.".user".\md5($password.$salt);
        $refresh_token_expire=3600*24*3;
        Cache::set($token,$userid,$token_expire);
        Cache::set($refresh_token,$userid,$refresh_token_expire);
        return [
            "token"=>$token,
            "token_expire"=>time()+$token_expire,
            "refresh_token"=>$refresh_token,
            "refresh_token_expire"=>time()+$refresh_token_expire
        ];
    }
    public static function del($key){
        Cache::del($key);
    }
}