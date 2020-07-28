# Ep11 - Introducción a Jenkins X

**¿Qué es?**

Jenkins X es una herramienta que nos permite crear un proceso CI/CD para hacer despliegues de nuestras aplicaciones en diferentes ambientes usando Kubernetes.

## Un poco de conceptos y teoría

**Características importantes**

* **Pipelines**: El proceso CI se basa en pipelines, que no es más que un conjunto de instrucciones o scripts para construir, verificar, probar, o lo que sea que necesito para hacer un *delivery* o *deployment* de nuestra aplicación.

* **Environments**: Los *environments* nos permiten segmentar los lugares donde nuestras aplicaciones serán desplegadas. Normalmente, un *environment* corresponse a un *namespace* de Kubernetes. Por defecto Jenkins X agregar los *environments* *staging* y *production*.

    * **Temporales/Permantentes**: Un environment temporal es donde nuestros *previews* se muestran al momento de tener un MR y que se eliminan al cerrarse el MR, mientras que un environment *permanente* es donde nuestras aplicaciones permanecerán hasta que decidamos eliminarla nosotros.

* **Preview**: Jenkins X nos genera *previews* o *previsualizaciones* de nuestras aplicaciones una vez se hace un MR a la rama *master*. Estas *previews* están disponibles mientras el MR esté abierto. De esta manera la persona encargada de aceptar el MR puede revisar y comprobar que todo funciona como debe antes de enviar el código a master.

* **Promotions**: Es el llevar una aplicación a un ambiente.

    * **Automáticos/Manuales**: Una promoción *automática* es aquella que despliega una aplicación apenas esta pasa a la rama master. Por otro lado, la *manual* requiere que una persona haga la *promoción* a un environment de manera explícita.


**¿Cuándo debo considerar Jenkins X una solución para mi?**

Sí usas Kubernetes, y sigues el patrón *GitOps*, Jenkins X puede ser una solución para hacer el proceso completo de *Continuous Integration* y *Delivery/Deployment Integration*.

**¿Diferencias entre Jenkins y Jenkins X?**
*Jenkins* y *Jenkins X* son soluciones diferentes. Algunas de las cosas a tener en cuenta:

* **Jenkins** es agnóstico a la infraestructura, **Jenkins X** es diseñado exclusivamente para Kubernetes.
* **Jenkins** sólo nos ayuda a construir *CI*, **Jenkins X** nos ayuda a construir *CI/CD*.
* **Jenkins X** cuando se usa como *static server* usa **Jenkins X** en su interior para hacer los pipelines.

## Demo

Para la presentación se tiene una cuenta Google Cloud con un proyecto *lhk-ep11-jenkinx*.

* **Instalación CLI jx**: Instalar el CLI de Jenkins X de acuerdo al sistema operativo que uses. https://jenkins-x.io/docs/install-setup/install-binary/

* **Cluster y Jenkins X**: Alternativas
    * Podemos agregar Jenkins X a un clúster con un Kubernetes listo.
    * Podemos instalar un clúster y agregar Jenkins X en un mismo paso, usando el CLI. (*opción elegida*).
    * Podemos instalar un clúster con Terraform dentro de un provider y luego instalar Jenkins X con el CLI. (*opción recomendada*).

    Links para instalar Jenkins X: https://jenkins-x.io/docs/ y https://jenkins-x.io/docs/install-setup/boot/

    Para la demostración se eligió usar el CLI por comidad y facilidad. Se usó el siguiente comando, donde ya definimos que ue proyecto de Google Cloud se creará el clúster y donde hago *skip* del login por que ya tengo mi cuenta iniciada en mi computador.

    `jx create cluster gke --project-id='lhk-ep11-jenkinx' --skip-login`

    Una vez comienza el proceso, algunas configuraciones requerirán nuestra atención, por ejemplo: Proveedor de Git que queramos usar, usuario de Git que queramos usar, region donde los clúster se crearán.

* **Crear aplicación**: Hay varias formas subir una aplicación.
    * **Crear**: Se puede crear una aplicación base, donde elegimos inclusive el lenguaje de programación que vamos a usar, esto nos generará un archivos bases donde podemos comenzar a programar.

    * **Importar**: Se puede importar un proyecto desde un repositorio de Git, o podemos importarlo desde código fuente, esta fue la opción elegida.

    Dentro de la carpeta *app_python* hay una aplicación Python que corre un servicio web. Dentro de esta carpeta se ejecuta el comando `jx import`, el prompt solicitará datos como proveedor de Git y nombre de la aplicación. Una vez hecho esto Jenkins X creará un repositorio (en caso de que se importe un proyecto desde archivos locales), detectará en que lenguaje funciona el proyecto y generara archivos *Helm*, *Dockerfile* y otros cuantos necesarios para hacer el deployment necesario, hace builds de los steps, luego hará un MR al repositorio *staging* añadiendo información de la nueva aplicación y finalmente la aplicación quedará dentro del namespace *jx-staging*.

    Podemos ver la aplicación ejecutando el comando `jx get applications`.

* **Promocionar aplicación**: Para promocionar la aplicación a *producción* debe hacerse de forma manual. Ejecutar el comando `jx promote --app=<nombre-aplicacion> --version=0.0.1 --env=production`
    
    Podemos verificar que ahora tenemos la aplicación en dos *environments* ejecutando el comando `jx get applications`.

## Recursos
* [Web Jenkins X: Components in Jenkins X](https://jenkins-x.io/docs/reference/components/)
* [Web Jenkins X: CLI](https://jenkins-x.io/docs/reference/commands/)
* [Web Jenkins X: Install and setup](https://jenkins-x.io/docs/install-setup/)
* [Web Jenkins X: Creating projects](https://jenkins-x.io/docs/create-project/)
* [Youtube: De 0 a 100 en la nube con Jenkins-X](https://www.youtube.com/watch?v=5EGOr1j8K7g).
* [Youtube: Jenkins X: Easy CI/CD for Kubernetes - James Strachan, CloudBees (Intermediate Skill Level)](https://www.youtube.com/watch?v=uHe7R_iZSLU)

## Autor
* Santiago Yepes <zetogk@gmail.com>