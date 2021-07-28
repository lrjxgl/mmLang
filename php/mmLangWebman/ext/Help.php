<?php
namespace ext;
class Help{
    public static $imgHost="http://oos.mmlang.com/";   
    public static function images_site($url){
        if(empty($url)){
            return "";
        }
        return self::$imgHost.$url; 
        
    }
    public static function umd5($str){
        return md5(md5($str));
    }

    public static function success($err,$msg){
        return json_encode(["error"=>$err,"message"=>$msg]);
    }
    public static function parseImgsData($imgsdata){
        if(empty($imgsdata)) return [];
        $ex=explode(",",$imgsdata);
        $rss=[];
        foreach($ex as $a){
            $rss[]=array(
                "imgurl"=>$a,
                "trueimgurl"=>self::images_site($a)
            );
        }
        return $rss;
    }
    public static function CheckImgsdata($imgsdata){
        $ex=explode(",",$imgsdata);
        $rss=[];
        foreach($ex as $a){
            $rss[]=self::safeFileName($a);
        }
        return implode(",",$rss);

    }
    public static function safeFileName($filename){
        $filename=substr($filename,strpos($filename,"/attach"));
        return $filename;
    }

}
