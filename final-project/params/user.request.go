package params

type UserRegisterRequest struct {
	Age      uint   `valid:"required~Age is required, range(8|80)~Age has to have minimum value above of 8 to 80"`
	Email    string `valid:"required~Email is required, email~Invalid format email"`
	Password string `valid:"required~Password is required, minstringlength(6)~Password has to have minimum length of 6 characters"`
	Username string `valid:"required~Username is required"`
}

type UserLoginRequest struct {
	Email    string `valid:"required~Email is required"`
	Password string `valid:"required~Password is required"`
}

type UserUpdateRequest struct {
	Email    string `valid:"required~Email is required"`
	Username string `valid:"required~Username is required"`
}
