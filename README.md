## 前言
        该项目目的在于简化Golang对数据库的操作，实现快速开发。
        项目实现较其它orm库有以下优点：  
         1、不使用反射实现;  
         2、数据库字段和结构体字段的映射关系可以自己定义
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
        InsertColumns:  []string{"icon","app_id","created_at"},
        // 表示查询操作的时候Scan到的字段
        QueryColumns:   []string{"id","icon","app_id","created_at"},
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
    	app:=T_app{
    		ID:model.GetInt64(fields[0],0),
    		Icon:model.GetString(fields[1]),
    		AppId:model.GetString(fields[2]),
    		CreatedAt:model.GetTime(fields[3],time.Now()),
    	}
    	return app
    }
    // 获取数据操作模型
    m,err:=model.GetModel(sc)
    app:=T_app{
    	Icon:"xxx",
        AppId:"AGJLJWEFJ298R5pF",
        CreatedAt:time.Now(),
        }
    // 插入到数据库
    err=m.Insert(app)
    // 查询t_app表记录数
    count,err:=m.CountAll()
    
    // 检索满足条件的记录，参考condition_test.go
    cond:=model.DbCondition{}.And2(">","id",0).And2("=","app_id","AGJLJWEFJ298R5pF")
    count,err=m.Count(cond)     // 满足条件的记录数
    res,err:=m.Query(cond)      // 满足条件的记录
    fmt.Println(obj.(T_app))    // 使用的时候要断言成指定类型
    err= m.Delete(cond)         // 删除满足条件的对象
    
    // 更新记录，参考setcondition_test.go
    setCond:=model.DbSetCondition{}
    // 表示将app_id like "AGJL"的记录设置为app_id=xxxxxxx
    setCond.Set2("app_id","xxxxxxx").And2("like","app_id","AGJL")
    err=m.Update(setCond)
    
```
