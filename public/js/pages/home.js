export default({
    template: 
        `<q-page>
            <div class="q-pa-md column items-center justify-center bg-primary text-white" style="min-height: 60vh;">
                <div class="text-h1 text-center text-bold q-mb-md">Simple ToDo</div>
                <q-btn label="Join Now" to="/register" color="white" class="text-primary"/>
            </div>
            <div class="q-pa-md column">
                <div class="text-h6 text-center">Project Description</div>
                <p class="text-center">
                    This is a basic project of a ToDo WebApp made with Go and VueJS.
                </p>
                <div class="text-h6">Backend</div>
                <p>
                    The backend is composed by Gin, Gin-JWT and Gorm for manage the server routes and the databases.
                </p>
                <div class="text-h6 text-right">Frontend</div>
                <p class="text-right">
                    The frontend section was made with VueJS, Quasar Framework and Axios.
                </p>
            </div>
        </q-page>`,
    data: () => ({
    })
});