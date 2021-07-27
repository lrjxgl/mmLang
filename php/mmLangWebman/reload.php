<?php

$dirs=array("./app/","./config","./ext");
$ftypes=array("php");
start();
$wii=0; 
$ctime=2;
if(file_exists("reload-lastTime.txt")){	
		$lastTime=file_get_contents("reload-lastTime.txt");
	}else{
		$lastTime=time();
	}
while(1){
	echo $wii++;
	
	file_put_contents("reload-lastTime.txt",time());
	$isnew=0;
	foreach($dirs as $dir){
		checkDir($dir);
	}
	
 
	$lastTime=time();
	if($isnew){
		start();
	}
	sleep($ctime);

}
function checkDir($dir){
	global $isnew,$lastTime,$ftypes,$ctime;
	$od=opendir($dir);
	while(false!==$file=readdir($od)){
		if($isnew) break;
		 
		if($file=="." || $file=="..") continue;
		if(is_dir($dir."/".$file)){
			checkDir($dir."/".$file);
		}else{
			$mtime=filemtime ($dir."/".$file);
			if($mtime>$lastTime){
				
				$ftype=substr($file,strpos($file,".")+1);
				if(in_array($ftype,$ftypes)){
					$isnew=1;
					echo $dir."/".$file."\r\n";
				break;
				}
				
			}
		}
	}
	 
}
 

function start(){
	
	$port="888";
	$tasklist = "netstat -ano | findstr $port ";
	@exec($tasklist,$arr);
	//print_r($arr);
	$cmd="PHP start.php start";
	//$cmd="ping 127.0.0.1";
	$ps=$arr[0]; 
	if(strpos($ps,":".$port)!==false){
		$pid=substr($ps,strpos($ps,"LISTENING")+9);
		$pid=intval($pid);
		echo $pid;
		if($pid){
			@exec("taskkill -F /pid  ".$pid);
		}
		 
		execInBackground('start cmd.exe @cmd /c "'.$cmd.'"');  
		 
	}else{
		execInBackground('start cmd.exe @cmd /c "'.$cmd.'"');  
		 
		
	}
	echo "restart success";
}
//execInBackground('start cmd.exe @cmd /k "ping 127.0.0.1"'); 
function execInBackground($cmd) { 
   if (substr(php_uname(), 0, 7) == "Windows"){ 
       pclose(popen("start /B ". $cmd."  ", "r"));  
   } 
   else { 
       exec($cmd . " > /dev/null &");   
   } 
}

?>