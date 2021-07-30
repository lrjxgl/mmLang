<?php
namespace app\middleware;

use Webman\MiddlewareInterface;
use Webman\Http\Response;
use Webman\Http\Request;
use ext\AdminAccess;
use ext\Help; 
use ext\DBS;
class AdminAccessMiddleware implements MiddlewareInterface
{
    public  function process(Request $request, callable $next) : Response
    {
        $path=$request->path();
        $arr=explode("/",$path);
        $m=$arr[2];
        $a=$arr[3];
        $p=$m.".".$a;
        if(!in_array($m,["login"])){
            $adminid=AdminAccess::checkAccess($request);
            if(!$adminid){
    
                return   json([
                    "error"=>1000,
                    "message"=>"请先登录",
                    
                ]);
                
            }
            $admin=DBS::MM("index","admin")->where("id",$adminid)->first();
            if($admin->isfounder!=1 ){
                $adminGroup=DBS::MM("index","adminGroup")->where("id",$admin->group_id)->first();
                $access=json_decode($adminGroup->content,true);
                
                $uns=[
                    "navbar.get"=>1
                ];
                if(!isset($access[$m][$a]) && !isset($uns[$p]) ){
                    return   json([
                        "error"=>1,
                        "message"=>"暂无权限",
                        
                    ]);
                }
            }
        }
        
         
        return $next($request);
    }
}