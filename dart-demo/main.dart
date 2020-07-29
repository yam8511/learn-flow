import 'dart:io';

main(List<String> args) {
  // 宣告變數
  String name = "Zuolar";
  num age = 25;
  num money = 66666.66;
  bool isStudent = true;

  // 抓取 stdin
  print("stdin: input your name.");
  name = stdin.readLineSync();
  print("stdin:  $name");

  // 印出到 stdout
  print("stdout: My Name is $name, $age years old.");

  // 印出到 stderr
  stderr.writeln("stderr: I pick ${money.toStringAsFixed(2)}.");

  // 判斷式
  if (isStudent) {
    print("if: I'm a life students.");
  } else {
    print("if: I'm a life person.");
  }

  // 宣告陣列
  const nums = [1, 2, 3, 4, 5];
  print("array: $nums");

  var foods = ["meat", "egg", "fish", "cake", "milk"];
  print("slice: $foods");
}
