package calculator

func Multiply(a float64, b float64) float64 {
	print("Multiply: ")
	return (a * b) + offset
}

// warn 同一パッケージ内で同じシグネチャを持つ関数の重複は禁止
func multiply(a float64, b float64) float64 {
	print("multiply: ")
	return (a * b) + offset
}

// memo: 頭小文字の変数、メソッドは同じパッケージ内で利用できる。(offset, multiply)
// memo: 頭大文字の変数、メソッドは外部パッケージでも利用できる。(Offset, Multiply)
