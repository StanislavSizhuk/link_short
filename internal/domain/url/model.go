package url

type URL struct {
	full  string
	short string
}

func (u *URL) Full() string {
	return u.full
}

func (u *URL) Short() string {
	return u.short
}
func (u *URL) SetShort(short string) {
	u.short = short
}
