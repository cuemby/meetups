(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["about"],{"51e2":function(t,e,a){"use strict";var r=a("d620"),s=a.n(r);s.a},"66c0":function(t,e,a){},"73cf":function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"text-center"},[a("div",{staticClass:"form-signin"},[t._m(0),a("h1",{staticClass:"h3 mb-3 font-weight-normal"},[t._v("Regitrarme")]),a("label",{staticClass:"sr-only"},[t._v("Nombre")]),a("input",{directives:[{name:"model",rawName:"v-model",value:t.name,expression:"name"}],staticClass:"form-control",attrs:{type:"text",placeholder:"Nombre",required:"",autofocus:""},domProps:{value:t.name},on:{input:function(e){e.target.composing||(t.name=e.target.value)}}}),a("label",{staticClass:"sr-only"},[t._v("Email")]),a("input",{directives:[{name:"model",rawName:"v-model",value:t.email,expression:"email"}],staticClass:"form-control",attrs:{type:"email",placeholder:"Email",required:"",autofocus:""},domProps:{value:t.email},on:{input:function(e){e.target.composing||(t.email=e.target.value)}}}),a("label",{staticClass:"sr-only"},[t._v("Contraseña")]),a("input",{directives:[{name:"model",rawName:"v-model",value:t.password,expression:"password"}],staticClass:"form-control",attrs:{type:"password",placeholder:"Contraseña",required:""},domProps:{value:t.password},on:{input:function(e){e.target.composing||(t.password=e.target.value)}}}),a("button",{staticClass:"btn btn-lg btn-primary btn-block",attrs:{disabled:t.loading},on:{click:t.register}},[t._v("Registrame")]),a("br"),a("span",{staticClass:"text-danger"},[t._v(t._s(t.error))]),a("br"),a("br"),a("router-link",{attrs:{to:{name:"Login"}}},[t._v("Ya tengo cuenta, iniciar sesión")])],1)])},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{width:"72",height:"72"}},[a("h1",[t._v("TODO App")]),a("br"),a("br")])}],n=(a("b0c0"),{name:"Register",data:function(){return{name:"",email:"",password:"",error:"",loading:!1}},methods:{register:function(){var t=this;this.loading=!0,this.error="",this.$http.post("/api/v1/accounts/register",{name:this.name,email:this.email,password:this.password}).then((function(e){t.$http.post("/api/v1/accounts/login",{email:t.email,password:t.password}).then((function(e){t.loading=!1,t.$http.defaults.headers["Authorization"]="Bearer "+e.data.token,localStorage.setItem("token","Bearer "+e.data.token),t.$router.push({name:"Home"})})).catch((function(e){t.loading=!1}))})).catch((function(e){t.loading=!1,t.error="Error al crear la cuenta, revisa tus datos e intenta nuevamente"}))}}}),o=n,i=(a("d478"),a("2877")),l=Object(i["a"])(o,r,s,!1,null,"5823af66",null);e["default"]=l.exports},a55b:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"text-center"},[a("div",{staticClass:"form-signin"},[t._m(0),a("h1",{staticClass:"h3 mb-3 font-weight-normal"},[t._v("Iniciar sesión")]),a("label",{staticClass:"sr-only"},[t._v("Email")]),a("input",{directives:[{name:"model",rawName:"v-model",value:t.email,expression:"email"}],staticClass:"form-control",attrs:{type:"email",placeholder:"Email",required:"",autofocus:""},domProps:{value:t.email},on:{input:function(e){e.target.composing||(t.email=e.target.value)}}}),a("label",{staticClass:"sr-only"},[t._v("Contraseña")]),a("input",{directives:[{name:"model",rawName:"v-model",value:t.password,expression:"password"}],staticClass:"form-control",attrs:{type:"password",placeholder:"Contraseña",required:""},domProps:{value:t.password},on:{input:function(e){e.target.composing||(t.password=e.target.value)}}}),a("button",{staticClass:"btn btn-lg btn-primary btn-block",on:{click:t.login}},[t._v("Ingresar")]),a("br"),a("span",{staticClass:"text-danger"},[t._v(t._s(t.error))]),a("br"),a("br"),a("router-link",{attrs:{to:{name:"Register"}}},[t._v("No tengo cuenta, crear una.")])],1)])},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{width:"72",height:"72"}},[a("h1",[t._v("TODO App")]),a("br"),a("br")])}],n={name:"Login",data:function(){return{email:"",password:"",error:""}},methods:{login:function(){var t=this;this.error="",this.$http.post("/api/v1/accounts/login",{email:this.email,password:this.password}).then((function(e){t.$http.defaults.headers["Authorization"]="Bearer "+e.data.token,localStorage.setItem("token","Bearer "+e.data.token),t.$router.push({name:"Home"})})).catch((function(e){t.error="Error al iniciar sesión, revisa tus credenciales."}))}}},o=n,i=(a("51e2"),a("2877")),l=Object(i["a"])(o,r,s,!1,null,"3674bf60",null);e["default"]=l.exports},b0c0:function(t,e,a){var r=a("83ab"),s=a("9bf2").f,n=Function.prototype,o=n.toString,i=/^\s*function ([^ (]*)/,l="name";r&&!(l in n)&&s(n,l,{configurable:!0,get:function(){try{return o.call(this).match(i)[1]}catch(t){return""}}})},d478:function(t,e,a){"use strict";var r=a("66c0"),s=a.n(r);s.a},d620:function(t,e,a){},f820:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement;t._self._c;return t._m(0)},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"about"},[a("h1",[t._v("This is an about page")])])}],n=a("2877"),o={},i=Object(n["a"])(o,r,s,!1,null,null,null);e["default"]=i.exports}}]);
//# sourceMappingURL=about.d7733678.js.map