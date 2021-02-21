package main

import "fmt"

type TaskInfo struct{
	task_id			int
	task_level		int
	task_name		string
	task_cmd		string
}

var g_all_tasks [][]TaskInfo
var g_low_tasks []TaskInfo
var g_mid_tasks []TaskInfo
var g_hig_tasks []TaskInfo

func init_task(){
	g_all_tasks = append(g_all_tasks, g_hig_tasks)
	g_all_tasks = append(g_all_tasks, g_mid_tasks)
	g_all_tasks = append(g_all_tasks, g_low_tasks)
}

func Add_task(task_add TaskInfo){
	fmt.Println(task_add.task_id)
	if task_add.task_level == 0 {
		g_low_tasks = append(g_low_tasks, task_add)
	}else if task_add.task_level == 1 {
		g_mid_tasks = append(g_mid_tasks, task_add)
	}else{
		g_hig_tasks = append(g_hig_tasks, task_add)
	}
}

func Pop_task(task_level int){

	if task_level == 0 {
		if len(g_low_tasks) >= 1{
			g_low_tasks = g_low_tasks[1: len(g_low_tasks)]
		}
	}else if task_level == 1 {
		if len(g_mid_tasks) >= 1{
			g_mid_tasks = g_mid_tasks[1: len(g_mid_tasks)]
		}
	}else{
		if len(g_hig_tasks) >= 1{
			g_hig_tasks = g_hig_tasks[1: len(g_hig_tasks)]
		}
	}
}


func Show_all_tasks(){
	fmt.Println("Show low level tasks:")
	for i := range g_low_tasks {
		fmt.Println("task level is:", g_low_tasks[i].task_id)
	}

	fmt.Println("Show mid level tasks:")
	for i := range g_mid_tasks {
		fmt.Println("task level is:", g_low_tasks[i].task_id)
	}

	fmt.Println("Show high level tasks:")
	for i := range g_hig_tasks {
		fmt.Println("task level is:", g_low_tasks[i].task_id)
	}
}



func main(){

	var task_1 TaskInfo
	task_1.task_id = 1
	task_1.task_level = 0
	task_1.task_name = "test"
	task_1.task_cmd = "echo test"

	init_task()
	Add_task(task_1)
	Show_all_tasks()
	Pop_task(0)
	Show_all_tasks()

}