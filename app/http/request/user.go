package request

// ChangePasswordRequest ...
type ChangePasswordRequest struct {
	OldPassword string `form:"old_password" json:"old_password" binding:"required" label:"old_password"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required,min=6,max=16" label:"new_password"`
}

// ChangeMobileRequest ...
type ChangeMobileRequest struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,len=11,phone" label:"mobile"`
	Password string `form:"old_password" json:"old_password" binding:"required" label:"old_password"`
	SmsCode  string `form:"sms_code" json:"sms_code" binding:"required,len=6,numeric" label:"sms_code"`
}
