use text_io::read;

// 宣告主程式
fn main() {
    // 宣告變數
    let name: String = String::from("Zuolar");
    let age: isize = 25;
    let money: f64 = 66666.0;
    let isStudent: bool = true;

    // 抓取 stdin
    println!("stdin: input your name.");
    let name: String = read!(); // 使用套件
    println!("stdin: {}", name);

    // 印出到 stdout
    println!("stdout: My Name is {}, {} years old.", name, age);

    // 印出到 stderr
    eprintln!("stderr: I pick ${}.", money);

    // 判斷式
    if isStudent {
        println!("if: I'm a life students.");
    } else {
        println!("if: I'm a life person.");
    }

    // 宣告陣列
    let nums = [1, 2, 3, 4, 5];
    println!("array: {:?}", nums);

    let mut foods = vec![
        String::from("meat"),
        String::from("egg"),
        String::from("fish"),
        String::from("cake"),
        String::from("milk"),
    ];

    // append元素到陣列
    foods.append(&mut vec![String::from("chocolate")]);

    // 陣列長度
    println!("slice: len: {}", foods.len());

    // 從陣列找出元素的index
    // index := -1
    // for i, v := range foods {
    // 	if v == "fish" {
    // 		index = i
    // 	}
    // }

    // // 從陣列remove元素
    // foods = append(foods[:index], foods[index+1:]...)

    // fmt.Println("slice:", foods)
}
