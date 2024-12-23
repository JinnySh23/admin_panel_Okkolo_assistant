$(function () {

    if(location.protocol !== "https:") {
        location.protocol = "https:";
    }

    // ===================================
    // 
    //                INIT
    // 
    // ===================================
    $("#box-error").hide();

    // Click handler anywhere
    $(document).click(function(event) {
        if (!$(event.target).closest('#box-error').length) {
            $('#box-error').hide();
        }
    });

    $('#box-message-close').click(function() {
        $('#box-error').hide();
    });


    // ===================================
    // 
    //                BUTTONS
    // 
    // ===================================

    // Button - Get information from dataset
    $("#button-get-dataset").on("click", function () {
        let datasetGetRequest = ajax_GET(CONFIG_APP_URL_BASE+"api/dataset", {}, {});
        handler_getRequest("get_dataset_json", datasetGetRequest);
        return false;
    });

    // Button - Add information to dataset
    $("#button-add-data-dataset").on("click", function () {

        let new_question = $("#input-question").val();
        let new_answer = $("#input-answer").val();

        // Are new_question and new_answer empty?
        if (new_question === "" && new_answer === "") {
            printMessage("error","Вопрос и ответ не указаны!");
            return false;
        } else if (new_question === "") {
            printMessage("error","Вопрос не указан!");
            return false;
        } else if (new_answer === "") {
            printMessage("error","Ответ не указан!");
            return false;
        } else {
            // Checking if the string ends with "?"
            if (!new_question.endsWith("?")) {
                printMessage("error","Вы забыли знак вопроса в конце вопроса!");
                return false;
            }
        }

        credentials = {
            "new_question": new_question,
            "new_answer": new_answer,
        };

        let datasetAddDataRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/dataset","POST", credentials, {});
        handler_postRequest("add_data_dataset", datasetAddDataRequest);
        return false;
    });

    // Button - Delete row from dataset
    $("#button-del-data-dataset").on("click", function () {

        let str_index = $("#input-str-index").val();
        if (str_index == "") {
            printMessage("error","Вы не ввели номер строки, которую хотите удалить.");
            return false;
        } else {
            if(confirm(`Вы действительно хотите удалить строку под номером ${str_index} из датасета?`)) {
                let datasetDelDataRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/dataset/?index=" + str_index,"DELETE", {}, {});
                handler_deleteRequest("del_data_dataset", datasetDelDataRequest);
            }
        }
        return false;
    });

    $("#button-logout").on("click", function (){
        let outSessionRequest = ajax_GET(CONFIG_APP_URL_BASE + "logout",{},{});
        handler_getRequest("out_session", outSessionRequest);
        return false;
    });

});

// -----------------------------------
//        Misc(other functions)
// -----------------------------------


// -----------------------------------
//      Views(data representation)
// -----------------------------------
function view_DatasetInfo(data) {
    let output = '';

    data.forEach((part,index) => {

        if (index == 0) {
            return;
        }
        output += `<div class="paragraph">
                       <span class="index">${index}.</span> 
                       <span class="bold">${part[0]}</span><br>
                       ${part[1]}
                   </div>`;
    });

    $('#box-dataset-info').html(output);
    sessionStorage.setItem("get_dataset","yes");
}


// -----------------------------------
//   Requests(response handlers)
// -----------------------------------


// -----------------------------------
//   Handlers(request handlers)
// -----------------------------------
// GET
function handler_getRequest(request_type, request) {
    request.always(function () {

        switch (request.status) {
            // Success
            case 200:
                switch (request_type) {
                    case "get_dataset_json":
                        // console.log(request.responseJSON.data);
                        view_DatasetInfo(request.responseJSON.data); 
                        break;

                    case "out_session":
                        window.location.replace("/login");
                        break;
                }
                break;

            case 404:
                printMessage("error","Ошибка 404");
                break;

            default:
                printMessage("error","Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}

// POST
function handler_postRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            // Success
            case 200:
                switch(request_type){
                    case "add_data_dataset":
                        printMessage("success","Ваше слово или фраза успешно добавлено!");
                        break;
                }     
                break;

            case 404:
                printMessage("error","Ошибка 404");
                break;

            default:
                printMessage("error","Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}

// DELETE
function handler_deleteRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            // Success
            case 200:
                switch(request_type){
                    case "del_data_dataset":
                        printMessage("success","Данная строка успешно удалена из датасета!");
                        let show_dataset = sessionStorage.getItem("get_dataset");
                        if (show_dataset == "yes") {
                            let datasetGetRequest = ajax_GET(CONFIG_APP_URL_BASE+"api/dataset", {}, {});
                            handler_getRequest("get_dataset_json", datasetGetRequest);
                        }
                        break;
                }     
                break;

            case 404:
                printMessage("error","Ошибка 404");
                break;

            default:
                printMessage("error","Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}