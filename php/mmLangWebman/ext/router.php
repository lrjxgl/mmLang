<?php
function RouteModuleDir($dir,$module,$group){
	
	Route::group($group,function() use ($dir,$module,$group){
		$files=glob($dir."/*.php");
		foreach($files as $file){
			$c=file_get_contents($file);
			preg_match_all("/@@(\w+)@@/iUs",$c,$arr);
			
			if(isset($arr[1])){
				$funs=$arr[1];
				foreach($funs as $fun){
					$bname=strtolower(str_replace(".php","",basename($file)));
					$fun=strtolower($fun);
					Route::any('/'.$bname.'/'.$fun,"app\module\\".$module."\\".$bname."@".$fun);
				}
			}
			unset($c);
			unset($funs);
			unset($arr);
		}
		unset($files);
	});
}