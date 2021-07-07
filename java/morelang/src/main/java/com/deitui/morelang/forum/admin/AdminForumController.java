package com.deitui.morelang.forum.admin;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import com.deitui.morelang.forum.model.ForumModel;

import java.util.List;
import  com.alibaba.fastjson.JSON;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
@RestController
public class AdminForumController {
    @RequestMapping("/admin/forum/index")
    public String Index(){
        ForumModel forumModel=new ForumModel();
        List recList=forumModel.getAll("select * from sky_mod_forum where id>%d order by %s  ",100,"id DESC");
        recList=forumModel.where("id>100").limit(0,2).select();
        Map<String,Object> data=new HashMap<String,Object>();
        data=forumModel.where("id=100").selectRow();
        Map<String,Object>  redata=new HashMap<String,Object>();
        String count=forumModel.fields(" id ").where("id>100").selectOne();
        List ids=forumModel.where("id=100").fields("id").limit(0,10000).selectCols();
        redata.put("error",0);
        redata.put("message","succcess");
        redata.put("count",count);
        redata.put("ids",ids);
        redata.put("forum",data);
        redata.put("recList",recList);


        return JSON.toJSONString(redata) ;
    }
    @RequestMapping("/admin/forum/show")
    public String Show(){
        return "论坛详情";
    }
    @RequestMapping("/admin/forum/add")
    public String Add(){
        return "发布帖子";
    }
}