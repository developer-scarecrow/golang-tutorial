package main

import "fmt"

// structを定義
type Person struct {
	name string
	age  int
}

// 二人の年齢を比較します。年齢が大きい方の人を返し、また年齢差も返します。
// structも値渡しです。
func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		// p1とp2の二人の年齢を比較します。
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom Person

	// 初期値を代入します。
	tom.name, tom.age = "Tom", 18

	// ２つのフィールドを明確に初期化します。
	bob := Person{age: 25, name: "Bob"}

	// structの定義の順番に従って初期化します。
	paul := Person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}

// func main() {
// 	// Pは現在Person型の変数
// 	var P Person

// 	P.name = "Astaxie"
// 	P.age = 25

// 	fmt.Printf("The person's name is %s", P.name) // Pのnameプロパティにアクセスします。
// }
