<?php
namespace ext;
class DBS{
    
	public static function MM($module,$table){
    	$model="\app\\".$module."\model\\".$table."Model";
    	return  new $model();
    }
}