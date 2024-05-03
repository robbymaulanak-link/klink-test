package dto

type MemberRequestDTO struct {
	NamaDepan    string `json:"nama_depan" form:"nama_depan" validate:"required"`
	NamaBelakang string `json:"nama_belakang" form:"nama_belakang" validate:"required"`
	Birthday     string `json:"birthday" form:"birthday" validate:"required"`
	Gender       string `json:"gender" form:"gender" validate:"required"`
	JoinDate     string `json:"join_date" form:"join_date" validate:"required"`
	Level        string `json:"level" form:"level" validate:"required"`
}
