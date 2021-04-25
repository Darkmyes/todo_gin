import tasksComponent from '../components/tasks_list.js'
import tasksRegisterComponent from '../components/tasks_register.js'

export default({
    template: 
        `<q-page>
            <tasks/>
            <q-page-sticky position="bottom-right" :offset="[18, 18]">
                <q-btn @click="newTaskModal = true" fab icon="add" color="positive" />
            </q-page-sticky>
            <q-dialog v-model="newTaskModal">
                <register-task v-on:register-done="registerDone"/>
                fsadas
            <q-dialog>
        </q-page>`,
    components: {
        'tasks' : tasksComponent,
        'register-task' : tasksRegisterComponent
    },
    data: () => ({
        tasks: [],
        newTaskModal: false
    }),
    mounted () {        
    },
    methods: {
        registerDone () {
            this.newTaskModal = false
        }
    }
});