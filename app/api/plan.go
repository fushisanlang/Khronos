/*
@Time : 2021/12/7 5:44 下午
@Author : fushisanlang
@File : plan
@Software: GoLand
*/
package api

import (
	"Khronos/app/model"
	"Khronos/app/service"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

func AddPlan(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		var NewPlan *model.PlanInfo
		r.Parse(&NewPlan)

		if NewPlan.PlanName == "" || NewPlan.Measure == "" || NewPlan.PlanAll == 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else if NewPlan.PlanAll <= 0 || NewPlan.Start < 1640966400 || NewPlan.Stop <= 1640966400 || NewPlan.PlanDone != 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息异常", "Code": 405,
			})
		} else {
			NewPlan.UserId = uid
			message, code := service.AddPlanByUserId(NewPlan)
			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}

}
func GetPlan(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		PlanList := service.GetAllPlanByUserId(uid)

		r.Response.WriteJson(PlanList)

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}
}

func DelPlan(r *ghttp.Request) {
	//验证session，并获取用户id
	planName := r.Get("planname")
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {

		if planName == "" {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else {
			message, code := service.DelPlanByPlanNameAndUserId(planName, uid)

			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}
}

func DonePlan(r *ghttp.Request) {
	//验证session，并获取用户id
	planName := r.Get("planname")
	planDone := r.Get("plandone")
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		if planName == "" || planDone == "" {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else {

			message, code := service.DonePlanByPlanNameAndUserId(planName, planDone, uid)

			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}
}

func ChangePlan(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		var NewPlan *model.PlanInfo
		r.Parse(&NewPlan)

		if NewPlan.PlanName == "" || NewPlan.Measure == "" || NewPlan.PlanAll == 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else if NewPlan.PlanAll <= 0 || NewPlan.Start >= NewPlan.Stop || NewPlan.PlanDone != 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息异常", "Code": 405,
			})
		} else {
			NewPlan.UserId = uid
			message, code := service.ChangePlanByUserId(NewPlan)
			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}

}
