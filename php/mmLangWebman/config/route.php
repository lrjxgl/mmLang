<?php
/**
 * This file is part of webman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author    walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link      http://www.workerman.net/
 * @license   http://www.opensource.org/licenses/mit-license.php MIT License
 */

use Webman\Route;
/**自动route
 * forum/index/forumController.php
 * forum/admin/forumController.php
 * 
 */
function RouteModuleDir($module){
	$apps=["index","admin"];
	foreach($apps as $appDir){
		$group="/";
		$dir="app/".$module."/".$appDir;
		if(!file_exists($dir)) continue;
		//处理group 
		if($appDir=="index"){
			$group="";
		}else{
			$group="/".$appDir;
		}
		$files=glob("app/".$module."/".$appDir."/*.php");
		foreach($files as $file){
			$c=file_get_contents($file);
			
			preg_match_all("/@@(\w+)@@/iUs",$c,$arr);
				
			if(isset($arr[1])){
				$funs=$arr[1];
				foreach($funs as $fun){
					$bname=str_replace(".php","",basename($file));
					$lowname=strtolower($bname);

					$fun=strtolower($fun);
					//app\forum\admin\\forum@index
					$ph="app\\".$module."\\".$appDir."\\".$lowname."@".$fun; 
					$routePath=$group.'/'.toUnderScore($bname).'/'.$fun;
					$routeLog="";
					$routeLog.=$ph."  ".$routePath;
					file_put_contents("log.txt",$routeLog."\r\n",FILE_APPEND);
					if($appDir=="admin"){
						Route::any($routePath,$ph)->middleware([
							app\middleware\AdminAccessMiddleware::class
						]);
					}else{
						Route::any($routePath,$ph);
					}
					  
				}
			}  
			unset($c);
			unset($funs);
			unset($arr);
		} 
		unset($files); 
	}
	
}

function toUnderScore($str){
          $dstr = preg_replace_callback('/([A-Z]+)/',function($matchs)
         {
             return '_'.strtolower($matchs[0]);
          },$str);
          return trim(preg_replace('/_{2,}/','_',$dstr),'_');
}

 

 
//应用命名
RouteModuleDir("index");  
RouteModuleDir("forum"); 

//Route::disableDefaultRoute();