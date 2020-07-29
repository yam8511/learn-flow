use text_io::read;

// 宣告主程式
fn main() {
    // 宣告變數
    let name: String;
    let age: isize = 25;
    let money: f64 = 66666.0;
    let is_student: bool = true;

    // 抓取 stdin
    println!("stdin: input your name.");
    name = read!(); // 使用套件
    println!("stdin: {}", name);

    // 印出到 stdout
    println!("stdout: My Name is {}, {} years old.", name, age);

    // 印出到 stderr
    eprintln!("stderr: I pick ${}.", money);

    // 判斷式
    if is_student {
        println!("if: I'm a life students.");
    } else {
        println!("if: I'm a life person.");
    }

    // 宣告陣列
    let nums = [1, 2, 3, 4, 5];
    println!("array: {:?}", nums);

    // 宣告切片
    let mut foods = vec![
        String::from("meat"),
        String::from("egg"),
        String::from("fish"),
        String::from("cake"),
        String::from("milk"),
    ];

    // 複製切片
    let copy_foods = foods.clone();

    // append元素到陣列
    foods.append(&mut vec![String::from("chocolate")]);

    // 陣列長度
    println!("slice: len: {}", foods.len());
    println!("slice: copy: {:?}", copy_foods);

    // 從陣列找出元素的index
    let index = foods.iter().position(|r| r == "fish").unwrap();
    println!("slice: index of fish: {}", index);

    // 從陣列remove元素
    foods.remove(index);
    println!("slice: remove fish: {:?}", foods);

    let arr2str: String = foods.join(", ");
    println!("slice: join: {}", arr2str);

    let split = arr2str.split(", ");
    println!("slice: split: {:?}", split); // 還不知道怎麼split
}
