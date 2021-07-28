<?php 
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\Help;
use ext\UserAccess;
class Index{
    /*@@Index@@*/
    public function Index(){
        $navList=DBS::MM("index","Ad")->listByNo("uniapp-nav");

        $redata=[
            "error" => 0, 
            "message" => "ok",
            "navList"=>$navList,
        ];
		return json($redata); 
    }
}