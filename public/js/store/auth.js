const tokenId = 'jwt-token'

export default ({
    namespaced: true,

    state: () => ({
        token: '',
        user: localStorage.getItem(tokenId) != null ? localStorage.getItem(tokenId) != null : ''
    }),
    getters: {

    },
    mutations: {
        setToken (state, token) {
            state.token = token
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
                        commit('setToken', token)
                        axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
                        resolve(res.data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        getUserInfo ({commit}, {axios}) {
            return new Promise((resolve, reject) => {
                axios.post('/auth/login',{
                    email: loginData.email,
                    password: loginData.password
                })
                    .then(res => {
                        commit('setToken', res.data.token)
                        resolve(res.data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        }
    }
})