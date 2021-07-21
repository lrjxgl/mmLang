package com.model;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import com.alibaba.fastjson.JSONObject;

public class Help {
	public static String error(int errno,String errmsg) {
		return "{\"error\":"+errno+",\"message\":\""+errmsg+"\"}";
	}
	
	public static String success(int errno,String errmsg) {
		return "{\"error\":"+errno+",\"message\":\""+errmsg+"\"}";
	}
	 
	public static Boolean in_array(String str,String[] arr) {
		List<String> tempList = Arrays.asList(arr);
		if(tempList.contains(str))
		{
			return true;
		} else {
		   return false;
		}
	 
	}
	
	public static String _implode(ArrayList arr) {
		int len=arr.size();
		String str="";
		for(int i=0;i<len;i++) {
			if(i>0) {
				str+=",";
			}
			str+="'"+arr.get(i)+"'";
		}
		return str;
	}
	
	public static Object getObjectByKey(List list,String key,String value) {
		
		for(int i=0;i<list.size();i++){
			System.out.println(list.get(i));
			JSONObject json=  (JSONObject) JSONObject.toJSON(list.get(i));
			String userid=json.getString(key);
		      if(userid.equals(value)){
		         return list.get(i);
		      }        
		 }
		return new Object();
	}
	
	public static String image_site(String url) {
		if(url!="") {
			return AppConfig.IMAGES_SITE+url;
		}else {
			return "";
		}
	}
	
	
}
