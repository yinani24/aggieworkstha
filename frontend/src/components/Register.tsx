import { FormEvent } from "react"

function Register(){
    const HandleRegister = (event: FormEvent) => {
        event.preventDefault()
        console.log(event)
    }

    return(
        <form onSubmit={HandleRegister}>
            <label>
                Username:
                <input type="text" name="username" />
            </label>
            <label>
                Email:
                <input type="email" name="email" />
            </label>
            <label>
                Password:
                <input type="password" name="password" />
            </label>
            <input type="submit" value="Submit" />
        </form>
    )
}

export default Register