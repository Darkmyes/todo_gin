export default({
    template: 
        `<q-card>
            Hola
        </q-card>`,
    data: () => ({
        tasks: []
    }),
    mounted () {
    },
    methods: {
        listTasks () {
            this.$store.dispatch('tasks/list', this.$http )
                .then(res => {
                    console.log(res)
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