package infrastructure

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf(".envが読み込み出来ませんでした: %v", err)
	}
}
