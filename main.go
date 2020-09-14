package main

func main() {
	engin := InitRouter()
	engin.Logger.Fatal(engin.Start(":8080"))
}
