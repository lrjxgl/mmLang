<?php
namespace app\index\admin;

use ext\DBS;
use ext\Help;
use support\Request;
use ext\AdminAccess;
class Article
{

    /*@@index@@*/
    public function index(Request $request)
    {
        $adminid=AdminAccess::checkAccess($request);
        if(!$adminid){
            return Help::success(1000,"暂无权限"); 
        } 
        $start = $request->get("per_page");
        $limit = 12;
        $fm = DBS::MM("index", "Article");
        $where = "status in(0,1,2) ";
        switch ($request->get("type", "")) {
            case "new":
                $where=" status=0 ";
                break;
            case "pass":
                $where=" status=1 ";
                break;
            case "forbid":
                $where=" status=2 ";
                break;
            default :
                $where="   status in(0,1,2) ";
                break;
        }
        switch($request->get("recommend","")){
            case "yes":
                $where.=" AND isrecommend=1";
                break;
            case "no":
                $where.=" AND isrecommend=0";
                break;
        }
        $id=intval($request->get("id","0"));
        if($id>0){
            $where.=" AND id=".$id;
        }
        $title=$request->get("title",""); 
        if($title!=""){
            $where.=" AND title like '%".Help::sql($title)."%' ";
        }
        $catid=intval($request->get("catid","0"));
        $dbsCat= DBS::MM("index", "category");
        if($catid>0){
            $fcatids=$dbsCat->id_family($catid);
            $where.=" AND catid in(".Help::_implode($fcatids).") ";
        }
        $list = $fm
            ->offset($start)
            ->limit($limit)
            ->whereRaw($where)
            ->orderBy("id", "desc")
            ->get();
        $list = $fm->Dselect($list);
        if (!empty($list)) {
            $ids = [];
            foreach ($list as $v) {
                $ids[] = $v->catid;

            }
            $cats =$dbsCat->getListByIds($ids, "catid,title,cname");
            foreach ($list as $k => $v) {
                $v->catid_name = $cats[$v->catid]->cname;

                $list[$k] = $v;
            }
        } 
        $rscount = $fm->whereRaw($where)->count();
        $per_page = $start + $limit;
        $per_page = $per_page > $rscount ? 0 : $per_page;
        $catList = DBS::MM("index", "Category")->children(0, "article", 1);
        $redata = [
            "error" => 0,
            "message" => "success",
            "list" => $list,
            "per_page" => $per_page,
            "rscount" => $rscount,
            "catList" => $catList,

        ];
        return json($redata);

    }

    /*@@add@@*/
    public function add(Request $request)
    {

        $id = intval($request->get("id"));
        $data = [];
        $imgList = [];
        if ($id) {
            $fm = DBS::MM("index", "Article");
            $data = $fm->find($id);
            $data->trueimgurl = Help::images_site($data->imgurl);
            $data->imgsList = Help::parseImgsData($data->imgsdata);
            $data->content = DBS::MM("index", "articleData")->where("id", $id)->value("content");
            $imgList = Help::parseImgsData($data->imgsdata);
        }
        $catList = DBS::MM("index", "Category")->children(0, "article", 1);
        $redata = [
            "error" => 0,
            "message" => "success",
            "data" => $data,
            "catList" => $catList,
            "imgList" => $imgList,
        ];
        return json($redata);
    }

    /*@@save@@*/
    public function save(Request $request)
    {

        $id = intval($request->post("id"));
        $data = [];
        $fm = DBS::MM("index", "Article");
        $indata = [];
        //处理发布内容

        $indata["title"] = $request->post("title", "");
        $indata["catid"] = intval($request->post("catid", "0"));
        $indata["catid_top"] = intval($request->post("catid_top", "0"));
        $indata["catid_2nd"] = intval($request->post("catid_2nd", "0"));
        $indata["description"] = $request->post("description", "");
        $indata["status"] = intval($request->post("status", "0"));
        $indata["imgurl"] = Help::safeFileName($request->post("imgurl", ""));
        $indata["tpl"] = $request->post("tpl", "");
        $indata["author"] = $request->post("author", "");
        $indata["price"] = floatval($request->post("price", "0"));
        $indata["market_price"] = floatval($request->post("market_price", "0"));
        $indata["total_num"] = intval($request->post("total_num", "0"));
        $indata["sold_num"] = intval($request->post("sold_num", "0"));
        $indata["grade"] = intval($request->post("grade", "0"));
        $indata["downurl"] = $request->post("downurl", "");
        $indata["downsize"] = floatval($request->post("downsize", "0"));
        $indata["orderindex"] = intval($request->post("orderindex", "0"));
        $indata["videourl"] = $request->post("videourl", "");
        $indata["tags"] = $request->post("tags", "");
        $indata["imgsdata"] = Help::CheckImgsdata($request->post("imgsdata", ""));
        if ($data["imgurl"] == "" && $data["imgsdata"] != "") {
            $ex = explode(",", $indata["imgsdata"]);
            $data["imgurl"] = $ex[0];
        }
        if ($id) {
            $row = $fm->find($id);

        }
        if ($id) {
            $indata["updatetime"] = date("Y-m-d H:i:s");
            $fm->where("id", $id)->update($indata);
            $content = $request->post("content", "h");
            DBS::MM("index", "articleData")->where("id", $id)->update([
                "content" => $content,
            ]);
        } else {

            $indata["createtime"] = date("Y-m-d H:i:s");
            $indata["updatetime"] = date("Y-m-d H:i:s");
            $indata["status"] = 0;
            $id = $fm->insertGetId($indata);
            $content = $request->post("content", "h");
            DBS::MM("index", "articleData")->insert([
                "id" => $id,
                "content" => $content,
            ]);
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
        $fm = DBS::MM("index", "Article");
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

    /*@@recommend@@*/
    public function recommend(Request $request)
    {

        $id = $request->get("id");
        $fm = DBS::MM("index", "Article");

        $row = $fm->where("id", $id)->first();

        if ($row->isrecommend == 1) {
            $isrecommend = 0;
        } else {
            $isrecommend = 1;
        }

        $row->isrecommend = $isrecommend;
        $row->save();
        $redata = [
            "error" => 0,
            "message" => "success",
            "isrecommend" => $isrecommend,
        ];
        return json($redata);
    }

    /*@@delete@@*/
    public function delete(Request $request)
    {

        $id = $request->get("id");
        $fm = DBS::MM("index", "Article");
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
