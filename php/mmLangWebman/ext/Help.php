<?php
namespace ext;
class Help{
    public static $imgHost="http://mmlang.com/";
    public static function images_site($url){
        if(empty($url)){
            return "";
        }
        return self::$imgHost.$url; 
        
    }
}
