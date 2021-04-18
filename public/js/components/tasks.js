export default({
    template: 
        `<q-list bordered>
            <q-item clickable v-ripple v-for="(task, index) in tasks" :key="index">
                <q-item-section avatar>
                    <q-checkbox v-model="task.state"/>
                </q-item-section>
                <q-item-section>
                    <q-item-label> {{ task.name }} </q-item-label>
                    <q-item-label caption lines="2"> ID: {{ task.id }} {{ formatDate(task.updated_at) }} {{ formatDate(task.created_at) }} </q-item-label>
                </q-item-section>
            </q-item>
        </q-list>`,
    data: () => ({
        tasks: []
    }),
    mounted () {
        this.listTasks()
    },
    methods: {
        listTasks () {
            this.$store.dispatch('tasks/list', this.$http )
                .then(res => {
                    this.tasks = res
                }) 
                .catch(err => {
                    console.error(err)
                })
        },
        formatDate(dateString) {
            let date = new Date(dateString)
            return Quasar.utils.date.formatDate(date,"YYYY-MM-DD")
        }
    }
});