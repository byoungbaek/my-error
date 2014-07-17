package my_error

func New(code int32, text string) error {
    return &MyError{code, text}
}
