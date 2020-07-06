<template>
    <div>
        <div v-if="user===null">
            Cargando...
        </div>
        <div v-else>
            <b-navbar toggleable="lg" type="dark" variant="primary">
                <b-navbar-brand href="#">TODO App</b-navbar-brand>

                <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

                <b-collapse id="nav-collapse" is-nav>
                    <b-navbar-nav>
                        <router-link :to="{name: 'Home'}" class="nav-link">Lista</router-link>
                        <router-link :to="{name: 'create'}" class="nav-link">Crear</router-link>
                    </b-navbar-nav>

                    <!-- Right aligned nav items -->
                    <b-navbar-nav class="ml-auto">
                        <b-nav-item-dropdown right>
                            <!-- Using 'button-content' slot -->
                            <template v-slot:button-content>
                                <em>{{user.name}}</em>
                            </template>
                            <b-dropdown-item @click="logout">Cerrar sesión</b-dropdown-item>
                        </b-nav-item-dropdown>
                    </b-navbar-nav>
                </b-collapse>
            </b-navbar>
            <br>
            <div class="container">
                <div class="row">
                    <div class="col-md-4">
                        <h3>Lista</h3>
                        <ul class="list-group" v-if="todos.length >0" v-for="row in todos">
                            <router-link :to="{name: 'show', params:{id: row.id}}" class="list-group-item">
                                {{row.name===''||row.name===null?'(Tarea sin nombre)':row.name}}
                            </router-link>
                        </ul>
                        <ul class="list-group" v-if="todos.length ===0">
                            <li class="list-group-item">
                                No hay tareas que mostrar.<br>
                                <router-link :to="{name:'create'}" class="btn-link">Crear nueva</router-link>
                            </li>
                        </ul>
                    </div>
                    <div class="col-md-8">
                        <router-view/>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

   export default {
      name: 'Home',
      data() {
         return {
            user: null,
            todos: [],
         }
      },
      methods: {
         logout() {
            this.$http.defaults.headers['Authorization'] = ''
            this.$router.push({name: 'Login'})
         },
         list() {
            this.$http.get('/api/v1/todo/').then(
               res => this.todos = res.data.todos
            )
         }
      },
      created() {
         this.$http.defaults.headers['Authorization'] = localStorage.getItem("token")
         this.$http.get('/api/v1/accounts/user/')
            .then(res => {
               this.user = res.data.user
               this.list()
            }).catch(err => {
            this.$router.push({name: 'Login'})
         })
      }
   }
</script>
<style scoped>
    .loader {
        position: fixed;
        left: 0px;
        top: 0px;
        width: 100%;
        height: 100%;
        z-index: 9999;
    }
</style>