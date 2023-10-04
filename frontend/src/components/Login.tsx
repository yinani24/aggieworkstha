import React, { FormEvent } from 'react';
import { useState } from 'react';

function Login(){
    const handleSubmit = (event: FormEvent) => {
        event.preventDefault()
        console.log(event)
    }
    
    return(
        <form onSubmit={handleSubmit}>
            <label> Username: <input type="text" name="username" /> </label>
            <label> Password: <input type="password" name="password" /> </label>
            <input type="submit" value="Submit" />
        </form>
    )
}

export default Login;