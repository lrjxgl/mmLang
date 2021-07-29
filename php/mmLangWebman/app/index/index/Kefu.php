<?php
namespace app\index\index;

use ext\DBS;
use ext\Help;
use ext\UserAccess;
use support\Request;

class Kefu
{
    /*@@index@@*/
    public function index(Request $request)
    {
        $ssuserid = UserAccess::checkAccess($request);
        if (!$ssuserid) {
            return Help . success(1000, "请先登录");
        }
        $start = $request->get("per_page");
        $limit = 12;
        $fm = DBS::MM("index", "KefuMsg");
        $where = " userid=" . $ssuserid . " AND  status in(0,1,2) ";
        $list = $fm
            ->offset($start)
            ->limit($limit)
            ->whereRaw($where)
            ->orderBy("id", "desc")
            ->get();
        $list = $fm->Dselect($list);
        if(!empty($list)){
            $dlist=[];
            foreach($list as $v){
                $ids[]=$v->id; 
                $dlist[]=$v;
            }
            array_multisort($dlist,$ids,SORT_ASC);
            $list=$dlist;
        }
        $rscount = $fm->whereRaw($where)->count();
        $per_page = $start + $limit;
        $per_page = $per_page > $rscount ? 0 : $per_page;
        $redata = [
            "error" => 0,
            "message" => "success",
            "list" => $list,
            "per_page" => $per_page,
            "rscount" => $rscount,

        ];
        return json($redata);
    }

    /*@@save@@*/
    public function save(Request $request)
    {

        $ssuserid = UserAccess::checkAccess($request);
        if (!$ssuserid) {
            return Help::success(1000, "请先登录");
        }

        $id = intval($request->post("id"));
        $data = [];
        $fm = DBS::MM("index", "KefuMsg");
        $indata = [];
        //处理发布内容

        $indata["userid"] = intval($request->post("userid", "0"));
        $indata["content"] = $request->post("content", "");
        $indata["status"] = intval($request->post("status", "0"));
        $indata["tablename"] = $request->post("tablename", "");
        if ($id) {
            $row = $fm->find($id);

            if (empty($row) || $row->userid != $ssuserid) {
                return Help::success(1, "暂无权限");
            }

        }
        if ($id) {

            $fm->where("id", $id)->update($indata);
        } else {
            $indata["userid"] = $ssuserid;
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
}
