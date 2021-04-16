export default({
    template: 
        `<q-page>
            <q-list bordered>
                <q-item clickable v-ripple v-for="(task, index) in tasks" :key="index">
                    <q-item-section avatar>
                        <q-icon color="primary" name="bluetooth" />
                    </q-item-section>
                    <q-item-section>
                        <q-item-label> {{ task.name }} </q-item-label>
                        <q-item-label caption lines="2"> ID: {{ task.id }} {{ formatDate(task.updated_at) }} {{ formatDate(task.created_at) }} </q-item-label>
                    </q-item-section>
                </q-item>
            </q-list>
        </q-page>`,
    data: () => ({
        tasks: []
    }),
    mounted () {
        this.listTodos()
    },
    methods: {
        listTodos() {
            axios.get('/api/tasks')
                .then(response => {
                    console.log(response)
                    this.tasks = response.data
                }).catch(err => {
                    console.log(err)
                })
        },
        formatDate(dateString) {
            let date = new Date(dateString)
            return Quasar.utils.date.formatDate(date,"YYYY-MM-DD")
        }
    }
});