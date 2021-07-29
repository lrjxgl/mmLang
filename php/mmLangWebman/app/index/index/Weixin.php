<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\UserAccess;
use ext\Help;
class Weixin
{
	
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","Weixin");
        $where="status in(0,1,2) ";
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
            "message" => "success",
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount

        ];
		return json($redata); 
         
		   
    }

	/*@@list@@*/    
    public function list(Request $request)
    {
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","Weixin");
        $where="status in(0,1,2) ";
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
            "message" => "success",
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount

        ];
		return json($redata); 
         
		   
    }

	/*@@show@@*/
    public function show(Request $request){
        $id=$request->get("id");
        $fm=DBS::MM("index","Weixin");
        $data=$fm->where("id",$id)->first();
        if(empty($data) || $data->status >1){
            return Help::success(1,"数据不存在");
        }
        $data->imgurl=Help::images_site($data->imgurl);
        $author=DBS::MM("index","user")->get($data->userid);
        $redata=[
            "error" => 0, 
            "message" => "success",
            "data"=>$data,
            "author"=>$author 
        ];
		return json($redata);       
    } 

	/*@@my@@*/    
    public function my(Request $request)
    {
		
			$ssuserid=UserAccess::checkAccess($request); 
			if(!$ssuserid){
				return Help::success(1000,"请先登录");
			}
		
	    $start=$request->get("per_page");
        $limit=12;
        $fm=DBS::MM("index","Weixin");
        $where="status in(0,1,2) ";
		$where.=" AND userid=".$ssuserid;
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
            "message" => "success",
            "list"=>$list,
            "per_page"=>$per_page,
            "rscount"=>$rscount

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
        $data=[];
        if($id){
            $fm=DBS::MM("index","Weixin");
            $data=$fm->find($id);
            
			if(empty($row) || $row->userid!=$ssuserid){
				return Help::success(1,"暂无权限");
			}
		

        }
        $redata=[
            "error" => 0, 
            "message" => "success",
            "data"=>$data 
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
        $fm=DBS::MM("index","Weixin");
        $indata=[];
        //处理发布内容
        
$indata["title"]=$request->post("title","");
$indata["status"]=intval($request->post("status","0"));
$indata["userid"]=intval($request->post("userid","0"));
$indata["domain"]=$request->post("domain","");
$indata["catid"]=intval($request->post("catid","0"));
$indata["logo"]=$request->post("logo","");
$indata["imgurl"]=$request->post("imgurl","");
$indata["imgsdata"]=$request->post("imgsdata","");
$indata["appid"]=$request->post("appid","");
$indata["appkey"]=$request->post("appkey","");
$indata["ysid"]=$request->post("ysid","");
$indata["shopid"]=intval($request->post("shopid","0"));
$indata["wx_username"]=$request->post("wx_username","");
$indata["wx_pwd"]=$request->post("wx_pwd","");
$indata["enkey"]=$request->post("enkey","");
$indata["mchid"]=$request->post("mchid","");
$indata["mchkey"]=$request->post("mchkey","");
$indata["sslcert_path"]=$request->post("sslcert_path","");
$indata["sslkey_path"]=$request->post("sslkey_path","");
$indata["openlogin"]=intval($request->post("openlogin","0"));
        if($id){
            $row=$fm->find($id);
            
			if(empty($row) || $row->userid!=$ssuserid){
				return Help::success(1,"暂无权限");
			}
		

        }
        if($id){
            
            $fm->where("id",$id)->update($indata);
        }else{       
            $indata["userid"]=$ssuserid;
            $indata["createtime"]=date("Y-m-d H:i:s");
            
			$indata["status"]=0;
            $id=$fm->insertGetId($indata);
        }
      
       
        $redata=[
            "error" => 0, 
            "message" => "保存成功",
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
        $fm=DBS::MM("index","Weixin");
        $row=$fm->where("id",$id)->first();
		
			if(empty($row) || $row->userid!=$ssuserid){
				return Help::success(1,"暂无权限");
			}
		

        if($row->status==1){
            $status=2;
        }else{
            $status=1;
        }
        $up=$fm->find($id);
        $up->status=$status;
        $up->save();
        $redata=[
            "error" => 0, 
            "message" => "success",
            "status"=>$status
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
        $fm=DBS::MM("index","Weixin");
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