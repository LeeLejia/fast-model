## 前言
        该项目目的在于简化Golang对数据库的操作，实现快速开发。
        项目实现较其它orm库有以下优点：  
            1、不使用反射实现;   
            2、数据表的字段类型和结构体字段类型映射方式可以自己定义;  
            3、调用安全方法，避免了sql注入。  

## 使用方法

导入包
```go
    import "github.com/cjwddz/fast-model"
```
```go
    // 以T_app结构体为例子
    type T_app struct {
    	ID            int64            `json:"id"`
    	Icon          string           `json:"icon"`
    	AppId         string           `json:"app_id"`
    	Name          string           `json:"name"`
    	Version       string           `json:"version"`
    	Describe      string           `json:"describe"`
    	Developer     int              `json:"developer"`
    	Valid         bool             `json:"valid"`
    	File          string           `json:"file"`
    	Src           string           `json:"src"`
    	Expend        *simplejson.Json `json:"expend"`
    	DownloadCount int              `json:"download_count"`
    	CreatedAt     time.Time        `json:"created_at"`
    }
```

```go
    // 首先配置数据库
    err:=model.InitDB("localhost", "5432", "postgres", "password", "dbName","postgres")
    
    // 以T_app结构体为例子,配置数据库模型
    sc:=model.SqlController {
        // 数据库表名
        TableName:      "t_app",    
        // 表示插入操作时插入的字段
        InsertColumns:  []string{"icon","app_id","version","name","describe","developer","valid","file","src","expend","download_count","created_at"},
        // 表示查询操作的时候Scan到的字段
        QueryColumns:   []string{"id","icon","app_id","version","name","describe","developer","valid","file","src","expend","download_count","created_at"},
        // 插入操作时调用，将结构体分解成字段数组，按InsertColumns顺序返回
        InSertFields:   InsertFields,
        // 查询操作时调用，将Scan到的字段数组合并成结构体，按QueryColumns排序组合
        QueryField2Obj: QueryField2Obj,
    }
    // 实现InSertFields
    func InsertFields(obj interface{}) []interface{} {
    	app:=obj.(T_app)
    	expend := []byte{}
    	if app.Expend != nil {
    		bs, err := app.Expend.MarshalJSON()
    		if err==nil{
    			expend = bs
    		}
    	}
    	return []interface{}{
    		app.Icon,app.AppId,app.Version,app.Name,app.Describe,app.Developer,app.Valid,app.File,app.Src,expend,app.DownloadCount,app.CreatedAt,
    	}
    }
    // 实现QueryField2Obj
    func QueryField2Obj(fields []interface{}) interface{} {
    	expend,_:=simplejson.NewJson(model.GetByteArr(fields[10]))
    	app:=T_app{
    		ID:model.GetInt64(fields[0],0),
    		Icon:model.GetString(fields[1]),
    		AppId:model.GetString(fields[2]),
    		Name:model.GetString(fields[4]),
    		Version:model.GetString(fields[3]),
    		Describe:model.GetString(fields[5]),
    		Developer:model.GetInt(fields[6],-1),
    		Valid:model.GetBool(fields[7],true),
    		File:model.GetString(fields[8]),
    		Src:model.GetString(fields[9]),
    		Expend:expend,
    		DownloadCount:model.GetInt(fields[11],0),
    		CreatedAt:model.GetTime(fields[12],time.Now()),
    	}
    	return app
    }
    // 获取数据操作模型
    m,err:=model.GetModel(sc)
    app:=T_app{Icon:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAADhUlEQVRoQ+1aTUwTQRj9ZkvZtkARhGi0RBHDT6IBE0JMSKhCOMihSBC5ICoxJurVxJP8yJEYL0QTD0qACxKCePBCFCXpBQ4UgyEQFVSKVNTYSmmXdnfMNE6zbFpYoGVbXG7tPmbe+973vkk6iyDO/1Ao/o39XDWAUBRb2hhbdx07JOW0TkBDv+cMI8BTQHA0tsj/Y4NhXmDgam+d/g3lFxQQII9hJCaJS0gJCM5SEUEBjX2euZitvLSqGOa76/XZ5OuAgIY+7jyDhMF4qD7lKGCmpreefR4Q0NjnaQUELfEkADC0ddfrW1UBirmmOqBY6YNngpoBZT1QM6Bs/UE9B5Q24P92oOFUwigCwD0TfrPUCUuBxmpKZXzS7987BP24nS9YXQMjeSYXF9bp7U6hY+lotqWCzSULP7CuTdoWhULxJp0WdjKFRYV/ODyJEML0WXIiFGEMznuvOcenXzhXLi7iAu6YtaM56cxBsrBjRVi+O+wrDSXgcr933d4GLTg7LewKAOCmAc5EBWyGi7iArgs61/Qyb7O7MFQeTyi7MeR10bYgm4UjRp7dOq19W5KlMd9+6bW3VLA/iFNSAVLcshsOhxSxnRYyZzNjTcWJJW2vuNkVDgwdVaxp+IN/tHfCX0Y3kSOAkJaLi6gD96vYsRQW9l0f5AIZIJ/T9HCItMRmAkh2msvZA2s8dpD/DycgMwnsHed0yRQXMQGGRHA9qtYZxRUnk6T2hLZUHGZKzMVhG908SQsZGgaZMAZX+wj37eNPnCcXFzEBZHSSnuf8eIbjwUMXNrKo6Mtv3krDTImNL/DBnz4IdmpJ0I8t8PmrPkgVZ2UzXMQEPKll7X4B3O+WhEXxovmZTBoJIw3zRr0tZ1rJPuG3EmI6+wemfNYX0/y6sZmzH800l7N5tLViUkB7pdaalao5eXPIi2kLiCv1uIadTWDAsNF8l1ZWrtAdtxAJ70OLDn918lPSQ4suTvNBwnytWAvh5rsiLURO0LxMZs7uxBnf3Tg4LsVkxJgUHbiNLPJMLG784zBpPTm4HTsgO1S7DdxKiHebm6z9VAGyyhRFkOpAFIsra2nVAVlliiJIdSCKxZW19J5yIO4v+Yhll5555hHAEVn2KQzCAJ97LuoDl/F756KbqCG39QhDV6w6QSqPEVwJ+aqBuCsCmYixlz0EYGzkYlvavSHfVlG4xbe0/V/rXqJPNUQmQgAAAABJRU5ErkJggg==",
         			AppId:"AGJLJWEFJ298R5pF",
         			Name:"appName1234",
         			Version:"1.0",
         			Describe:"a app for test",
         			Developer:15,
         			Valid:true,
         			File:"XXXXXXXXXXXXX",
         			Src:"XXXXXX",
         			Expend:simplejson.New(),
         			DownloadCount:124,
         			CreatedAt:time.Now(),}
    // 插入到数据库
    err=m.Insert(app)
    // 查询t_app表记录数
    count,err:=m.CountAll()
    
    // 检索满足条件的记录，参考condition_test.go
    cond:=model.DbCondition{}.And2(">","id",0).And2("=","name","appName1234")
    count,err=m.Count(cond)     // 满足条件的记录数
    res,err:=m.Query(cond)      // 满足条件的记录
    fmt.Println(obj.(T_app))    // 使用的时候要断言成指定类型
    err= m.Delete(cond)         // 删除满足条件的对象
    
    // 更新记录，参考setcondition_test.go
    setCond:=model.DbSetCondition{}
    // 表示将name like "appName456"的记录设置为name=modifyName
    setCond.Set2("name","modifyName").And2("like","name","appName456")
    err=m.Update(setCond)
    
```
