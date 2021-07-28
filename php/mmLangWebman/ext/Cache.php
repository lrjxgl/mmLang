<?php
namespace ext;
use ext\DBS; 
class Cache{
     
    public static function set($k,$v,$expire){
        $fm=DBS::MM("index","dbcache");
        $row=$fm->where("k",$k)->first();
        if(!empty($row)){
            $fm->where("k",$k)->update([
                    "v"=>$v,
                    "expire"=>$expire
            ]);
        }else{
            $fm->insert([
                "k"=>$k,
                "v"=>$v,
                "expire"=>$expire
            ]);
        }

    }

    public static function get($k){
        $fm=DBS::MM("index","dbcache");
        $id=$fm->where("k",$k)->value("v");
        
        return $id;
    }

}