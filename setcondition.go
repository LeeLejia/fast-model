package model

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

type DbSetCondition struct{
	DbCondition
	fields    []string
	newValues []interface{}
}

/**
添加一个设置项
 */
func (dbSet *DbSetCondition) Set(r *http.Request, t_key string) *DbSetCondition {
	if len(t_key)<=2 || t_key[1]!='_'{
		// todo 写到系统日志
		fmt.Println("是否错误调用了GetCondition？t_key格式为类型首写和列名，如int类型id则为i_id,再如：s_name,b_valid")
		return dbSet
	}
	t:= t_key[2:]
	value :=r.PostFormValue(t)
	if value ==""{
		return dbSet
	}
	field:=t
	var newValue interface{}
	switch t_key[0] {
	case 'b':
		if strings.ToLower(value)=="true"{
			newValue = true
		}else{
			newValue = false
		}
	case 'i':
		i,err:=strconv.Atoi(value)
		if err!=nil{
			// todo 写到系统日志
			fmt.Println(fmt.Sprintf("类型转化出错！key=%s,value=%s,err=%s", t_key,value,err.Error()))
			return dbSet
		}
		newValue = i
	default:
		newValue = value
	}
	return dbSet.Set2(field,newValue)
}

/**
添加一个设置项
 */
func (dbSet *DbSetCondition) Set2(field string,newValue interface{}) *DbSetCondition {
	if dbSet.fields == nil{
		dbSet.fields = make([]string,0)
		dbSet.newValues = make([]interface{},0)
	}
	dbSet.fields = append(dbSet.fields, field)
	dbSet.newValues = append(dbSet.newValues, newValue)
	return dbSet
}

/**
重置条件和设置字段
 */
func (dbSet *DbSetCondition) Reset() *DbSetCondition {
	dbSet.fields = make([]string,0)
	dbSet.newValues = make([]interface{},0)
	dbSet.DbCondition.Reset()
	return dbSet
}

/**
重置条件
 */
func (dbSet *DbSetCondition) ResetCondition() *DbSetCondition {
	dbSet.DbCondition.Reset()
	return dbSet
}


/**
获取表达式
 */
func (dbSet *DbSetCondition) GetSetCondStr()string{

	// 设置语句
	if dbSet.fields == nil{
		return ""
	}
	setCount := len(dbSet.fields)
	if setCount ==0{
		return ""
	}
	setStr:= "SET "
	for i,f:=range dbSet.fields{
		setStr = fmt.Sprintf("%s%s=$%d,",setStr,f,i+1)
	}

	// 条件语句
	condStr := setStr[:len(setStr)-1]
	if dbSet.condCount >0{

		is:=make([]interface{},dbSet.condCount)
		for i:=0;i<dbSet.condCount;i++{
			is[i] = i+1 + setCount
		}
		rs := fmt.Sprintf(dbSet.condStr,is...)
		condStr += fmt.Sprintf(" WHERE %s ",rs)
	}
	condStr +=dbSet.order
	if dbSet.limit_pos>0 && dbSet.limit_len>0{
		return fmt.Sprintf(" %s limit $%d offset $%d", condStr,setCount+dbSet.condCount+1,setCount+dbSet.condCount+2)
	}else if dbSet.limit_pos>0{
		return fmt.Sprintf(" %s offset $%d", condStr,setCount+dbSet.condCount+1)
	}else if dbSet.limit_len>0{
		return fmt.Sprintf(" %s limit $%d", condStr,setCount+dbSet.condCount+1)
	}
	return condStr
}

/**
获取参数
 */
func (dbSet *DbSetCondition) GetSetCondParams() []interface{}{
	rs:= dbSet.newValues
	rs= append(rs,dbSet.args...)
	if dbSet.limit_len>0{
		rs=append(rs,dbSet.limit_len)
	}
	if dbSet.limit_pos>0{
		rs=append(rs,dbSet.limit_pos)
	}
	return rs
}
