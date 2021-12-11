const password = document.getElementById("psw")
    , confirm_password = document.getElementById("psw-repeat");

function validatePassword(){
    if(password.value !== confirm_password.value) {
        confirm_password.setCustomValidity("Passwords Don't Match");
    } else {
        confirm_password.setCustomValidity('');
    }
}

password.onchange = validatePassword;
confirm_password.onkeyup = validatePassword;

const regBtn = document.getElementById("regBtn");
regBtn.addEventListener("click", function(){
    if (document.getElementById("psw").value && document.getElementById("psw-repeat").value && document.getElementById("email").value) {
        window.location.href("/login")
    } else {
        alert('All fields are required!');
    }
})