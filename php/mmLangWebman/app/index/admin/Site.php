<?php
namespace app\index\admin;

use ext\DBS;
use support\Request;
use ext\Help;
class Site
{

    /*@@index@@*/
    public function index(Request $request)
    {

        $data = DBS::MM("index", "site")->get();
        $data->truelogo=Help::images_site($data->logo);  
        $redata = [
            "error" => 0,
            "message" => "success",
            "data" => $data,
        ];
        return json($redata);
    }

    /*@@save@@*/
    public function save(Request $request)
    {

        $siteid = intval($request->post("siteid"));
     
        $fm = DBS::MM("index", "Site");
        $indata = [];
        //处理发布内容
        $data = $fm->get();
        $indata["sitename"] = $request->post("sitename", "");
        $indata["domain"] = $request->post("domain", "");
        $indata["title"] = $request->post("title", "");
        $indata["keywords"] = $request->post("keywords", "");
        $indata["description"] = $request->post("description", "");
        $indata["close_why"] = $request->post("close_why", "");
        $indata["logo"] = $request->post("logo", "");
        $indata["icp"] = $request->post("icp", "");
        $indata["status"] = intval($request->post("status", "0"));
        $indata["template"] = $request->post("template", "");
        $indata["statjs"] = $request->post("statjs", "");
        $indata["index_template"] = $request->post("index_template", "");
        $indata["wap_template"] = $request->post("wap_template", "");
        $indata["wapbg"] = $request->post("wapbg", "");
        
        $fm->where("siteid", $data->siteid)->update($indata);
        $redata = [
            "error" => 0,
            "message" => "保存成功",
            "insert_id" => $siteid,
        ];
        return json($redata);
    }

    /*@@status@@*/
    public function Status(Request $request)
    {

        $siteid = $request->get("siteid");
        $fm = DBS::MM("index", "Site");
        $row = $fm->where("siteid", $siteid)->first();

        if ($row->status == 1) {
            $status = 2;
        } else {
            $status = 1;
        }
        $up = $fm->find($siteid);
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

        $siteid = $request->get("siteid");
        $fm = DBS::MM("index", "Site");
        $row = $fm->find($siteid);

        $row->status = 11;
        $row->save();
        $redata = [
            "error" => 0,
            "message" => "success",
        ];
        return json($redata);
    }

}
