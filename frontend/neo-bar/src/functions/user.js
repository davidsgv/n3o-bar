function getPayload(){
    var token = localStorage.getItem("login.token");

    if(token == null){
        return null
    }

    if(token.length > 1){
        var base64Url = token.split('.')[1];
        var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
            return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        }).join(''));

        return JSON.parse(jsonPayload);
    }
    return null
}


export function getUserName(){
    var payload = getPayload();
    if(payload) return payload.usuario;
    return null
}

export function getUserRol(){
    var payload = getPayload();
    if(payload) return payload.rol;
    return null
}