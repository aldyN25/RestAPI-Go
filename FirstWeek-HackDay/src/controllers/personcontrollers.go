package personcontrollers

// import (
// 	"html/template"
// 	"net/http"
// 	"rest-api-crud-gin/src/entitas"
// 	"rest-api-crud-gin/src/models"
// 	"strconv"
// )

// func Index(response http.ResponseWriter, request *http.Request) {
// 	var PersonModel models.PersonModel
// 	products, _ := PersonModel.Findpersons()
// 	data := map[string]interface{}{
// 		"products": products,
// 	}

// 	tmp, _ := template.ParseFiles("views/person/index.html")
// 	tmp.Execute(response, data)
// }

// func Add(response http.ResponseWriter, request *http.Request) {
// 	tmp, _ := template.ParseFiles("views/product/index.html")
// 	tmp.Execute(response, nil)
// }

// func ProcessAdd(response http.ResponseWriter, request *http.Request) {
// 	request.ParseForm()
// 	var product entitas.Person
// 	product.Username = request.Form.Get("username")
// 	product.FirstName = request.Form.Get("first_name")
// 	product.LastName = request.Form.Get("last_name")
// 	product.Password = request.Form.Get("password")
// 	var PersonModel models.PersonModel
// 	PersonModel.Createperson(&person)
// 	http.Redirect(response, request, "/product", http.StatusSeeOther)
// }

// func Delete(response http.ResponseWriter, request *http.Request) {
// 	query := request.URL.Query()
// 	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
// 	var PersonModel models.PersonModel
// 	PersonModel.Deleteperson(id)
// 	http.Redirect(response, request, "/product", http.StatusSeeOther)
// }

// func Edit(response http.ResponseWriter, request *http.Request) {
// 	query := request.URL.Query()
// 	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
// 	var PersonModel models.PersonModel
// 	product, _ := PersonModel.Findperson(id)
// 	data := map[string]interface{}{
// 		"product": product,
// 	}
// 	tmp, _ := template.ParseFiles("view/product/view.html")
// 	tmp.Execute(response, data)
// }

// func Update(response http.ResponseWriter, request *http.Request) {
// 	request.ParseForm()
// 	var product entitas.Person
// 	product.ID, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
// 	product.Username = request.Form.Get("username")
// 	product.FirstName = request.Form.Get("first_name")
// 	product.LastName = request.Form.Get("last_name")
// 	product.Password = request.Form.Get("password")
// 	var PersonModel models.PersonModel
// 	PersonModel.Updateperson(product)
// 	http.Redirect(response, request, "/product", http.StatusSeeOther)
// }
