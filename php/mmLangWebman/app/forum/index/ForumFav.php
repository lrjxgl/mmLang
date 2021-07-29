<?php
namespace app\forum\index;

use ext\DBS;
use support\Request;

class ForumFav
{
    /*@@index@@*/
    public function index(Request $request)
    {
        $start = $request->get("per_page");
        $limit = 12;
        $fm = DBS::MM("index", "Fav");
        $where = " tablename='mod_forum' ";
        $list = $fm
            ->offset($start)
            ->limit($limit)
            ->whereRaw($where)
            ->orderBy("id", "desc")
            ->get();
        $list = $fm->Dselect($list);
        if (!empty($list)) {
            foreach ($list as $v) {
                $ids[] = $v->objectid;
            }
            $res = DBS::MM("forum", "forum")->getListByIds($ids);
            $nlist = [];
            foreach ($list as $k => $v) {
                if (isset($res[$v->objectid])) {
                    $nlist[] = $res[$v->objectid];
                }

            }
            $list = $nlist;
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
}
