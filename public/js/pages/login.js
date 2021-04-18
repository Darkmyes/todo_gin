export default({
    template: 
        `<q-page class="q-pa-md column items-center justify-center bg-deep-purple-8 text-white">
            <div class="text-center">
                <div class="text-h3 q-mb-md">Login</div>
                <div>
                    or Go to
                    <q-btn label="Register" to="/register" color="white" no-caps flat dense/>
                </div>
            </div>
            <div class="q-pa-md" style="display: grid; grid-template-columns: auto; gap: .5rem;"> 
                <q-input type="email" id="email" label="Email" v-model="loginForm.email" outlined dark color="white">
                    <template v-slot:prepend>
                        <q-icon name="mail" />
                    </template>
                </q-input>
                <q-input type="password" id="password" label="Password" v-model="loginForm.password" outlined dark color="white">
                    <template v-slot:prepend>
                        <q-icon name="lock" />
                    </template>
                </q-input>
            </div>

            <q-btn label="Login" @click="loginBtn" color="white" class="text-deep-purple-8"/>

            <q-dialog v-model="errorModal">
                <q-card>
                    <q-card-section class="text-deep-purple-8 text-h6">
                        Eror, Credentials not Match
                    </q-card-section>
                    <q-card-actions align="right">
                        <q-btn label="OK" v-close-popup color="deep-purple-8"/>
                    </q-card-actions>
                </q-card>
            </q-dialog>
        </q-page>`,
    data: () => ({
        loginForm: {
            name: '',
            email: '',
            password: ''
        },
        errorModal: false
    }),
    mounted () {
    },
    methods: {
        loginBtn () {
            this.$store.dispatch('auth/login', { axios: this.$http, loginData: this.loginForm })
                .then(res => {
                    this.$router.push({ path: '/home' })
                }) 
                .catch(err => {
                    this.errorModal = true
                    console.error(err)
                })
        }
    }
});