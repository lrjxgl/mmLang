<?php
namespace app\index\admin;
use support\Request;
use support\DB;
use ext\DBS;
class Navbar
{ 
	/*@@index@@*/    
    public function index(Request $request)
    {
	    $start=$request->get("per_page");  
        $limit=4;
        $fm=DBS::MM("index","Navbar");
        $group_id=$request->get("group_id",2);
        $where=" status in(0,1,2) AND group_id=".$group_id;
		$list=$fm
                ->offset($start)
                ->limit($limit)
                ->whereRaw($where)
                ->orderByRaw("orderindex ASC")
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
    /*@@add@@*/
    public function add(Request $request){
        $id=$request->get("id");
        $fm=DBS::MM("index","Navbar");
        $data=$fm->where("id",$id)->first();
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "data"=>$data 
        ];
		return json($redata);       
    } 
    /*@@save@@*/
    public function save(Request $request){
        $fm=DBS::MM("index","Navbar");
        $fm->title="aaaa";  
        $fm->save();
        $id=$fm->id; 
        $redata=[
            "error" => 0, 
            "message" => "save ok",
            "insert_id"=>$id
        ];
		return json($redata); 
    }
    /*@@status@@*/
    public function Status(Request $request){
        $id=$request->get("id");
        $fm=DBS::MM("index","Navbar");
        $row=$fm->where("id",$id)->first();
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
            "message" => "ok",
            "status"=>$status,
            "row"=>$row
        ];
		return json($redata); 
    }

    /*@@recommend@@*/
    public function recommend(Request $request){
        $id=$request->get("id");
       $fm=DBS::MM("index","Navbar");
        
        $row=$fm->where("id",$id)->first();
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
        $id=$request->get("id");
        $fm=DBS::MM("index","Navbar");
        $up=$fm->find($id);
        $up->status=11;
        $up->save();
        $redata=[
            "error" => 0, 
            "message" => "ok"
        ];
		return json($redata); 
    }
     
    /*@@get@@*/
    public function get(Request $request){
        $group_id=$request->get("group_id");
        $fm=DBS::MM("index","Navbar");
        $list=$fm->whereRaw("group_id=".$group_id)->orderBy("orderindex","ASC")->get();
        $parent=[];
        if(!empty($list)){
            
            foreach($list as $k=>$v){
                if($v->pid==0){
                    $parent[]=$v;
                    unset($list[$k]);
                }
            }
            foreach($parent as $k=>$v){
                $child=[];
                foreach($list as $kk=>$vv){
                    if($v->id==$vv->pid){
                        $child[]=$vv;
                        unset($list[$kk]);
                    }

                }
                $v["child"]=$child;
                $parent[$k]=$v;
            }

        }
        $redata=[
            "error" => 0, 
            "message" => "ok",
            "list"=>$parent
        ];
		return json($redata); 
    }
}

?>