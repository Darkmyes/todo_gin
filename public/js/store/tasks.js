export default ({
    namespaced: true,

    state: () => ({
    }),
    getters: {

    },
    mutations: {
    },
    actions: {
        list ({commit}, axios) {
            return new Promise((resolve, reject) => {
                axios.get('/api/tasks')
                    .then(res => {
                        let tasks = res.data
                        resolve(tasks)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        register ({commit}, {axios}) {
            /*return new Promise((resolve, reject) => {
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
            })*/
        }
    }
})