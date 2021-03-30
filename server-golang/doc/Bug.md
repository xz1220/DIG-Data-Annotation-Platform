# Bug Reports
## Controller
### AdminUser
<details>
<summary><strong>删除用户时，未判断是否用户是否已经被分配给了其他任务，导致查询任务时发生错误</strong></summary>

- 涉及到的代码
```go
// controller/AdminUser.go
// DeleteUser
func DeleteUser(ctx *gin.Context) {
	/**
    根据用户ID查询到用户
     */

	// 查找是否已有任务分配给用户
	// adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(common.GetDB())
	if user.Authorities == "ROLE_ADMIN" {
        // 判断是否Admin用户被分配给了任务
        // 如果有，则删除失败
	} else if user.Authorities == "ROLE_USER" {
        // 判断是否User用户被分配给了任务
        // 如果有，则删除失败
	} else if user.Authorities == "ROLE_REVIEWER" {
        // 判断是否Reviewer用户被分配给了任务
        // 如果有，则删除失败
	}

	err = adminUserReposityInstance.DeleteUser(user.UserID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Delete User Error!!!")
		return
	}

	log.Println("Delete User Success!!!")
	util.Success(ctx, gin.H{}, "SUCCESS")

}

```

</details>

## repository

## model

