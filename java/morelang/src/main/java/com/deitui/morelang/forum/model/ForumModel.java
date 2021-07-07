package com.deitui.morelang.forum.model;
import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.model.AppConfig;
import com.model.Model;

import java.util.List;
import java.util.Map;

public class ForumModel extends  Model {
    public ForumModel(){
        this.table="mod_forum";
        this.preTable=this.table_pre+this.table;
    }
    public List Dselect(){

        List list=this.select();
        int len=list.size();
        for(int i=0;i<len;i++){
            Object obj=list.get(i);
            JSONObject json= JSONObject.parseObject(JSONObject.toJSONString(obj));
            json.put("imgurl", AppConfig.IMAGES_SITE +json.get("imgurl"));
            String imgsdata=json.get("imgsdata")+"";
            String imgList[]=imgsdata.split(",");
            for(int ii=0;ii<imgList.length;ii++){
                imgList[ii]=AppConfig.IMAGES_SITE+imgList[ii];
            }
            json.put("imgList",imgList);
            list.set(i,json);
        }
        return list;
    }
}
