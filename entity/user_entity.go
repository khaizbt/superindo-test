package entity

type (
	DataUserInput struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	TransferInput struct {
		Balance  int    `json:"balance" binding:"required"`
		Username string `json:"username" binding:"required"`
	}
)
