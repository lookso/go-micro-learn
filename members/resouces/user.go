package resouces

type UserInfoResponse struct {
	Uid      uint32 `json:"uid"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}
