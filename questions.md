### *1. Какой самый эффективный способ конкатенации строк?*
Самый эффективный способ это использовать strings.Builder. Он предоставляет методы для работы с строкой без необходимости создания временных строк. Это приводит к повышению эффективности и избегает излишнего копирования.
### *2. Что такое интерфейсы, как они применяются в Go?*
Интерфейс в Go — это абстракция, которая определяет набор методов, но не их реализацию. Тип в Go считается соответствующим интерфейсу, если он реализует все методы этого интерфейса. Если говорить простыми словами, то интерфейс — это некий контракт, согласно которому компоненты системы ожидают друг от друга определенного поведения, например в части обмена информацией.
В Go типы реализуют интерфейс неявно. Это означает, что вам не нужно явно указывать, что тип реализует интерфейс. Тип соответствует интерфейсу, если он имеет все методы, которые интерфейс требует.
### *3. Чем отличаются RWMutex от Mutex?*
Mutex обеспечивает взаимное исключение, позволяя только одной горутине за раз обращаться к разделяемому ресурсу. Горутины должны захватывать мьютекс перед доступом к ресурсу и освобождать его после использования. 
RWMutex это расширенная версия мьютекса, которая разделяет доступ на операции чтения и записи. Она позволяет множеству горутин безопасно читать данные, пока никто не пишет.
### *4. Чем отличаются буферизированные и не буферизированные каналы?*
Буферезированный и не буферезированный канал используется для связи между горутинами, но они отличаются тем, что в не буферезированном канале нет предварительного определенного места для хранения значение. Когда данные посылаются в такой канал, горутина блокируется до тех пор, пока другая горутина не прочитает данные из этого канала, и аналогично, горутина которая читает данные блокируется до тех пор, пока другая горутина не отправит туда данные.
Что касается буферезированного канала, горутина не будет блокироваться пока в таком канале есть свободное место для записи, если свободного места нет, то блокировка пропадет после появления хотя бы 1 свободного места в канале. Аналогично, при чтении данных горутина будет заблокирована если в канале нет хотя бы 1 значения.
### *5. Какой размер у структуры struct{}{}?*
0 байт
### *6. Есть ли в Go перегрузка методов или операторов?*
Нет 
### *7. В какой последовательности будут выведены элементы map[int]int?*
Пример:
m[0]=1
m[1]=124
m[2]=281*
Элементы выведутся в рандомном порядкея 
### *8. В чем разница make и new?*
make используется для аллокации и инициализации слайсов, мап и каналов и возвращает инициализированный экземпляр типа.
new используется для любых структур и возвращает указатель на тип.
### *9. Сколько существует способов задать переменную типа slice или map?*
slice:
- с помощью функции make (s := make([]Тип, длина, вместимость))
- с помощью литерала (s := []Тип{значение1, значение2, значение3})
- с помощью другого слайса (s := другойСлайс[начало:конец])
- с помощью создание пустого слайса (var s []int)
map:
- с помощью функции make (m := make(map[ТипКлюча]ТипЗначения))
- с помощью литерала (m := map[ТипКлюча]ТипЗначения{ключ1: значение1, ключ2: значение2})
- с помощью создания пустой карты (var m map[string]int)
### *10. Что выведет данная программа и почему?*
```go
func update(p *int) {
b := 2
p = &b
}
func main() {
var (
a = 1
p = &a
)
fmt.Println(*p)
update(p)
fmt.Println(*p)
}
```
Программа выведет
1
1
Потому что в первом вызове fmt.Println(*p) выводится указатель на переменную (a), то есть 1.
Далее вызывается функция (update) в которую передается переменая p (=1), но изменение адреса, которое происходит в функции не влияет на исходное значение указателя, т.к. передается только копия указателя, следовательно второй вызов fmt.Println(*p) выведет тоже 1.
### *11. Что выведет данная программа и почему?*
```go
func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg sync.WaitGroup, i int) {
fmt.Println(i)
wg.Done()
}(wg, i)
}
wg.Wait()
fmt.Println("exit")
}
```
В данном коде ошибка, которая заключается передачи WaitGroup, здесь она передается не по ссылке, а по значени, т.е. в горутину добавляют копию WaitGroup, и когда вызывается wg.Wait() ожидает завершение горутин, он никогда этого не дождется, так как ожидание выполняется на оригинале, а не на копии. Программа по итогу завершится ошибкой с deadlock.
### *12. Что выведет данная программа и почему?*
```go
func main() {
n := 0
if true {
n := 1
n++
}
fmt.Println(n)
}
```
Программа выведет 0, т.к. изначальное значение n = 0, а внутри условного блока создается локальная переменная с таким же названием, но действия с этой переменной никак не повлияют на переменную внутри функции main.
### 13. *Что выведет данная программа и почему?*
```go
func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b)
}
func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}
```
Программа выведет [100 2 3 4 5] и вот почему:
Как видно функция someAction принимает некий слайс и какое-то значение, внутри функции выполняется замена элемента с индексом 0 на значение 100, затем выполняется попытка добавить новое значение в срез, и вывод этого среза. Но у нас не изначальное capacity не указано, поэтому оно =5 и при попытке добавить новое значение, оно не вмещается в рамки вместимости среза, следовательно в функции создается новый срез, который никак не влияет на исходный.
### 14. *Что выведет данная программа и почему?*
```go
func main() {
slice := []string{"a", "a"}
func(slice []string) {
slice = append(slice, "a")
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}
```
Данная программа выведет [b b a][a a], потому что в начальном срезе только 2 значения, и вместимость не указано (следовательно capacity=2), потом вызывается анонимная функция, которая добавляет в конец слайса новое значение "а", и меняет первые 2 элемента на "b", а далее в функции вызывается fmt.Println, что выведет [b b a], после выхода из функции вызывается еще раз fmt.Println, что выведет [a a], т.е. исходный слайс, т.к. внутри функции из-за того что вместимость слайса была ограничена 2 элементами, была попытка добавить 3-й элемент, из-за чего создалась копия слайса и все изменения уже происходили с копией, а не с исходным слайсом.
