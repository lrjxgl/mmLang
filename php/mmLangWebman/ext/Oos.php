<?php
namespace ext;
class Oos{
    public static function Upload($from){
    
        $filename=substr($from,strpos("/attach",$from));
        $to="D:/phpstudy_pro/WWW/oos.mmlang.com/".$filename;
        $dir=dirname($to);
        if(!file_exists($dir)){
            mkdir($dir,0777,true);
        }
        copy($from,$to);
        

    }
}