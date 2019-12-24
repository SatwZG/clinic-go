package api

type LoginRequest struct {
	Username string  `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
}


type DoctorSearchRequest struct {
	Name string 	  `json:"name"`
	Department string `json:"department"`
	Sex string 		  `json:"sex"`
	Age []int 		  `json:"age"`
	Page int 		  `json:"page"`
}

type Doctor struct {
	ID int              `json:"id"`
	Name string 		`json:"name"`
	Sex string 			`json:"sex"`
	Age int 			`json:"age"`
	Department string   `json:"department"`
	Avatar string       `json:"avatar"`
	Introduction string `json:"introduction"`
}

type DoctorSearchResponse struct {
	NowPage int	 	 `json:"nowPage"`
	TotalPage int	 `json:"totalPage"`
	Doctors []Doctor `json:"doctors"`
}