import { useState, useContext, ReactNode, createContext } from "react";

const AuthCont = createContext({
    auth: false,
    user: "",
    login: (user: any) => {},
    logout: () => {}
})

function AuthProvider(children: ReactNode){
    const [auth, setAuth] = useState(false);
    const [user, setUser] = useState("");

    const login = (user: any) => {
        setAuth(true);
        setUser(user);
    }

    const logout = () => {
        setAuth(false);
        setUser("");
    }

    return(
        <AuthCont.Provider value={{auth, user, login, logout}}>
            {children}
        </AuthCont.Provider>
    )
}

export default AuthProvider;