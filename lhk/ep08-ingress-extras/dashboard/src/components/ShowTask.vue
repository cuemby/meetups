<template>
    <div>
        <div class="row">
            <div class="col-md-11">
                <h3>Visualizar tarea</h3>
            </div>
            <div class="col-md-1">
                <b-button @click="remove" size="sm" variant="danger">
                    <b-icon icon="trash"></b-icon>
                </b-button>
            </div>
            <div class="form-group col-12">
                <label>Nombre:</label>
                <input type="email" v-model="task.name" class="form-control" aria-describedby="emailHelp">
            </div>
            <div class="form-group col-12">
                <label>Descripción:</label>
                <textarea class="form-control" v-model="task.description" rows="4"/>
            </div>
            <div class="form-group col-6">
                <label>Prioridad:</label>
                <select class="form-control" v-model="task.priority">
                    <option>Baja</option>
                    <option selected>Normal</option>
                    <option>Alta</option>
                </select>
            </div>
            <div class="form-group col-12">
                <b-button-group>
                    <b-button variant="success" @click="edit">Editar</b-button>
                </b-button-group>
            </div>
        </div>
    </div>
</template>

<script>
   export default {
      name: "ShowTask",
      watch: {
         '$route.params.id': function (id) {
            console.log(id)
            this.load(id)
         }
      },
      data() {
         return {
            task: {
               id: '',
               name: '',
               description: '',
               priority: '',
            }
         }
      },
      methods: {
         load(id) {
            this.$http.get(`/api/v1/todo/${id}`).then(res => {
               console.log(res)
               this.task.name = res.data.name
               this.task.id = res.data.id
               this.task.description = res.data.description
               this.task.priority = res.data.priority
            }).catch(err => console.log(err))
         },
         edit() {
            this.$http.put('/api/v1/todo/', this.task).then(res => {
               console.log(res)
               this.$parent.list()
               this.load(res.data.id)
            })
         },
         remove() {
            this.$http.delete(`/api/v1/todo/${this.task.id}`).then(res => {
               this.$parent.list()
               this.$router.push({name: 'Home'})
            }).catch(err => console.log(err))
         }
      },
      created() {
         this.load(this.$route.params.id)
      },
      mounted() {
         //this.load()
      }

   }
</script>

<style scoped>

</style>