import tasksComponent from '../components/tasks_list.js'

export default({
    template: 
        `<q-page>
            <tasks/>
            <q-page-sticky position="bottom-right" :offset="[18, 18]">
                <q-btn fab icon="add" color="positive" />
            </q-page-sticky>
        </q-page>`,
    modules: {
        tasks : tasksComponent
    },
    data: () => ({
        tasks: []
    }),
    mounted () {
        console.log("tasks mounted")
    },
    methods: {
        
    }
});