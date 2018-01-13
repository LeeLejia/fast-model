package test

import (
	"testing"
	"net/http"
	"net/url"
	".."
	"github.com/bmizerany/assert"
)

func TestCondition(t *testing.T){
	r:=&http.Request{}
	r.PostForm = url.Values{}
	r.PostForm.Add("name","stupi")
	r.PostForm.Add("id","123")
	r.PostForm.Add("valid","TrUe")
	r.PostForm.Add("pos","10")
	r.PostForm.Add("len","50")

	cond:=model.DbCondition{}
	assert.Equal(t,cond.GetCondStr(),"")
	assert.Equal(t,len(cond.GetParams()),0)

	cond.And(r,"=","s_name")
	assert.Equal(t,cond.GetCondStr(),"WHERE name = $1 ")
	assert.Equal(t,cond.GetParams()[0],"stupi")

	cond.And(r,"like","s_name")
	assert.Equal(t,cond.GetCondStr(),"WHERE name = $1 AND name like $2 ")
	assert.Equal(t,cond.GetParams()[1],"%stupi%")

	cond.And(r,">","i_id").Or(r,"!=","b_valid")
	assert.Equal(t,cond.GetCondStr(),"WHERE name = $1 AND name like $2 AND id > $3 OR valid != $4 ")
	assert.Equal(t,cond.GetParams()[3],true)

	cond.Limit(r,"pos","len").Order("Order by id desc")
	assert.Equal(t,cond.GetCondStr()," WHERE name = $1 AND name like $2 AND id > $3 OR valid != $4 Order by id desc limit $5 offset $6")
	assert.Equal(t,cond.GetParams()[5],10)

	cond.Reset().And2("like","name","cjwddz").Or2("=","name","cjwddz123").Limit2(11,2)
	assert.Equal(t,cond.GetCondStr()," WHERE name like $1 AND name = $2  limit $3 offset $4")
	assert.Equal(t,cond.GetParams()[3],2)
}
