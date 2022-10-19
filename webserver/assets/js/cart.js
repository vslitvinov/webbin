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