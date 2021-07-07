package com.deitui.morelang.forum.index;
import com.alibaba.fastjson.JSON;
import com.deitui.morelang.forum.model.ForumTagsModel;
import com.deitui.morelang.index.model.AdModel;
import com.model.Model;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import com.model.Login;
@RestController
@CrossOrigin("*")
public class ForumController {
    @RequestMapping("/forum/index")
    public String Index(){
        ForumTagsModel tags=new ForumTagsModel();
        List recList=tags.getForumByKey("index");
        AdModel ad=new AdModel();
        List navList= ad.listByNo("uniapp-forum-nav",100);
        List adList= ad.listByNo("uniapp-forum-ad",100);
        List flashList= ad.listByNo("uniapp-forum-index",100);
        //返回json设置
        Map<String,Object> redata=new HashMap<String,Object>();
        redata.put("error",0);
        redata.put("message","succcess");
        redata.put("recList",recList);
        redata.put("navList",navList);
        redata.put("flashList",flashList);
        redata.put("adList",adList);
        return JSON.toJSONString(redata);
    }
    @RequestMapping("/forum/show")
    public String Show(){
        return "论坛详情";
    }
    @RequestMapping("/forum/add")
    public String Add(@RequestParam(value="token",defaultValue="") String token){
    	
        return "发布帖子";
    }
}