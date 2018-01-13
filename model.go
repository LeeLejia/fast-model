package model

import (
	"fmt"
	"strings"
	"time"
	sql2 "database/sql"
)

type DbModel struct {
	sc SqlController
}
// 配置数据库模型，其中
// InsertColumns  表示对数据库进行插入操作时插入的字段
// QueryColumns   表示对数据库进行查询操作的时候获取的字段
// InSertFields   函数签名 func(obj interface{}) []interface{}
// 				  obj指定结构，按InsertColumns的顺序解析返回字段数组
// QueryField2Obj 函数签名 func QueryField2Obj(fields []interface{}) interface{}
//				  处理过程和InSertFields相反，按QueryColumns顺序将字段合并成一个结构体返回
// 以上两个函数的作用：
// 				1、避免反射处理，2、能手动的对字段进行转变,即数据库字段和结构体字段不需要一一对应
type SqlController struct {
	// 表名
	TableName string
	// 插入操作的列
	InsertColumns []string
	insertColumns string
	// 插入操作的占位符
	insertPlaceHold string
	// 查找操作获取到的列名
	QueryColumns []string
	queryColumns string
	// 查找操作列占位符
	queryPlaceHold string
	// 返回需要插入的列的集合obj，对应InsertColumns
	InSertFields func(obj interface{}) []interface{}
	// 将queryColumns返回的对象赋值到具体model对象,对应QueryColumns
	QueryField2Obj func(fields []interface{}) interface{}
}

// 获取一个model对象
func GetModel(sqlController SqlController) (DbModel,error){
	if sqlController.TableName==""{
		return DbModel{},fmt.Errorf("请配置TableName！")
	}
	if sqlController.InsertColumns==nil{
		return DbModel{},fmt.Errorf("InsertColumns为空")
	}
	if sqlController.QueryColumns==nil{
		return DbModel{},fmt.Errorf("QueryColumns为空")
	}
	if sqlController.InSertFields ==nil{
		return DbModel{},fmt.Errorf("CFMap为空")
	}
	sqlController.insertColumns = strings.Join(sqlController.InsertColumns,",")
	sqlController.queryColumns = strings.Join(sqlController.QueryColumns,",")
	ics:=make([]string,len(sqlController.InsertColumns))
	for i:=range sqlController.InsertColumns{
		ics[i]=fmt.Sprintf("$%d",i+1)
	}
	sqlController.insertPlaceHold = strings.Join(ics,",")

	qcs:=make([]string,len(sqlController.QueryColumns))
	for i:=range sqlController.QueryColumns{
		qcs[i]=fmt.Sprintf("$%d",i+1)
	}
	sqlController.queryPlaceHold = strings.Join(qcs,",")
	return DbModel{sc:sqlController},nil
}

// 获取表名
func (m *DbModel) GetTableName() string {
	return m.sc.TableName
}

// 设置插入列
func (m *DbModel) SetInsertColumns(columns []string,insertFileds func(obj interface{}) []interface{}){
	ics:=make([]string,len(columns))
	for i:=range columns{
		ics[i]=fmt.Sprintf("$%d",i+1)
	}
	m.sc.InsertColumns = columns
	m.sc.insertPlaceHold = strings.Join(ics,",")
	m.sc.insertColumns = strings.Join(columns,",")
	m.sc.InSertFields = insertFileds
}

// 设置搜索列
func (m *DbModel) SetQueryColumns(columns []string,queryFileds2Obj func(fields []interface{}) interface{}){
	qcs :=make([]string,len(columns))
	for i:=range columns{
		qcs[i]=fmt.Sprintf("$%d",i+1)
	}
	m.sc.QueryColumns = columns
	m.sc.queryPlaceHold = strings.Join(qcs,",")
	m.sc.queryColumns = strings.Join(m.sc.QueryColumns,",")
	m.sc.QueryField2Obj = queryFileds2Obj
}

// 插入操作
func (m *DbModel) Insert(obj interface{}) (err error) {
	stmt, err := Session.Prepare(fmt.Sprintf("INSERT INTO %s(%s) "+
		"VALUES(%s)", m.GetTableName(), m.sc.insertColumns, m.sc.insertPlaceHold))
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.sc.InSertFields(obj)...)
	return
}

/**
获取数据
conditionAndLimit：where id > $1 order by  limit $2 offset $3
 */
func (m *DbModel) Query(cond DbCondition) (result []interface{}, err error) {
	sql:=fmt.Sprintf("SELECT %s FROM %s %s", m.sc.queryColumns, m.GetTableName(), cond.GetCondStr())
	fmt.Println(sql)
	stmt, err := Session.Prepare(sql)
	if err != nil {
		return result, err
	}
	defer stmt.Close()
	rows,err:=stmt.Query(cond.GetParams()...)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		refs := make([]interface{},0, len(m.sc.QueryColumns))
		for  range m.sc.QueryColumns {
			var ref interface{}
			refs = append(refs, &ref)
		}
		err = rows.Scan(refs...)
		row:=make([]interface{},len(m.sc.QueryColumns))
		for i:=range row{
			row[i] = *refs[i].(*interface{})
		}
		if err == nil {
			result = append(result, m.sc.QueryField2Obj(row))
		}
	}
	return result, err
}

/**
获取记录数量
condition:where id > $1 or name == $2
args: 5 'cjwddz'
 */
func (m *DbModel) Count(cond DbCondition) (count int, err error) {
	count = 0
	stmt,err:= Session.Prepare(fmt.Sprintf("SELECT COUNT(*) FROM %s %s", m.GetTableName(), cond.GetCondStr()))
	if err!=nil{
		return 0,err
	}
	defer stmt.Close()
	err = stmt.QueryRow(cond.GetParams()...).Scan(&count)
	return
}
func (m *DbModel) CountAll() (count int, err error){
	count = 0
	stmt,err:= Session.Prepare(fmt.Sprintf("SELECT COUNT(*) FROM %s", m.GetTableName()))
	if err!=nil{
		return 0,err
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&count)
	return
}

/**
更新数据
setAndCondition: SET id = $1,name=$2 where id = $3
args: 3 'cjwddz' 5
 */
func (m *DbModel) Update(setCond DbSetCondition) (err error) {
	stmt, err := Session.Prepare(fmt.Sprintf("UPDATE %s %s", m.GetTableName(),setCond.GetSetCondStr()))
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(setCond.GetSetCondParams()...)
	return
}

/**
删除数据
condition: where id = $3
args: 3 'cjwddz' 5
 */
func (m *DbModel) Delete(cond DbCondition)error{
	stmt,err:= Session.Prepare(fmt.Sprintf("DELETE FROM %s %s", m.GetTableName(), cond.GetCondStr()))
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(cond.GetParams()...)
	return err
}

/**
	执行语句
 */
func (m *DbModel) Exe(sql string,args ...interface{})error{
	stmt,err:= Session.Prepare(sql)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(args...)
	return err
}

/**
	执行语句，并返回结果
 */
func (m *DbModel) ExeForResult(sql string,args ...interface{})(*sql2.Rows, error){
	stmt,err:= Session.Prepare(sql)
	if err!=nil{
		return nil,err
	}
	defer stmt.Close()
	return stmt.Query(args)
}

/**
安全断言
 */
func GetInt(field interface{},def int)int{
	if field == nil{
		return def
	}
	if rs,ok:=field.(int);ok{
		return rs
	}
	if rs,ok:=field.(int64);ok{
		return int(rs)
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return def
}
func GetInt64(field interface{},def int64)int64{
	if field == nil{
		return def
	}
	if rs,ok:=field.(int64);ok{
		return int64(rs)
	}
	if rs,ok:=field.(int);ok{
		return int64(rs)
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return def
}
func GetByteArr(field interface{})[]byte{
	if field == nil{
		return nil
	}
	if rs,ok:=field.([]byte);ok{
		return rs
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return nil
}
func GetString(field interface{})string{
	if field ==nil{
		return ""
	}
	if rs,ok:=field.(string);ok{
		return rs
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return ""
}
func GetBool(field interface{},def bool)bool{
	if field==nil{
		return def
	}
	if rs,ok:=field.(bool);ok{
		return rs
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return def
}
func GetFloat(field interface{},def float32)float32{
	if field == nil{
		return def
	}
	if rs,ok:=field.(float32);ok{
		return rs
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return def
}
func GetFloat64(field interface{},def float64)float64{
	if field == nil{
		return def
	}
	if rs,ok:=field.(float64);ok{
		return rs
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return def
}
func GetTime(field interface{},def time.Time)time.Time{
	if field == nil{
		return def
	}
	if rs,ok:=field.(time.Time);ok{
		return rs
	}
	fmt.Println(fmt.Println("注意，该处强转可能出现错误！请在编译前检查！值："))
	fmt.Println(fmt.Println(field))
	return def
}