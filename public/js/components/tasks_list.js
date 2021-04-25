export default({
    template: 
        `<div>
            <q-list bordered class="q-pt-md q-mb-md">
                <q-item-label class="text-h6 q-pl-md q-pb-md">
                    Pending
                </q-item-label>
                <q-expansion-item
                    v-for="task in tasks.filter(task => !task.state)" :key="task.id"
                    expand-separator
                    icon="perm_identity"
                    :label="task.name"
                    :caption="'ID: ' + task.id"
                >
                <template v-slot:header>
                    <q-item-section avatar>
                        <q-checkbox v-model="task.state" v-on:click.native="changeState(task.id, task.state)"/>
                    </q-item-section>
                    <q-item-section>
                        <q-item-label> {{ task.name }} </q-item-label>
                        <q-item-label caption lines="2">
                            ID: {{ task.id }},
                            Created at: {{ formatDate(task.created_at) }}
                        </q-item-label>
                    </q-item-section>
                </template>
                    <q-card>
                        <q-card-section>
                            {{ task.description }}
                        </q-card-section>
                    </q-card>
                </q-expansion-item>
            </q-list>
            <q-list bordered class="q-mt-md q-pt-md q-mb-md">
                <q-item-label class="text-h6 q-pl-md q-pb-md">
                    Completed
                </q-item-label>
                <q-expansion-item
                    v-for="task in tasks.filter(task => task.state)" :key="task.id"
                    expand-separator
                    icon="perm_identity"
                    :label="task.name"
                    :caption="'ID: ' + task.id"
                >
                <template v-slot:header>
                    <q-item-section avatar>
                        <q-checkbox v-model="task.state" v-on:click.native="changeState(task.id, task.state)"/>
                    </q-item-section>
                    <q-item-section>
                        <q-item-label> {{ task.name }} </q-item-label>
                        <q-item-label caption lines="2">
                            ID: {{ task.id }},
                            Created at: {{ formatDate(task.created_at) }}
                        </q-item-label>
                    </q-item-section>
                </template>
                    <q-card>
                        <q-card-section>
                            {{ task.description }}
                        </q-card-section>
                    </q-card>
                </q-expansion-item>
            </q-list>
        </div>`,
    data: () => ({}),
    mounted () {
        this.listTasks()
    },
    computed: {
        tasks () {
            return this.$store.state.tasks.tasks
        }
    },
    methods: {
        listTasks () {
            this.$store.dispatch('tasks/list', this.$http )
        },
        changeState (taskID, state) {
            this.$store.dispatch('tasks/setState', { 
                axios: this.$http,
                taskID: taskID,
                state: state
            } )
                .then(res => {
                    console.log(res)
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