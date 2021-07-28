<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Follow
{
	
	/*@@toggle@@*/    
    public function toggle(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $t_userid=intval($request->get("t_userid"));
        $row=DBS::MM("index","follow")->whereRaw("userid=".$ssuserid." AND t_userid=".$t_userid)->first();
        if($row){
            $isFollow=0;
            DBS::MM("index","follow")->whereRaw("userid=".$ssuserid." AND t_userid=".$t_userid)->delete();
        }else{
            $isFollow=1;
            DBS::MM("index","follow")->insert(array(
                "userid"=>$ssuserid, 
                "t_userid"=>$t_userid
            ));
        }
        $reData=[
            "error"=>0,
            "message"=>"success",
            "isFollow"=>$isFollow
        ];
        return json($reData);
    }
      
}

?>