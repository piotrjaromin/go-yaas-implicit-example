$(document).ready(() => {

    const path = window.location.pathname;
    if (path.indexOf('/callback') !== -1) {
        const urlFragment = window.location.hash.substr(1);
        let text = "";

        urlFragment.split("&").forEach(element => {
            console.log(element);
            const splited = element.split("=");
            text += splited[0] + " : " + decodeURIComponent(splited[1]) + "<br>"
        });
        $("#tokenField").html(text);
    }
});

function onSubmitForm() {
    //set authorize endpoint url when form is submitted
    document.tokenForm.action = $("#authorizeUrl option:selected").text();
    return true;
}