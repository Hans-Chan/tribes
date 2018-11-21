function serializeObject(form){
    var data = form.serializeArray();
    var object  = new Object();
    for (i in data) {
        object[data[i].name] = data[i].value;
    }
    return object;
}