<template>
    <div>
        <div class="row">
            <div class="col-md-12">
                <h3>Nueva tarea</h3>
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
                <b-button type="primary" @click="create">Guardar</b-button>
            </div>
        </div>
    </div>
</template>

<script>
   export default {
      name: "CreateTask",
      data() {
         return {
            task: {
               name: '',
               description: '',
               priority: 'Normal',
            }
         }
      },
      methods: {
         create() {
            this.$http.post('/api/v1/todo/', this.task).then(res => {
               console.log(res)
               this.$parent.list()
               this.$router.push({name: 'Home'})
            })
         }
      }

   }
</script>

<style scoped>

</style>