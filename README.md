# ekakao

### Example
> ```go
> package main
> 
> import "github.com/app6460/ekakao"
> 
> func main() {
>   client := ekakao.New("email", "password")
>   client.Login()
>   client.SendEmoji("shaky-animals-2", 1)
> }
> ```