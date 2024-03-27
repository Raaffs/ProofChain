import { ChangeEvent,useEffect,useState } from "react";
import { Login,Register } from "../../../wailsjs/go/main/App";
function RegisterUser() {
    const [input, setInput] = useState({
        privateKey: "",
        username: "",
        password: ""
    });

    const handleClick = () => {
        Register(input.privateKey, input.username, input.password)
            .then(() => {
                // Registration successful, do something if needed
            })
            .catch((err) => {
                console.log(err);
                // Handle error if needed
            });
    };

    const handleChange = (event) => {
        setInput({
            ...input,
            [event.target.name]: event.target.value
        });
    };

    return (
        <div>
            <input
                name="privateKey"
                value={input.privateKey}
                type="text"
                onChange={handleChange}
            />
            <input
                name="username"
                value={input.username}
                type="text"
                onChange={handleChange}
            />
            <input
                name="password"
                value={input.password}
                type="text"
                onChange={handleChange}
            />
            <button onClick={handleClick}>Submit</button>
        </div>
    );
}

export default RegisterUser