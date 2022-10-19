function Request(method, url, data) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, url);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify(data));
    return xhr
}




function AddCartItem(uuid) {
    var data = {
        uuid: uuid,
    };

    req = Request("POST","/cart/additem", data)

    req.onload = () => {
        uuid = JSON.parse(req.response).uuid
        document.cookie = "cart_token="+uuid
    };

}
function DeleteItemFromCart(uuid) {
    var data = {
        uuid: uuid,
    };

    req = Request("POST","/cart/deleteitem", data)

    req.onload = () => {
        uuid = JSON.parse(req.response).uuid
        document.cookie = "cart_token="+uuid
    };

}
document.addEventListener("DOMContentLoaded", GetCartInfo);



function GetCartInfo() {
    let cartHeader = document.getElementById("cartHeader")

    console.log(cartHeader)
    let tempO = '<div class="pb-4"><a href="{Url}" class="d-flex align-items-center"><img src="assets/images/icons/{Icon}" class="avatar avatar-ex-sm me-2 style="max-height: 64px;" alt=""><div class="flex-1 text-start ms-3"><h6 class="text-dark mb-0">{Name}</h6><p class="text-muted mb-0">$ {Price} X {Count}</p></div><h6 class="text-dark mb-0">$ {TotalPrice}</h6></a></div>'
    
    Request("GET","/cart/getcart", "").onload = (e) =>{
        data = JSON.parse(e.srcElement.response)
        let tempCart = ""
        for (key in data.Cart.Items) {
            var item = data.Cart.Items[key]
            temp = tempO
            temp = temp.replace("{Count}",item.Count)
            temp = temp.replace("{Name}",item.Product.name)
            temp = temp.replace("{Price}",item.Product.price)
            temp = temp.replace("{Url}",item.Product.url)
            temp = temp.replace("{Icon}",item.Product.icon)
            temp = temp.replace("{TotalPrice}",item.Product.price*item.Count)
            tempCart = tempCart + temp
        }   

        cartHeader.innerHTML = tempCart



    }

}