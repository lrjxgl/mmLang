<?php
namespace app\index\admin;

use ext\DBS;
use support\Request;
use ext\Help;
class Ad
{
 
    /*@@index@@*/
    public function index(Request $request)
    {
        $start = $request->get("per_page");
        $limit = 100;  
        $fm = DBS::MM("index", "Ad");
        $where = "status in(0,1,2) ";
        $list = $fm
            ->offset($start)
            ->limit($limit)
            ->whereRaw($where)
            ->get();
        $list = $fm->Dselect($list);
        if(!empty($list)){
            $ids=[]; 
            foreach($list as $v){
                $ids[]=$v->tag_id;
                $ids[]=$v->tag_id_2nd;
            }
            $tags=DBS::MM("index","adTags")->getListByIds($ids,"tag_id,title");
            foreach($list as $k=>$v){
                $v->tag_id_name=$tags[$v->tag_id]->title;
                $v->tag_id_2nd_name=$tags[$v->tag_id_2nd]->title;
                $list[$k]=$v; 
            }
        }
        $rscount = $fm->whereRaw($where)->count();
        $per_page = $start + $limit;
        $per_page = $per_page > $rscount ? 0 : $per_page;
        $redata = [
            "error" => 0,
            "message" => "success",
            "list" => $list,
            "per_page" => $per_page,
            "limit"=>$limit,
            "rscount" => $rscount,

        ];
        return json($redata);

    }

    /*@@add@@*/
    public function add(Request $request)
    {

        $id = intval($request->get("id"));
        $row = [];
        $tag_id=$request->get("tag_id","0");
        if ($id) {
            $fm = DBS::MM("index", "Ad");
            $row = $fm->find($id);
            $row->trueimgurl=Help::images_site($row->imgurl);
            $row->trueimgurl2=Help::images_site($row->imgurl2); 
            $tag_id=$row->tag_id;
        }
        $fm2=DBS::MM("index","adTags");
        $tagList=$fm2->whereRaw("status=1 AND pid=0")->get();
        $tagList_2nd=[];
        if($tag_id){
            $tagList_2nd=$fm2->whereRaw("status=1 AND pid=".$tag_id)->get(); 
        } 
        $redata = [
            "error" => 0,
            "message" => "success",
            "data" => $row,
            "tagList"=>$tagList,
            "tagList_2nd"=>$tagList_2nd   
        ];
        return json($redata);
    }
    /*@@getTags@@*/
    public function getTags(Request $request){
        $pid=intval($request->get("tag_id","0"));
        $fm2=DBS::MM("index","adTags");
        $list=$fm2->whereRaw("status=1 AND pid=".$pid)->get();
        $redata = [
            "error" => 0,
            "message" => "success",
            "list"=>$list
        ];  
        return json($redata);
    }
    /*@@save@@*/
    public function save(Request $request)
    {

        $id = intval($request->post("id"));
        $data = [];
        $fm = DBS::MM("index", "Ad");
        $indata = [];
        //处理发布内容
        $indata["tag_id"]=intval($request->post("tag_id","0"));
        $indata["tag_id_2nd"]=intval($request->post("tag_id_2nd","0"));
        $indata["title"]=$request->post("title","");
        $indata["info"]=$request->post("info","");
        $indata["link1"]=$request->post("link1","");
        $indata["link2"]=$request->post("link2","");
        $indata["starttime"]=$request->post("starttime","");
        $indata["endtime"]=$request->post("endtime","");
        $indata["imgurl"]=$request->post("imgurl","");
        $indata["imgurl2"]=$request->post("imgurl2","");
        $indata["orderindex"]=intval($request->post("orderindex","0"));
        $indata["status"]=intval($request->post("status","0"));
        $indata["price"]=floatval($request->post("price","0"));
        $indata["object_id"]=intval($request->post("object_id","0"));
        if ($id) {
            $row = $fm->find($id);

        }
        if ($id) {
          
            $fm->where("id", $id)->update($indata);
        } else {

            $indata["createtime"] = date("Y-m-d H:i:s");
          
            $indata["status"] = 0;
            $id = $fm->insertGetId($indata);
        }

        $redata = [
            "error" => 0,
            "message" => "保存成功",
            "insert_id" => $id,
        ];
        return json($redata);
    }

    /*@@status@@*/
    public function Status(Request $request)
    {

        $id = $request->get("id");
        $fm = DBS::MM("index", "Ad");
        $row = $fm->where("id", $id)->first();

        if ($row->status == 1) {
            $status = 2;
        } else {
            $status = 1;
        }
        $up = $fm->find($id);
        $up->status = $status;
        $up->save();
        $redata = [
            "error" => 0,
            "message" => "success",
            "status" => $status,
        ];
        return json($redata);
    }

    /*@@delete@@*/
    public function delete(Request $request)
    {

        $id = $request->get("id");
        $fm = DBS::MM("index", "Ad");
        $row = $fm->find($id);

        $row->status = 11;
        $row->save();
        $redata = [
            "error" => 0,
            "message" => "success",
        ];
        return json($redata);
    }

}
