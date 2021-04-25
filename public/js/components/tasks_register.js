export default({
    template: 
        `<q-card>
            <q-card-section class="bg-primary text-white row items-center justify-center">
                <div class="text-center text-h6 text-bold"> New Task </div>
            </q-card-section>
            <q-card-section class="column">
                <q-input outlined type="text" v-model="task.name" label="Name" class="q-mb-sm"/>
                <q-input outlined type="textarea" v-model="task.description" label="Description"/>
            </q-card-section>
            <q-card-actions align="center">
                <q-btn @click="registerTask" label="Register" color="positive"/>
            </q-card-actions>
        </q-card>`,
    data: () => ({
        task: {
            name: "",
            description: ""
        }
    }),
    mounted () {
        console.log("tasks register mounted")
    },
    methods: {
        registerTask () {
            this.$store.dispatch('tasks/register', { axios: this.$http, task: this.task} )
                .then(res => {
                    console.log(res)
                    this.$emit('register-done')
                }) 
                .catch(err => {
                    console.error(err)
                })
        }
    }
});