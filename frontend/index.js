const serverURL = "http://localhost:4000/check";

const passwordInput = document.getElementById("password-input");
const submitForm = document.getElementById("submit-form");
const familyPrint = document.getElementById("family-print");
console.log(passwordInput.value);

// send to server to check for Family
const submitChecker = async (e) => {
    e.preventDefault();
    
    let inputString = passwordInput.value;
    
    if (inputString != "") {

        let checkedPassword = await axios
        .post(serverURL, inputString);
        let percentFamily = checkedPassword.data * 100;
        familyPrint.innerText = `Your password is ${percentFamily}% Family.`;
        passwordInput.value = "";
        
    } else {
        alert("please enter a valid password string");
    }
}

submitForm.addEventListener("submit", submitChecker);

