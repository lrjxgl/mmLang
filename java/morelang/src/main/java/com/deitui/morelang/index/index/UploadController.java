package com.deitui.morelang.index.index;

import java.io.File;
import java.io.IOException;
import java.util.Map;

import org.springframework.util.ResourceUtils;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import com.alibaba.fastjson.JSONObject;
import com.model.Help;

@RestController
@CrossOrigin("*")
public class UploadController {
	@RequestMapping("/upload/img")
	public String img(
			@RequestParam("upimg") MultipartFile file
	) {
		if (file.isEmpty()) {
            return "上传失败，请选择文件";
        }

        String fileName = file.getOriginalFilename();
        File rootPath;
        String filePath="attach/images/";
        try {
        	rootPath = new File(ResourceUtils.getURL("classpath:").getPath());
            if(!rootPath.exists()) rootPath = new File("");
           
    		File upload = new File(rootPath.getAbsolutePath(),filePath);
    		if(!upload.exists()) upload.mkdirs();
        }catch(Exception e) {
        	 return Help.error(1, "上传失败");
        }
         
     
        File imgurl = new File(rootPath+"/"+filePath + fileName);
        try {
            file.transferTo(imgurl);
            JSONObject redata= JSONObject.parseObject( Help.success(0, "上传成功"));
            redata.put("imgurl", imgurl);
            return redata.toJSONString();
             
        } catch (IOException e) {
        	 return Help.error(1, "上传失败");
        }

		
	}
}
