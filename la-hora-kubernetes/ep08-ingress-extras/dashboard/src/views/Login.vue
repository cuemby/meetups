<template>
    <div class="text-center">
        <div class="form-signin">
            <div width="72"
                 height="72">
                <h1>TODO App</h1>
                <br>
                <br>
            </div>
            <h1 class="h3 mb-3 font-weight-normal">Iniciar sesión</h1>
            <label class="sr-only">Email</label>
            <input type="email" v-model="email" class="form-control" placeholder="Email" required autofocus>
            <label class="sr-only">Contraseña</label>
            <input type="password" v-model="password" class="form-control" placeholder="Contraseña" required>

            <button class="btn btn-lg btn-primary btn-block" @click="login">Ingresar</button>
            <br>
            <span class="text-danger">{{error}}</span>
            <br>
            <br>
            <router-link :to="{name: 'Register'}">No tengo cuenta, crear una.</router-link>
        </div>
    </div>
</template>

<script>
   export default {
      name: "Login",
      data() {
         return {
            email: '',
            password: '',
            error: ''
         }
      },
      methods: {
         login() {
            this.error = ''
            this.$http.post('/api/v1/accounts/login', {email: this.email, password: this.password})
               .then(res => {
                  this.$http.defaults.headers['Authorization'] = 'Bearer ' + res.data.token
                  localStorage.setItem("token", 'Bearer ' + res.data.token)
                  this.$router.push({name: 'Home'})
               })
               .catch(err => {
                  this.error = "Error al iniciar sesión, revisa tus credenciales."
               })
         }
      }
   }
</script>

<style scoped>
    .bd-placeholder-img {
        font-size: 1.125rem;
        text-anchor: middle;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
    }

    @media (min-width: 768px) {
        .bd-placeholder-img-lg {
            font-size: 3.5rem;
        }
    }

    html,
    body {
        height: 100%;
    }

    body {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-align: center;
        align-items: center;
        padding-top: 40px;
        padding-bottom: 40px;
        background-color: #f5f5f5;
    }

    .form-signin {
        width: 100%;
        max-width: 330px;
        padding: 15px;
        margin: auto;
    }

    .form-signin .checkbox {
        font-weight: 400;
    }

    .form-signin .form-control {
        position: relative;
        box-sizing: border-box;
        height: auto;
        padding: 10px;
        font-size: 16px;
    }

    .form-signin .form-control:focus {
        z-index: 2;
    }

    .form-signin input[type="email"] {
        margin-bottom: -1px;
        border-bottom-right-radius: 0;
        border-bottom-left-radius: 0;
    }

    .form-signin input[type="password"] {
        margin-bottom: 10px;
        border-top-left-radius: 0;
        border-top-right-radius: 0;
    }
</style>