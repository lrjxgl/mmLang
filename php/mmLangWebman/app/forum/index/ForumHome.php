<?php
namespace forum\index;
namespace app\forum\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class ForumHome{
    
    /*@@index@@*/
    public function index(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        $userid=intval($request->get("userid"));   
        $user=DBS::MM("index","user")->get($userid,"userid,nickname,user_head,follow_num,followed_num,description");
        if($userid!==ssuserid){
            $follow=DBS::MM("index","follow")->whereRaw("userid=".$ssuserid." AND t_userid=".$userid)->first();
            if($follow){
                $user->isFollow=1;
            }else{
                $user->isFollow=0;
            }
        }
        $fm=DBS::MM("forum","forum");
        $start=$request->get("per_page");
        $limit=4;
        $where=" userid=".$userid." AND status=1 ";
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
        $reData=[
            "error"=>0,
            "message"=>"success",
            "user"=>$user,
            "per_page"=>$per_page,
            "rscount"=>$rscount,
            "list"=>$list

        ];
        return json($reData);
    }
}

?>