<?php
namespace app\forum\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Forum
{ 
	/*@@index@@*/    
    public function index(Request $request)
    {
        $navList=DBS::MM("index","ad")->listByNo("uniapp-forum-nav",123);
        $flashList=DBS::MM("index","ad")->listByNo("uniapp-forum-index",4);
        $adList=DBS::MM("index","ad")->listByNo("uniapp-forum-ad",3);
        $recList=DBS::MM("forum","forumTags")->GetForumByKey("index"); 
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "navList"=>$navList,
            "flashList"=>$flashList,
            "adList"=>$adList,
            "recList"=>$recList,
        ];
		return json($redata); 
         
		   
    }
    /*@@list@@*/
    public function list(Request $request){
        $gid=intval($request->get("gid"));
        $where=" gid=".$gid." AND status=1 ";
        $group=DBS::MM("forum","ForumGroup")->where("gid",$gid)->first();
        
        if(!empty($group)){
            $group->logo=Help::images_site($group->logo);
        }
        $catid=intval($request->get("catid"));
        if($catid){
            
            $where.=" AND catid=".$catid;
        }
        $catList=DBS::MM("forum","forumCategory")->where("gid",$gid)->get();
        $start=$request->get("per_page");
        $limit=4;
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
            "group"=>$group,
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount,
            "catList"=>$catList 

        ];
		return json($redata);
    }

     /*@@new@@*/
     public function new(Request $request){
        
        $where=" status=1 ";
        $start=$request->get("per_page");
        $limit=4;
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
            "rscount"=>$rscount
        ];
		return json($redata);
    }
    /*@@show@@*/
    public function show(Request $request){
        $id=$request->get("id");
        $fm=DBS::MM("forum","Forum");
        $data=$fm->where("id",$id)->first();
        if(empty($data) || $data->status >1){
            return Help::success(1,"数据不存在");  
        }
        $data->imgurl=Help::images_site($data->imgurl);
        $data->imgsList=Help::parseImgsData($data["imgsdata"]);
        $data->content=DBS::MM("forum","forumData")->where("id",$id)->value("content"); 
        $author=DBS::MM("index","user")->get($data->userid);
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "data"=>$data,
            "author"=>$author 
        ];
		return json($redata);       
    } 
     /*@@my@@*/
     public function my(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("forum","Forum");
        $where="status in(0,1,2) ";
        $where.=" AND userid=".$ssuserid;
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
                ->orderby("id","desc")
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

    /*@@add@@*/
    public function add(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $id=intval($request->get("id"));
        $row=[]; 
        if($id){
            $fm=DBS::MM("forum","Forum");
            $row=$fm->find($id);
            if(empty($row) || $row->userid!=$ssuserid){
                return Help::success(1,"暂无权限");
            }
            $imgList=Help::parseImgsData($row["imgsdata"]);
            $row->content=DBS::MM("forum","forumData")->where("id",$id)->value("content");  
        }
        $gid=intval($request->get("gid"));
        $groupList=DBS::MM("forum","forumGroup")->where("status",1)->orderby("orderindex","asc")->get();
        $catList=DBS::MM("forum","forumCategory")->where("status",1)->orderby("orderindex","asc")->get();
        if(!empty($groupList)){
            
            foreach($groupList as $k=>$group){
                $child=[];
                if(!empty($catList)){
                    foreach($catList as $kk=>$v){
                        if($v->gid==$group->gid){
                            $child[]=$v;
                            unset($catList[$kk]);
                        }
                    }
                }
                $group->child=$child;
                $groupList[$k]=$group;
            }
        }
        
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "groupList"=>$groupList,
            "data"=>$row,
            "imgList"=>$imgList 
        ];
		return json($redata);       
    } 
    /*@@save@@*/
    public function save(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $id=intval($request->post("id"));
        $data=[];
        $fm=DBS::MM("forum","Forum"); 
        $indata=[];
        //处理发布内容
        $indata["title"]=$request->post("title","");
        
        $indata["gid"]=intval($request->post("gid","0"));
        $indata["catid"]=intval($request->post("catid","0"));
        $indata["description"]=$request->post("description","");
        $indata["imgurl"]=Help::safeFileName($request->post("imgurl",""));
        $indata["videourl"]=Help::safeFileName($request->post("videourl",""));
        $indata["imgsdata"]=Help::CheckImgsdata($request->post("imgsdata",""));
        if($id){
            $fm=DBS::MM("forum","Forum");
            $row=$fm->find($id);
            if(empty($row) || $row->userid!=$ssuserid){
                return Help::success(1,"暂无权限");
            }
        }
        $conent=$request->post("content");
        if($id){
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $fm->where("id",$id)->update($indata);
            DBS::MM("forum","forumData")->where("id",$id)->update(["content"=>$conent]);
        }else{       
            $indata["userid"]=$ssuserid;
            $indata["createtime"]=date("Y-m-d H:i:s");
            $indata["updatetime"]=date("Y-m-d H:i:s");
            $indata["status"]=0;      
            $id=$fm->insertGetId($indata);
            DBS::MM("forum","forumData")->insert(["content"=>$conent,"id"=>$id]);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "save ok",
            "insert_id"=>$id
        ];
		return json($redata); 
    }
    /*@@status@@*/
    public function Status(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $id=$request->get("id");
        $fm=DBS::MM("forum","Forum");
     
        $row=$fm->find($id);
        if(empty($row) || $row->userid!=$ssuserid){
            return Help::success(1,"暂无权限");
        }
        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
     
        $row->status=$status;
        $row->save();
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "status"=>$status,
            "row"=>$row
        ];
		return json($redata); 
    }

    /*@@recommend@@*/
    public function recommend(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        }
        $id=$request->get("id");
        $fm=DBS::MM("forum","Forum");
        
        $row=$fm->find($id);
        if(empty($row) || $row->userid!=$ssuserid){
            return Help::success(1,"暂无权限");
        }
        if($row->isrecommend==1){
            $isrecommend=0;
        }else{
            $isrecommend=1;
        }
         
        $row->isrecommend=$isrecommend;
        $row->save();
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "isrecommend"=>$isrecommend
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
        $fm=DBS::MM("forum","Forum");
   
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

    /*@@user@@*/
    public function user(Request $request){
        $ssuserid=UserAccess::checkAccess($request); 
        if(!$ssuserid){
            return Help::success(1000,"请先登录");
        } 
        $user=DBS::MM("index","user")->get($ssuserid,"userid,gold,money,grade,nickname,user_head,follow_num,followed_num,description");
        $topic_num=DBS::MM("forum","forum")->where("userid",$ssuserid)->count();
        $comment_num=DBS::MM("forum","forum")->where([
            ["userid",$ssuserid]
        ])->count();
        $redata=[
            
            "error" => 0, 
            "message" => "success",
            "user"=>$user,
            "topic_num"=>$topic_num,
            "comment_num"=>$comment_num
        ];
		return json($redata); 
    }

    /*@@addclick@@ */
    public function addclick(Request $request){
        return Help::success(1,"success");
    }
      
}

?>