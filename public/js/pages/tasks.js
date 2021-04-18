import tasksComponent from '../components/tasks.js'

export default({
    template: 
        `<q-page>
            <tasks/>
        </q-page>`,
    modules: {
        tasks : tasksComponent
    },
    data: () => ({
        tasks: []
    }),
    mounted () {
    },
    methods: {
        
    }
});