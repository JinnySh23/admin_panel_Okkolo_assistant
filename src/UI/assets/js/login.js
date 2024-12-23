$(function () {

    if(location.protocol !== "https:") {
        location.protocol = "https:";
    }

    $("#box-error").hide();

    // Click on the login button
    $("#button-login").on("click", function () {
 
        // We collect information from the email and password fields
        let login = $("#input-login").val();
        let password = $("#input-password").val();

        if(login == "" || password == "") {
            printMessage("error","Все поля должны быть заполнены!");
            return false;
        };

        let credentials = {
            "login": login,
            "password": password,
        };

        loginRequest = ajax_JSON(CONFIG_APP_URL_BASE + "login", "POST", credentials, {});
        handler_sendLoginRequest(loginRequest);
        return false;
    });

    $("#box-message-close").on("click", function () {
        $("#box-error").hide();
        return false;
    });
    
});



// -----------------------------------
// 
//          AJAX HANDLERS
// 
// -----------------------------------

function handler_sendLoginRequest(request) {
    request.always(function () {
        switch (request.status) {
            // Success
            case 200:
                $("#box-error").hide();

                window.location.replace("/lord-panel");
                break;

            case 404:
                printMessage("error","Неверный логин или пароль! Код ошибки: " + request.responseJSON.status.error_id);
                console_RequestError("Invalid auth!", request);
                break;

            case 401:
                printMessage("error","Неверный логин или пароль! Код ошибки: " + request.responseJSON.status.error_id);
                console_RequestError("Invalid auth!", request);
                break;

            case 400:
                switch(request.responseJSON.status.code){

                    case 3:
                        printMessage("error","Неверные данные в запросе. Код ошибки: " + request.responseJSON.status.error_id);
                        console_RequestError("Invalid JSON!",request);

                    case 8:
                        printMessage("error","Доступ вашей организации в личный кабинет запрещён. Код ошибки: " + request.responseJSON.status.error_id);
                        console_RequestError("Permission denied!",request);
                        break;

                    default:
                        printMessage("error","Неизвестная ошибка! Код ошибки: " + request.responseJSON.status.error_id);
                        console_RequestError("Error!", request);
                        break;
                }
                break;

            default:
                printMessage("error","Неизвестная ошибка! Код ошибки: " + request.responseJSON.status.error_id);
                console_RequestError("Error!", request);
                break;
        }
    });
}
