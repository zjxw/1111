package controller

import "fmt"

//View 命令视图
func View() {
	fmt.Println(`
	*********************************
	|	add	添加用户	|
	|	del	删除用户	|
	|	modify	修改用户	|
	|	query	查找用户	|
	|	print	打印用户	|	
	|	exit	退出		|
	|	help	帮助		|
	********************************* 
	`)
}
