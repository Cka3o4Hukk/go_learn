package main

import "fmt"

type UserProfile struct {
	username   string
	email      string
	postsCount int
}

func NewProfile(username, email string, postCount int) UserProfile {
	return UserProfile{
		username:   username,
		email:      email,
		postsCount: postCount,
	}
}

func (u *UserProfile) AddPost(newPostsCount int) {
	u.postsCount += newPostsCount
}

func (u UserProfile) Display() {
	fmt.Printf("Пользователь: %s, Email: %s, Постов: %d", u.username, u.email, u.postsCount)
}

func main() {
	var name string
	var email string
	var postsCount int
	var newPostsCount int

	fmt.Print("Введите имя пользователя: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Print("Ошибка с именем")
	}
	fmt.Print("Введите email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		fmt.Print("Ошибка с почтой")
	}
	fmt.Print("Введите начальное количество постов: ")
	_, err = fmt.Scan(&postsCount)
	if err != nil {
		fmt.Print("Ошибка с числом")
	}

	alex := NewProfile(name, email, postsCount)

	fmt.Print("Введите количество новых постов: ")

	_, err = fmt.Scan(&newPostsCount)
	if err != nil {
		fmt.Print("Ошибка с числом")
	}

	alex.AddPost(newPostsCount)

	fmt.Printf("Пользователь %s, Email: %s, Постов: %d", alex.username, alex.email, alex.postsCount)
}
