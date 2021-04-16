export default({
    template: 
        `<q-page class="q-pa-md column items-center justify-center bg-deep-purple-8 text-white">
            <div class="text-center">
                <div class="text-h3 q-mb-md">Register</div>
                <div>
                    or Go to
                    <q-btn label="Login" to="/login" color="white" no-caps flat dense/>
                </div>
            </div>
            <div class="q-pa-md" style="display: grid; grid-template-columns: auto; gap: .5rem;">
                <q-input type="text" label="Name" v-model="registerForm.name" outlined dark color="white">
                    <template v-slot:prepend>
                        <q-icon name="text_fields" />
                    </template>
                </q-input>    
                <q-input type="email" label="Email" v-model="registerForm.email" outlined dark color="white">
                    <template v-slot:prepend>
                        <q-icon name="mail" />
                    </template>
                </q-input>
                <q-input type="password" label="Password" v-model="registerForm.password" outlined dark color="white">
                    <template v-slot:prepend>
                        <q-icon name="lock" />
                    </template>
                </q-input>
            </div>

            <q-btn label="Register" @click="register" color="white" class="text-deep-purple-8"/>

            <q-dialog v-model="errorModal">
                <q-card>
                    <q-card-section class="text-deep-purple-8 text-h6">
                        Error while Register
                    </q-card-section>
                    <q-card-actions align="right">
                        <q-btn label="OK" v-close-popup color="deep-purple-8"/>
                    </q-card-actions>
                </q-card>
            </q-dialog>
            <q-dialog v-model="errorModal" persistent>
                <q-card>
                    <q-card-section class="text-deep-purple-8 text-h6">
                        You're register was sucessfull
                    </q-card-section>
                    <q-card-actions align="right">
                        <q-btn label="Go to Login" to="/login" color="deep-purple-8"/>
                    </q-card-actions>
                </q-card>
            </q-dialog>
        </q-page>`,
    data: () => ({
        registerForm: {
            name: '',
            email: '',
            password: ''
        },
        errorModal: false,
        sucessModal: false
    }),
    mounted () {
    },
    methods: {
        register () {
            // llamar a la API que registra
            this.errorModal = true;
            this.succcesModal = true;
        }
    }
});