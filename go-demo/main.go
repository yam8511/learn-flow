package main

// 引入套件
import (
	"demo/demo"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 宣告主程式
func main() {
	// 宣告變數
	var name string
	var age int = 25
	var money float64 = 66666
	var isStudent bool = true

	// 抓取 stdin
	fmt.Println("stdin: input your name.")
	stdin, err := fmt.Scanf("%s", &name)
	fmt.Println("stdin: ", name, stdin, err)

	// 印出到 stdout
	fmt.Printf("stdout: My Name is %s, %d years old.\n", name, age)

	// 印出到 stderr
	log.Printf("stderr: I pick $%.2f.\n", money)

	// 判斷式
	if isStudent {
		fmt.Println("if: I'm a life students.")
	} else {
		fmt.Println("if: I'm a life person.")
	}

	// 宣告陣列
	nums := [...]int{1, 2, 3, 4, 5}
	fmt.Println("array:", nums)

	// 宣告切片
	foods := []string{"meat", "egg", "fish", "cake", "milk"}

	// 複製陣列
	copyFoods := make([]string, len(foods))
	copy(copyFoods, foods)

	// append元素到陣列
	foods = append(foods, "chocolate")

	// 陣列長度
	fmt.Println("slice: len:", len(foods))
	fmt.Println("slice: copy:", copyFoods)

	// 從陣列找出元素的index
	index := -1
	for i, v := range foods {
		if v == "fish" {
			index = i
			break
		}
	}
	fmt.Println("slice: index of fish:", index)

	// 從陣列remove元素
	foods = append(foods[:index], foods[index+1:]...)

	fmt.Println("slice: remove fish:", foods)

	// 陣列 join
	arr2str := strings.Join(foods, ", ")
	fmt.Println("slice: join:", arr2str)

	// 字串 split
	fmt.Println("slice: split:", strings.Split(arr2str, ", "))

	// 宣告字典
	dicts := map[string]interface{}{
		"num":  1,
		"name": "a",
		"ok":   true,
	}
	fmt.Println("dicts:", dicts)

	// 確認key存不存在
	value, ok := dicts["a"]
	fmt.Println("dicts: check key: a ->", value, ok)
	value, ok = dicts["name"]
	fmt.Println("dicts: check key: name ->", value, ok)

	// 迴圈 - for
	for i := 0; i < len(foods); i++ { // 陣列長度
		fmt.Println("for:", i)
	}

	for i, food := range foods {
		fmt.Println("for range array:", i, food)
	}

	for k, v := range dicts {
		fmt.Println("for range dicts:", k, v)
	}

	// 迴圈 - while
	for {
		break
	}

	// 引入自定義模組
	fmt.Println("call self package:", demo.Sum2Num(1, 2))

	// 數字轉字串
	num2Str := strconv.FormatInt(123, 10)
	fmt.Println("str to num:", num2Str)
	fmt.Printf("str to num: type of 123 is %T\n", num2Str)

	// 字串轉數字
	str2Num, err := strconv.ParseInt("123", 10, 64)
	fmt.Println("num to str:", str2Num, err)
	fmt.Printf("num to str: type of 123 is %T\n", str2Num)

	// 字串替換
	fmt.Println("str replace:", strings.ReplaceAll("hello world", "world", "programming"))

	// 字串去空白
	fmt.Println("str trim space:", strings.TrimSpace(" hello world "))

	// 字串包含
	fmt.Println("str contains:", strings.Contains(" hello world ", "hello"))

	// 隨機數字
	fmt.Println("rand: [0,10):", rand.Intn(10))

	// 環境變數
	_ = os.Environ()
	fmt.Println("env:", os.Getenv("ENV"))

	// 指令參數
	fmt.Println("args:", os.Args)

	// 現在時間
	now := time.Now()
	fmt.Println("time: now:", now)
	fmt.Println("time: timestamp:", now.Unix())

	// 時間格式化
	fmt.Println("time: format: [YEAR]-[MONTH]-[DATE]T[HOUR]:[MINUTE]:[SECOND]Z :", now.Format(time.RFC3339))
	fmt.Println("time: format: [YEAR]-[MONTH]-[DATE] [HOUR]:[MINUTE]:[SECOND]  :", now.Format("2006-01-02 15:04:05"))

	// JSON編碼
	data, err := json.Marshal(dicts)
	if err != nil {
		panic(err) // 噴出錯誤
	}
	fmt.Println("json: encode:", string(data))

	// JSON解碼
	dicts2 := map[string]interface{}{}
	err = json.Unmarshal(data, &dicts2)
	if err != nil {
		panic(err) // 噴出錯誤
	}
	fmt.Println("json: decode:", dicts2)

	// 開啟/建立檔案
	f, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 讀取檔案資訊
	stat, err := f.Stat()
	fmt.Println("file: open:", stat.IsDir(), stat.Size(), err)

	// 寫入檔案
	n, err := f.WriteString("demo")
	if err != nil {
		panic(err)
	}
	fmt.Println("file: write:", n, err)

	// 設定目前檔案的cursor
	ret, err := f.Seek(0, 0)
	fmt.Println("file: seek:", ret, err)

	// 讀取檔案
	line := make([]byte, 1000)
	for {
		n, err := f.Read(line)
		if err != nil {
			fmt.Println("file: read:", n, err)
			break
		}

		fmt.Println("file: read:", n, err)
		fmt.Println("file: read:", string(line[:n]))
	}

	err = os.Remove("demo.txt")
	if err != nil {
		panic(err)
	}

	// 非同步: 併發
	ch := make(chan float64)
	go func() {
		fmt.Println("async: goroutine:")
		money += 100
		ch <- money
		fmt.Println("channel: send:", money)
	}()
	money = <-ch
	fmt.Println("channel: receive:", money)
}

// 宣告自定義型態
type Sex string

// 宣告Enum
const (
	SexBoy  Sex = "boy"
	SexGirl Sex = "girl"
)

// 宣告結構(物件)
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Boy  Sex    `json:"boy"`
}

// 宣告結構的函式
func (p *Person) Hello() string {
	return "hello, my name is " + p.Name + "."
}

// interface (契約)
type Animal interface {
	Name() string
	AsPet() bool
	Feed(string)
}
