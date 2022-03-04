package main

import ("fmt"; "net/http" ;"html/template")

type User struct {
  name string
  age uint16
  money int16
  grades float64
}
//func Wpage (page http.ResponseWriter, r *http.Request){

    //tmpl, _ := template.ParseFiles("templates/page.html")
    //tmpl.Execute(page, bob)
    //go run main.go  }



func home(page http.ResponseWriter, r *http.Request){
    //  fmt.Fprintf(page, `<h1>About Page</h1>
                          //<b>Main page</b>`)
      bob := User{ "Bob", 25, -50, 4.5}
      //ob.setNewName("Alex")
    //fmt.Fprint( page, " user name is " + bob.name)
    tmpl, _ :=template.ParseFiles("templates/page.html")
    tmpl.Execute(page, bob)
    }

    func news(page http.ResponseWriter, r *http.Request){
bob := User{ "Bob", 25, -50, 4.5}
        t, _ :=template.ParseFiles("templates/news.html")

        //ob.setNewName("Alex")
        //fmt.Fprint( page, " user name is " + bob.name)

        t.Execute(page, bob)

  }
func reg(page http.ResponseWriter, r *http.Request){
            //  fmt.Fprintf(page, `<h1>About Page</h1>
                                  //<b>Main page</b>`)
bob := User{ "Bob", 25, -50, 4.5}
              //ob.setNewName("Alex")
            //fmt.Fprint( page, " user name is " + bob.name)
tmp, _ :=template.ParseFiles("templates/reg.html")
tmp.Execute(page, bob)
            }
  func info(page http.ResponseWriter, r *http.Request){
                        //  fmt.Fprintf(page, `<h1>About Page</h1>
                                              //<b>Main page</b>`)
        bob := User{ "Bob", 25, -50, 4.5}
                          //ob.setNewName("Alex")
                        //fmt.Fprint( page, " user name is " + bob.name)
      tm, _ :=template.ParseFiles("templates/product.html")
      tm.Execute(page, bob)
                        }

func(u User) getInfo() string {

  return fmt.Sprintf("user name is %s/ He is %d and he has monet equal %d", u.name, u.age, u.money)
}

func( u User) setNewName(newName string){
  u.name= newName
}

func HandleRequest(){
http.Handle("/static/",http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

http.HandleFunc("/about",home)
http.HandleFunc("/news/",news)
http.HandleFunc("/reg/",reg)
http.HandleFunc("/info/",info)
http.ListenAndServe("localhost:8080", nil)
}


func main ()  {

//bob := User{ "Bob", 25, -50, 4.5}
HandleRequest()

}
