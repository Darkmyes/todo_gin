export default({
    template: 
        `<q-page class="row justify-center items-center">
            <q-card>
                <q-card-section class="row justify-between">
                    <b class="q-mr-md">Name: </b> {{ user.name }}
                </q-card-section>
                <q-card-section class="row justify-between">
                    <b class="q-mr-md">Email: </b> {{ user.email }}
                </q-card-section>
            </q-card>
        </q-page>
        `,
    data: () => ({}),
    computed: {
        user () {
            return this.$store.state.auth.user
        }
    },
    methods: {
    }
});