type CreditCard struct {
  OwnerName string
  Cash      float64
  Limit     float64
}

Функция
NewCard(ownerName string, cash, limit float64)

Методы CreditCard
ShowRUB() // cash
ShowUSDT() // cash, переведенный в USDT
IsExpired() bool // смотрит, если у тебя cash > limit, тогда все ок и не expired

Еще один метод CreditCard
ShowInfo(currentHour int)
Если у чела не просрочена кредитка, то вернуть "Hello (Name), it's (Hour) o'clock"
Если же просрочена, то "(Name), жди коллекторов сука"

Для перевода в USDT объяви константу USDTtoRUB = 80 вне функций

Потом создай в main две кредитки
первая, которая не просрочена
вторая просрочена
выведи для первой ShowUSDT и ShowInfo
для второй ShowRUB и ShowInfo

Методы Show... у тебя должны через Printf выводить в терминал сообщения. 
Для долларов и рублей округление до второго знака после запятой