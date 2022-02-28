package productcontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"training/src/entitas"
	"training/src/models"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	data := map[string]interface{}{
		"products": products,
	}

	tmp, _ := template.ParseFiles("views/product/index.html")
	tmp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/product/add.html")
	tmp.Execute(response, nil)
}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entitas.Product
	product.Nama = request.Form.Get("nama")
	product.Harga, _ = strconv.ParseFloat(request.Form.Get("harga"), 64)
	product.Jumlah, _ = strconv.ParseInt(request.Form.Get("jumlah"), 10, 64)
	product.Deskripsi = request.Form.Get("deskripsi")
	var productModel models.ProductModel
	productModel.Create(&product)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var productModel models.ProductModel
	productModel.Delete(id)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}

func Edit(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var productModel models.ProductModel
	product, _ := productModel.Find(id)
	data := map[string]interface{}{
		"product": product,
	}
	fmt.Println(product)
	// fmt.Println(product)
	tmp, _ := template.ParseFiles("views/product/edit.html")
	tmp.Execute(response, data)
}

func Update(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entitas.Product
	query := request.URL.Query()
	product.Id, _ = strconv.ParseInt(query.Get("id"), 10, 64)
	product.Nama = request.Form.Get("nama")
	product.Harga, _ = strconv.ParseFloat(request.Form.Get("harga"), 64)
	product.Jumlah, _ = strconv.ParseInt(request.Form.Get("jumlah"), 10, 64)
	product.Deskripsi = request.Form.Get("deskripsi")
	fmt.Println(product)
	var productModel models.ProductModel
	productModel.Update(product)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}
