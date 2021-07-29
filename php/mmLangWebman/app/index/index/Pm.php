<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Pm
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
	    $start=$request->get("per_page");
        $limit=4;
        $fm=DBS::MM("index","PmIndex");
        $where="userid=".$ssuserid." AND status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
                ->get();
        $list=$fm->Dselect($list);
        if($list){
            foreach($list as $v){
                $uids[]=$v["t_userid"];
            }
            $us=DBS::MM("index","user")->getListByIds($uids);
            foreach($list as &$v){
                $v["user"]=$us[$v["t_userid"]];
            }
        }
        $rscount=$fm->whereRaw($where)->count();
        $per_page=$start+$limit;
        $per_page=$per_page>$rscount?0:$per_page;
        $redata=[
            "error" => 0, 
            "message" => "success",
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount

        ];
		return json($redata); 
         
		   
    }

	 

	/*@@detail@@*/
    public function detail(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $userid=intval($request->get("t_userid"));
        if($ssuserid==$userid){
            return Help::success(1,"不能发给自发"); 
        }
        $u=DBS::MM("index","user");
        $ssuser=$u->get($ssuserid);
        $user=$u->get($userid);
        $start=$request->get("per_page");
        $limit=24;
        $fm=DBS::MM("index","Pm");
        DBS::MM("index","PmIndex")->whereRaw("userid=".$ssuserid." AND t_userid=".$userid)->update(["status"=>1]);
        $where="userid=".$ssuserid." AND t_userid=".$userid." AND status in(0,1,2) ";
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
                ->orderBy("id","desc")
                ->get();
        $list=$fm->Dselect($list);
        $dlist=[];
        if($list){
            foreach($list as $k=>$v){
                $ids[]=$v["id"];
                $dlist[]=$v;
            }
            array_multisort($dlist,$ids,SORT_ASC); 
        }
        ;
        $rscount=$fm->whereRaw($where)->count();
        $per_page=$start+$limit;
        $per_page=$per_page>$rscount?0:$per_page;
        $redata=[
            "error" => 0, 
            "message" => "success",
            "user"=>$user,
            "ssuser"=>$ssuser,
            "list"=>$dlist,
            "per_page"=>$per_page,
            "rscount"=>$rscount
        ];
		return json($redata);       
    } 

	 
    /*@@save@@*/
	public function save(Request $request){
       
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $userid=intval($request->post("t_userid","0"));
        
        $data=[];
        $fm=DBS::MM("index","Pm");
        $fmIndex=DBS::MM("index","PmIndex");
        $indata=[];
        $createtime=date("Y-m-d H:i:s");
        //处理发布内容
        
        $indata["userid"]=$ssuserid;
        
        $indata["t_userid"]=$userid;
        $indata["status"]=0;
        $indata["content"]=$request->post("content","");
        $indata["createtime"]=$createtime;
         
        //发布人 
        $indata["type_id"]=1;    
        $fm->insert($indata); 
        $fmIndex->updateOrInsert(["userid"=>$ssuserid,"t_userid"=>$userid],$indata);
       //接收人
        $indata["status"]=0; 
        $indata["userid"]=$userid;
        $indata["type_id"]=1;
        $indata["t_userid"]=$ssuserid;
        $fm->insert($indata); 
        $fmIndex->updateOrInsert(["userid"=>$userid,"t_userid"=>$ssuserid],$indata);
        $redata=[
            "error" => 0, 
            "message" => "保存成功"
        ];
		return json($redata); 
    }
 
    /*@@delete@@*/
    public function delete(Request $request){
		
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
		
        $id=$request->get("id");
        $fm=DBS::MM("index","Pm");
        $row=$fm->find($id); 
        
			if(empty($row) || $row->userid!=$ssuserid){
				return Help::success(1,"暂无权限");
			}
		

        $row->status=11;
        $row->save();
        $redata=[
            "error" => 0, 
            "message" => "success"
        ];
		return json($redata); 
    }
      
}

?>