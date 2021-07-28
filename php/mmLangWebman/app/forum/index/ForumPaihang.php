<?php
namespace app\forum\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class ForumPaihang{
    /*@@index@@*/
    public function index(){
        $wzList=DBS::MM("forum","forum")->where("status",1)->orderBy("comment_num","desc")->get();
        $wzList=DBS::MM("forum","forum")->Dselect($wzList);
        $fsList=DBS::MM("index","user")->where("status",1)->orderBy("followed_num","desc")->get();
        $fsList=DBS::MM("index","user")->Dselect($fsList);
        $reData=[
            "error"=>0,
            "message"=>"success",
            "wzList"=>$wzList,
            "fsList"=>$fsList
        ];
        return json($reData);
    }
}