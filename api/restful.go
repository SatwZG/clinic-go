package api

type LoginRequest struct {
	Username string  `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
}

type Admin struct {
	ID int              `json:"id"`
	Username string 	`json:"username"`
	Name string 		`json:"name"`
}

type Doctor struct {
	ID int              `json:"id"`
	Username string 	`json:"username"`
	Name string 		`json:"name"`
	Sex string 			`json:"sex"`
	Age int 			`json:"age"`
	Department string   `json:"department"`
	Avatar string       `json:"avatar"`
	Introduction string `json:"introduction"`
}

type GetOpTypeRequest struct {
}

type GetOpTypeResponse struct {
	Doctor Doctor `json:"doctor"`
	Admin Admin   `json:"admin"`
	Type int	  `json:"type"`
}


type DoctorSearchRequest struct {
	Name string 	  `json:"name"`
	Department string `json:"department"`
	Sex string 		  `json:"sex"`
	Age []int 		  `json:"age"`
	Page int 		  `json:"page"`
}

type DoctorSearchResponse struct {
	NowPage int	 	 `json:"nowPage"`
	TotalPage int	 `json:"totalPage"`
	Doctors []Doctor `json:"doctors"`
}