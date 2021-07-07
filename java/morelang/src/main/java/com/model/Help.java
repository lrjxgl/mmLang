package com.model;

public class Help {
	public static String error(int errno,String errmsg) {
		return "{\"error\":"+errno+",\"message\":\""+errmsg+"\"}";
	}
	
	public static String success(int errno,String errmsg) {
		return "{\"error\":"+errno+",\"message\":\""+errmsg+"\"}";
	}
	 
	
}
