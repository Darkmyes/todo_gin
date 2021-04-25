export default ({
    namespaced: true,

    state: () => ({
        tasks: []
    }),
    getters: {

    },
    mutations: {
        setTasks (state, tasks) {
            state.tasks = tasks
        },
    },
    actions: {
        list ({commit}, axios) {   
            console.log("aqui")         
            axios.get('/api/tasks')
                .then(res => {
                    commit("setTasks", res.data)
                })
                .catch(err => {
                    console.error(err)
                })
        },
        register ({ dispatch }, {axios, task}) {
            return new Promise((resolve, reject) => {
                axios.post('/api/tasks',{
                    'name': task.name,
                    'description': task.description
                })
                    .then(res => {
                        dispatch("list", axios)
                        resolve(res.data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        setState ({ dispatch }, {axios, taskID, state}) {
            return new Promise((resolve, reject) => {
                axios.put('/api/tasks/' + taskID ,{
                    'state': state
                })
                    .then(res => {
                        dispatch("list", axios)
                        resolve(res.data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        }
    }
})