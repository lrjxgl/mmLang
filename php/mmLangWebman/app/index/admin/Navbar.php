<?php
namespace app\index\admin;

use ext\DBS;
use support\Request;
use ext\AdminAccess; 
use ext\Help; 
class Navbar
{
    /*@@index@@*/
    public function index(Request $request)
    {
        $start = $request->get("per_page");
        $limit = 4;
        $fm = DBS::MM("index", "Navbar");
        $group_id = $request->get("group_id", 2);
        $where = " status in(0,1,2) AND group_id=" . $group_id;
        $list = $fm
            ->offset($start)
            ->limit($limit)
            ->whereRaw($where)
            ->orderByRaw("orderindex ASC")
            ->get();
        $list = $fm->Dselect($list);
        $rscount = $fm->whereRaw($where)->count();
        $per_page = $start + $limit;
        $per_page = $per_page > $rscount ? 0 : $per_page;
        $redata = [
            "error" => 0,
            "message" => "ok",
            "list" => $list,
            "per_page" => $per_page,
            "rscount" => $rscount,
            "limit"=>$limit 

        ];
        return json($redata);

    }
    /*@@add@@*/
    public function add(Request $request)
    {
        $id = $request->get("id");
        $pid = $request->get("pid");
        $fm = DBS::MM("index", "Navbar");
        $data = [];
        $group_id = 0;
        if ($id != 0) {
            $data = $fm->where("id", $id)->first();
            $pid = $data->pid;
            $group_id = $data->group_id;
        } elseif ($pid != 0) {
            $parent = $fm->where("id", $pid)->first();
            $group_id = $parent->group_id;
        }

        $groupList = $fm->groupList();
        $parentList = $fm->whereRaw("pid=0 AND status=1")->get();
        $redata = [
            "error" => 0,
            "message" => "ok",
            "data" => $data,
            "groupList" => $groupList,
            "parentList" => $parentList,
            "group_id" => $group_id,
            "pid" => $pid,
        ];
        return json($redata);
    }
    /*@@save@@*/
    public function save(Request $request)
    {

        $id = intval($request->post("id"));
        $data = [];
        $fm = DBS::MM("index", "Navbar");
        $indata = [];
        //处理发布内容

        $indata["title"] = $request->post("title", "");
        $indata["orderindex"] = intval($request->post("orderindex", "0"));
        $indata["link_url"] = $request->post("link_url", "");
        $indata["target"] = $request->post("target", "");
        $indata["pid"] = intval($request->post("pid", "0"));
        $indata["group_id"] = intval($request->post("group_id", "0"));
        $indata["m"] = $request->post("m", "");
        $indata["a"] = $request->post("a", "");
        $indata["status"] = 1;
        $indata["logo"] = $request->post("logo", "");
        $indata["icon"] = $request->post("icon", "");
     
        if ($id) {

            $fm->where("id", $id)->update($indata);
        } else {

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
        $fm = DBS::MM("index", "Navbar");
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
            "message" => "ok",
            "status" => $status,
            "row" => $row,
        ];
        return json($redata);
    }

    /*@@recommend@@*/
    public function recommend(Request $request)
    {
        $id = $request->get("id");
        $fm = DBS::MM("index", "Navbar");

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
            "message" => "ok",
            "isrecommend" => $isrecommend,
        ];
        return json($redata);
    }

    /*@@delete@@*/
    public function delete(Request $request)
    {
        $id = $request->get("id");
        $fm = DBS::MM("index", "Navbar");
        $up = $fm->find($id);
        $up->status = 11;
        $up->save();
        $redata = [
            "error" => 0,
            "message" => "ok",
        ];
        return json($redata);
    }

    /*@@get@@*/
    public function get(Request $request)
    {
        $adminid=AdminAccess::checkAccess($request);
        if(!$adminid){
            return Help::success(1000,"暂无权限"); 
        }
        $admin=DBS::MM("index","admin")->where("id",$adminid)->first();
        $adminGroup=DBS::MM("index","adminGroup")->where("id",$admin->group_id)->first();
        $access=json_decode($adminGroup->content,true);
        $group_id = $request->get("group_id");
        $fm = DBS::MM("index", "Navbar");
        $list = $fm->whereRaw("group_id=" . $group_id)->orderBy("orderindex", "ASC")->get();
        $parent = []; 
        if (!empty($list)) {
            foreach ($list as $k => $v) {
                if(!isset($access[$v->m][$v->a]) && $admin->isfounder!=1 && $v->pid!=0){
                    unset($list[$k]);
                }
            }
            foreach ($list as $k => $v) {

                if ($v->pid == 0) {
                    $parent[] = $v;
                    unset($list[$k]);
                }
            }
            foreach ($parent as $k => $v) {
                $child = [];
                foreach ($list as $kk => $vv) {
                    if ($v->id == $vv->pid) {
                        $child[] = $vv;
                        unset($list[$kk]);
                    }

                }
                $v["child"] = $child;
                $parent[$k] = $v;
                if(empty($child) && $group_id==2){
                    unset($parent[$k]);

                } 
               
            }

        }
        $redata = [
            "error" => 0,
            "message" => "ok",
            "list" => $parent,
        ];
        return json($redata);
    }
}
