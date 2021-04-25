const tokenId = 'jwt-token'

export default ({
    namespaced: true,

    state: () => ({
        token: localStorage.getItem(tokenId) != null ? localStorage.getItem(tokenId) : '',
        user: null
    }),
    getters: {
        getUser (state) {
            return state.user
        },
        isLoggedIn: state => {
            if(state.token !== '' && state.token !== null){
                return true
            }
            return false
        },
    },
    mutations: {
        setToken (state, token) {
            state.token = token
        },
        setUser (state, user) {
            state.user = user
        }
    },
    actions: {
        register ({commit}, {axios, registerData}) {
            return new Promise((resolve, reject) => {
                axios.post('/auth/register',{
                    name: registerData.name,
                    email: registerData.email,
                    password: registerData.password
                })
                    .then(res => {
                        resolve(res.data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        login ({commit}, {axios, loginData}) {
            return new Promise((resolve, reject) => {
                axios.post('/auth/login',{
                    email: loginData.email,
                    password: loginData.password
                })
                    .then(res => {
                        let token = res.data.token
                        localStorage.setItem(tokenId,token)
                        commit('setToken', token)
                        axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
                        resolve(res.data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        getUserData ({commit}, axios ) {
            return new Promise((resolve, reject) => {
                let token = localStorage.getItem(tokenId) != null ? localStorage.getItem(tokenId) : ''                

                axios.get('/api/users/data',{
                    headers: {
                        'Authorization': 'Bearer ' + token
                    }
                })
                    .then(res => {
                        axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
                        commit('setUser', res.data)
                        resolve(res.data)
                    })
                    .catch(err => {
                        delete axios.defaults.headers.common['Authorization']
                        commit('setToken', null)
                        commit('setUser', null)
                        console.error(err)
                        reject(err)
                    })
            })
        }

        /* GETTERS
        isLoggedIn: state => {
            if(state.token !== '' && state.token !== null){
                return true
            }
            return false
        },
        getUsuario: state => {
            return state.arrayUser
        }
        */

        /*
            ACTIONS
            obtenerUserData: async function({commit}, credentials){
            credentials.axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('access_token');
            return new Promise((resolve, reject) =>{
                credentials.axios.post(process.env.API_URL + '/api/details')
                .then(response =>{
                    const dataUser = response.data.success;
                    //localStorage.setItem('access_token', token);
                    //console.log("obtenerUserDataAction");
                    commit('obtenerUserDataMut', dataUser);
                    //console.log(response);
                    //console.log("obtenerDataAction");
                    resolve(response);
                })
                .catch(error =>{
                    delete credentials.axios.defaults.headers.common["Authorization"];
                    localStorage.removeItem('access_token');
                    commit('destroyTokenMut');
                    
                    console.log("Error GetData Store Action");
                    console.log(error);
                    reject(error);
                });
            })
        },
        */
    }
})