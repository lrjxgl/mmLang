<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Fav
{ 
	 /*@@get@@*/
     public function get(Request $request){
        $userid=UserAccess::checkAccess($request);  
        if($userid==0){
            return Help::success(1,"请先登录");
        }
        $objectid=intval($request->get("objectid"));
        $tablename=$request->get("tablename","");
        $where=[
            ["userid",$userid],
            ["objectid",$objectid],
            ["tablename",$tablename]

        ];
        $action="delete";
        $row=DBS::MM("index","fav")->where($where)->first();
        if(!empty($row)){
            $action="add";
        }
        $reData=[
            "error"=>0,
            "message"=>"success",
            "action"=>$action

        ];
        return json($reData);
    }
    /*@@toggle@@*/
    public function toggle(Request $request){
        $userid=UserAccess::checkAccess($request);  
        if($userid==0){
            return Help::success(1,"请先登录");
        }
        $objectid=intval($request->get("objectid"));
        $tablename=$request->get("tablename","");
        $where=[
            ["userid",$userid],
            ["objectid",$objectid],
            ["tablename",$tablename]

        ];
        $indata=[
            "userid"=>$userid,
            "objectid"=>$objectid,
            "tablename"=>$tablename
        ];
        $row=DBS::MM("index","fav")->where($where)->first();
        if($row){
            $action="delete";
            DBS::MM("index","fav")->where($where)->delete();
        }else{
            DBS::MM("index","fav")->insert($indata);
            $action="add";
        }
        $reData=[
            "error"=>0,
            "message"=>"success",
            "action"=>$action

        ]; 
        return json($reData);
    } 
 
      
}

?>