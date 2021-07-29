<?php
namespace app\index\index;
use support\Request;
use support\DB;
use ext\DBS;
use ext\Help;
use Intervention\Image\ImageManagerStatic as Image;
use ext\Oos;
use ext\UserAccess;
 
class Upload{
    /*@@img@@*/
    function img(Request $request){
        $ssuserid=UserAccess::checkAccess($request);
        if(!$ssuserid){
             
        }
        $file = $request->file('upimg');
        $dir="/attach/".date("Y/m/d");
        
        if(!file_exists($dir)){
            mkdir($dir,0777,true);
        }
        $ftype=$file->getUploadExtension(); 
        $filename=$dir."/".$file->getUploadName();
        $rootFile= public_path()."/".$filename;
        if ($file && $file->isValid()) {
            $file->move($filename);
            DBS::MM("index","attach")->add([
                "url"=>$filename,
                "userid"=>$ssuserid
            ]);
            //处理图片
            $a=$filename.".100x100.png";
            $b=$filename.".small.png";
            $c=$filename.".middle.png";
            $img = Image::make($filename)->resize(160, 160)->save($a);
            $img = Image::make($filename)->resize(480, 480)->save($b);
            $img = Image::make($filename)->resize(750, 750)->save($c);
            Oos::upload($filename);
            Oos::upload($a);
            Oos::upload($b);
            Oos::upload($c);
            return json(['error' => 0, 'imgurl' => $filename,"trueimgurl"=>Help::images_site($filename)]);
        }
        return json(['error' => 1, 'message' => 'file not found']); 
    }
}