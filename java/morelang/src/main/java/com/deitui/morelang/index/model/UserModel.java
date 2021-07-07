package com.deitui.morelang.index.model;
import com.alibaba.fastjson.JSONObject;
import com.model.AppConfig;
import com.model.Model;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
public class UserModel extends Model {
	 public UserModel(){
		 this.preTable=this.table_pre+"user";
	}
	public Map get(int userid) {
		Map<String,Object> user=this.where("userid="+userid).selectRow();
		user.put("user_head",AppConfig.IMAGES_SITE+user.get("user_head")+"");
		return user;
	}
}
