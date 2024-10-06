package main

func main() {
	ctx, err := Initialize()
	if err != nil {
		panic(err)
	}

	ctx.echo.Logger.Fatal(ctx.echo.Start(":8080"))
}
