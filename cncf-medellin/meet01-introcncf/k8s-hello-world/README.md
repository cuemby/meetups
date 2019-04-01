## K8s Hello World

### **Descripción**

proyecto de ejemplo en NodeJS para probar kubernetes


###  **Dependencias**

- [Docker](https://docs.docker.com/install/)
- [Oracle Vbox](https://www.virtualbox.org/)
- [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) 
- Editor de texto o IDE

### **Instrucciones**

- subir minikube `minikube start`
- revisar que efectivamente minikube esté arriba `minikube status`
- `eval $(minikube docker-env)` para hacer switch al docker de la VM de minikube
- construye tu imagen de  docker  `docker build -t k8s-hello-world .`
- kubectl create namespace `meetup-medellin`
- creamos nuestras variables de ambiente `kubectl apply -f .k8s/create config-map-hello-world.yaml`
- desplegamos nuesta aplicacion `kubectl apply -f .k8s/deployment-hello-world.yaml`
- damos accesibilidad a nuestra aplicación `kubectl apply -f .k8s/service-hello-world.yaml`
- obtenemos la ip de nuestro cluster `minikube ip`
- obtenemos el puerto de nuestro servicio en el cluster `kubectl get services hello-world --namespace meetup-medellin`
- probemos con `curl http://{minikubeIp}:{servicePort}/` o en el navegador
