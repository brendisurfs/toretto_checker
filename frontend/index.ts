import axios, {AxiosRequestConfig, AxiosResponse} from "axios";

const serverURL = "http://localhost:4000/check";

let passwordInput = <HTMLInputElement>document.getElementById("password-input");
const submitForm = document.getElementById("submit-form");
const familyPrint = document.getElementById("family-print");

// send to server to check for Family
const submitChecker = async (e: Event) => {
    e.preventDefault();
    
    if (passwordInput.value != "") {

        let checkedPassword = await axios
        .post(serverURL, passwordInput.value);
        let percentFamily = checkedPassword.data * 100;
        familyPrint.innerText = `Your password is ${percentFamily}% Family.`;
        passwordInput.value = "";
        
    } else {
        alert("please enter a valid password string");
    }
    
}

submitForm.addEventListener("submit", submitChecker);


