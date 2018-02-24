package test

import (
	"testing"
	"fmt"
	"github.com/bitly/go-simplejson"
	"time"
	"github.com/cjwddz/fast-model"
	"github.com/bmizerany/assert"
)
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
func TestT(t *testing.T) {
	// 配置数据库
	err:=model.InitDB("localhost", "5432", "postgres", "password", "dbName","postgres")
	assert.Equal(t,err,nil)
	// 配置数据库模型，其中
	// InsertColumns  表示对数据库进行插入操作时插入的字段
	// QueryColumns   表示对数据库进行查询操作的时候获取的字段
	// InSertFields   函数签名 func(obj interface{}) []interface{}
	// 				  obj指定结构，按InsertColumns的顺序解析返回字段数组
	// QueryField2Obj 函数签名 func QueryField2Obj(fields []interface{}) interface{}
	//				  处理过程和InSertFields相反，按QueryColumns顺序将字段合并成一个结构体返回
	// 以上两个函数的作用：
	// 				1、避免反射处理，2、能手动的对字段进行转变,即数据库字段和结构体字段不需要一一对应
	sc:=model.SqlController {
		TableName:      "t_app",
		InsertColumns:  []string{"icon","app_id","version","name","describe","developer","valid","file","src","expend","download_count","created_at"},
		QueryColumns:   []string{"id","icon","app_id","version","name","describe","developer","valid","file","src","expend","download_count","created_at"},
		InSertFields:   InsertFields,
		QueryField2Obj: QueryField2Obj,
	}
	// 获取数据操作模型
	m,err:=model.GetModel(sc)
	assert.Equal(t,err,nil)
	// 测试数据
	apps:=[]T_app{
		{Icon:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAADhUlEQVRoQ+1aTUwTQRj9ZkvZtkARhGi0RBHDT6IBE0JMSKhCOMihSBC5ICoxJurVxJP8yJEYL0QTD0qACxKCePBCFCXpBQ4UgyEQFVSKVNTYSmmXdnfMNE6zbFpYoGVbXG7tPmbe+973vkk6iyDO/1Ao/o39XDWAUBRb2hhbdx07JOW0TkBDv+cMI8BTQHA0tsj/Y4NhXmDgam+d/g3lFxQQII9hJCaJS0gJCM5SEUEBjX2euZitvLSqGOa76/XZ5OuAgIY+7jyDhMF4qD7lKGCmpreefR4Q0NjnaQUELfEkADC0ddfrW1UBirmmOqBY6YNngpoBZT1QM6Bs/UE9B5Q24P92oOFUwigCwD0TfrPUCUuBxmpKZXzS7987BP24nS9YXQMjeSYXF9bp7U6hY+lotqWCzSULP7CuTdoWhULxJp0WdjKFRYV/ODyJEML0WXIiFGEMznuvOcenXzhXLi7iAu6YtaM56cxBsrBjRVi+O+wrDSXgcr933d4GLTg7LewKAOCmAc5EBWyGi7iArgs61/Qyb7O7MFQeTyi7MeR10bYgm4UjRp7dOq19W5KlMd9+6bW3VLA/iFNSAVLcshsOhxSxnRYyZzNjTcWJJW2vuNkVDgwdVaxp+IN/tHfCX0Y3kSOAkJaLi6gD96vYsRQW9l0f5AIZIJ/T9HCItMRmAkh2msvZA2s8dpD/DycgMwnsHed0yRQXMQGGRHA9qtYZxRUnk6T2hLZUHGZKzMVhG908SQsZGgaZMAZX+wj37eNPnCcXFzEBZHSSnuf8eIbjwUMXNrKo6Mtv3krDTImNL/DBnz4IdmpJ0I8t8PmrPkgVZ2UzXMQEPKll7X4B3O+WhEXxovmZTBoJIw3zRr0tZ1rJPuG3EmI6+wemfNYX0/y6sZmzH800l7N5tLViUkB7pdaalao5eXPIi2kLiCv1uIadTWDAsNF8l1ZWrtAdtxAJ70OLDn918lPSQ4suTvNBwnytWAvh5rsiLURO0LxMZs7uxBnf3Tg4LsVkxJgUHbiNLPJMLG784zBpPTm4HTsgO1S7DdxKiHebm6z9VAGyyhRFkOpAFIsra2nVAVlliiJIdSCKxZW19J5yIO4v+Yhll5555hHAEVn2KQzCAJ97LuoDl/F756KbqCG39QhDV6w6QSqPEVwJ+aqBuCsCmYixlz0EYGzkYlvavSHfVlG4xbe0/V/rXqJPNUQmQgAAAABJRU5ErkJggg==",
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
			CreatedAt:time.Now(),},
		{Icon:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAADhUlEQVRoQ+1aTUwTQRj9ZkvZtkARhGi0RBHDT6IBE0JMSKhCOMihSBC5ICoxJurVxJP8yJEYL0QTD0qACxKCePBCFCXpBQ4UgyEQFVSKVNTYSmmXdnfMNE6zbFpYoGVbXG7tPmbe+973vkk6iyDO/1Ao/o39XDWAUBRb2hhbdx07JOW0TkBDv+cMI8BTQHA0tsj/Y4NhXmDgam+d/g3lFxQQII9hJCaJS0gJCM5SEUEBjX2euZitvLSqGOa76/XZ5OuAgIY+7jyDhMF4qD7lKGCmpreefR4Q0NjnaQUELfEkADC0ddfrW1UBirmmOqBY6YNngpoBZT1QM6Bs/UE9B5Q24P92oOFUwigCwD0TfrPUCUuBxmpKZXzS7987BP24nS9YXQMjeSYXF9bp7U6hY+lotqWCzSULP7CuTdoWhULxJp0WdjKFRYV/ODyJEML0WXIiFGEMznuvOcenXzhXLi7iAu6YtaM56cxBsrBjRVi+O+wrDSXgcr933d4GLTg7LewKAOCmAc5EBWyGi7iArgs61/Qyb7O7MFQeTyi7MeR10bYgm4UjRp7dOq19W5KlMd9+6bW3VLA/iFNSAVLcshsOhxSxnRYyZzNjTcWJJW2vuNkVDgwdVaxp+IN/tHfCX0Y3kSOAkJaLi6gD96vYsRQW9l0f5AIZIJ/T9HCItMRmAkh2msvZA2s8dpD/DycgMwnsHed0yRQXMQGGRHA9qtYZxRUnk6T2hLZUHGZKzMVhG908SQsZGgaZMAZX+wj37eNPnCcXFzEBZHSSnuf8eIbjwUMXNrKo6Mtv3krDTImNL/DBnz4IdmpJ0I8t8PmrPkgVZ2UzXMQEPKll7X4B3O+WhEXxovmZTBoJIw3zRr0tZ1rJPuG3EmI6+wemfNYX0/y6sZmzH800l7N5tLViUkB7pdaalao5eXPIi2kLiCv1uIadTWDAsNF8l1ZWrtAdtxAJ70OLDn918lPSQ4suTvNBwnytWAvh5rsiLURO0LxMZs7uxBnf3Tg4LsVkxJgUHbiNLPJMLG784zBpPTm4HTsgO1S7DdxKiHebm6z9VAGyyhRFkOpAFIsra2nVAVlliiJIdSCKxZW19J5yIO4v+Yhll5555hHAEVn2KQzCAJ97LuoDl/F756KbqCG39QhDV6w6QSqPEVwJ+aqBuCsCmYixlz0EYGzkYlvavSHfVlG4xbe0/V/rXqJPNUQmQgAAAABJRU5ErkJggg==",
			AppId:"AGJLJWEFJ298R5pF",
			Name:"appName456",
			Version:"1.0",
			Describe:"a app for test",
			Developer:14,
			Valid:true,
			File:"XXXXXXXXXXXXX",
			Src:"XXXXXX",
			Expend:simplejson.New(),
			DownloadCount:124,
			CreatedAt:time.Now(),},
		{Icon:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAADhUlEQVRoQ+1aTUwTQRj9ZkvZtkARhGi0RBHDT6IBE0JMSKhCOMihSBC5ICoxJurVxJP8yJEYL0QTD0qACxKCePBCFCXpBQ4UgyEQFVSKVNTYSmmXdnfMNE6zbFpYoGVbXG7tPmbe+973vkk6iyDO/1Ao/o39XDWAUBRb2hhbdx07JOW0TkBDv+cMI8BTQHA0tsj/Y4NhXmDgam+d/g3lFxQQII9hJCaJS0gJCM5SEUEBjX2euZitvLSqGOa76/XZ5OuAgIY+7jyDhMF4qD7lKGCmpreefR4Q0NjnaQUELfEkADC0ddfrW1UBirmmOqBY6YNngpoBZT1QM6Bs/UE9B5Q24P92oOFUwigCwD0TfrPUCUuBxmpKZXzS7987BP24nS9YXQMjeSYXF9bp7U6hY+lotqWCzSULP7CuTdoWhULxJp0WdjKFRYV/ODyJEML0WXIiFGEMznuvOcenXzhXLi7iAu6YtaM56cxBsrBjRVi+O+wrDSXgcr933d4GLTg7LewKAOCmAc5EBWyGi7iArgs61/Qyb7O7MFQeTyi7MeR10bYgm4UjRp7dOq19W5KlMd9+6bW3VLA/iFNSAVLcshsOhxSxnRYyZzNjTcWJJW2vuNkVDgwdVaxp+IN/tHfCX0Y3kSOAkJaLi6gD96vYsRQW9l0f5AIZIJ/T9HCItMRmAkh2msvZA2s8dpD/DycgMwnsHed0yRQXMQGGRHA9qtYZxRUnk6T2hLZUHGZKzMVhG908SQsZGgaZMAZX+wj37eNPnCcXFzEBZHSSnuf8eIbjwUMXNrKo6Mtv3krDTImNL/DBnz4IdmpJ0I8t8PmrPkgVZ2UzXMQEPKll7X4B3O+WhEXxovmZTBoJIw3zRr0tZ1rJPuG3EmI6+wemfNYX0/y6sZmzH800l7N5tLViUkB7pdaalao5eXPIi2kLiCv1uIadTWDAsNF8l1ZWrtAdtxAJ70OLDn918lPSQ4suTvNBwnytWAvh5rsiLURO0LxMZs7uxBnf3Tg4LsVkxJgUHbiNLPJMLG784zBpPTm4HTsgO1S7DdxKiHebm6z9VAGyyhRFkOpAFIsra2nVAVlliiJIdSCKxZW19J5yIO4v+Yhll5555hHAEVn2KQzCAJ97LuoDl/F756KbqCG39QhDV6w6QSqPEVwJ+aqBuCsCmYixlz0EYGzkYlvavSHfVlG4xbe0/V/rXqJPNUQmQgAAAABJRU5ErkJggg==",
			AppId:"AGJLJWEFJ298R5pF",
			Name:"appName456",
			Version:"1.0",
			Describe:"a app for test",
			Developer:14,
			Valid:true,
			File:"XXXXXXXXXXXXX",
			Src:"XXXXXX",
			Expend:simplejson.New(),
			DownloadCount:124,
			CreatedAt:time.Now(),},
	}
	// 插入数据
	for _,app:=range apps{
		err=m.Insert(app)
		assert.Equal(t,err,nil)
	}
	// 查询t_app表记录数
	count,err:=m.CountAll()
	assert.Equal(t,err,nil)
	fmt.Println(fmt.Sprintf("count:%d",count))

	// 新建条件
	cond:=model.DbCondition{}
	// 如果是从request中获取条件，参考condition_test.go
	cond.And(">","id",0).And("=","name","appName1234")
	count,err=m.Count(cond)
	assert.Equal(t,err,nil)
	fmt.Println(fmt.Sprintf("count:%d",count))

	// 查看满足条件的对象
	res,err:=m.Query(cond)
	assert.Equal(t,err,nil)
	for _,obj:=range res{
		fmt.Println(obj.(T_app))
	}
	// 删除满足条件的对象
	err= m.Delete(cond)
	assert.Equal(t,err,nil)

	// 更新数据
	setCond:=model.DbSetCondition{}
	setCond.Set("name","modifyName").And("like","name","appName456")
	err=m.Update(setCond)
	assert.Equal(t,err,nil)
}
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
