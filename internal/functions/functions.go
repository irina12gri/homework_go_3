package functions

import "fmt"

type User struct {
	ID     int
	Name   string
	Active bool
}

// 01. SafeDivide выполняет целочисленное деление.
// TODO: при b == 0 верните 0 и ошибку "division by zero".
func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 02. FindUserByID ищет пользователя по ID.
// TODO: найденный пользователь возвращается с true; отсутствие — User{} и false.
func FindUserByID(users []User, id int) (User, bool) {
	for _, u := range users {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

// 03. FindActiveUser ищет пользователя с нужным ID только среди активных.
// TODO: неактивный пользователь считается ненайденным.
func FindActiveUser(users []User, id int) (User, bool) {
	for _, u := range users {
		if u.ID == id && u.Active == true {
			return u, true
		}
	}
	return User{}, false
}

// 04. SumAll складывает произвольное количество чисел.
// TODO: вызов без аргументов должен вернуть 0.
func SumAll(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// 05. Apply вызывает переданную функцию op для a и b.
// TODO: верните результат op без изменения аргументов.
func Apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 06. ApplyIf применяет transform к value только при allowed=true.
// TODO: при false верните исходное value и не вызывайте transform.
func ApplyIf(value int, allowed bool, transform func(int) int) int {
	if allowed == true {
		return transform(value)
	}
	return value
}

// 07. NewCounter создаёт замыкание-счётчик.
// TODO: первый вызов возвращённой функции должен вернуть start+1,
// каждый следующий — ещё на единицу больше.
func NewCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}

// 08. NewAccumulator создаёт замыкание с накопленной суммой.
// TODO: каждый вызов добавляет аргумент к текущему состоянию и возвращает сумму.
func NewAccumulator(initial int) func(int) int {
	sum := initial
	return func(n int) int {
		sum += n
		return sum
	}
}

// 09. MakeMultiplier возвращает функцию умножения на factor.
// TODO: factor должен сохраняться внутри возвращённой функции.
func MakeMultiplier(factor int) func(int) int {
	return func(voz int) int {
		return factor * voz
	}
}

// 10. DeferOrder возвращает строку, показывающую порядок defer.
// TODO: тело добавляет "body", затем зарегистрированы defer "first" и "second".
// Итоговая строка должна быть "body-second-first".
func DeferOrder() (result string) {
	result += "body"
	defer func() {
		result += "-first"
	}()
	defer func() {
		result += "-second"
	}()
	return result
}

// 11. CaptureDeferArgument демонстрирует вычисление аргумента defer сразу.
// TODO: сначала value="first", затем зарегистрируйте defer с value как параметром,
// после регистрации поменяйте value на "second". Верните захваченное значение.
func CaptureDeferArgument() (result string) {
	value := "first"
	defer func(value string) { // вычислится сразу
		result = value
	}(value)
	value = "second"
	return value
}

// 12. ReadDeferredVariable демонстрирует чтение переменной deferred-замыканием.
// TODO: deferred-функция без параметров должна прочитать value после того,
// как оно изменилось с "first" на "second".
func ReadDeferredVariable() (result string) {
	value := "first"
	defer func() { // вычислится позже
		result = value
	}()
	value = "second"
	return value
}

// 13. IncrementNamedResult возвращает value, увеличенный deferred-функцией на 1.
// TODO: используйте именованное возвращаемое значение.
func IncrementNamedResult(value int) (result int) {
	defer func() {
		result++
	}()
	return value
}

// 14. RunWithCleanup выполняет action, а cleanup откладывает до выхода.
// TODO: верните две строки в порядке фактического выполнения: action, cleanup.
func RunWithCleanup(action, cleanup func() string) (events []string) {
	defer func() {
		events = append(events, cleanup())
	}()
	events = append(events, action())
	return events
}

// 15. ChooseOperation возвращает функцию для "add", "sub" или "mul".
// TODO: для неизвестного имени верните nil и false.
func ChooseOperation(name string) (func(int, int) int, bool) {
	addFunc := func(a int, b int) int {
		return a + b
	}
	subFunc := func(a int, b int) int {
		return a - b
	}
	mulFunc := func(a int, b int) int {
		return a * b
	}
	switch name {
	case "add":
		return addFunc, true
	case "sub":
		return subFunc, true
	case "mul":
		return mulFunc, true
	}
	return nil, false
}
