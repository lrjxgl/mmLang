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

    public static function timeago($time){
        if(!is_int($time)){
            $time=strtotime($time);
        }
        $ee=time()-$time;
        if($ee>3600*24){
            $a=$ee/(3600*24);
            $day=intval($a);
            return $day."天前";
        }elseif($ee>3600){
            $h=$ee/3600;
            $h=intval($h);
            return $h."小时前";
        }else{
            $i=intval($ee/60);
            $s=$ee-$s*60;
            return $i."分".$s."秒";
        }
    }

    public static function sql($str){
        $str=addslashes($str);
        return $str;
    }
    public static function _implode($array) {
        if(!empty($array)) {
            return "'".implode("','", is_array($array) ? $array : array($array))."'";
        } else {
            return '';
        }
    }

}
