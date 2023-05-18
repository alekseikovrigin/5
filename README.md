## Structocaster

```go
type User struct {
	ID          int
	Name        string
	Surname     string
	Certificate Passport
}
type Passport struct {
	Serial      string
	CityOfBirth CityOfBirth
}
type CityOfBirth struct {
	Name string
}

type UserToReport struct {
	UID  int
	Name string
	G    string
	From City
}

type City struct {
	Title string `out:"Certificate.CityOfBirth.Name"`
}

func main() {

	UserFromDB := User{
		ID:      12,
		Name:    "Aleksei",
		Surname: "Kovrigin",
		Certificate: Passport{
			Serial: "DF374-23479",
			CityOfBirth: CityOfBirth{
				Name: "Kirov",
			},
		},
	}

	As := City{Title: ""}
	UserToReport := UserToReport{
		UID:  13,
		Name: "",
		From: As,
	}

	structocaster.Cast(&UserFromDB, &UserToReport)



```
