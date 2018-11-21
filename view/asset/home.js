function login_request(form){
    var form_data   = serializeObject($('#form_data'));
    url     = $(form).attr("action");
    var request = $.ajax({
        url:    url,
        method: "POST",
        data:   form_data,
    }); 
    request.done(login_response);
    return false;
}

function login_response(result) {
    location.reload();
}