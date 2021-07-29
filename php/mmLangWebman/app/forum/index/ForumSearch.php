<?php
namespace app\forum\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class forumSearch{
    /**@@index@@**/
    public function index(Request $request){
        $where=" status=1 ";
        $keyword=$request->get("keyword");
        if($keyword){
            $where.=" AND title like '%".$keyword."%' ";
        }else{

            $where.=" AND isrecommend=1 ";
        }  
        $start=intval($request->get("per_page"));
        $limit=12;
        $fm=DBS::MM("forum","Forum");
       
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
                ->orderBy("id","desc")
                ->get();
        $list=$fm->Dselect($list);
        $rscount=$fm->whereRaw($where)->count();
        $per_page=$start+$limit;
        $per_page=$per_page>$rscount?0:$per_page;
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount,

        ];
		return json($redata);
    }
}