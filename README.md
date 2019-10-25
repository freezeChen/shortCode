# shortCode
## 固定长度短码生成器

### 使用方法
```
go get -u github.com/freezeChen/shortCode
```

```go
	var id int64 = 592044735
	shortCode, _ := NewShortCode(CodeLength(6)) 
	
	code, err := shortCode.Encode(id) //生成的短码 pH25YW
	

	sid, _ := shortCode.Decode(code)    //sid = 592044735
```

### 变量
```
CodeLength() 生成后字符串的长度 默认7位
Seeds() 种子 生成的字符串基于该值 建议使用10组1-0,A-Z,a-z 62个不重复字符随机打乱作为种子
SeedsIndex() 种子下标   生成后的字符串为pH25YW 其中第3位的5为种子,
```

