package com.deitui.morelang.index.index;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import com.alibaba.fastjson.JSON;
import com.model.Model;
@RestController
@CrossOrigin("*")
public class IndexController {
	@RequestMapping("/")
	public String Index() {
		return "hello Index aa ss";
	}
}
