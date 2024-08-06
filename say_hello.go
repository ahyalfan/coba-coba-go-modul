package coba_coba_golang_modul

// jangan lupa pakaihuruf besar awlanya
// misal disini kita bikin sebuah method
func SayHello() string {
	return "hello world"
}

// Untuk membuat module baru, kita bisa menggunakan perintah : go mod init nama-module
// Go-Lang akan secara otomatis membuat file go.mod yang berisikan informasi nama module dan juga versi Go-Lang yang kita gunakan

// Go-Lang terintegrasi baik dengan Git
// Untuk merilis module, kita hanya perlu membuat Tag di Git

// nah jika ingin bisa terintegrasi bisa tambahkan tag
// git tag v1.0.0 // usahkan versinya mengugakan awalan v
// git push origin v1.0.0
