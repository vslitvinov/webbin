document.addEventListener("DOMContentLoaded", GetCartInfo);


function Request(method, url, data) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, url);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify(data));
    return xhr
}



function GetCartInfo() {
    Request("GET", "/cart/getcart", "").onload = (e) => {
        PaintFullCart(JSON.parse(e.srcElement.response))
    }
}

function PaintFullCart(data) {
	let cartHeader = document.getElementById("cartHeader")
    // let totalHeaderCart = document.getElementById("totalHeaderCart")
    let tempO = `<tr class="shop-list"><td class="h6 text-center"><a href="javascript:void(0)" class="text-danger"><i class="uil uil-times"></i></a></td><td><div class="d-flex align-items-center"><img src="/assets/images/icons/{Icon}" class="img-fluid avatar avatar-small rounded shadow" style="height:auto;" alt=""><h6 class="mb-0 ms-3">{Name}</h6></div></td><td class="text-center">$ {Price}</td><td class="text-center qty-icons"><button onclick="this.parentNode.querySelector('input[type=number]').stepDown()" class="btn btn-icon btn-soft-primary minus">-</button><input min="0" name="quantity" value="1" type="number" class="btn btn-icon btn-soft-primary qty-btn quantity"><button onclick="this.parentNode.querySelector('input[type=number]').stepUp()" class="btn btn-icon btn-soft-primary plus">+</button></td><td class="text-end fw-bold pe-4">$ {Total}</td></tr>`
    var total = 0
    let tempCart = ""
    for (key in data.Cart.Items) {
        var item = data.Cart.Items[key]
        temp = tempO
        temp = temp.replace("{Count}", item.Count)
        temp = temp.replace("{Name}", item.Product.name)
        temp = temp.replace("{Price}", item.Product.price)
        temp = temp.replace("{Url}", item.Product.url)
        temp = temp.replace("{Icon}", item.Product.icon)
        temp = temp.replace("{TotalPrice}", item.Product.price * item.Count)
        tempCart = tempCart + temp
        total += item.Product.price * item.Count
    }

    // totalHeaderCart.innerText = total
    cartHeader.innerHTML = tempCart

}